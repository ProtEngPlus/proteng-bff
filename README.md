# proteng-bff

proteng-bff เป็น API gateway ตัวเดียวที่หน้าเว็บเรียก ทำหน้าที่ตรวจ JWT และสิทธิ์ของผู้เรียก แล้วส่ง request ต่อไปยัง [proteng-user-mgmt](https://github.com/ProtEngPlus/proteng-user-mgmt) หรือ [proteng-conductor](https://github.com/ProtEngPlus/proteng-conductor) ตาม path ตัวมันเองไม่มี database และไม่ต่อ RabbitMQ Swagger ของ bff จึงเป็น API docs ของทั้งระบบ ภาพรวมของระบบอยู่ที่ [manual-guides-2023](https://github.com/ProtEngPlus/manual-guides-2023/blob/main/reference/architecture.md)

## เริ่มใช้

ถ้ายังไม่เคยตั้งเครื่อง ให้ทำตาม [tutorials/01-local-setup.md](https://github.com/ProtEngPlus/manual-guides-2023/blob/main/tutorials/01-local-setup.md) ของ hub ซึ่งตั้งทุก repo พร้อมกัน ถ้าจะตั้งเฉพาะ repo นี้ ให้รันใน Git Bash

```sh
make setup
make -C ../manual-guides-2023 local-keys
make run
```

- `make setup` ติดตั้ง git hook สร้าง `.env.local` จาก `.env.example` (ถ้ามีอยู่แล้วจะไม่ทับ) และดาวน์โหลด Go module
- `make -C ../manual-guides-2023 local-keys` ใส่ `ACCESS_TOKEN_PUBLIC_KEY` ที่เป็นคู่กับ private key ของ user-mgmt ในเครื่อง ถ้าไม่มีค่านี้ bff จะตรวจ token ไม่ผ่านทุกครั้ง
- `make run` สร้าง Swagger ใหม่แล้วรันด้วย `ENV=local` ผ่านเมื่อเห็น `proteng-bff is running on :8080` ดู API ทั้งหมดได้ที่ <http://localhost:8080/swagger/index.html>

bff ต้องมี user-mgmt และ conductor รันอยู่จึงจะตอบ request ได้จริง ต้องใช้ Go 1.22 ขึ้นไป และ `pip install pre-commit`

## คำสั่ง

| คำสั่ง | ทำอะไร |
| --- | --- |
| `make setup` | ติดตั้ง hook, สร้าง `.env.local` และดาวน์โหลด module รันซ้ำได้ |
| `make run` | สร้าง Swagger ใหม่แล้วรันในเครื่อง |
| `make swagger` | สร้าง `docs/` ใหม่จาก annotation ของ handler (ติดตั้ง swag v1.16.4 ให้ถ้ายังไม่มี) |
| `make check` | gofmt, `go vet` และ `go test` เหมือนกับ CI |
| `make fmt` | จัด format ด้วย gofmt |
| `make build` และ `make docker-build` | compile ทุก package และ build image ในเครื่อง |

## Config

bff อ่าน `.env.local` เมื่อรันด้วย `ENV=local` ส่วน dev และ production ได้ค่าจาก ConfigMap และ SealedSecret ใน devops-k8s

| ตัวแปร | ค่าตอนรัน local | ใช้ทำอะไร |
| --- | --- | --- |
| `HTTP_PORT` | `8080` | port ที่ bff ฟัง |
| `USER_MGMT_URL` | `http://localhost:8082` | ที่อยู่ของ user-mgmt |
| `CONDUCTOR_URL` | `http://localhost:8081` | ที่อยู่ของ conductor |
| `FRONTEND_URLS` | `http://localhost:5173` | origin ที่อนุญาตสำหรับ CORS |
| `ACCESS_TOKEN_PUBLIC_KEY` | ใส่ด้วย `make -C ../manual-guides-2023 local-keys` | base64 ของ public key แบบ PEM ใช้ตรวจ JWT ที่ user-mgmt เซ็น |

## โครงสร้างโค้ด

| ที่อยู่ | มีอะไร |
| --- | --- |
| `apis/router.go` | ทุก endpoint พร้อม middleware ของแต่ละตัว แบ่งเป็นสองกลุ่มคือ `/proteng-user-mgmt/*` และ `/proteng-conductor/*` |
| `apis/*_handler.go` | handler ของแต่ละ endpoint และ swagger annotation ของ endpoint นั้น |
| `apis/forwarder.go` | ฟังก์ชันที่ส่ง request ต่อไปยัง service ปลายทาง เช่น `Forward`, `ForwardAddParam`, `ForwardAddBody`, `ForwardStrict` และ `ForwardWithRawDataResponse` สำหรับไฟล์ที่ดาวน์โหลด |
| `middleware/authentication.go` | `Authenticate()` ตรวจ JWT และ `Authorize(roles...)` ตรวจ role |
| `utils/token.go` | `ValidateToken` ตรวจลายเซ็นด้วย public key |
| `models/` | struct ของ request และ response ที่ใช้ใน swagger และ `ForwardStrict` |
| `docs/` | ไฟล์ที่ swag สร้าง (`docs.go`, `swagger.json`, `swagger.yaml`) |

### เพิ่ม endpoint

1. เพิ่ม handler ใน `apis/<กลุ่ม>_handler.go` โดยเรียก forwarder ที่เหมาะสม ถ้าต้องมี struct ใหม่ให้เพิ่มใน `models/`
2. ใส่ swagger annotation (`@Summary`, `@Router`, `@Success` และอื่น ๆ) แบบเดียวกับ handler เดิม
3. เพิ่ม route ใน `apis/router.go` พร้อม `middleware.Authenticate()` และ `middleware.Authorize(...)` ถ้า endpoint ต้องการสิทธิ์
4. รัน `make swagger` แล้ว commit ไฟล์ใน `docs/` ไปพร้อมกับโค้ด

## ข้อควรระวัง

- `docs/` เป็นของ swag ทั้งโฟลเดอร์ ห้ามแก้ด้วยมือ และห้ามวางไฟล์อื่นไว้ในนั้น
- public key ต้องเป็นคู่กับ private key ของ user-mgmt ใน environment เดียวกัน และ base64 ต้องไม่มี CR ถ้าสร้างเองบน Windows ให้ใช้ `openssl rsa -in private.pem -pubout | tr -d '\r' | openssl base64 -A`
- endpoint ใหม่ต้องใส่ `Authenticate()` เสมอ ยกเว้น endpoint ที่ตั้งใจให้เรียกได้โดยไม่ login เช่น สมัครสมาชิกและ login ปัญหาเรื่อง endpoint ที่เปิดโล่งอยู่ในหัวข้อ Security ของ [explanation/known-issues.md](https://github.com/ProtEngPlus/manual-guides-2023/blob/main/explanation/known-issues.md)

## Deploy

push เข้า `dev` จะ build image และ deploy ขึ้น dev ส่วน push เข้า `main` จะ deploy ขึ้น production ทั้งสองแบบเกิดขึ้นทันทีทุกครั้งที่ push แม้จะแก้แค่ docs วิธีตรวจและ rollback อยู่ที่ [how-to/deploy-app.md](https://github.com/ProtEngPlus/manual-guides-2023/blob/main/how-to/deploy-app.md)

## ลิงก์

- กติกาการทำงานและ hook ของ repo นี้: [CONTRIBUTING.md](./CONTRIBUTING.md)
- เอกสารของทั้งระบบ: [manual-guides-2023](https://github.com/ProtEngPlus/manual-guides-2023/blob/main/README.md)
