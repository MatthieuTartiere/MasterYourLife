#!/usr/bin/env bash
#MAINTAINER "Matthieu TARTIERE"

function kafka__create_topic(){
    topicName=$1 ; partitionCount=$2 ; retention=$3
    dockerExecString="kafka-topics.sh --create --if-not-exists --zookeeper $KAFKA_URI:2181 --replication-factor 1 --partitions $partitionCount --config retention.ms=$(($retention*24*60*60*1000)) --topic $topicName"
    docker run -ti --rm ovhcom/queue-kafka-topics-tools $dockerExecString
}