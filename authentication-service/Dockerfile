FROM golang:1.21.4-alpine as builder

WORKDIR /app

COPY . /app/

RUN CGO_ENABLED=0 go build -o authApp ./cmd/api

RUN chmod +x ./authApp

FROM scratch

WORKDIR /app

COPY --from=builder /app/authApp .

CMD [ "/app/authApp" ]
