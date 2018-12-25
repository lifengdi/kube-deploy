FROM docker.io/mbrt/golang-vim-dev

USER root

WORKDIR /go/src/kube-deploy

COPY . /go/src/kube-deploy

ENV TERM=xterm

EXPOSE 8080

ENTRYPOINT ["sh", "/go/src/kube-deploy/docker/entrypoint.sh"]