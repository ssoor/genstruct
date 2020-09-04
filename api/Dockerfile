FROM alpine:latest
RUN apk update && \
  apk add --no-cache ca-certificates \
   tzdata

WORKDIR /app

COPY ./api ./

ENV TZ=Asia/Shanghai
EXPOSE 80

ENTRYPOINT ["./api","--addr=:80"]