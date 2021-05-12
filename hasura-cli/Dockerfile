FROM ubuntu:latest

RUN apt-get update -y && apt-get install -y curl

# Create Hasura CLI user
# RUN adduser --system --home /hasuracli --disabled-password --group hasuracli
WORKDIR /hasuracli
# USER hasuracli

# Install HasuraCLI
RUN curl -L https://github.com/hasura/graphql-engine/raw/master/cli/get.sh | INSTALL_PATH=/hasuracli bash

ENV PATH="/hasuracli"

ADD run.sh /hasuracli/run.sh

# Run Hasura on start
ENTRYPOINT ["/hasuracli/run.sh"]
CMD ["version"]
