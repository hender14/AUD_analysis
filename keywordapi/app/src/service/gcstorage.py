import base64
from google.cloud import storage
import json
import os

sa_key =  os.environ['GOOGLE_APPLICATION_CREDENTIALS']

def make_keyfile():
  # sa_key =  os.environ['GOOGLE_APPLICATION_CREDENTIALS']
  data_encode_bytes = os.environ['GCP_KEYFILE_JSON']
  res_data = base64.b64decode(data_encode_bytes)
  key = res_data.decode()
  with open(sa_key, mode='w') as f:
    f.write(key)
  return

def upload_blob_normal( filename, response, bucketname, foldername):
  client = storage.Client.from_service_account_json(sa_key)
  # https://console.cloud.google.com/storage/browser/[bucket-id]/
  bucket = client.get_bucket(bucketname)

  blob = bucket.blob( foldername + "/" + filename + ".txt" )
  # blob.upload_from_file(response)
  # print(type(response))
  blob.upload_from_string(response, content_type="text/plain")
  # blob.upload_from_string(io.BytesIO(response).read())

  return

def upload_blob_json( filename, response, bucketname, foldername):  
  client = storage.Client.from_service_account_json(sa_key)
  # https://console.cloud.google.com/storage/browser/[bucket-id]/
  bucket = client.get_bucket(bucketname)

  blob = bucket.blob( foldername + "/" + filename + ".json" )
  # blob.upload_from_file(response)
  # print(type(response))
  blob.upload_from_string(json.dumps(response), content_type="text/plain")
  # blob.upload_from_string(io.BytesIO(response).read())

  return

def download_blob_into_memory(bucketname, filename):
  client = storage.Client.from_service_account_json(sa_key)
  bucket = client.bucket(bucketname)

  blob = bucket.blob( filename )
  # blob = bucket.blob( username + "/" + filename + ".json")
  contents = blob.download_as_bytes()
  # print(type(contents))
  print(
      "Downloaded storage object {} from bucket {} as the following string: {}.".format(
          blob, bucketname, contents
      )
  )
  # detail = contents.decode('unicode-esca')

  return contents

def list_blobs_with_prefix(bucket_name, prefix, delimiter=None):
  """Lists all the blobs in the bucket."""
  # list = []
  namelist = []
  genlist = []
  leng = len(prefix) + 1
  print(leng)
  # storage_client = storage.Client()
  client = storage.Client.from_service_account_json(sa_key)
  # Note: Client.list_blobs requires at least package version 1.17.0.
  blobs = client.list_blobs(bucket_name, prefix=prefix, delimiter=delimiter)

  for blob in blobs:
    print(blob.name)
    namelist.append(blob.name[leng:])
    # list.append(blob.name[leng:])
    # genlist.append(str(blob.generation)[:-3])
    genlist.append(int(blob.generation/1000))
  # if delimiter:
  #   print("Prefixes:")
  #   for prefix in blobs.prefixes:
  #     print(prefix)

  json = {"name": namelist, "generation": genlist}
  # json = {"list": list}
  print(json)

  return json

def delete_blob(bucket_name, blob_name):
  """Deletes a blob from the bucket."""
  client = storage.Client.from_service_account_json(sa_key)

  bucket = client.bucket(bucket_name)
  blob = bucket.blob(blob_name)
  blob.delete()

  print("Blob {} deleted.".format(blob_name))
  
  return