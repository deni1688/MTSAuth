FROM ubuntu:latest

WORKDIR /app

COPY ./authapi .

EXPOSE 9000
   
CMD ["./authapi"]