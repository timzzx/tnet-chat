# tnet-chat （项目设计参考go-zero，tcp基于tnet）

## 测试

启动sever
```
root@tdev:/home/code/tnet-chat# make dev
go run chat.go -f etc/chat.yaml
TCP服务启动成功...
```

启动客户单
```
root@tdev:/home/code/tnet-chat/Client# go run main.go 
{"code":500,"msg":"invalid character 'e' in literal true (expecting 'r')"}
{"code":500,"msg":"密码错误"}
{"code":200,"msg":"登录成功"}
```

服务端的打印
```
root@tdev:/home/code/tnet-chat# make dev
go run chat.go -f etc/chat.yaml
TCP服务启动成功...
连接建立成功

2023/02/28 16:51:52 /home/code/tnet-chat/chatmodel/dao/query/user.gen.go:234
[0.948ms] [rows:1] SELECT * FROM `user` WHERE `user`.`name` = 'timzzx' ORDER BY `user`.`id` LIMIT 1
123456

2023/02/28 16:51:52 /home/code/tnet-chat/chatmodel/dao/query/user.gen.go:234
[0.718ms] [rows:1] SELECT * FROM `user` WHERE `user`.`name` = 'timzzx' ORDER BY `user`.`id` LIMIT 1
消息解析失败： EOF
```

