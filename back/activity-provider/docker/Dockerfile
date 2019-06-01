FROM registry.tellmeplus.com/docker/alpine:master
LABEL maintainer="Matthieu tartiere <matthieu.tartiere@mines-ales.org>"

ENV \
    BINARY=activity-provider \
    CONF_PATH=/etc/master-your-life


WORKDIR /opt/master-your-life/activity-provider/

COPY target/${BINARY} ${BINARY}

EXPOSE 80
CMD ["./activity-provider"]
