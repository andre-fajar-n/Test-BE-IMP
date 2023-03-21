FROM golang:1.20.2-alpine as builder

WORKDIR /app

RUN apk add --no-cache make

RUN wget "https://github.com/go-swagger/go-swagger/releases/download/v0.30.3/swagger_linux_amd64" && \
    mv swagger_linux_amd64 /usr/local/bin/swagger && \
    chmod a+x /usr/local/bin/swagger

COPY . .

RUN make all

FROM alpine:3.17

WORKDIR /app

RUN apk update && apk add tzdata

# set timezone
ENV TZ="Asia/Jakarta"

COPY --from=builder ./app/build/main .

COPY --from=builder ./app/config.yaml .

EXPOSE 3000

CMD [ "./main", "start" ]
