#This is golang project for learning

**Prerequisite**
- [golang](https://go.dev/doc/install) this project is using version 1.21
- [air ( live reload )](https://github.com/cosmtrek/air#prefer-installsh)

**Config in ~/.zshrc**
```
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$(go env GOPATH)/bin
alias air="$(go env GOPATH)/bin/air"
```

**Library**
- REST API: [gofiber v2](https://docs.gofiber.io/)
- Websocket: [gofiber/contrib](https://github.com/gofiber/contrib/blob/main/websocket/README.md)
- Entity Framework: [entgo](https://entgo.io/docs/getting-started)
- Database Connection (MySQL): [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- JSON Converter: [go-json](https://github.com/goccy/go-json)

**Project Start**
```
$ cd scripts
$ ./regenerate-schema.sh
```

**Start Service ( live reload )**
```
$ air run dev
```

---
**Remark**
Initial for air (live reload)
```
$ air init
```