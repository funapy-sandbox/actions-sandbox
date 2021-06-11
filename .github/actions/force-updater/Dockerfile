# FROM ubuntu
#
# ADD entrypoint.sh /entrypoint.sh
# RUN chmod 711 /entrypoint.sh
#
# ENTRYPOINT ["/entrypoint.sh"]

FROM alpine:3.7
RUN apk update
RUN apk add curl

ADD entrypoint.sh /entrypoint.sh
RUN chmod 711 /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
