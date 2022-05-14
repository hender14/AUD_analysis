# from requests.api import request
from requests import api
from flask import Flask, jsonify, request
from flask_cors import CORS
from dotenv import load_dotenv
import os
import controller.data_analysis as data_analysis
import controller.transcripts as transcripts
import controller.comprehend as comprehend
import service.gcstorage as gcstorage

app = Flask(__name__)
app.config['JSON_AS_ASCII'] = False
CORS(app)
load_dotenv()
gcstorage.make_keyfile()

@app.route('/', methods=["GET"])
def index():
#   return jsonify({
#     "message": "テスト!!"
#   })
  return "root"
  # return jsonify({'message': 'parameter is not enough'}), 200

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
# @app.route('/analysis', methods=["GET"])
def analysis():
  if request.method == "POST":
  # if request.method == "GET":
    # username = request.args.get('username', default="", type=str)
    # filename = request.args.get('filename', default="", type=str)
    username = request.json["username"]
    filename = request.json["filename"]
    if not (username and filename):
      return jsonify({'message': 'parameter is not enough'}), 400
    # else:
    #   return jsonify({'message': 'OK'}), 200

  sepa = filename.split(".")
  filenameo = sepa[len(sepa)-2]

  transcripts.transcribe_gcs(filename, filenameo)
  keyword = comprehend.sample_analyze_entities(filenameo)
  # payload = request.json
  # name = payload.get('name')
  # age = payload.get('age')

  data_analysis.youtube_analysis(keyword, filenameo)

  return jsonify({'message': 'analysis has completed'}), 200

# youtube解析用APIの解析結果ﾘｽﾄ送付
@app.route('/list', methods=["GET"])
def list():
  if request.method == "GET":
    username = request.args.get('username', default="", type=str)
    if not (username):
      return jsonify({'message': 'parameter is not enough'}), 400

  bucket_name = os.environ['GCP_STORAGE_BUCKET']
  prefix = os.environ['GCP_ANALYSIS_FOLDER'] + "/" + username
  blobs = gcstorage.list_blobs_with_prefix(bucket_name, prefix)
  return  jsonify(blobs), 200

# youtube解析用APIの解析結果詳細送付
@app.route('/detail', methods=["GET"])
def detail():
  if request.method == "GET":
    username = request.args.get('username', default="", type=str)
    filename = request.args.get('filename', default="", type=str)
    if not (username and filename):
      return jsonify({'message': 'parameter is not enough'}), 400
  
  bucket_name = os.environ['GCP_STORAGE_BUCKET']
  foldername = os.environ['GCP_ANALYSIS_FOLDER']
  filename = foldername + "/" + username + '/' + filename
  print(bucket_name)
  print(filename)
  contents = gcstorage.download_blob_into_memory(bucket_name, filename)
  return contents

if __name__ == '__main__':
  app.debug = True
  aport = os.environ['PORT']
  if not aport:
    port = 8080
  app.run(host='0.0.0.0', port=aport)
