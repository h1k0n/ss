docker run -d --name some-mongo \
    -e MONGO_INITDB_ROOT_USERNAME='xsec' \
    -e MONGO_INITDB_ROOT_PASSWORD='xatxsecio' \
	-e MONGO_INITDB_DATABASE="passive_scan" \
	-v /data/mongodata:/data/db \
	-p 27017:27017 \
    mongo:4.0.9-xenial
