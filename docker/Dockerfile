FROM docker.io/mbrt/golang-vim-dev

USER root

WORKDIR /go/src/b2c-deploy

COPY . /go/src/b2c-deploy

ENV TERM=xterm

EXPOSE 8080

ENTRYPOINT ["sh", "/go/src/b2c-deploy/docker/entrypoint.sh"]