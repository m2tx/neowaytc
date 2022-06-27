#!/bin/sh
docker build -t m2tx/neowaytc-frontend:0.0.5 ./frontend/
docker push m2tx/neowaytc-frontend:0.0.5
docker build -t m2tx/neowaytc-backendgo:0.0.3 ./backendgo/
docker push m2tx/neowaytc-backendgo:0.0.3
