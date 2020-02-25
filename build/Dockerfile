FROM registry.access.redhat.com/ubi7-dev-preview/ubi-minimal:7.6

LABEL maintainer="LitmusChaos"

ENV RUNNER=/usr/local/bin/litmus-sidecar

COPY build/_output/bin/litmus-sidecar ${RUNNER}

ENTRYPOINT ["/usr/local/bin/litmus-sidecar"]
