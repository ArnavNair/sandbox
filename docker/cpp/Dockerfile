FROM iron/gcc:dev

RUN apk add --no-cache bash

ARG PATH_TO_CODE
WORKDIR /sandbox
COPY docker-entrypoint.sh .
COPY ${PATH_TO_CODE} code.cpp
RUN g++ code.cpp -o executable
ENTRYPOINT ["bash", "docker-entrypoint.sh"]
