FROM registry.access.redhat.com/ubi7-dev-preview/ubi-minimal:7.6

LABEL maintainer="LitmusChaos"

ENV SIDECAR=/usr/local/bin/litmus-sidecar

COPY build/_output/bin/litmus-sidecar ${SIDECAR}

ENTRYPOINT ["/usr/local/bin/litmus-sidecar"]
