FROM golang:1.21.4-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@v0.2.543

COPY . .
RUN templ generate ./
RUN GOOS=linux go build -o frameserver ./cmd/frameserver

FROM alpine:latest
RUN apk --no-cache add ca-certificates

ARG PORT
ARG APP_URL
ARG HUB_URL
ARG STATIC_DIR

ENV APP_URL=$APP_URL
ENV PORT=$PORT
ENV HUB_URL=$HUB_URL
ENV STATIC_DIR=$STATIC_DIR

WORKDIR /root/

COPY --from=builder /app/frameserver .
COPY --from=builder /app/static ./static

CMD [ "/root/frameserver" ]
