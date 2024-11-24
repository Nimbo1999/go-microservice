FROM golang:1.23.2-alpine as builder

ARG ALLOWED_ORIGIN
ENV ALLOWED_ORIGIN=$ALLOWED_ORIGIN

WORKDIR /app

COPY . /app/

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

RUN chmod +x ./brokerApp

FROM scratch

WORKDIR /app

COPY --from=builder /app/brokerApp .

CMD [ "/app/brokerApp" ]
