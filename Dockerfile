FROM golang:1.23-alpine

WORKDIR /app

COPY . .
#RUN go mod download
#RUN go install github.com/air-verse/air@latest
RUN go install
RUN go build

CMD TrainTracking start