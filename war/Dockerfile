# Get golang image
FROM golang:1.11-stretch

RUN apt-get update && \
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY . $GOPATH/src/github.com/evermos/war

# Set WORKDIR
WORKDIR $GOPATH/src/github.com/evermos/war

RUN  dep ensure -vendor-only -v && go build -o goBinary

CMD ["./goBinary"]