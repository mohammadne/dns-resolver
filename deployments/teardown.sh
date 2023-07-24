#!/usr/bin/env bash

set -e

cmd=$1

if [ -z "$cmd" ]; then
    echo "Error: cmd should be given as an argument"
    exit 9999 # die with error code 9999
fi

echo "teardown dns-resolver-$cmd ..."

helm secrets uninstall "dns-resolver-$cmd"
