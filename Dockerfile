FROM golang:1.17-alpine
RUN apk add gcc
RUN apk add musl-dev

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o back.exe .
CMD [ "./back.exe" ]