# Setup

Setup ทั้งระบบครั้งแรกดูที่ [Guidebook](https://github.com/ProtEngPlus/manual-guides-2023/blob/main/README.md) ไฟล์นี้มีแค่รายละเอียดเฉพาะของ `proteng-bff`

## รันบนเครื่อง

ขั้นตอนหลัก (`cp .env.example .env.local` → `go mod tidy` → `./run.sh`) อยู่ใน Guidebook §4.4
ที่ต้องรู้เพิ่มเฉพาะ bff:

- `.env.local` ต้องมี `ACCESS_TOKEN_PUBLIC_KEY` เป็น public key คู่เดียวกับ
  `ACCESS_TOKEN_PRIVATE_KEY` ของ `proteng-user-mgmt` (bff แค่ verify token, user-mgmt เป็นคนเซ็น)
  แปลงจาก private key:

  ```sh
  openssl rsa -in private.pem -pubout | tr -d '\r' | openssl base64 -A
  ```

- `./run.sh` ลง `swag` ถ้ายังไม่มี แล้ว regenerate Swagger doc ก่อนตั้ง `ENV=local` แล้ว
  `go run main.go` (ไม่มี `.env.dev` แล้ว รันมือด้วย `ENV=local go run main.go` ก็ได้)
- bff ไม่ต่อ DB/queue เอง คุยผ่าน HTTP ไป conductor กับ user-mgmt เท่านั้น
- เสร็จเมื่อ terminal พิมพ์ `proteng-bff is running on :8080` (หรือ `HTTP_PORT` ที่ตั้ง) แล้วไม่ crash

## Format & lint

`gofmt` autofix ตอน save/commit, `go vet` รายงานอย่างเดียวต้องแก้เอง รันมือทั้ง repo:

```sh
gofmt -l -w .
go vet ./...
```

ทั้งคู่รันเป็น pre-commit hook ให้อัตโนมัติ (ดู [CONTRIBUTING.md](./CONTRIBUTING.md)) และรันใน CI
ทุก push ด้วย

## API docs

bff เป็น API surface เดียวที่ frontend คุยด้วย (ไม่เรียก conductor/user-mgmt ตรง) เพราะงั้น
Swagger ของ bff คือ API docs ของทั้ง project ไม่ต้องมี doc service แยก

- **ดู**: รัน bff (`./run.sh`) เปิด `http://localhost:8080/swagger/index.html`
- **regenerate**: อัตโนมัติ `./run.sh` รัน `swag init` ทุกครั้งที่สตาร์ท (~0.5s) แค่ใส่ annotation
  `@Router` / `@Success` / ฯลฯ ให้ handler ใหม่ (แบบเดียวกับ handler เดิม) แล้วรัน app ไฟล์
  `docs/docs.go` / `swagger.json` / `swagger.yaml` อัปเดตเอง commit ไฟล์ที่ regenerate ไปพร้อมโค้ด
  handler

## Build (ถ้าจะทดสอบ deploy)

env var ไม่ถูก bake เข้า image ส่งตอน run:

```sh
docker build -t proteng-bff .
docker run -d --name proteng-bff --env-file .env.local -p 8080:8080 proteng-bff
```
