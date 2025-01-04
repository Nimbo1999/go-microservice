FROM golang:1.23.2-alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o authenticationApp ./cmd/api

RUN chmod +x ./authenticationApp

FROM scratch

WORKDIR /app

COPY --from=builder /app/authenticationApp .

CMD [ "/app/authenticationApp" ]
