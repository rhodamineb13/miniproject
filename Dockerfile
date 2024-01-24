FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o /miniproject

EXPOSE 8080

CMD [ "/miniproject" ]