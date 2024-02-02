# gin + gorm usage
## 请求流程
request -> router -> controller -> service -> dao

接收到请求后通过路由绑定将请求发送到 controller
- controller 进行参数验证,未通过验证返回 badRequest, 通过则调用对应的service
- service 接受controller 发送的数据进行业务逻辑处理, 若需要操作/查询数据库 则调用 dao层, 完成后返回数据给controller
- dao 操作和查询数据库中的数据

路由中间件:
进行token 校验, 常用的参数绑定

# 生成 model
编辑 `gen.yml`文件中的 dsn 以设置数据库连接

```shell
go install gorm.io/gen/tools/gentool@latest

gentool -c ./gen.yml
```

生成的文件位于 `./dao/model` `./dao/query`, 这两个文件夹中的内容不能手动编辑


# 小坑
gin.Engine.Set(k, v) 后获取有些问题, c.GetTxx(k) 可能获取不到值, 使用 `v, exists := c.Get(k)` 获取, 之后使用 `v.(T)` 转换类型