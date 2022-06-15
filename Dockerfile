FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go get .

COPY *.go ./

RUN go build -o back.exe .

CMD [ "./back.exe" ]