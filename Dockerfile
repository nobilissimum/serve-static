FROM docker.io/golang:1.23.3-alpine AS builder

LABEL org.opencontainers.image.authors="ronn.angelo.lee@gmail.com"

ENV LANG=C.UTF-8

WORKDIR /app
COPY . ./

RUN go build -o app main.go


FROM docker.io/golang:1.23.3-alpine

LABEL org.opencontainers.image.authors="ronn.angelo.lee@gmail.com"

ENV LANG=C.UTF-8

WORKDIR /app
COPY --from=builder /app/app /app/app

CMD ["./app"]
