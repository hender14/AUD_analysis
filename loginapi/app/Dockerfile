FROM golang:1.18-alpine

# RUN go get github.com/cespare/reflex
# ENV CGO_ENABLED=0

RUN mkdir /go/src/app
WORKDIR /go/src/app
COPY . ./

# pkg更新の際は-uｵﾌﾟｼｮﾝを追加
RUN go get github.com/gin-gonic/gin@v1.8.1 \
&& go get github.com/dgrijalva/jwt-go@v3.2.0+incompatible \
&& go get github.com/go-playground/validator/v10@v10.9.0 \
&& go get github.com/leodido/go-urn@v1.2.1 \
&& go get github.com/go-playground/universal-translator@v0.18.0 \
&& go get github.com/go-playground/locales@v0.14.0 \
&& go get github.com/json-iterator/go@v1.1.12 \
&& go get github.com/mattn/go-isatty@v0.0.14 \
&& go get github.com/modern-go/concurrent@v0.0.0-20180306012644-bacd9c7ef1dd \
&& go get golang.org/x/crypto@v0.0.0-20211215153901-e495a2d5b3d3 \
&& go get gopkg.in/yaml.v2@v2.4.0 \
&& go get	github.com/sendgrid/sendgrid-go \
&& go get	github.com/sendgrid/rest \
&& go get	github.com/guregu/dynamo@v1.16.0 \
&& go get	github.com/aws/aws-sdk-go@v1.44.86 \
&& go get github.com/gin-gonic/gin/binding@v1.7.7 \
&& go get github.com/golang/mock@v1.6.0

# RUN apk add build-base
RUN apk add --no-cache gcc
RUN apk add --no-cache musl-dev
RUN go get github.com/go-playground/assert/v2

EXPOSE 8080

ENTRYPOINT ["go", "run", "main.go"]