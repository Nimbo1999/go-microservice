FROM golang:1.23.2-alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o listenerApp .

RUN chmod +x ./listenerApp

FROM scratch

WORKDIR /app

COPY --from=builder /app/listenerApp .

CMD [ "/app/listenerApp" ]
