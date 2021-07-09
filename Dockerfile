FROM golang:latest
LABEL maintainer="Karan Khunti"
WORKDIR /grpcapp
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["cmd/server/"]