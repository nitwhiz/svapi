#!/bin/bash

curl -o /tmp/svapi-data.zip "$1"
unzip /tmp/svapi-data.zip ../internal/data/embedded/
rm /tmp/svapi-data.zip
