FROM golang:latest
WORKDIR /gamecomp-backend
COPY . .
RUN go mod download
RUN go build -o main ./main.go
EXPOSE 8888
CMD ["./main"]