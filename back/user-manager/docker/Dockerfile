FROM registry.tellmeplus.com/docker/alpine:master
LABEL maintainer="Matthieu tartiere <matthieu.tartiere@mines-ales.org>"

ENV \
    BINARY=user-manager \
    CONF_PATH=/etc/master-your-life


WORKDIR /opt/master-your-life/user-manager/

COPY target/${BINARY} ${BINARY}

EXPOSE 80
CMD ["./user-manager"]
