FROM iron/java:1.8-dev

RUN apk add --no-cache bash

ARG PATH_TO_CODE
WORKDIR /sandbox
COPY docker-entrypoint.sh .
COPY ${PATH_TO_CODE} Code.java
RUN javac Code.java
ENTRYPOINT ["bash", "docker-entrypoint.sh"]
