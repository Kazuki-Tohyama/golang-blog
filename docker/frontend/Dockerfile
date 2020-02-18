FROM node:12.16.0-alpine3.10
ENV TZ Asia/Tokyo

WORKDIR /go/src/github.com/Kazuki-Tohyama/go-nuxt-blog/frontend

RUN apk update && \
    apk add git && \
    npm install -g npm && \
    yarn global add @vue/cli

ENV HOST 0.0.0.0
EXPOSE 3000
