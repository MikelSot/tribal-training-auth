FROM ubuntu:latest
LABEL authors="mikel"

ENTRYPOINT ["top", "-b"]