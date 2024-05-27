Nmid-Discovery UI  
base on [e3w](https://github.com/soyking/e3w),it's etcd v3 web ui
===

## Quick Start

```
docker-compose up
# open http://localhost:8080

这个启动一个e3w现成的docker服务
```

## Overview

KEY/VALUE

![](./images/kv.png)

MEMBERS

![](./images/members.png)

ROLES

![](./images/roles.png)

USERS

![](./images/users.png)

SETTING

![](./images/setting.png)

## Usage

1.Fetch the project `go get github.com/HughNian/nmid-discovery`

2.frontend

```
cd static
npm install
npm run publish
```

3.backend

a. Start etcd, such as [goreman](https://github.com/coreos/etcd/#running-a-local-etcd-cluster)

b. Edit conf/config.default.ini if needed, `go build && ./etcdui`