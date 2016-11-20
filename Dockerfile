FROM alpine:3.4

ADD https://storage.googleapis.com/kubernetes-release/release/v1.4.6/kubernetes-client-linux-amd64.tar.gz /tmp
RUN tar zxf /tmp/kubernetes-client-*.tar.gz && \
	mv /kubernetes/client/bin/kubectl /usr/local/bin/

ADD watcher /watcher
ENTRYPOINT ["/watcher"]
