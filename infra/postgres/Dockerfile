FROM alpine:latest

RUN apk add --no-cache postgresql-client bash

COPY scripts/ /scripts/

ENTRYPOINT ["bash", "/scripts/check_slave.sh"]
