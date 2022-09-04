from apiclient.discovery import build
import requests
from lxml import etree
import numpy as np
import os

class Analysis:
  def __init__(self):
    self.youtube = build("youtube","v3",developerKey=os.environ['GCP_API_KEY']) # APIｷｰの設定

  def youtube_analysis(self, keywordlist):
    yjson = dict()

    keylen = len(keywordlist)
    for i in range(keylen):
      # 関連ﾜｰﾄﾞの抽出
      keyword = keywordlist[i]
      keylist = self.extract_words(keyword)

      # 再生数による動画ﾘｽﾄを作成
      response_kotlin_buffer, view_count_buffer, view_count_sum_buffer = self.video_views(keyword, keylist)

      # ﾘｽﾄの並べ替えを実施
      sort_list = np.array(view_count_sum_buffer)
      index_list = np.argsort(sort_list)

      index_i = 0
      listnum = len(response_kotlin_buffer) - 1
      # 再生数上位の項目を作成
      for j in range(3):
        if (len(response_kotlin_buffer[listnum - j]["items"]) == 0):
          listnum -= 1
        index_i = index_list[listnum - j]
        print(index_i)
        knum = len(response_kotlin_buffer[index_i]["items"])
        for k in range(3):
          yjson["key" + str(i) + str(j) + str(k)] = {"size": {"key": j+1, "play": k+1}, 
            "keyword": keyword + " " + keylist[index_i],
            "title": response_kotlin_buffer[index_i]["items"][k]["snippet"]["title"], 
            "概要": response_kotlin_buffer[index_i]["items"][k]["snippet"]["description"], 
            "サムネイル": response_kotlin_buffer[index_i]["items"][k]["snippet"]["thumbnails"]["default"]["url"],
            "動画url": "https://www.youtube.com/watch?v=" + response_kotlin_buffer[index_i]["items"][k]["id"]["videoId"], 
            "再生数": view_count_buffer[index_i][k] 
          }
          if k == knum - 1:
            break

    return yjson

  def extract_words(self, keyword):
    tlen = len(keyword)
    r = requests.get("http://clients1.google.com/complete/search", params={'q':keyword, 'hl':'ja', 'ie':'utf_8', 'oe':'utf_8', 'output': 'toolbar'})
    root = etree.XML(r.text)
    sugs = root.xpath('//suggestion')
    sugstrs = [s.get('data') for s in sugs]
    squares = []
    for ss in sugstrs:
        if ss[tlen:tlen+1] == ' ':
          squares.append(ss[tlen:].lstrip())
        else :
          squares.append(ss)
    keylist = squares[1:6]
    return keylist

  def video_views(self, keyword, keylist):
    # ﾘｽﾄの初期化
    view_count_buffer = []
    view_count_sum_buffer = []
    response_kotlin_buffer = []
    view_count = []

    # ﾊﾟﾗﾒｰﾀの設定
    MAX_RESULT = 5
    ORDER = "viewCount"

    # ｷｰﾜｰﾄﾞの動画を分析
    for ss in keylist:
      request_kotlin = []
      request_kotlin = self.youtube.search().list(
        part="id, snippet",
        type="video",
        q = keyword + " " + ss,
        maxResults=MAX_RESULT,
        order=ORDER,
        regionCode="JP",
        publishedAfter="2021-10-01T00:00:00Z"
      )

      # youtube data APIの実行
      response_kotlin = request_kotlin.execute()
      response_kotlin_buffer.append(response_kotlin)

      # 動画の再生数を算出
      for search_result in response_kotlin.get("items", []):
        id = search_result["id"]["videoId"]
        view_count.append(self.youtube.videos().list(part = 'statistics', id = id).execute()['items'][0]['statistics']['viewCount'])
      view_count_buffer.append(view_count)

      # 再生数の合計値を算出
      view_count_sum = sum(list(map(int, view_count)))
      view_count_sum_buffer.append(view_count_sum)

    return response_kotlin_buffer, view_count_buffer, view_count_sum_buffer