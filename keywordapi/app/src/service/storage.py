import boto3
import json
import os

class S3():
  def __init__(self):
    self.client = boto3.client('s3')
    self.bucketname = os.environ['AWS_STORAGE_BUCKET']
  
  def Get_object(self, filepath):
    response = self.client.get_object(Bucket=self.bucketname, Key=filepath)
    return response

  def Put_object_json(self, body, filepath):
    body_json = json.dumps(body)
    response = self.client.put_object(Bucket=self.bucketname, Body= body_json, Key=filepath)
    return response

  def Put_object(self, body, filepath):
    response = self.client.put_object(Bucket=self.bucketname, Body= body, Key=filepath)
    return response