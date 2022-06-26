#!/bin/sh
docker build -t m2tx/neowaytc-frontend:0.0.2 ./frontend/
docker push m2tx/neowaytc-frontend:0.0.2
docker build -t m2tx/neowaytc-backend:0.0.2 ./backend/
docker push m2tx/neowaytc-backend:0.0.2
docker build -t m2tx/neowaytc-backendgo:0.0.1 ./backendgo/
docker push m2tx/neowaytc-backendgo:0.0.1
