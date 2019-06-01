FROM registry.tellmeplus.com/docker/alpine:master
LABEL maintainer="Matthieu tartiere <matthieu.tartiere@mines-ales.org>"

ENV \
    BINARY=goals-manager \
    CONF_PATH=/etc/master-your-life


WORKDIR /opt/master-your-life/goals-manager/

COPY target/${BINARY} ${BINARY}

EXPOSE 80
CMD ["./goals-manager"]
