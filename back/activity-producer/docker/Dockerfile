FROM registry.tellmeplus.com/docker/alpine:master
LABEL maintainer="Matthieu tartiere <matthieu.tartiere@mines-ales.org>"

ENV \
    BINARY=activity-producer \
    CONF_PATH=/etc/master-your-life


WORKDIR /opt/master-your-life/activity-producer/

COPY target/${BINARY} ${BINARY}

EXPOSE 80
CMD ["./activity-producer"]
