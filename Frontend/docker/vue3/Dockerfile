FROM node:16-alpine3.15
# FROM node:14.9.0-alpine3.10

CMD ["/bin/sh"]

WORKDIR /app

RUN apk update && \
    npm install -g vue-cli

ADD ./ui/package.json /app

RUN npm install