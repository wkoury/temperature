FROM golang:1.24-alpine as builder

RUN apk add --no-cache make

WORKDIR /app

# COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download

COPY . .

RUN make build

# EXPOSE 8080

FROM scratch as runner

COPY --from=builder /app/temp /usr/bin/temp

CMD ["temp"]
