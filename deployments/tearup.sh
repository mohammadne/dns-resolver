#!/usr/bin/env bash

set -e

cmd=$1

if [ -z "$cmd" ]; then
    echo "Error: cmd should be given as an argument"
    exit 9999 # die with error code 9999
fi

current_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
path_to_helm="$current_dir/$cmd"
path_to_secrets="$current_dir/secrets.yaml"

if [ ! -d "$path_to_helm" ]; then
    echo "Error: command (directory) $1 does not exists."
    exit 9999 # die with error code 9999
fi

echo "tearup dns-resolver-$cmd ..."

helm secrets upgrade --install "dns-resolver-$cmd" "$path_to_helm" -f "$path_to_secrets"
