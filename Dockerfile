FROM zeqi/micro:k8s

WORKDIR /go/src/vc.cli
RUN rm -rf ./*
COPY . .
CMD [ "sh", "/go/src/vc.cli/entrypoint.sh" ]
# ENTRYPOINT /go/src/app/vc
# ENTRYPOINT [ "/go/src/app/vc" ]
# ENTRYPOINT ['entrypoint.sh']