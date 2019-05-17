#!/usr/bin/env bash

if [ ! -f install.sh ]; then
echo 'install.sh must be run within its container folder' 1>&2
exit 1
fi

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

gofmt -w src

go install erouser

export GOPATH="$OLDGOPATH"

echo 'finished'
