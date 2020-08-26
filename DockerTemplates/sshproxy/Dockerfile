FROM ubuntu:latest

COPY exec.sh /usr/share/cloudhorseman/exec.sh
RUN chmod +x /usr/share/cloudhorseman/exec.sh

RUN apt-get update && apt-get install -y openssh-server
RUN mkdir /var/run/sshd
RUN useradd -ms /bin/bash cloudhorseman
RUN echo 'cloudhorseman:94ytg249fge9uyG3(#G%#(rgeuiofoueb' | chpasswd

# SSH login fix. Otherwise user is kicked off after login
RUN sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd

ENV NOTVISIBLE "in users profile"
RUN echo "export VISIBLE=now" >> /etc/profile

EXPOSE 22
ENTRYPOINT ["/usr/share/cloudhorseman/exec.sh"]

#Run kubectl get services to get the External IP (actually a domain) and then use that as @domain below
#Locally: ssh -D 12345 -C -q -N cloudhorseman@acfc42d54981644a1bfd5d7496ce705c-1453548692.us-east-1.elb.amazonaws.com