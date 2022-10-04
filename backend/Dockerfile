FROM golang:alpine AS builder
WORKDIR /app
COPY . .
ENV GO111MODULE=on

RUN go build -mod=vendor -o main . 

FROM golang:alpine
WORKDIR /app
COPY --from=builder ./app/main /app/
COPY --from=builder ./app/result.xlsx /app/

ENTRYPOINT [ "./main", "api"]
EXPOSE 8080

