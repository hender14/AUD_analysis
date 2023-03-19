import json
import boto3
import datetime

client = boto3.client('s3')

def lambda_handler(event, context):
   
    username = event["queryStringParameters"]['username']
    bucketname = "python-install-sample-12"
    prefix = username + "/analysis"
    try:
        namelist = []
        genlist = []
        leng = len(prefix) + 1
        obj = client.list_objects_v2(Bucket=bucketname, Prefix=prefix)
        print(obj['Contents'])
    
        for content in obj.get('Contents'):
          print(content.get('Key'))
          namelist.append(content.get('Key')[leng:])
          genlist.append(content.get('LastModified').strftime('%Y%m%d-%H%M%S'))
    
        json2 = {"name": namelist, "generation": genlist}
        return {
            'statusCode': 200,
            'body': json.dumps(json2),
            'isBase64Encoded': False,
            'headers': {}
        }
    
    except Exception as e:
        print(e)
        raise e