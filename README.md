# DinoNotify

A Go notification library that supports multiple providers: **LINE Messaging API**, **Slack**, and **Discord**.

---

## Installation

```bash
go get github.com/ukritfongsomboon/DinoNotify
```

---

## Supported Providers

| Provider | Text | Image | File (PDF/CSV) | Video |
|----------|------|-------|----------------|-------|
| LINE Messaging API | ✅ | ✅ | ✅ (URL) | ✅ |
| Slack | ✅ | ✅ | ✅ (URL) | ✅ (URL) |
| Discord | ✅ | ✅ | ✅ (URL) | ✅ (URL) |

---

## Providers

### LINE Messaging API

**Credentials ที่ต้องการ:**
- `Channel Access Token` — ใช้ยืนยันตัวตนกับ LINE API
- `To` — User ID หรือ Group ID ที่ต้องการส่งข้อความไป

**วิธีขอ Credential:**
1. เข้า [LINE Developers Console](https://developers.line.biz)
2. สร้าง Provider และ Channel ประเภท **Messaging API**
3. ไปที่ **Messaging API** → **Channel access token** → กด **Issue**
4. สำหรับ Group ID ให้เพิ่มบอทเข้ากลุ่ม แล้วใช้ Webhook รับ `groupId` จาก event

**ตัวอย่างการใช้งาน:**
```go
line := dinonotify.NewLineMessageAPIProvider(
    "CHANNEL_ACCESS_TOKEN",
    "GROUP_OR_USER_ID",
)
```

---

### Slack

**Credentials ที่ต้องการ:**
- `Webhook URL` — URL สำหรับส่งข้อความเข้า Slack channel

**วิธีขอ Credential:**
1. เข้า [Slack API](https://api.slack.com/apps) → **Create New App** → **From scratch**
2. เลือก workspace และตั้งชื่อ App
3. ไปที่ **Incoming Webhooks** → เปิด **Activate Incoming Webhooks**
4. กด **Add New Webhook to Workspace** → เลือก channel → **Allow**
5. คัดลอก Webhook URL (`https://hooks.slack.com/services/xxx/yyy/zzz`)

**ตัวอย่างการใช้งาน:**
```go
slack := dinonotify.NewSlackProvider(
    "https://hooks.slack.com/services/xxx/yyy/zzz",
)
```

---

### Discord

**Credentials ที่ต้องการ:**
- `Webhook URL` — URL สำหรับส่งข้อความเข้า Discord channel

**วิธีขอ Credential:**
1. เปิด Discord → ไปที่ channel ที่ต้องการ
2. คลิก **Edit Channel** (ไอคอนฟันเฟือง) → **Integrations** → **Webhooks**
3. กด **New Webhook** → ตั้งชื่อ → **Copy Webhook URL**

**ตัวอย่างการใช้งาน:**
```go
discord := dinonotify.NewDiscordProvider(
    "https://discord.com/api/webhooks/xxx/yyy",
)
```

---

## การใช้งาน

### Text Message

ทุก provider รองรับ 4 ระดับ severity โดย `MessagePayload` มี 3 field:

| Field | คำอธิบาย |
|-------|----------|
| `Title` | หัวข้อหลัก เช่น ชื่อ project หรือ service |
| `Subtitle` | หัวข้อรอง เช่น environment หรือ stage |
| `Message` | รายละเอียดของเหตุการณ์ |

**Info** — แจ้งสถานะทั่วไป เช่น service เริ่มทำงาน
```go
provider.Info(dinonotify.MessagePayload{
    Title:    "Project : worker-job",
    Subtitle: "Stage   : UAT",
    Message:  "Service started successfully",
})
```

**Warning** — เตือนเหตุการณ์ที่อาจส่งผลกระทบ แต่ยังไม่ถึงขั้น error
```go
provider.Warning(dinonotify.MessagePayload{
    Title:    "Project : worker-job",
    Subtitle: "Stage   : UAT",
    Message:  "Queue delay > 10s",
})
```

**Error** — แจ้งข้อผิดพลาดที่เกิดขึ้นในระบบ
```go
provider.Error(dinonotify.MessagePayload{
    Title:    "Project : Payment API",
    Subtitle: "Stage   : UAT",
    Message:  "Cannot connect to PostgreSQL on port 5432",
})
```

**Success** — แจ้งเมื่อการทำงานสำเร็จหรือระบบกลับมาปกติ
```go
provider.Success(dinonotify.MessagePayload{
    Title:    "Project : worker-job",
    Subtitle: "Stage   : UAT",
    Message:  "Queue delay back to normal",
})
```

รูปแบบข้อความที่ได้:
```
⚠️ WARNING (2026-04-23 10:35)
──────────────
Project : worker-job
Stage   : UAT
──────────────
Queue delay > 10s
```

### Send File

```go
// Image (แสดงเป็นรูปภาพโดยตรง)
provider.SendFile(dinonotify.FilePayload{
    Name: "screenshot.png",
    URL:  "https://example.com/screenshot.png",
})

// Image URL ที่ไม่มี extension ให้ระบุ Type
provider.SendFile(dinonotify.FilePayload{
    Name: "photo",
    URL:  "https://example.com/photo/123",
    Type: dinonotify.FileTypeImage,
})

// PDF / CSV / ไฟล์อื่น (ส่งเป็น URL)
provider.SendFile(dinonotify.FilePayload{
    Name: "report.pdf",
    URL:  "https://example.com/report.pdf",
})

// Video (LINE แสดง native player, ต้องมี PreviewURL)
provider.SendFile(dinonotify.FilePayload{
    Name:       "demo.mp4",
    URL:        "https://example.com/demo.mp4",
    PreviewURL: "https://example.com/preview.jpg",
    Type:       dinonotify.FileTypeVideo,
})
```

---

## Severity Levels

| Constant | Label | Emoji |
|----------|-------|-------|
| `SeverityInfo` | INFO | ℹ️ |
| `SeverityWarning` | WARNING | ⚠️ |
| `SeverityError` | ERROR | 🚨 |
| `SeveritySuccess` | SUCCESS | ✅ |
| `SeverityCritical` | CRITICAL | 🔴 |

---

## ตัวอย่างแบบเต็ม

```go
package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/ukritfongsomboon/DinoNotify/dinonotify"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    line := dinonotify.NewLineMessageAPIProvider(
        os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
        os.Getenv("LINE_TO"),
    )

    line.Error(dinonotify.MessagePayload{
        Title:    "Project : Payment API",
        Subtitle: "Stage   : Production",
        Message:  "Cannot connect to database",
    })
}
```

**.env**
```env
LINE_CHANNEL_ACCESS_TOKEN=your_token_here
LINE_TO=Cxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

SLACK_WEBHOOK_URL=https://hooks.slack.com/services/xxx/yyy/zzz

DISCORD_WEBHOOK_URL=https://discord.com/api/webhooks/xxx/yyy
```

---

## License

MIT License — © 2026 Ukrit Fongsomboon
