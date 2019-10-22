#!/usr/bin/env bash

set -e

exec_cmd () {
    echo Exec : $@
    bash -c "$@"
}

PROJECT_PATH=$(cd $(dirname $(dirname $0)) && pwd)

# Copy the default config
exec_cmd "cp ${PROJECT_PATH}/config.yml.dist ${PROJECT_PATH}/config.yml"