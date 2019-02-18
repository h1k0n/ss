#!/bin/bash
docker run --name MysqlContainer -e MYSQL_ROOT_PASSWORD=xxx -p 3306:3306 -v /data/docker/mysql_5.7.22/lib:/var/lib/mysql -d mysql:5.7.22 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci --skip-character-set-client-handshake


docker run -it --link MysqlContainer:mysql --rm mysql:5.7.22 sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p'

create database xops;

create user 'paas'@'%' identified by '123456';

GRANT ALL PRIVILEGES ON xops.* TO 'paas'@'%' WITH GRANT OPTION;

CREATE TABLE `userinfo` ( `uid` INTEGER NOT NULL AUTO_INCREMENT, `username` VARCHAR(64) NULL DEFAULT NULL, `departname` VARCHAR(64) NULL DEFAULT NULL, `created` DATE NULL DEFAULT NULL, PRIMARY KEY (`uid`) );

docker run -it --link MysqlContainer:mysql --rm mysql:5.7.22 sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -upaas -p'

docker run --name MysqlContainer8015 -e MYSQL_ROOT_PASSWORD=A1234xxx -v /data/docker/mysql_8.0.15/lib:/var/lib/mysql -d mysql:8.0.15 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci --skip-character-set-client-handshake
docker run -it --link MysqlContainer8015:mysql --rm mysql:8.0.15 sh -c 'exec mysql -h "$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p'
