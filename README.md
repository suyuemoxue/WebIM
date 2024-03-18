### 基于Gin+Gorm+MySQL+Redis+WebSocket实现类QQ即时通讯系统

### go代理

```
GOPROXY=https://goproxy.cn,direct
```

### gin框架下载

```
go get -u github.com/gin-gonic/gin
```

### gorm框架下载

```
 go get -u gorm.io/gorm
```

```
go get -u gorm.io/driver/mysql
```

### yaml配置文件

```
go get gopkg.in/yaml.v2
```

### 配置文件编写

```yaml
# mysql数据库
mysql:
  host: 118.178.135.40
  port: 3306
  db: personal
  user: root
  password: 123456
  log_level: dev

# 系统运行
system:
  host: "0.0.0.0"
  port: 8080
  env: dev

# 日志
logger:
  level: info
  prefix: "[gvb]"
  director: log
  show_line: true
  log_in_console: true
```

```go
type Config struct {
MySQL  MySQL  `yaml:"mysql"`
System System `yaml:"system"`
Logger Logger `yaml:"logger"`
}

type MySQL struct {
Host     string `yaml:"host"`
Port     int    `yaml:"port"`
DB       string `yaml:"db"`
User     string `yaml:"users"`
Password string `yaml:"password"`
LogLevel string `yaml:"log_level"` //日志等级
}

type System struct {
Host string `yaml:"host"`
Port int    `yaml:"port"`
Env  string `yaml:"env"`
}

type Logger struct {
Level        string `yaml:"level"`
Prefix       string `yaml:"prefix"`
Director     string `yaml:"director"`
Showline     bool   `yaml:"show_line"`      //是否显示行号
LogInConsole bool   `yaml:"log_in_console"` //是否显示打印的路径
}
```

