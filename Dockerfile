FROM ubuntu:latest
LABEL authors="typetrait"

ENTRYPOINT ["top", "-b"]