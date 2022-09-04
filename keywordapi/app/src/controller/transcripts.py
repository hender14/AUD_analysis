import os
import time
import boto3

class Transcribe:
  def __init__(self, username, filename):
    self.bucketname = os.environ['AWS_STORAGE_BUCKET']
    self.foldernameIn = os.environ['AWS_TRANSLATE_FOLDER']
    self.foldernameOut = os.environ['AWS_COMPLEHEND_FOLDER']
    self.username = username
    self.filename = filename
    self.media_file_url = 's3://' + self.bucketname + '/' + self.username + '/' + self.foldernameIn + '/' + self.filename
    self.client = boto3.client('transcribe')

  def start_transcription(self, job_name):
    response = self.client.start_transcription_job(
      TranscriptionJobName=job_name,
      LanguageCode="ja-JP",
      Media={
      'MediaFileUri': self.media_file_url
      },
      OutputBucketName=self.bucketname,
      OutputKey = self.username + '/' + self.foldernameOut + '/'
    )
    return

  def wait_handler(self, job_name):
      # client = boto3.client('transcribe')
      current_status = self.get_status(job_name)
      while current_status != 'COMPLETED':
        print("Waiting for operation to complete... Status: {}".format(current_status))
        time.sleep(60)
        current_status = self.get_status(job_name)
        print(current_status)
      return

  def get_status(self,filenameo):
    status = self.client.get_transcription_job(
        TranscriptionJobName=filenameo
    )
    return status['TranscriptionJob']['TranscriptionJobStatus']

  def delete_job(self):
    self.client.delete_transcription_job(TranscriptionJobName=self.filename)
    return