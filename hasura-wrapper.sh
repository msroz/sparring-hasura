#/bin/sh

# docker build -t hasuracli:latest hasura-cli/
docker run --net="host" --rm -v "$(PWD)/hasura:/tmp" hasuracli:latest $@
