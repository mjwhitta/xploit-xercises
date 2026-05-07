#alpine# FROM alpine:edge
#alpine# RUN apk --no-cache --update add bash && \
#alpine#     rm -f -r /tmp/* /var/{cache/apk,tmp}/*
#debian# FROM debian:stable
#slim# FROM debian:stable-slim
#ubuntu# FROM ubuntu:latest
#rolling# FROM ubuntu:rolling
SHELL ["/bin/bash", "-c"]
ADD entrypoint /src /
RUN chmod -R u=rwx,go=rx /entrypoint
WORKDIR /
ENTRYPOINT ["/entrypoint"]
