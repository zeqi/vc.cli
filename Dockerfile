FROM zeqi/micro:0.1.0

WORKDIR /go/src/vc.cli
RUN rm -rf ./*
COPY . .
RUN make build-linux-server
ENTRYPOINT /go/src/vc.cli/vc-cli
# CMD [ "sh", "/go/src/vc.cli/entrypoint.sh" ]