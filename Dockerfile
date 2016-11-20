FROM alpine:3.4

RUN apk add --update --no-cache curl

RUN curl -fsSL https://storage.googleapis.com/kubernetes-release/release/v1.4.6/kubernetes-client-linux-amd64.tar.gz | \
	tar zxf - && mv /kubernetes/client/bin/kubectl /usr/local/bin/

ADD watcher /watcher
ENTRYPOINT ["/watcher"]
