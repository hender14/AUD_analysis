*作成途中であり、日本語/英語が入り混じっています。ご了承ください。

# AUD analysis


## Service abstract
動画・音声ファイルのｷーﾜーﾄﾞを算出するサービスにです。必要な操作は非常に容易で、動画をアップロードするのみ。動画をアップロードして頂ければ、本アプリケーションが内容を解析して動画に関するアドバイスを提供します。


![](static/sample.jpeg)

## URL

[https://gold-cycling-307817.an.r.appspot.com](https://gold-cycling-307817.an.r.appspot.com)

## Background

動画配信サービスは、スマートフォンやインタｰネットインターネットの普及によって年々利用者が増えています。その中でもYoutubeは基本使用料が無料であることもあり、突出して人気があります。​

ただし、動画配信者目線で見た場合インフルエンサーや動画が飽和しており再生数を伸ばすことが困難となってきています。​

その中で、動画の再生数を増加させるためにはユーザの検索に引っかかる必要があり、そのためにも検索上位を狙いたいｷーﾜーﾄﾞを決定する事が重要となっていますが、このｷーﾜーﾄﾞを算出する手法が確立されていない事が現状になります。そのため、ｷｰﾜｰﾄﾞをデータ分析によって算出する​事を目的として、本サービスを作成しました。

上記を実施することで、再生数を増加させる事も勿論ですが、データ分析にて自動化する事でクリエイターの負担を軽減する事にも繋がります。​

# Architecture

Adopts microservice architecture. The backend is managed individually by deploying it in the form of API for each service.

User registration and login information are controlled between FrontEnd-BackEnd: loginAPI, video, audio information are analyzed and keyword information is estimated between FrontEnd-BackEnd: keywordAPI.

![](static/Architecture.svg)
** There are some parts that are in the process of being created, and there are parts that do not reflect the above architecture.

Please refer to the following API specifications for the details of the created API.

* [loginAPI](https://github.com/hender14/AUD_analysis/tree/main/loginapi)

* [keywordAPI](https://github.com/hender14/AUD_analysis/tree/main/keywordapi)

## Security

* Password encryption
* Access control by CORS between FrontEnd and BackEnd 
* Login authentication with JWT token
* Application permission management with GCP IAM to access between Gateway and BackEnd
* Access control to BackEnd with GCP API GATEWAY
* Manage Google Cloud Storage bucket Access rules


## CI/CD

test/build/uploading registry are automatically done by github action.

deploy is done by terraform.

# Apprication detail


## FrontEnd
### skill
* vue.js framework (typescript)

### deploy environment
* Google App Engine (Google PaaS service)

### DFD structure
under creating

### Directry figure
![](static/Component/Vue.js.svg)


## BackEnd


## loginAPI

### skill
* gin framework (golang)

### deploy environment
* Google Cloud Run (Google container managed service)

### DFD figure

![](static/DFD/Login-Forgot.svg)
![](static/DFD/Login-Auth.svg)
*user: email, name, password, time, ID items

### Component structure

![](static/Component/login.svg)

Please refer to the following API specifications for the details of the created API.

* [loginAPI](https://github.com/hender14/AUD_analysis/tree/main/loginapi)

## keywordAPI

### skill
* flask framework (Python)

### deploy environment
* Google Cloud Run (Google container managed service)

### DFD figure
![](static/DFD/keyword.svg)

### Component structure
![](static/Component/keyword.svg)

Please refer to the following API specifications for the details of the created API.

* [keywordAPI](https://github.com/hender14/AUD_analysis/tree/main/keywordapi)