rm -r /data/docker/influxdb/lib
mkdir -p /data/docker/influxdb/lib

#执行创建用户，创建数据库操作
docker run  -p 8086:8086 \
      -e INFLUXDB_DB=db0 -e INFLUXDB_HTTP_AUTH_ENABLED=true \
      -e INFLUXDB_ADMIN_USER=xxx1 -e INFLUXDB_ADMIN_PASSWORD=xxx \
      -e INFLUXDB_USER=xxx2 -e INFLUXDB_USER_PASSWORD=xxx \
      -v /data/docker/influxdb/lib:/var/lib/influxdb \
      influxdb:1.5.4 /init-influxdb.sh
#用环境变量覆盖配置文件
docker run -d -p 8086:8086 \
      -e INFLUXDB_HTTP_AUTH_ENABLED=true \
      -v /data/docker/influxdb/lib:/var/lib/influxdb \
      influxdb:1.5.4

#docker exec -it 716f93c2cfc influx