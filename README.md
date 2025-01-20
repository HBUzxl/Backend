# 在线会诊系统后端

## 项目简介
在线会诊系统后端服务，基于Go语言和Gin框架开发，提供病例管理、用户认证、专家诊断等功能的API服务。

## 技术栈
- Go 1.20+
- Gin Web Framework
- GORM
- MySQL
- JWT Authentication
- Excelize (Excel导出)

## 目录结构
```
backend/
├── main.go          // 程序入口文件
├── config/
│   └── config.go    // 数据库配置和其他全局配置
├── handlers/        // HTTP请求处理器
│   ├── auth_handler.go    // 认证相关处理
│   ├── case_handler.go    // 病例相关处理
│   └── user_handler.go    // 用户相关处理
├── middleware/
│   └── jwt.go            // JWT认证中间件
├── models/              // 数据模型
│   ├── case.go          // 病例模型
│   ├── user.go          // 用户模型
│   └── slice.go         // 切片模型
├── routes/
│   └── routes.go        // 路由配置
├── services/           // 业务逻辑层
│   ├── auth_service.go  // 认证服务
│   ├── case_service.go  // 病例服务
│   └── user_service.go  // 用户服务
└── README.md
```

## API接口

### 认证相关 (/api/auth)
```
POST /api/auth/login              - 用户登录
GET  /api/auth/current-user       - 获取当前用户信息
POST /api/auth/change-password    - 修改密码
```

### 病例管理 (/api/case)
```
POST /api/case/unsubmitted        - 获取未提交的病例
POST /api/case/pendingdiagnosis   - 获取待诊断的病例
POST /api/case/diagnosed          - 获取已诊断的病例
POST /api/case/returned           - 获取已退回的病例
POST /api/case/withdraw           - 获取已撤回的病例
GET  /api/case/all               - 获取所有病例
GET  /api/case/:caseID          - 根据病例ID获取病例
POST /api/case/submit           - 提交病例
GET  /api/case/excel            - 导出Excel报表

// 状态更新接口
POST /api/case/toPendingdiagnosis/:caseID  - 更新状态：到待诊断
POST /api/case/toDiagnosed/:caseID         - 更新状态：到已诊断
POST /api/case/toReturned/:caseID          - 更新状态：到已退回
POST /api/case/toWithdraw/:caseID          - 更新状态：到已撤回
POST /api/case/:caseID/print               - 增加打印次数
```

### 专家管理 (/api/expert)
```
GET  /api/expert/                 - 获取专家列表
```

### 文件上传 (/api/slice & /api/attachment)
```
POST /api/slice/upload           - 上传切片
POST /api/attachment/upload      - 上传附件
```

## 数据模型

### Case 模型
```go
type Case struct {
    ID              uint      `gorm:"primarykey"`
    CaseID          string    // 病理号
    ConsultationID  string    // 会诊编号
    PathologyType   string    // 病例类型
    PatientName     string    // 患者姓名
    PatientGender   string    // 患者性别
    PatientAge      int       // 患者年龄
    Hospital        string    // 医院
    Department      string    // 科室
    CaseStatus      string    // 状态(unsubmitted/pendingdiagnosis/diagnosed/returned/withdraw)
    DiagnosisContent string   // 诊断内容
    ExpertID        uint      // 专家ID
    Expert          User      // 专家信息
    PrintCount      int       // 打印次数
    SubmitAt        time.Time // 提交时间
    CreatedAt       time.Time
    UpdatedAt       time.Time
}
```

### User 模型
```go
type User struct {
    ID        uint      `gorm:"primarykey"`
    Username  string    `gorm:"unique"`
    Password  string
    Role      string    // 角色
    NickName  string    // 昵称
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

## 运行项目

1. 安装依赖
```bash
go mod download
```

2. 配置数据库
编辑 `config/config.go` 文件，设置数据库连接信息

3. 运行项目
```bash
go run main.go
```

## 注意事项
- 所有API请求需要在Header中携带JWT Token（除了登录接口）
- Token格式：`Authorization: Bearer <token>`
- 文件上传大小限制：50MB
- Excel导出功能会自动格式化日期和状态信息

```
uploads/
└── slices/
    └── case_{caseID}/
        ├── original/            # 存放原始SVS文件
        │   └── {sliceID}.svs
        └── dzi/                # 存放转换后的DZI文件
            └── {sliceID}/      # DZI文件夹
                ├── {sliceID}.dzi
                └── {sliceID}_files/
    
```