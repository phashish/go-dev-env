FROM ubuntu:20.04

ARG MYGHKEY

RUN apt update && apt upgrade -y
RUN apt install -y sudo vim git curl rsync

RUN groupadd -g 1000 ashish
RUN useradd -u 1000 -g 1000 -m ashish
RUN echo 'ashish   ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

RUN mkdir /home/ashish/repos
RUN mkdir /home/ashish/.ssh && chmod 700 /home/ashish/.ssh
RUN echo $MYGHKEY | base64 --decode > /home/ashish/.ssh/id_rsa && chmod 400 /home/ashish/.ssh/id_rsa

RUN chown -R ashish.ashish /home/ashish/

# Things to do as ashish user
USER ashish
WORKDIR /home/ashish

RUN git config --global --add user.name "Ashish Disawal"
RUN git config --global --add user.email "ashish.disawal@gmail.com"
RUN ssh-keyscan -H gitlab.com >> ~/.ssh/known_hosts
RUN git clone git@github.com:shifu137/go-dev-env.git ~/repos/go-dev-env
