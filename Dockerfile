FROM alpine:latest

WORKDIR /app

COPY ./motusauth .

EXPOSE 9000
   
CMD ["./motusauth"]