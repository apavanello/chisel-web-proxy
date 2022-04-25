#!/bin/bash

echo $1

if [[ ! -f $1 ]]
then
    echo "$1 does not exist on your filesystem. EXITING"
    exit 1
fi

cp $1 ./app/backend/dataset.json

echo "Starting Backend  ..."
./app/backend/main -dataset /app/backend/dataset.json &

echo "Starting Frontend ..."
nginx -g "daemon off;"