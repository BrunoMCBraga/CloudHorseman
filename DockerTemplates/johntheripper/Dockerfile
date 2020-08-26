FROM library/centos:latest

COPY wordlist.txt /usr/share/cloudhorseman/wordlist.txt
COPY exec.sh /usr/share/cloudhorseman/exec.sh
RUN chmod +x /usr/share/cloudhorseman/exec.sh

RUN yum -y update && yum install -y openssl openssl-devel nss-devel gmp-devel krb5-devel git @development
RUN cd /usr/share/ && git clone https://github.com/magnumripper/JohnTheRipper.git && cd JohnTheRipper/src && ./configure && make clean && make -s

ENTRYPOINT ["/usr/share/cloudhorseman/exec.sh"]