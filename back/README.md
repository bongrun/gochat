go get github.com/astaxie/beego
go get github.com/beego/bee
go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson


go get github.com/Masterminds/glide
go get github.com/beego/bee
glide install

bee run


2017/09/11 00:31:58 SUCCESS  ▶ 0005 Built Successfully!
2017/09/11 00:31:58 INFO     ▶ 0006 Restarting 'backend'...
2017/09/11 00:31:58 SUCCESS  ▶ 0007 './backend' is running...
2017/09/11 00:31:58 [I] [asm_amd64.s:2337] http server Running on http://:9012
[beego] 2017/09/11 - 00:32:14 |      127.0.0.1| 200 |     2.0027ms|   match| GET      /     r:/
[beego] 2017/09/11 - 00:32:14 |      127.0.0.1| 200 |    11.5288ms|   match| GET      /static/js/reload.min.js
2017/09/11 00:51:28 ERROR    ▶ 0008 Failed to open file on 'D:/go/src/gochat/backend/controllers/default.go': open D:/go/src/gochat/backend/controllers/default.go: The system cannot find the file specified.