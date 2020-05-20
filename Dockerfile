FROM scratch
MAINTAINER mateuszmierzwinski@gmail.com

COPY domain.crt /
COPY domain.key /
COPY http2push-linux-amd64 /
COPY static /static/

EXPOSE 8443
ENTRYPOINT ["/http2push-linux-amd64"]