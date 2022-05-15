# testAPI


<a name="overview"></a>
## Overview
Sample API on API Gateway with a Cloud Run backend


### Version information
*Version* : 1.0.0


### URI scheme
*Schemes* : HTTPS


### Produces

* `application/json`




<a name="paths"></a>
## Paths

<a name="hello"></a>
### Cloud Run hello world
```
POST /app/login
```


#### Parameters

|Type|Name|Schema|
|---|---|---|
|**Body**|**user**  <br>*optional*|[User](#user)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|A successful response|string|


#### Example HTTP request

##### Request path
```
/app/login
```


##### Request body
```json
{
  "email" : "string",
  "password" : "string"
}
```


#### Example HTTP response

##### Response 200
```json
"string"
```


<a name="corshelloworld"></a>
### OPTIONS /app/login

#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**204**|A successful response|No Content|


#### Example HTTP request

##### Request path
```
/app/login
```




<a name="definitions"></a>
## Definitions

<a name="user"></a>
### User

|Name|Description|Schema|
|---|---|---|
|**email**  <br>*required*|**Example** : `"string"`|string|
|**password**  <br>*required*|**Example** : `"string"`|string|





