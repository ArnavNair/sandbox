FROM iron/gcc:dev

RUN apk add --no-cache bash

WORKDIR /sandbox
COPY docker-entrypoint.sh .
ENTRYPOINT ["bash", "docker-entrypoint.sh"]
