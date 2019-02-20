FROM iron/python

RUN apk add --no-cache bash

ARG PATH_TO_CODE
WORKDIR /sandbox
COPY docker-entrypoint.sh .
COPY ${PATH_TO_CODE} code.py
ENTRYPOINT ["bash", "docker-entrypoint.sh"]
