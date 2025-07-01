
# 🛒 Go E-Commerce Backend

Dự án này là hệ thống web bán hàng được xây dựng bằng Golang với các công nghệ hiện đại như:

- 🌐 Gin-Gonic (web framework)
- 🛢️ GORM (ORM cho PostgreSQL)
- 🔐 JWT Authentication & Role-based Authorization
- 🧠 Dependency Injection (Google Wire)
- 📦 Redis Caching
- 📨 RabbitMQ (Queue giả lập gửi email khi tạo đơn hàng)
- 📊 OpenTelemetry (Tracing & Metrics)
- 📄 Swagger Documentation
- 🐳 Docker & Docker Compose
- 📦 Logging tập trung (Fluent Bit → Elasticsearch hoặc Grafana Loki)

---

## 🚀 Chức năng chính

- Đăng ký / Đăng nhập người dùng (JWT)
- Phân quyền: người dùng thường / quản trị viên
- CRUD sản phẩm (admin)
- Tạo đơn hàng (user)
- Xem đơn hàng cá nhân hoặc toàn bộ (admin)
- Gửi message vào RabbitMQ khi tạo đơn hàng
- Caching Redis danh sách sản phẩm
- Swagger UI tại `/swagger/index.html`

---

## 🛠️ Cài đặt

### 1. Clone dự án

```bash
git clone https://github.com/quocphong204/go-ecommerce.git
cd go-ecommerce
```

### 2. Cấu hình biến môi trường (Docker tự động dùng)

```env
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=ecommerce

REDIS_HOST=redis
REDIS_PORT=6379

RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
```

### 3. Chạy Docker Compose

```bash
docker-compose up --build
```

App sẽ chạy tại: [http://localhost:8080](http://localhost:8080)

Swagger: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## 📦 Cấu trúc thư mục

```
.
├── cmd/                  # Main entrypoint
├── internal/
│   ├── api/              # Handlers, routes, middleware
│   ├── config/           # Kết nối DB, Redis, RabbitMQ
│   ├── model/            # Structs cho entity
│   ├── repository/       # Giao tiếp DB
│   ├── service/          # Logic nghiệp vụ
│   ├── middleware/       # Xác thực và phân quyền
│   ├── producer/         # RabbitMQ publisher
│   ├── consumer/         # RabbitMQ consumer (giả lập xử lý)
│   └── logger/           # Zap logger
├── go.mod
├── Dockerfile
├── docker-compose.yml
└── README.md
```

---

## 🧪 Swagger Test API

- GET /products
- POST /admin/products
- POST /auth/login
- GET /me
- POST /orders

---

## ✅ Kỹ thuật nổi bật

| Tính năng               | Trạng thái |
|------------------------|------------|
| Auth với JWT           | ✅         |
| Phân quyền (admin/user)| ✅         |
| CRUD sản phẩm          | ✅         |
| Order & RabbitMQ       | ✅         |
| Redis Cache            | ✅         |
| Wire DI                | ✅         |
| Swagger UI             | ✅         |
| Docker + Postgres      | ✅         |
| RabbitMQ container     | ✅         |
| Central Logging        | ⏳         |
| OpenTelemetry          | ⏳         |

---

## 📜 Ghi chú thêm

- Tài khoản `admin` có thể gọi các route `/admin/*`
- Redis đang cache danh sách sản phẩm `/products`
- Khi tạo đơn hàng, message gửi vào RabbitMQ (giả lập xử lý gửi email)
- Logging dùng Zap (sẽ mở rộng sang Fluent Bit hoặc Loki)

---

## ✍️ Tác giả

**Phong Liêu**  
Dự án thực hiện cho mục đích học tập & thực tập backend Golan
