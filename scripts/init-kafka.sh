#!/usr/bin/env bash
#MAINTAINER "Matthieu TARTIERE"

#### LOAD LIBS ####

    source ./lib/utils.sh
    source ./lib/kafka.sh

    loadEnv

# init topics

    echo  "Initialize kafka:"
    # Params topicname/partitionCount/Retention
    kafka__create_topic "master-your-life-ingest"        "20" "1"
