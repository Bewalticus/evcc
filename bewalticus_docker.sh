#!/bin/bash
docker buildx build --platform=linux/amd64 -t bewalticus/evcc --push .