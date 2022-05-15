# keywordAPI


<a name="overview"></a>
## Overview
keyword API on API Gateway with a Cloud Run backend


### Version information
*Version* : 1.0.1


### URI scheme
*Schemes* : HTTPS


### Produces

* `application/json`




<a name="paths"></a>
## Paths

<a name="analysis"></a>
### analysis files
```
POST /analysis
```


#### Parameters

|Type|Name|Schema|
|---|---|---|
|**Body**|**analysis**  <br>*optional*|[Analysis](#analysis)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|A successful response|string|


#### Example HTTP request

##### Request path
```
/analysis
```


##### Request body
```json
{
  "username" : "string",
  "filename" : "string"
}
```


#### Example HTTP response

##### Response 200
```json
"string"
```


<a name="authcorsregister"></a>
### cors authentication
```
OPTIONS /analysis
```


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**204**|A successful response|No Content|


#### Example HTTP request

##### Request path
```
/analysis
```


<a name="detail"></a>
### get detail infomation
```
GET /detail
```


#### Parameters

|Type|Name|Schema|
|---|---|---|
|**Body**|**detail**  <br>*optional*|[Detail](#detail)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|A successful response|string|


#### Example HTTP request

##### Request path
```
/detail
```


##### Request body
```json
{
  "username" : "string",
  "filename" : "string"
}
```


#### Example HTTP response

##### Response 200
```json
"string"
```


<a name="authcorslogout"></a>
### cors authentication
```
OPTIONS /detail
```


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**204**|A successful response|No Content|


#### Example HTTP request

##### Request path
```
/detail
```


<a name="getlist"></a>
### get list infomation
```
GET /list
```


#### Parameters

|Type|Name|Schema|
|---|---|---|
|**Body**|**list**  <br>*optional*|[List](#list)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|A successful response|No Content|


#### Example HTTP request

##### Request path
```
/list
```


##### Request body
```json
{
  "username" : "string"
}
```


<a name="authcorsuser"></a>
### cors authentication
```
OPTIONS /list
```


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**204**|A successful response|No Content|


#### Example HTTP request

##### Request path
```
/list
```




<a name="definitions"></a>
## Definitions

<a name="analysis"></a>
### Analysis

|Name|Description|Schema|
|---|---|---|
|**filename**  <br>*required*|**Example** : `"string"`|string|
|**username**  <br>*required*|**Example** : `"string"`|string|


<a name="detail"></a>
### Detail

|Name|Description|Schema|
|---|---|---|
|**filename**  <br>*required*|**Example** : `"string"`|string|
|**username**  <br>*required*|**Example** : `"string"`|string|


<a name="list"></a>
### List

|Name|Description|Schema|
|---|---|---|
|**username**  <br>*required*|**Example** : `"string"`|string|





