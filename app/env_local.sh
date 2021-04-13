#!/bin/sh

export DB_READ_USER="docker:docker"
export DB_WRITE_USER="docker:docker"
export READ_CONNECTION="tcp(172.19.0.3:3306)/test?parseTime=true&loc=Asia%2FTokyo"
export WRITE_CONNECTION="tcp(172.19.0.3:3306)/test?parseTime=true&loc=Asia%2FTokyo"
export JWT_SIGNINGKEY="SAMPLE"
export GIN_MODE="debug"
