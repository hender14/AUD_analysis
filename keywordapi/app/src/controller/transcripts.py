import os
from google.cloud import speech_v1p1beta1 as speech
import service.gcstorage as gcstorage

# GCP APIﾃｽﾄ用
def transcribe_gcs(filename, filenameo):
  foldername = os.environ['GCP_TRANSLATE_FOLDER']
  client = speech.SpeechClient()

  audio = {"uri": "gs://" + os.environ['GCP_STORAGE_BUCKET'] + "/" + foldername + "/" + filename }
  config = {
    "encoding": "MP3",
    "sample_rate_hertz": 22050,
    "language_code": "ja-JP",
  }
  request = {
    "config": config,
    "audio": audio,
    # "output_config": output_config,
  }

  operation = client.long_running_recognize(request=request)
  # operation = client.long_running_recognize(config=config, audio=audio, output_config=output_config)

  print("Waiting for operation to complete...")
  response = operation.result(timeout=90)

  for result in response.results:
    # The first alternative is the most likely one for this portion.
    print(u"Transcript: {}".format(result.alternatives[0].transcript))
    print("Confidence: {}".format(result.alternatives[0].confidence))

  uploadfile = result.alternatives[0].transcript
  foldername = os.environ['GCP_COMPLEHEND_FOLDER']
  gcstorage.upload_blob_normal(filenameo, uploadfile, os.environ['GCP_STORAGE_BUCKET'], foldername)

  return