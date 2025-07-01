# Cache Admin 缓存管理系统

本项目是一个基于 Go + Gin + Redis 的缓存管理系统，配套 Vue3 + Element Plus 前端界面，支持缓存的可视化管理、查询与删除。

## 功能特性
- 用户数据缓存查询与分页展示
- 缓存按前缀/参数删除、模糊删除
- 支持接口式访问和可视化前端操作

## 后端接口说明

### 1. 获取用户缓存
- **接口**：`GET /getUser`
- **参数**：
  - `page`（int，页码）
  - `limit`（int，每页数量）
- **返回**：用户列表

### 2. 删除指定缓存
- **接口**：`GET /cache/del`
- **参数**：
  - `prefix`（string，缓存前缀）
  - `params`（object，参数对象）
- **返回**：删除结果

### 3. 按前缀模糊删除缓存
- **接口**：`GET /cache/delByPrefix`
- **参数**：
  - `prefix`（string，缓存前缀）
- **返回**：删除结果

## 后端启动说明

1. 安装 Go 1.18 及以上版本，并确保已安装 Redis 服务。
2. 配置 Redis 连接参数（见 `config/dev.config.yaml`）。
3. 在项目根目录下安装依赖并启动后端服务：
   ```bash
   go mod tidy
   go run cmd/main.go
   ```
4. 默认后端接口监听 8080 端口，用户缓存接口监听 8081 端口。

## 前端启动说明

1. 进入 `web` 目录，安装依赖：
   ```bash
   cd web
   npm install
   ```
2. 启动前端开发服务器：
   ```bash
   npm run dev
   ```
3. 访问 [http://localhost:3000](http://localhost:3000) 使用可视化界面。

前端已自动代理接口请求到后端，无需额外配置。

## 目录结构
- `cmd/`         后端主程序入口
- `handler/`     Gin 路由处理
- `pkg/`         缓存核心逻辑
- `config/`      配置文件
- `examples/`    示例与测试
- `web/`         前端 Vue3 项目

## 适用场景
- 开发环境缓存调试
- 本地接口测试
- 业务缓存可视化管理

---
如需扩展功能或有其他需求，欢迎提 issue 或联系作者。