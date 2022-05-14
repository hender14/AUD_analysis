from apiclient.discovery import build
import requests
from lxml import etree
from argparse import ArgumentParser
import numpy as np
import os
import service.gcstorage as gcstorage

API_KEY = os.environ['GCP_API_KEY']

def youtube_analysis(keywordlist, filename):
  foldername = os.environ['GCP_ANALYSIS_FOLDER']
  # APIｷｰの設定
  my_youtube = build("youtube","v3",developerKey=API_KEY)
  # json = []
  yjson = dict()

  keylen = len(keywordlist)
  for i in range(keylen):
    # 関連ﾜｰﾄﾞの抽出
    keyword = keywordlist[i]
    keylist = extract_words(keyword)

    # 再生数による動画ﾘｽﾄを作成
    response_kotlin_buffer, view_count_buffer, view_count_sum_buffer = video_views(keyword, keylist, my_youtube)

    # ﾘｽﾄの並べ替えを実施
    sort_list = np.array(view_count_sum_buffer)
    index_list = np.argsort(sort_list)
    # np.argsort(A)[::-1]
    # index = [index_list[4], index_list[3], index_list[2]]

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

  gcstorage.upload_blob_json(filename, yjson, os.environ['GCP_STORAGE_BUCKET'], foldername)

  return

def extract_words(keyword):
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

def video_views(keyword, keylist, my_youtube):
  # ﾊﾟﾗﾒｰﾀの設定
  MAX_RESULT = 5
  ORDER = "viewCount"

  # ｷｰﾜｰﾄﾞの動画を分析
  # ﾘｽﾄの初期化
  view_count_buffer = []
  view_count_sum_buffer = []
  response_kotlin_buffer = []
  view_count = []

  for ss in keylist:
    request_kotlin = []
    # int = 0
    request_kotlin = my_youtube.search().list(
      part="id, snippet",
      type="video",
      q = keyword + " " + ss,
      maxResults=MAX_RESULT,
      order=ORDER,
      regionCode="JP",
      # videoCategoryId,
      # videoDuration,
      # publishedBefore="2021-01-01T00:00:00Z",  #from  after  to  before
      publishedAfter="2021-10-01T00:00:00Z"
    )

    # youtube data APIの実行
    response_kotlin = request_kotlin.execute()
    response_kotlin_buffer.append(response_kotlin)

    # 動画の再生数を算出
    for search_result in response_kotlin.get("items", []):
      id = search_result["id"]["videoId"]
      view_count.append(my_youtube.videos().list(part = 'statistics', id = id).execute()['items'][0]['statistics']['viewCount'])
    view_count_buffer.append(view_count)

    # 再生数の合計値を算出
    view_count_sum = sum(list(map(int, view_count)))
    view_count_sum_buffer.append(view_count_sum)

  return response_kotlin_buffer, view_count_buffer, view_count_sum_buffer