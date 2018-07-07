FROM alpine

RUN apk add -U bash curl git sudo

ADD https://raw.githubusercontent.com/fishworks/gofish/master/scripts/install.sh install.sh
RUN bash install.sh

ENV USER root
ENV HOME /root
