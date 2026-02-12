# JWT 认证中间件使用指南

## 概述

JWT 认证中间件用于验证请求中的 JWT token，并从中提取 `userid` 存储到 Echo Context 中，供后续的 handler 使用。

## 目录结构

```
internal/
├── middleware/
│   └── jwt.go              # JWT 认证中间件
├── route/
│   └── user.go             # 用户路由（应用了中间件）
└── controller/
    └── user/
        └── user.go         # 用户控制器（使用 userid）
```

## 使用方法

### 1. 中间件功能

中间件会执行以下操作：

1. 从请求头 `Authorization` 中提取 token（格式：`Bearer <token>`）
2. 验证 token 的有效性
3. 从 token 中提取 `userid`
4. 将 `userid` 存储到 Echo Context 中

### 2. 在路由中应用中间件

```go
import (
    "go-echo-demo/internal/middleware"
    "github.com/labstack/echo/v5"
)

func SetupUserRoutes(e *echo.Group) {
    userPath := e.Group("/user")

    // 应用 JWT 认证中间件
    userPath.Use(middleware.JWTAuth())

    // 所有该组下的路由都会经过 JWT 验证
    userPath.GET("/info", user_handler.GetUserInfoController)
    userPath.GET("/profile", user_handler.GetUserProfileController)
}
```

### 3. 在 Handler 中获取 userid

```go
import (
    "go-echo-demo/internal/middleware"
    "github.com/labstack/echo/v5"
)

func GetUserInfoController(c *echo.Context) error {
    // 从 context 中获取已验证的 userid
    userid, ok := middleware.GetUserID(c)
    if !ok {
        return c.JSON(500, map[string]interface{}{
            "code": 500,
            "message": "无法获取用户 ID",
        })
    }

    // 使用 userid 进行后续操作
    // 例如：查询数据库、业务逻辑处理等
    return c.JSON(200, map[string]interface{}{
        "userid": userid,
        "data": "用户数据",
    })
}
```

## API 测试示例

### 1. 生成 Token（假设有登录接口）

```bash
# 登录获取 token
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "test", "password": "123456"}'

# 响应示例
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 2. 使用 Token 访问受保护的接口

```bash
# 成功请求（携带有效 token）
curl -X GET http://localhost:8080/api/user/info \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."

# 响应示例
{
  "code": 200,
  "message": "success",
  "data": {
    "userid": 123,
    "info": "这里是用户信息"
  }
}
```

### 3. 错误情况

#### 缺少 token

```bash
curl -X GET http://localhost:8080/api/user/info

# 响应
{
  "code": 401,
  "message": "缺少认证令牌"
}
```

#### Token 格式错误

```bash
curl -X GET http://localhost:8080/api/user/info \
  -H "Authorization: InvalidToken"

# 响应
{
  "code": 401,
  "message": "认证令牌格式错误"
}
```

#### Token 无效或过期

```bash
curl -X GET http://localhost:8080/api/user/info \
  -H "Authorization: Bearer invalid_token_here"

# 响应
{
  "code": 401,
  "message": "认证令牌无效或已过期",
  "error": "token is malformed: ..."
}
```

## 高级用法

### 1. 为特定路由禁用中间件

如果某个组内的部分路由不需要认证，可以单独定义：

```go
func SetupUserRoutes(e *echo.Group) {
    userPath := e.Group("/user")


    // 受保护的接口
    protectedPath := userPath.Group("")
    protectedPath.Use(middleware.JWTAuth())
    protectedPath.GET("/info", user_handler.GetUserInfoController)
    protectedPath.GET("/profile", user_handler.GetUserProfileController)
}
```

### 2. 多层中间件

```go
func SetupUserRoutes(e *echo.Group) {
    userPath := e.Group("/user")

    // 先验证 JWT
    userPath.Use(middleware.JWTAuth())

    // 再进行权限检查
    userPath.Use(middleware.PermissionCheck())

    userPath.GET("/admin", user_handler.AdminController)
}
```

### 3. 自定义错误响应

如需自定义错误响应格式，可以修改 `middleware/jwt.go` 中的错误返回部分。

## 注意事项

1. **Secret Key 安全性**：当前代码中 JWT 的 secret 硬编码为 `"secret"`，生产环境应从环境变量或配置文件中读取
2. **Token 过期时间**：当前设置为 72 小时，可根据业务需求调整
3. **Token 存储位置**：客户端应将 token 安全存储（如 localStorage、sessionStorage 等）
4. **HTTPS**：生产环境务必使用 HTTPS 传输 token

## 代码改进建议

### 1. 将 Secret Key 配置化

```go
// pkg/jwt.go
var jwtSecret = os.Getenv("JWT_SECRET")

func GenerateToken(id uint) (string, error) {
    // ...
    return token.SignedString([]byte(jwtSecret))
}
```

### 2. 添加 Token 刷新机制

可以添加一个 refresh token 端点，允许客户端在 access token 过期前刷新。

### 3. 添加日志记录

在中间件中添加日志，记录认证失败的请求，便于排查问题和安全审计。
