#!/bin/bash
git clone git://git.cypherpunks.ru/govpn.git
cd govpn
git submodule init
git submodule update
export GOPATH=$PWD
export GOROOT=/usr/local/go1.11.2
export PATH=$GOROOT/bin:$PATH
make