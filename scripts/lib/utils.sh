#!/usr/bin/env bash
#MAINTAINER "Matthieu TARTIERE"

function getCurrentDir(){
    echo "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
}

source "`getCurrentDir`/lib/kafka.sh"
###################################################################
### Load ENV
#####

function loadEnv {
    if [[ "$1" != "" ]]; then
        env_file="$1"
    else
        env_file=".env"
    fi


    NOT_FOUND=0

    if [ -f $env_file ]; then
        source ../$env_file
    else
        NOT_FOUND=1
    fi

}

