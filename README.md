

## 目录结构
```
backend/
├── main.go          // 程序入口文件
├── config/
│   └── config.go    // 配置文件
├── handlers/        
│   ├── admin.go     // 管理员处理
│   ├── expert.go    // 专家处理
│   ├── allocator.go // 分配员处理
│   └── auth.go      // 登录认证处理
├── middleware/
│   └── auth.go      // 认证中间件
├── models/
│   ├── case.go      // 病例模型
│   ├── user.go      // 用户模型
│   └── ...
├── routes/
│   └── routes.go    // 路由文件
├── services/
│   ├── admin.go     // 管理员服务
│   ├── expert.go    // 专家服务
│   ├── allocator.go // 分配员服务
│   └── auth.go      // 登录认证服务
└── README.md        // 项目说明
```

## API结构

1. 认证相关 API (/api/v1/auth)
```
POST /api/v1/auth/login          - 用户登录
POST /api/v1/auth/logout         - 用户登出
GET  /api/v1/auth/profile        - 获取当前用户信息
PUT  /api/v1/auth/password       - 修改密码
```

2. 用户管理相关 API (/api/v1/users) (管理员权限)
```
GET    /api/v1/users             - 获取用户列表
POST   /api/v1/users             - 创建新用户
GET    /api/v1/users/:id         - 获取特定用户信息
PUT    /api/v1/users/:id         - 更新用户信息
DELETE /api/v1/users/:id         - 删除用户
GET    /api/v1/users/experts     - 获取专家列表
```

3. 诊断管理相关 API (/api/v1/cases) (专家权限)
```
GET    /api/v1/diagnoses              - 获取诊断列表
GET    /api/v1/diagnoses/assigned     - 获取分配给我的诊断
POST   /api/v1/diagnoses/:caseId      - 提交诊断结果
PUT    /api/v1/diagnoses/:id          - 更新诊断结果
GET    /api/v1/diagnoses/:id/history  - 获取诊断历史记录
```

4. 病例管理相关 API (/api/v1/applications) (分配员权限)
```
GET    /api/v1/cases             - 获取病例列表
POST   /api/v1/cases             - 创建新病例
GET    /api/v1/cases/:id         - 获取特定病例详情
PUT    /api/v1/cases/:id         - 更新病例信息
DELETE /api/v1/cases/:id         - 删除病例
POST   /api/v1/cases/:id/assign  - 分配病例给专家
```

5. 统计报告相关 API (/api/v1/reports) (管理员权限)
```
GET /api/v1/reports/diagnosis   - 获取诊断统计报告
GET /api/v1/reports/assignment  - 获取分配统计报告
```

6. 文件上传相关 API (/api/v1/files) (管理员权限)
```
POST   /api/v1/files/upload      - 上传文件（病例图片等）
DELETE /api/v1/files/:id         - 删除文件
GET    /api/v1/files/:id         - 获取文件
```

## 数据库表结构

1. User表

```
CREATE TABLE users (
    id          SERIAL PRIMARY KEY,
    username    VARCHAR(50) UNIQUE NOT NULL,
    password    VARCHAR(255) NOT NULL,
    role        VARCHAR(20) NOT NULL,  -- admin/expert/allocator
    name        VARCHAR(100),
    department  VARCHAR(100),
    title       VARCHAR(50),
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP
);
```

2. Case表

```
CREATE TABLE cases (
    id          SERIAL PRIMARY KEY,
    patient_name VARCHAR(100) NOT NULL,
    age         INT,
    gender      VARCHAR(10),
    symptoms    TEXT,
    description TEXT,
    status      VARCHAR(20),  -- pending/assigned/diagnosed
    created_by  INT REFERENCES users(id),
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP
);
```

3. Diagnosis表

```
CREATE TABLE diagnoses (
    id          SERIAL PRIMARY KEY,
    case_id     INT REFERENCES cases(id),
    expert_id   INT REFERENCES users(id),
    diagnosis   TEXT,
    comments    TEXT,
    status      VARCHAR(20),  -- pending/completed
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP
);
```

4. File表

```
CREATE TABLE files (
    id          SERIAL PRIMARY KEY,
    case_id     INT REFERENCES cases(id),
    file_type   VARCHAR(50),
    file_path   VARCHAR(255),
    created_at  TIMESTAMP
);
```
5