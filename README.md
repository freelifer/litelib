# litelib


### 函数

```
litelib := NewLiteLib()
litelib.SetConfigPath("testdata/conf.ini")
litelib.Run()
```

### 配置

#### 默认配置
```
[DEFAULT]
name=项目名称
port=http监听端口

```

#### 模块配置

##### 数据库模块配置
```
[database]
type = sqlite3
host = 127.0.0.1:3306
name = blog
user = root
passwd = rootroot
path = data.db
```