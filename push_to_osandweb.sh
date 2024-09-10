#!/bin/bash
env CC=arm-linux-musleabihf-gcc GOOS=linux GOARCH=arm GOARM=7  go build .
mv backend docker/backend
mkdir docker/assets
cp -r assets docker/assets
cd docker
docker buildx build --platform=linux/arm/v7 . -t 192.168.0.210:5000/tobias_backend:latest --load
docker push 192.168.0.210:5000/tobias_backend:latest
sudo rm -rf assets
cd ..