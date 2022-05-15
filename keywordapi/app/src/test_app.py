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

###############
## TEST END ###
###############

# def test_clean():
#   clean()
#   assert True


########################
## Utility. Not tests ##
########################

# def clean():
#   r = get(f'{BASEURL}/keys/')
#   for key in r.json():
#     delete(f'{BASEURL}/keys/{key}')
#   num_keys = len(get(f'{BASEURL}/keys/').json())
#   assert 0 == num_keys

# def clean_and_add_keys():
#   clean()
#   r = put(f'{BASEURL}/keys/apple', data='red')
#   assert r.status_code == 200
#   r = put(f'{BASEURL}/keys/banana', data='yellow')
#   assert r.status_code == 200