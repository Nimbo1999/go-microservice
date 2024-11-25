FROM golang:1.23.2-alpine as builder

ARG ALLOWED_ORIGIN
ENV ALLOWED_ORIGIN=$ALLOWED_ORIGIN
ARG AUTHENTICATION_SERVICE_BASE_URL
ENV AUTHENTICATION_SERVICE_BASE_URL=$AUTHENTICATION_SERVICE_BASE_URL

WORKDIR /app

COPY . /app/

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

RUN chmod +x ./brokerApp

FROM scratch

WORKDIR /app

COPY --from=builder /app/brokerApp .

CMD [ "/app/brokerApp" ]
