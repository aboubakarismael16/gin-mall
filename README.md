# gin-mall

Use gin+gorm+Mysql read/write to build the e-mall project.

This project is modified from the author Congz's e-mall, which removes some functions such as third-party login, extreme verification, third-party payment, and adds MySQL read/write separation, ELK log system, AES symmetric encryption for data desensitization, etc. Thank the author for open source!

This project is comprehensive and suitable for beginner to learn ``web development``.

# The main functions of the project

- User Registration Login (JWT-Go Authentication)
- User basic information modification, unbound mailbox, change password
- Products release, browse, etc.
- create, delete, update, etc. shopping carts
- Order creation, deletion, payment, etc.
- Address additions, deletions, modifications, etc.
- Number of visits to each item and ranking of some items
- Set payment password, symmetrically encrypt user's amount
- You can upload pictures to the object store or switch branches to upload to the local static directory


# Project's primary dependency

Golang v1.16+
 -    [gin](https://github.com/gin-gonic/gin)
 -    [gorm](https://gorm.io/gorm)
 -    [mysql](https://github.com/go-sql-driver/mysql)
 -    [redis](https://github.com/go-redis/redis)
 -    [ini](https://gopkg.in/ini.v1)
 -    [jwt-go-go](https://github.com/golang-jwt/jwt/v4)
 -    crypto
 -    [logrus](https://github.com/sirupsen/logrus)
 -    [dbresolver](https://github.com/go-gorm-v1/dbresolver)

# Project structure

```
gin-mall/
├── api
├── cache
├── conf
├── dao
├── doc
├── middleware
├── model
├── pkg
│  ├── e
│  └── util
├── routes
├── serializer
└── service
```

- api: Used to define interface functions
- cache: Place redis cache
- conf: Used to store configuration files
- dao: Operate on persistence layer
- doc: store interface documents
- middleware: Application Middleware
- model: Apply database model
- pkg/e: encapsulation error code
- pkg/util: Tool function
- routes: Routing logic processing
- serializer: A function that serializes data to JSON
- service: implementation of interface function

