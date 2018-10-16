# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN apk add make git gcc libc-dev
RUN cd /src && make install && make build

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/bin/shop_api /app/
COPY --from=build-env /src/db/shops.db /app/db/
EXPOSE 8080
CMD ./shop_api
