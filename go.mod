module github.com/tiptok/examples_gincomm

go 1.12

require (
	github.com/astaxie/beego v1.10.0
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/tiptok/gocomm v0.0.0-20190902031746-1511efde9b13
)

replace github.com/tiptok/gocomm => ../gocomm
