# GoFrame框架
参考 https://goframe.org
## gf-web
### gen ctrl
> 通过定义api层参数生成controller
>
```shell
# 生成对应的ctrl并且相同包下的合并为一个ctrl
gf gen ctrl --merge=true
```
### gen dao
> 通过定义数据库表生成对应entity/do/dao
> 
hack/config.yaml文件中定义生成的路径(如果有分模块的需求可以使用)
```shell
gf gen dao 
```

### gen service
> 通过定义logic中的业务逻辑生成对应的service并注入
>
hack/config.yaml文件中定义生成的路径(如果有分模块的需求可以使用)
```shell
gf gen service 
```

## 支持常见的CURD、PAGE、TX、ORM操作