FROM alpine:latest

RUN mkdir /app

COPY /brokerApp /app

WORKDIR /app

CMD [ "/brokerApp" ]