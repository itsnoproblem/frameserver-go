FROM golang:1.21.4-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@v0.2.543

COPY . .
RUN templ generate ./
RUN GOOS=linux go build -o fc-frame ./cmd/fc-frame

FROM alpine:latest
RUN apk --no-cache add ca-certificates

ARG PORT
ARG APP_URL

ENV APP_URL=$APP_URL
ENV LISTEN_PORT=$PORT

WORKDIR /root/

COPY --from=builder /app/fc-frame .
COPY --from=builder /app/www ./www

CMD [ "./fc-frame" ]
