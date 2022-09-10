# from requests.api import request
from requests import api
import requests
from flask import Flask, jsonify, request
from flask_cors import CORS
from dotenv import load_dotenv
import os

from controller.transcripts import Transcribe
from controller.comprehend import Comprehend
from controller.data_analysis import Analysis
from service.storage import S3

app = Flask(__name__)
app.config['JSON_AS_ASCII'] = False
# CORS(app)
load_dotenv()

#CORS設定用
@app.after_request
def after_request(response):
  response.headers.add('Access-Control-Allow-Origin', os.environ['CORS_ADDRESS'])
  response.headers.add('Access-Control-Allow-Headers', 'Content-Type,Authorization')
  response.headers.add('Access-Control-Allow-Methods', 'GET,PUT,POST,DELETE,OPTIONS')
  response.headers.add('Access-Control-Allow-Credentials', 'true')
  return response

# youtube解析用API
@app.route('/analysis', methods=["POST"])
def analysis():
  if request.method == "POST":
    file = request.files["file"]
    username = request.form["username"]
    filename = request.form["filename"]
    if not (file and username and filename):
      print('message: parameter is not enough: {}, {}'.format(username, filename))
      return jsonify({'message': 'parameter is not enough'}), 400

  sepa = filename.split(".")
  filenameo = sepa[len(sepa)-2]
  # Transcribe
  transcribe(file, username, filename, filenameo)
  # Comprehend
  keyword = comprehend(username, filenameo)
  # Analysis
  analysis(keyword, username, filenameo)

  return jsonify({'message': 'analysis has completed'}), 200


# youtube解析用APIの解析結果ﾘｽﾄ送付
@app.route('/list', methods=["GET"])
def list():
  if request.method == "GET":
    username = request.args.get('username', default="", type=str)
    if not (username):
      print('message: parameter is not enough: {}'.format(username))
      return jsonify({'message': 'parameter is not enough: {}'.format(username)}), 400

  # bucket_name = os.environ['GCP_STORAGE_BUCKET']
  # prefix = os.environ['GCP_ANALYSIS_FOLDER'] + "/" + username
  url = os.environ['AWS_LAMBDA_LIST']
  param = {'username':username}
  response = requests.get(url, params=param)
  return  response.json(), 200

# youtube解析用APIの解析結果詳細送付
@app.route('/detail', methods=["GET"])
def detail():
  if request.method == "GET":
    username = request.args.get('username', default="", type=str)
    filename = request.args.get('filename', default="", type=str)
    if not (username and filename):
      print('message: parameter is not enough: {}, {}'.format(username, filename))
      return jsonify({'message': 'parameter is not enough'}), 400
  
  # bucket_name = os.environ['GCP_STORAGE_BUCKET']
  # foldername = os.environ['GCP_ANALYSIS_FOLDER']
  # filename = foldername + "/" + username + '/' + filename
  url = os.environ['AWS_LAMBDA_DETAIL']
  param = {'username':username, 'filename':filename}
  response = requests.get(url, params=param)
  print(response.text)
  return  response.text, 200

# # youtube解析用APIのオブジェクト削除
# @app.route('/delete', methods=["GET"])
# def delete():
#   if request.method == "GET":
#     username = request.args.get('username', default="", type=str)
#     filename = request.args.get('filename', default="", type=str)
#     if not (username and filename):
#       return jsonify({'message': 'parameter is not enough'}), 400

#   bucket_name = os.environ['GCP_STORAGE_BUCKET']
#   foldername = os.environ['GCP_ANALYSIS_FOLDER']
#   filename = foldername + "/" + username + '/' + filename
#   gcstorage.delete_blob(bucket_name, filename)
#   return  jsonify({'message': 'delete is completed'}), 200

def transcribe(file, username, filename, filenameo):
  s3 = S3()
  # Transcribe Start
  filepath_transcribe = username + '/' + os.environ['AWS_TRANSLATE_FOLDER'] + '/' + filename
  s3.Put_object(file, filepath_transcribe)
  transcribe = Transcribe(username, filename)
  transcribe.start_transcription(filenameo)
  transcribe.wait_handler(filenameo)
  print('Transcribe finished')

  return

def comprehend(username, filenameo):
  s3 = S3()
  # Comprehend Start
  filepath_comprehend = username + '/' + os.environ['AWS_COMPLEHEND_FOLDER'] + '/' + filenameo + '.json'
  filepath_entities = username + '/' + os.environ['AWS_ENTITIES_FOLDER'] + '/' + filenameo + '.json'
  transcription = s3.Get_object(filepath_comprehend)

  comprehend = Comprehend()
  phrases = comprehend.detect_key_phrases(transcription)
  s3.Put_object_json(phrases, filepath_entities)
  keyword = comprehend.key_phrases_list(phrases)
  print('Comprehend finished')

  return keyword

def analysis(keyword, username, filenameo):
  s3 = S3()
  # Analysis Start
  filepath_analysis = username + '/' + os.environ['AWS_ANALYSIS_FOLDER'] + '/' + filenameo + '.json'
  analysis = Analysis()
  result = analysis.youtube_analysis(keyword)
  s3.Put_object_json(result, filepath_analysis)

  return

if __name__ == '__main__':
  app.debug = True
  portnum = os.environ['PORT']
  if not portnum:
    portnum = 8080
  app.run(host='0.0.0.0', port=portnum)
