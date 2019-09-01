#!/bin/bash

export GOOS=linux; 
export GOARCH=amd64; 
echo "Building for linux...";
go build -o motusauth; 
echo "Building completed!";
echo "Shutting down running containers...";
docker-compose down; 
echo "Containers shut down!";
echo "Rebuilding and bringing containers up!";
docker-compose up --build -d    