import os
import pytest
from app import app
from dotenv import load_dotenv
import service.gcstorage as gcstorage
load_dotenv()
gcstorage.make_keyfile()

#############
## KEY GET ##
#############
def test_key_get_exist():
    app.config['TESTING'] = True
    client = app.test_client() 
    result = client.get('/')
    assert b'root' == result.data

def test_key_get_status():
    app.config['TESTING'] = True
    client = app.test_client() 
    result = client.get('/')
    assert result.status_code == 200

def test_list_get_status():
    app.config['TESTING'] = True
    client = app.test_client()
    result = client.get('/list', query_string={'username': os.environ['TESTUSER_ID']})
    # ans = {'name': ['1650899873639-h4b_serverless_4.json'], 'generation': [1650899887475]}
    assert result.status_code == 200
    # assert ans == result.json