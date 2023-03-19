import json
import numpy as np
import boto3

class Comprehend:
  def __init__(self):
    self.client = boto3.client('comprehend')

  def detect_key_phrases(self, response):
    # response = self.client_s3.get_object(Bucket=self.bucketname, Key=self.filepath)
    body = json.load(response['Body'])
    transcript = body['results']['transcripts'][0]['transcript']

    response = self.client.detect_key_phrases(
        Text=transcript, LanguageCode='ja')
    phrases = response['KeyPhrases']
    return phrases

  def key_phrases_list(self, phrases):
    score = []
    item = []
    keyword = []

    for entity in phrases:
      item.append(entity['Text'])
      score.append(entity['Score'])

    length = len(score) - 1
    sort_list = np.array(score)
    index_list = np.argsort(sort_list)

    for index in range(length-2, length):
      keyword.append(item[index_list[index]])
    print(keyword)
    return keyword