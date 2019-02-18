#!/bin/bash

docker run --name redis-two  -p 6379:6379 -d redis:4.0.10

docker run -p 6379:6379 -v /data/docker/redis/conf/redis.conf:/usr/local/etc/redis/redis.conf --name redis-two redis:4.0.10 redis-server /usr/local/etc/redis/redis.conf

docker run -it --link redis-two:redis --rm redis:4.0.10 redis-cli -h redis -p 6379
