# For solc binary
FROM ethereum/solc:0.5.12 as solc-builder
# We use a multistage build to avoid bloating our deployment image with build dependencies
FROM golang:1.15-alpine3.12 as builder

RUN apk add --no-cache --update git bash make musl-dev gcc libc6-compat

ARG REPO=/src/hsc
COPY . $REPO
WORKDIR $REPO

# Build purely static binaries
RUN make build

# This will be our base container image
FROM alpine:3.11

# Variable arguments to populate labels
ARG USER=hivesmartchain
ARG INSTALL_BASE=/usr/local/bin

# Fixed labels according to container label-schema
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.name = "Hive Smart Chain"
LABEL org.label-schema.vendor="Hyperledger HiveSmartChain Authors"
LABEL org.label-schema.description="Hive Smart Chain is a permissioned Ethereum smart-contract side chain node for HIVE."
LABEL org.label-schema.license="Apache-2.0"
LABEL org.label-schema.vcs-url="https://github.com/klyed/hivesmartchain"

# Run hsc as hsc user; not as root user
ENV HSC_PATH /home/$USER
RUN addgroup -g 101 -S $USER && adduser -S -D -u 1000 $USER $USER
WORKDIR $HSC_PATH
ADD --chown=$USER /testnet/ $HSC_PATH

# Copy binaries built in previous stage
COPY --from=builder /src/hsc/bin/hsc $INSTALL_BASE/
COPY --from=builder /src/hsc/bin/hsc-debug $INSTALL_BASE/
COPY --from=solc-builder /usr/bin/solc $INSTALL_BASE/

# Expose ports for 26656:peer; 26658:info; 10997:grpc
EXPOSE 26656
EXPOSE 26658
EXPOSE 10997

USER $USER:$USER
ENTRYPOINT [ "hsc" ]
CMD [ "start" ]
