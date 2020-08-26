FROM golang:1.14

COPY exec.sh /usr/share/cloudhorseman/exec.sh
RUN chmod +x /usr/share/cloudhorseman/exec.sh && cd /usr/share/ && git clone https://github.com/zayotic/udpflood
RUN cd /usr/share/udpflood && export GOPATH=`pwd` && export GOBIN=$GOPATH/bin && go get ./... && go build

ENTRYPOINT ["/usr/share/cloudhorseman/exec.sh"]