FROM golang:1.21.4-alpine as builder

WORKDIR /app

COPY . /app/

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

RUN chmod +x ./brokerApp

FROM scratch

WORKDIR /app

COPY --from=builder /app/brokerApp .

CMD [ "/app/brokerApp" ]
