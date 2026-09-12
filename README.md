# 「我也要」(Me Too) - 靈感與行動社群平台

本專案將 Figma 原型設計重構為**現代化電腦版優先 (Desktop-First) 兼具行動裝置自適應 (Mobile RWD)** 的全端 Web 應用。

---

## 🛠️ 技術架構 (Fullstack Architecture)

- **入口網關反向代理 (Gateway / Proxy)**：**Nginx** (gzip 壓縮、靜態快取、統一路由分發 `/` 與 `/api/`)
- **前端 (Frontend)**：**Vue 3** (Composition API / `<script setup>`) + 原生 CSS3 響應式系統 + Vite
- **後端 (Backend)**：**Go (Golang)** + Gin Web Framework + GORM ORM
- **資料庫 (Database)**：**PostgreSQL 16** (自動遷移 Migration 與 Figma 種子資料注入)
- **容器化與編排 (Containerization)**：**Docker & Docker Compose** (四容器架構：`nginx` + `frontend` + `backend` + `db`)

```text
[使用者瀏覽器 (Desktop / Mobile RWD)]
                │
                ▼
       [ Nginx 反向代理網關 ] (Port 80 / 3000)
         │                │
         ▼ (/)            ▼ (/api/*)
  [ Vue 3 前端服務 ]    [ Go 後端 REST API ] (Port 8080)
                          │
                          ▼
                   [ PostgreSQL 16 ] (Port 5432)
```

---

## 📱 響應式與重構設計重點 (Responsive & Layout Design)

1. **電腦版 (Desktop >= 1024px)**：
   - **左側固定選單 (Sidebar)**：品牌識別、主要導航 (貼文區、靈感清單、我的任務、個人主頁)、高亮行動按鈕 (✨ 我也要發文、今日任務規劃)、使用者個人微型卡片。
   - **中間主內容流 (Main Flow)**：最適閱讀寬度 (640px)，提供貼文分類過濾、一鍵「我也要」認領任務、打卡成果與討論串折疊展開 (kb_q_p, 同學A, 同學B)。
   - **右側生產力小工具 (Right Panel)**：今日任務即時勾選與快速新增、靈感隨手速記盒、社群熱門挑戰推薦。
2. **行動裝置 (Mobile < 768px)**：
   - 完美還原 Figma 390px 經典外觀與觸控手感。
   - 頂部模擬手機狀態欄 (9:41 / 電量 100%)。
   - 底部浮動五鍵導航列，中間保留 Figma 特色黑色凸起膠囊「我也要<br/>發文」按鈕。

---

## 🚀 快速啟動指南 (Quick Start)

### 方式 A：使用 Docker Compose 一鍵部署 (推薦)

確保本機已啟動 Docker Desktop，在專案根目錄下執行：

```bash
# 構建並啟動所有容器 (Nginx 網關 + 前端 + Go 後端 + PostgreSQL)
docker compose up --build -d
```

啟動後即可在瀏覽器訪問：
- **Web 應用入口 (Nginx 網關)**：[http://localhost:3000](http://localhost:3000) 或 [http://localhost](http://localhost)
- **後端 API 直接訪問 (可選)**：[http://localhost:8080/api/health](http://localhost:8080/api/health)
- **PostgreSQL 資料庫**：`localhost:5432` (帳號/密碼: `postgres` / `postgres`)

查看容器運行狀態：
```bash
docker compose ps
```

停止並移除容器：
```bash
docker compose down
```

---

### 方式 B：本機開發模式 (Local Development)

#### 1. 啟動後端 (Go)
```bash
cd backend
go run main.go
```
*預設運行於 http://localhost:8080*

#### 2. 啟動前端 (Vue 3)
```bash
cd frontend
npm install
npm run dev
```
*預設運行於 http://localhost:5173，已配置 API Proxy 自動轉發 `/api` 至後端*

---

## 📂 專案目錄結構

```text
c:/專題/
├── docker-compose.yml         # 四容器編排配置 (nginx, frontend, backend, db)
├── README.md                  # 專案說明文件
├── nginx/                     # Nginx 專屬反向代理網關
│   ├── Dockerfile
│   └── default.conf           # 路由轉發與 Gzip 緩存配置
├── backend/                   # Go 後端工程
│   ├── Dockerfile             # Go 輕量多階段編譯鏡像
│   ├── go.mod / go.sum
│   ├── main.go                # API 入口與路由初始化
│   ├── database/
│   │   └── database.go        # PostgreSQL 連線與 Figma 種子資料初始化
│   ├── handlers/
│   │   └── handlers.go        # 業務邏輯控制器 (每日任務、貼文、成果、靈感、個人)
│   └── models/
│       └── models.go          # GORM 資料庫實體定義
└── frontend/                  # Vue 3 前端工程
    ├── Dockerfile             # Node 打包 + Web 服務鏡像
    ├── nginx.conf
    ├── package.json
    ├── vite.config.js
    ├── index.html
    └── src/
        ├── main.js
        ├── App.vue            # 主視圖與跨元件狀態排版
        ├── assets/
        │   └── style.css      # Figma 設計系統、調色盤與 RWD 斷點
        ├── api/
        │   └── client.js      # REST API 請求客戶端
        └── components/        # 核心組件
            ├── Sidebar.vue            # 電腦版左側導航欄
            ├── RightPanel.vue         # 電腦版右側功能小工具
            ├── BottomNav.vue          # 行動版底部導航欄
            ├── FeedView.vue           # 貼文區與成果時間軸 (Figma Screen 2 & 3)
            ├── TasksView.vue          # 我的任務 (Figma Screen 5)
            ├── InspirationsView.vue   # 靈感清單與速記 (Figma Screen 6)
            ├── ProfileView.vue        # 個人主頁與九宮格作品 (Figma Screen 7)
            ├── CreatePostModal.vue    # 我也要發文彈窗 (Figma Screen 4)
            └── DailyTaskModal.vue     # 今天是否有任務彈窗 (Figma Screen 1 & 8)
```
