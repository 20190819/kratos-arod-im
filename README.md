### 添加 Proto 文件
```code
kratos proto add api/logic/v1/logic.proto
kratos proto add api/connector/v1/connector.proto
```

### 生成 service
```code
kratos proto server api/logic/v1/logic.proto -t app/logic/internal/service
kratos proto server api/job/v1/job.proto -t app/job/internal/service
kratos proto server api/connector/v1/connector.proto -t app/connector/internal/service
```

### 直接通过 make 命令生成
```code
make api
```

### Reference
```code
git@github.com:gqzcl/arod-im.git (项目不完整)
```