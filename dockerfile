FROM registry.lemonvd.com/coco-server/docker/go-runtime:latest

ARG ARG_PROJECT_NAME
ARG ARG_CI_BUILD_INFO

ENV PROJECT_NAME=$ARG_PROJECT_NAME
ENV CI_BUILD_INFO=$ARG_CI_BUILD_INFO

LABEL CI_BUILD_INFO=$ARG_CI_BUILD_INFO

COPY ./bin/${PROJECT_NAME} /app/bin/
COPY ./config.json /app/bin/

WORKDIR /app/bin
# ENTRYPOINT "./${PROJECT_NAME} -f ./config.json"
