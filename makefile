# 命令
help:
	@echo 'Usage:'
	@echo '     db 生成sql执行代码'
	@echo '     api 根据api文件生成go-zero api代码'
	@echo '     dev 运行'
db:
	gentool -dsn 'root:123456@tcp(192.168.1.13:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local' -outPath './chatmodel/dao/query'
dev:
	go run chat.go -f etc/chat.yaml