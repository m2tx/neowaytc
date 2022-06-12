#!/bin/sh
docker build -t m2tx/neowaytc-frontend:0.0.1 ./frontend/
docker push m2tx/neowaytc-frontend:0.0.1
docker build -t m2tx/neowaytc-backend:0.0.1 ./backend/
docker push m2tx/neowaytc-backend:0.0.1