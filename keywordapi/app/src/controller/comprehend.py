import os
from google.cloud import language_v1
import numpy as np
import service.gcstorage as gcstorage

def sample_analyze_entities(filename):
  foldername = os.environ['GCP_COMPLEHEND_FOLDER']
  client = language_v1.LanguageServiceClient()

  # Available types: PLAIN_TEXT, HTML
  type_ = language_v1.Document.Type.PLAIN_TEXT

  language = "ja"
  filepath = "gs://" + os.environ['GCP_STORAGE_BUCKET'] + "/" + foldername + "/" + filename + ".txt"
  document = {"gcs_content_uri": filepath, "type_": type_, "language": language}

  # Available values: NONE, UTF8, UTF16, UTF32
  encoding_type = language_v1.EncodingType.UTF8

  response = client.analyze_entities(request = {'document': document, 'encoding_type': encoding_type})

  cjson = []
  score = []
  item = []
  index_list = []
  keyword = []
  confirm = []

  # Loop through entitites returned from the API
  for entity in response.entities:
    print(u"Representative name for the entity: {}".format(entity.name))
    # Get entity type, e.g. PERSON, LOCATION, ADDRESS, NUMBER, et al
    print(u"Entity type: {}".format(language_v1.Entity.Type(entity.type_).name))
    # Get the salience score associated with the entity in the [0, 1.0] range
    print(u"Salience score: {}".format(entity.salience))
    cjson = {"name": entity.name, "type": language_v1.Entity.Type(entity.type_).name, "score": entity.salience}
    item.append(entity.name)
    score.append(entity.salience)
    confirm.append(cjson)

    for metadata_name, metadata_value in entity.metadata.items():
        print(u"{}: {}".format(metadata_name, metadata_value))

    for mention in entity.mentions:
      print(u"Mention text: {}".format(mention.text.content))
      # Get the mention type, e.g. PROPER for proper noun
      print(
          u"Mention type: {}".format(language_v1.EntityMention.Type(mention.type_).name)
      )
      # print("print: " + str(entity.salience))
  length = len(score) - 1
  sort_list = np.array(score)
  index_list = np.argsort(sort_list)

  print(item)

  # np.argsort(A)[::-1]
  for index in range(length-2, length-1):
    keyword.append(item[index_list[index]])
    print(keyword)

  print(u"Language of the text: {}".format(response.language))
  foldername = os.environ['GCP_ENTITIES_FOLDER']
  gcstorage.upload_blob_json(filename, cjson, os.environ['GCP_STORAGE_BUCKET'], foldername)

  return keyword