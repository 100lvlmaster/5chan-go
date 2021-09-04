# Build Stage
FROM golang:alpine as build

WORKDIR /src/5chan-go

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /app/5chan-go

# Final Stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/5chan-go /app/

EXPOSE 8080

CMD ./5chan-go
