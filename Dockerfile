FROM ubuntu:22.04

ARG MYGHKEY

RUN apt update && apt upgrade -y && \
    apt install -y sudo vim git curl rsync

RUN groupadd -g 1000 ashish \
    && useradd -u 1000 -g 1000 -m ashish \
    && echo 'ashish   ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers \
    && mkdir -v /home/ashish/repos /home/ashish/.ssh && chmod -v 700 /home/ashish/.ssh \
    && echo $MYGHKEY | base64 --ignore-garbage --decode > /home/ashish/.ssh/id_rsa \
    && chmod -v 400 /home/ashish/.ssh/id_rsa \
    && chown -Rv ashish.ashish /home/ashish

# Things to do as ashish user
USER ashish
WORKDIR /home/ashish

RUN git config --global --add user.name "Ashish Disawal" \
    && git config --global --add user.email "ashish.disawal@gmail.com" \
    && ssh-keyscan -H github.com >> ~/.ssh/known_hosts

RUN ls -al /home/ashish/.ssh \
    && mkdir -pv ~/repos/ \
    && git clone git@github.com:shifu137/go-dev-env.git ~/repos/go-dev-env
