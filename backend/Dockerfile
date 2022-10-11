FROM golang:1.19-alpine AS builder

WORKDIR /backend

RUN apk update && apk add --no-cache git curl 

COPY go.mod ./
COPY go.sum ./

COPY . .

RUN go mod tidy
RUN go build -o main main.go
EXPOSE 3000

FROM frolvlad/alpine-glibc
COPY --from=builder /backend .


CMD [ "./main" ]