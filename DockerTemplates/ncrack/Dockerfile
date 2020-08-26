FROM debian:latest

ENV NCRACK_VERSION 0.7

COPY usernames.txt /usr/share/cloudhorseman/usernames.txt
COPY passwords.txt /usr/share/cloudhorseman/passwords.txt
COPY exec.sh /usr/share/cloudhorseman/exec.sh
RUN chmod +x /usr/share/cloudhorseman/exec.sh

RUN apt-get update && apt-get install -y wget build-essential libssl-dev zlib1g libz-dev

RUN cd /usr/share/ && wget https://nmap.org/ncrack/dist/ncrack-${NCRACK_VERSION}.tar.gz && tar -xzf ncrack-${NCRACK_VERSION}.tar.gz && cd ncrack-${NCRACK_VERSION} && ./configure && make && make install


ENTRYPOINT ["/usr/share/cloudhorseman/exec.sh"]