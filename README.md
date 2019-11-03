### 服务池—发号器

#### 项目前置

#### 项目配置
```golang
    需要配置 wb-impulse-sender/config/app.dev的mysql和redis信息
    mysql：用于存储和生成发号器服务的NodeId
    redis：用于存储发号器生成的号码
```

#### 运行
```golang
    go run main.go --conf=$GOPATH/src/wb-impulse-sender/config/app.dev.toml
```

#### 获取号码
```golang
    http://127.0.0.1:8991/pond/number
```

#### 项目目录信息
```golang
    D03B4CDA-2CA8-4691-9D15-C2DCB263283A.jpeg
```