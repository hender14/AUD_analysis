import os
import pytest
from app import app
from dotenv import load_dotenv
import service.gcstorage as gcstorage
load_dotenv()
gcstorage.make_keyfile()

##########
## KEYS ##
##########

# def test_keys_get():
#   r = get(f'{BASEURL}/keys/')
#   assert r.status_code == 200

# def test_keys_get_wrongurl():
#   r = get(f'{BASEURL}/key/')
#   assert r.status_code == 404

# def test_keys_post():
#   r = post(f'{BASEURL}/keys/')
#   assert r.status_code == 405

# def test_keys_put():
#   r = put(f'{BASEURL}/keys/')
#   assert r.status_code == 405

# def test_keys_delete():
#   r = delete(f'{BASEURL}/keys/')
#   assert r.status_code == 405

# 一時的
# def test_version():
#   assert pytest.__version__ == '5.2.0'


#############
## KEY GET ##
#############
# 一時的
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

# 一時的
# def test_key_get_exist():
#   # clean_and_add_keys()
#   # info = {'keyword1':'テスト', 'keyword2':'結果'}
#   info = {'keyword1': 'test'}
#   r = get(f'{BASEURL}/analysis', params=info)
#   # r = get(f'{BASEURL}?keyword1=test')
#   assert r.status_code == 200

# 一時的
# def test_key_get_exist_confirm():
#   # clean_and_add_keys()
#   # info = {'keyword1':'テスト', 'keyword2':'結果'}
#   info = {'keyword1': "test"}
#   r = get(f'{BASEURL}/analysis', params=info)
#   print(r.url)
#   # print(r.params)
#   # r = get(f'{BASEURL}?keyword1=test')
#   assert r.status_code == 200
#   # assert r.json() == {'apple':'red'}
#   # user = json.loads(json_data, object_hook=ObjectLike)

#   # print(type(user))            # <class '__main__.ObjectLike'>
#   # print(user.name)             # "taro"
#   # print(user.home.city)        # "Osaka"
#   # print(user.home.prefecture)  # None (存在しないキーなので)

# 一時的
# def test_key_get_exist_check():
#   # clean_and_add_keys()
#   info = {'keyword1': 'test'}
#   r = get(f'{BASEURL}/test', params=info)
#   assert r.status_code == 200

# def test_key_get_notalnum1():
#   # clean_and_add_keys()
#   r = get(f'{BASEURL}/analysis?keyword1="test"')
#   assert r.status_code == 200

# def test_key_get_notalnum2():
#   clean_and_add_keys()
#   r = get(f'{BASEURL}/keys/あいうえお')
#   assert r.status_code == 400

##############
## KEY POST ##
##############

# def test_key_post_notexist():
#   clean_and_add_keys()
#   r = post(f'{BASEURL}/keys/grape', data='purple')
#   assert r.status_code == 200
#   assert r.json() == {'grape':'purple'}

# def test_key_post_exist():
#   clean_and_add_keys()
#   r = post(f'{BASEURL}/keys/apple', data='green')
#   assert r.status_code == 409

# def test_key_post_notalnum1():
#   clean_and_add_keys()
#   r = post(f'{BASEURL}/keys/gr_pe', data='purple')
#   assert r.status_code == 400

# def test_key_post_notalnum2():
#   clean_and_add_keys()
#   r = post(f'{BASEURL}/keys/grape', data='pu_ple')
#   assert r.status_code == 400

# def test_key_post_notalnum3():
#   clean_and_add_keys()
#   r = post(f'{BASEURL}/keys/あいうえお', data='purple')
#   assert r.status_code == 400

# def test_key_post_notalnum4():
#   clean_and_add_keys()
#   r = post(f'{BASEURL}/keys/grape', data='あいうえお'.encode())
#   assert r.status_code == 400

#############
## KEY PUT ##
#############

# def test_key_put_notexist():
#   clean_and_add_keys()
#   r = put(f'{BASEURL}/keys/grape', data='purple')
#   assert r.status_code == 200
#   assert r.json() == {'grape':'purple'}

# def test_key_put_exist():
#   clean_and_add_keys()
#   r = put(f'{BASEURL}/keys/apple', data='green')
#   assert r.status_code == 200
#   assert r.json() == {'apple':'green'}

# def test_key_put_notalnum1():
#   clean_and_add_keys()
#   r = put(f'{BASEURL}/keys/gr_pe', data='purple')
#   assert r.status_code == 400

# def test_key_put_notalnum2():
#   clean_and_add_keys()
#   r = put(f'{BASEURL}/keys/grape', data='pu_ple')
#   assert r.status_code == 400

# def test_key_put_notalnum3():
#   clean_and_add_keys()
#   r = put(f'{BASEURL}/keys/あいうえお', data='purple')
#   assert r.status_code == 400

# def test_key_put_notalnum4():
#   clean_and_add_keys()
#   r = put(f'{BASEURL}/keys/grape', data='あいうえお'.encode())
#   assert r.status_code == 400

################
## KEY DELETE ##
################

# def test_key_delete_exist():
#   clean_and_add_keys()
#   r = delete(f'{BASEURL}/keys/apple')
#   assert r.status_code == 200
#   assert r.json() == {}

# def test_key_delete_notexist():
#   clean_and_add_keys()
#   r = delete(f'{BASEURL}/keys/grape')
#   assert r.status_code == 404

# def test_key_delete_notalnum1():
#   clean_and_add_keys()
#   r = delete(f'{BASEURL}/keys/ap_le')
#   assert r.status_code == 400

# def test_key_delete_notalnum1():
#   clean_and_add_keys()
#   r = delete(f'{BASEURL}/keys/あいうえお')
#   assert r.status_code == 400


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