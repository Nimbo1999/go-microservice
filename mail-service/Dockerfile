FROM golang:1.23.2-alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o mailApp ./cmd/api

RUN chmod +x ./mailApp

FROM scratch

WORKDIR /app

COPY --from=builder /app/mailApp .
COPY templates templates/

CMD [ "/app/mailApp" ]
