etcdui base on e3w
===
## Quick Start

```
docker-compose up
# open http://localhost:8080
```

## Usage

1.Fetch the project `go get github.com/HughNian/etcdui`


2.frontend

```
cd static
npm install
npm run publish
```

3.backend

a. Start etcd, such as [goreman](https://github.com/coreos/etcd/#running-a-local-etcd-cluster)

b. Edit conf/config.default.ini if needed, `go build && ./e3w`

c. For auth:

```
ETCDCTL_API=3 etcdctl auth enable
# edit conf/config.default.ini[app]#auth
./e3w
# you could set your username and password in SETTING page
```

4.build image

Install dependencies in 3.b, then run `docker build -t soyking/e3w .`