FROM golang:alpine
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
ENV PORT=3001
RUN apk update && apk add git && go get github.com/labstack/echo/v4
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
EXPOSE 3001
ENTRYPOINT ["/app/heroku-test"]
