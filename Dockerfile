FROM zeqi/micro:k8s

WORKDIR /go/src/vc.cli
RUN rm -rf ./*
COPY . .
ENTRYPOINT /go/src/vc.cli/vc-cli
# CMD [ "sh", "/go/src/vc.cli/entrypoint.sh" ]