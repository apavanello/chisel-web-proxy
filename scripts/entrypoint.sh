#!/bin/bash

echo $1

cp $1 ./app/backend/dataset.json

echo "Starting Backend  ..."
./app/backend/main -dataset /app/backend/dataset.json &

echo "Starting Frontend ..."
nginx -g "daemon off;"