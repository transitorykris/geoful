all:
	docker build -t transitorykris/geoful .

push:
	docker push transitorykris/geoful

run:
	docker run -it -p 8080:8080 transitorykris/geoful

data:
	mkdir -p data
	curl http://geolite.maxmind.com/download/geoip/database/GeoLiteCountry/GeoIP.dat.gz | gzip -d > data/GeoIP.dat
	curl http://geolite.maxmind.com/download/geoip/database/GeoLiteCity.dat.gz | gzip -d > data/GeoLiteCity.dat
	curl http://download.maxmind.com/download/geoip/database/asnum/GeoIPASNum.dat.gz | gzip -d > data/GeoIPASNum.dat

clean:
	docker rmi transitorykris/geoful
	rm -rf data

up:
	docker stack deploy geoful --compose-file docker-compose.yml

down:
	docker stack rm geoful
