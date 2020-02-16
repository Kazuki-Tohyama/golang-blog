FROM node:12.16.0-alpine3.10

WORKDIR /go/src/github.com/Kazuki-Tohyama/go-nuxt-blog/frontend

RUN apk update && \
    apk add git && \
    npm install -g npm && \
    npm install -g vue-cli

ENV HOST 0.0.0.0
EXPOSE 3000
