FROM golang:1.23.2-alpine as builder

ARG ALLOWED_ORIGIN
ENV ALLOWED_ORIGIN=$ALLOWED_ORIGIN

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o authenticationApp ./cmd/api

RUN chmod +x ./authenticationApp

FROM scratch

WORKDIR /app

COPY --from=builder /app/authenticationApp .

CMD [ "/app/authenticationApp" ]
