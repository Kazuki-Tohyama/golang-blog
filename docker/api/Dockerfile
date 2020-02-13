FROM golang:1.12-alpine
RUN apk add --no-cache git
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
RUN apk add curl
RUN apk add --no-cache mysql-client

WORKDIR /go/src/github.com/Kazuki-Tohyama/go-nuxt-blog/api
ENV GO111MODULE=on

COPY ./api/wait.sh . 
COPY ./api/go.mod .
COPY ./api/go.sum .

RUN go mod download
COPY . .
RUN curl -fLo /go/bin/air https://git.io/linux_air \
	&& chmod +x /go/bin/air

EXPOSE 8080

CMD ["sh", "wait.sh", "mysql", "root", "root", "air"]
