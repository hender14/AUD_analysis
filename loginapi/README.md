# loginAPI


<a name="overview"></a>
## Overview
login API on API Gateway with a Cloud Run backend


### Version information
*Version* : 1.0.1


### URI scheme
*Schemes* : HTTPS


### Produces

* `application/json`


### Test

test is automatically done by github action


<a name="paths"></a>
## Paths

<a name="forgot"></a>
### send email for resetting password
```
POST /app/forgot
```


#### Parameters

|Type|Name|Schema|
|---|---|---|
|**Body**|**forgot**  <br>*optional*|[Forgot](#forgot)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|A successful response|string|


#### Example HTTP request

##### Request path
```
/app/forgot
```


##### Request body
```json
{
  "email" : "string"
}
```


#### Example HTTP response

##### Response 200
```json
"string"
```


<a name="authcorsforgot"></a>
### cors authentication
```
OPTIONS /app/forgot
```


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**204**|A successful response|No Content|


#### Example HTTP request

##### Request path
```
/app/forgot
```


<a name="login"></a>
### login on web
```
POST /app/login
```


#### Parameters

|Type|Name|Schema|
|---|---|---|
|**Body**|**login**  <br>*optional*|[Login](#login)|


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


<a name="authcorslogin"></a>
### cors authentication
```
OPTIONS /app/login
```


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**204**|A successful response|No Content|


#### Example HTTP request

##### Request path
```
/app/login
```


<a name="logout"></a>
### logout on web
```
GET /app/logout
```


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|A successful response|string|


#### Example HTTP request

##### Request path
```
/app/logout
```


#### Example HTTP response

##### Response 200
```json
"string"
```


<a name="authcorslogout"></a>
### cors authentication
```
OPTIONS /app/logout
```


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**204**|A successful response|No Content|


#### Example HTTP request

##### Request path
```
/app/logout
```


<a name="register"></a>
### register account on web
```
POST /app/register
```


#### Parameters

|Type|Name|Schema|
|---|---|---|
|**Body**|**register**  <br>*optional*|[Register](#register)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|A successful response|string|


#### Example HTTP request

##### Request path
```
/app/register
```


##### Request body
```json
{
  "first_name" : "string",
  "last_name" : "string",
  "email" : "string",
  "password" : "string",
  "password_connfirm" : "string"
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
OPTIONS /app/register
```


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**204**|A successful response|No Content|


#### Example HTTP request

##### Request path
```
/app/register
```


<a name="reset"></a>
### reset password
```
POST /app/reset
```


#### Parameters

|Type|Name|Schema|
|---|---|---|
|**Body**|**reset**  <br>*optional*|[Reset](#reset)|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|A successful response|string|


#### Example HTTP request

##### Request path
```
/app/reset
```


##### Request body
```json
{
  "token" : "string",
  "password" : "string",
  "password_confirm" : "string"
}
```


#### Example HTTP response

##### Response 200
```json
"string"
```


<a name="authcorsreset"></a>
### cors authentication
```
OPTIONS /app/reset
```


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**204**|A successful response|No Content|


#### Example HTTP request

##### Request path
```
/app/reset
```


<a name="getuser"></a>
### get user infomation
```
GET /app/user
```


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|A successful response|No Content|


#### Example HTTP request

##### Request path
```
/app/user
```


<a name="authcorsuser"></a>
### cors authentication
```
OPTIONS /app/user
```


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**204**|A successful response|No Content|


#### Example HTTP request

##### Request path
```
/app/user
```




<a name="definitions"></a>
## Definitions

<a name="forgot"></a>
### Forgot

|Name|Description|Schema|
|---|---|---|
|**email**  <br>*required*|**Example** : `"string"`|string|


<a name="login"></a>
### Login

|Name|Description|Schema|
|---|---|---|
|**email**  <br>*required*|**Example** : `"string"`|string|
|**password**  <br>*required*|**Example** : `"string"`|string|


<a name="register"></a>
### Register

|Name|Description|Schema|
|---|---|---|
|**email**  <br>*required*|**Example** : `"string"`|string|
|**first_name**  <br>*required*|**Example** : `"string"`|string|
|**last_name**  <br>*required*|**Example** : `"string"`|string|
|**password**  <br>*required*|**Example** : `"string"`|string|
|**password_connfirm**  <br>*optional*|**Example** : `"string"`|string|


<a name="reset"></a>
### Reset

|Name|Description|Schema|
|---|---|---|
|**password**  <br>*required*|**Example** : `"string"`|string|
|**password_confirm**  <br>*required*|**Example** : `"string"`|string|
|**token**  <br>*required*|**Example** : `"string"`|string|