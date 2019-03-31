# Gin Web

[文档 | Gin Web Framework](https://gin-gonic.com/zh-cn/docs/#asciijson)

## install gin library

```bash
go get -u github.com/gin-gonic/gin

# postgresql
# go get github.com/astaxie/beego/orm

go get -u github.com/jinzhu/gorm
go get -u github.com/lib/pq
```

## usage

### router path parameters

```go
// :parameter -> could not be null
// *parameter -> could be null

parameter := c.Param("parameter")
```

### url query parameter

```go
// query
firstname := c.DefaultQuery("firstname", "Guest")
lastname := c.Query("lastname")

ids := c.QueryMap("ids")

// form
message := c.PostForm("message")

names := c.PostFormMap("names")
```

## orm

[GORM Guides | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](http://gorm.io/docs/)

## config

### env file

```bash
go get -u github.com/joho/godotenv

# for production
# should pre-prepare .env file and then map .env file
```

### database

```go
db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=mydb password=postgres sslmode=disable")
defer db.Close()
```