#!/bin/bash

# allow specifying different destination directory
DIR="${DIR:-"$HOME/.local/bin"}"

# map different architecture variations to the available binaries
ARCH=$(uname -m)
case $ARCH in
    i386|i686) ARCH=386 ;;
    armv6*) ARCH=armv6 ;;
    armv7*) ARCH=armv7 ;;
    aarch64*) ARCH=arm64 ;;
    x86_64*) ARCH=amd64 ;;
esac

# prepare the download URL
GITHUB_LATEST_VERSION=$(curl -L -s -H 'Accept: application/json' https://github.com/Spacelocust/dahl/releases/latest | sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/')
GITHUB_FILE="dahl_${GITHUB_LATEST_VERSION//v/}_$(uname -s | tr '[:upper:]' '[:lower:]')_${ARCH}.tar.gz"
GITHUB_URL="https://github.com/Spacelocust/dahl/releases/download/${GITHUB_LATEST_VERSION}/${GITHUB_FILE}"

# install/update the local binary
curl -L -o dahl.tar.gz $GITHUB_URL
tar -xzvf dahl.tar.gz dahl

if [ "$(uname -s)" == "Darwin" ]; then
    mkdir -p "$DIR" && cp dahl "$DIR" && chmod 755 "$DIR/dahl"
else
    install -Dm 755 dahl -t "$DIR"
fi

rm dahl dahl.tar.gz

echo "dahl installation complete !"
