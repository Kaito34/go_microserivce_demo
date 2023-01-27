# Ubuntuの公式コンテナを軸に環境構築
# 22.04ではaptからpython3.8が入っていなかったので20.04で固定する
FROM ubuntu:20.04

ENV ARCH amd64
ENV GOVERSION 1.15.6
# ENV HOME /root
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH /go_microservice_demo/src/work

# インタラクティブモードにならないようにする
ARG DEBIAN_FRONTEND=noninteractive
# タイムゾーンを日本に設定
ENV TZ=Asia/Tokyo

# インフラを整備
RUN apt-get update && \
    apt-get install -y time tzdata tree git curl expect zip

# C++, Python3, PyPy3, Goの4つの環境想定
RUN apt-get update \
  && apt-get install -y vim gcc-9 g++-9 gdb
RUN apt-get install -y python3.8 python3-pip pypy3
RUN apt-get install -y nodejs npm 
RUN pip3 install -U pip
RUN pip3 install online-judge-tools

RUN curl -s -o /tmp/go.tar.gz https://storage.googleapis.com/golang/go$GOVERSION.linux-$ARCH.tar.gz
RUN tar -C /usr/local -xzf /tmp/go.tar.gz
RUN rm /tmp/go.tar.gz 



# 一般的なコマンドで使えるように設定
# e.g. python3.8 main.py => python main.py
RUN update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-9 30 && \
    update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-9 30 && \
    update-alternatives --install /usr/bin/python python /usr/bin/python3.8 30 && \
    update-alternatives --install /usr/bin/pip pip /usr/bin/pip3 30 && \
    update-alternatives --install /usr/bin/pypy pypy /usr/bin/pypy3 30 && \
    update-alternatives --install /usr/bin/node node /usr/bin/nodejs 30

# # AHC用のRustのinstall
RUN curl https://sh.rustup.rs -sSf | sh -s -- -y
ENV PATH $PATH:/home/root/.cargo/bin

# Go
ENV GO111MODULE off

# install go tools
RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs
RUN go get github.com/ramya-rao-a/go-outline
RUN go get github.com/nsf/gocode
RUN go get github.com/acroca/go-symbols
RUN go get github.com/fatih/gomodifytags
RUN go get github.com/josharian/impl
RUN go get github.com/haya14busa/goplay/cmd/goplay
RUN go get github.com/go-delve/delve/cmd/dlv
# RUN go get golang.org/x/lint/golint
# RUN go get golang.org/x/tools/gopls
