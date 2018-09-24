FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 8080

ADD shop_api /bin/shop_api

CMD ["shop_api"]
