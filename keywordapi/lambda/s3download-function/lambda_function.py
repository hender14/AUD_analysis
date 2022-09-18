import json
import boto3
import datetime

client = boto3.client('s3')

def lambda_handler(event, context):
   
    username = event["queryStringParameters"]['username']
    filename = event["queryStringParameters"]['filename']
    bucketname = "python-install-sample-12"
    filepath = username + "/analysis/" + filename
    # prefix = username + "/analysis"
    try:
        response = client.get_object(Bucket=bucketname, Key=filepath)
        body = response['Body'].read()
        print(body)
        return {
            'statusCode': 200,
            'body': body,
            'isBase64Encoded': False,
            'headers': {}
        }
    
    except Exception as e:
        print(e)
        raise e