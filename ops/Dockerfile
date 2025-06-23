FROM golang:1.21.5-alpine

WORKDIR /app

COPY . .
#RUN go mod download
#RUN go install github.com/air-verse/air@latest
RUN go install
RUN go build

CMD AirplaneTracking start