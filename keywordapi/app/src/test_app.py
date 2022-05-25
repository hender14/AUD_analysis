import os
import pytest
from app import app
from dotenv import load_dotenv
import service.gcstorage as gcstorage
load_dotenv()
gcstorage.make_keyfile()

###############
## TEST CASE ##
###############
def test_analysis_get_status():
	app.config['TESTING'] = True
	client = app.test_client()
	result = client.post('/analysis', json={'username': os.environ['TESTUSER_ID'], 'filename': os.environ['TESTUSER_TRANSLATE_FILENAME']})
	assert result.status_code == 200

def test_list_get_status():
	app.config['TESTING'] = True
	client = app.test_client()
	result = client.get('/list', query_string={'username': os.environ['TESTUSER_ID']})
	ans = [os.environ['TESTUSER_ANALYSIS_FILENAME']]
	assert result.status_code == 200
	assert ans == result.json["name"]


def test_detail_get_status():
	app.config['TESTING'] = True
	client = app.test_client()
	result = client.get('/detail', query_string={'username': os.environ['TESTUSER_ID'], 'filename': os.environ['TESTUSER_ANALYSIS_FILENAME']})
	print(result.json)
	assert result.status_code == 200

def test_delete_get_status():
	app.config['TESTING'] = True
	client = app.test_client()
	result = client.get('/delete', query_string={'username': os.environ['TESTUSER_ID'], 'filename': os.environ['TESTUSER_ANALYSIS_FILENAME']})
	assert result.status_code == 200