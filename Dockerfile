# ROCKETSHIP

### Stage 1
FROM golang:latest as builder
WORKDIR /gotest
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

### Stage 2
FROM scratch  
WORKDIR /root/
COPY --from=builder /gotest/main .
EXPOSE 8080
CMD ["./main"] 