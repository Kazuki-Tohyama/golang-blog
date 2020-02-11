FROM golang:latest

WORKDIR /go/src/github.com/Kazuki-Tohyama/go-nuxt-blog/api

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN curl -fLo /go/bin/air https://git.io/linux_air \
	&& chmod +x /go/bin/air

CMD air
