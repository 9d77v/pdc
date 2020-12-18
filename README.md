# pdc
Personal Data Center.

# Quickstart
1. set host

add `your_ip domain.local` to host file

win10:C:\Windows\System32\drivers\etc\hosts

linux/mac: /etc/hosts

2. run app
```
git clone git@github.com:9d77v/pdc.git
cd pdc

# arm64 machine
docker-compose up -d
```

3. install crt
```
docker cp caddy:/data/caddy/pki/authorities/local/root.crt root.crt

docker cp caddy:/data/caddy/pki/authorities/local/intermediate.crt intermediate.crt

docker cp caddy:/data/caddy/certificates/local/domain.local/domain.local.crt domain.local.crt

docker cp caddy:/data/caddy/certificates/local/oss.domain.local/oss.domain.local.crt oss.domain.local.crt
```
install crts to your system

4. visit homepage
homepage https://domain.local

docs https://domain.local/docs

# Develop 
## run server
```
make gen
make dev
```
## run web
```
cd web
yarn start
```
## reference docs
gqlgen https://gqlgen.com/

gorm https://gorm.io/zh_CN/docs/index.html

apollographql https://www.apollographql.com/docs/react/

ant.design https://ant.design/docs/react/introduce-cn

minio https://docs.min.io/

