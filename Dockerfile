FROM ubuntu:20.04

ARG MYGHKEY

RUN apt update && apt upgrade -y && \
    apt install -y sudo vim git curl rsync

RUN groupadd -g 1000 ashish \
    && useradd -u 1000 -g 1000 -m ashish \
    && echo 'ashish   ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers \
    && mkdir /home/ashish/repos \
    && mkdir /home/ashish/.ssh && chmod 700 /home/ashish/.ssh \
    && echo $MYGHKEY | base64 --ignore-garbage --decode > /home/ashish/.ssh/id_rsa \
    && chmod 400 /home/ashish/.ssh/id_rsa \
    && chown -R ashish.ashish /home/ashish/

# Things to do as ashish user
USER ashish
WORKDIR /home/ashish

RUN git config --global --add user.name "Ashish Disawal" \
    && git config --global --add user.email "ashish.disawal@gmail.com" \
    && ssh-keyscan -H gitlab.com >> ~/.ssh/known_hosts

RUN ssh git@github.com
#    && git clone git@github.com:shifu137/go-dev-env.git ~/repos/go-dev-env
