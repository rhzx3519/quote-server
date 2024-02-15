FROM ubuntu:latest

WORKDIR /app

COPY quote-server ./
COPY .env ./

EXPOSE 80

ENTRYPOINT ["./quote-server"]
