package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"me-too-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "me_too_db")
	sslmode := getEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Taipei",
		host, port, user, password, dbname, sslmode)

	var db *gorm.DB
	var err error

	// Retry connection for up to 30 seconds (useful during docker compose startup)
	for i := 1; i <= 15; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			break
		}
		log.Printf("Connecting to PostgreSQL attempt %d/15 failed: %v. Retrying in 2 seconds...", i, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto-migrate tables
	err = db.AutoMigrate(
		&models.User{},
		&models.DailyTask{},
		&models.Post{},
		&models.PostJoin{},
		&models.TaskSubmission{},
		&models.MyTask{},
		&models.Inspiration{},
		&models.PortfolioItem{},
	)
	if err != nil {
		return nil, fmt.Errorf("auto migration failed: %w", err)
	}

	DB = db
	SeedData(db)
	return db, nil
}

func SeedData(db *gorm.DB) {
	var userCount int64
	db.Model(&models.User{}).Count(&userCount)
	if userCount > 0 {
		return // Data already seeded
	}

	log.Println("Seeding initial data based on Figma specifications...")

	// 1. Primary User
	primaryUser := models.User{
		Username:   "lavender_nation",
		Password:   "123456",
		Nickname:   "薰衣草國度",
		AvatarURL:  "",
		Bio:        "把靈感變成行動，把行動留下成果",
		WeeklyInfo: "本週完成 5 個任務，收藏 8 則靈感",
	}
	db.Create(&primaryUser)

	// 2. Daily Tasks (Screen 1 & 8)
	todayStr := time.Now().Format("2006-01-02")
	dailyTasks := []models.DailyTask{
		{UserID: primaryUser.ID, Title: "完成專題使用流程圖", IsCompleted: false, TaskDate: todayStr},
		{UserID: primaryUser.ID, Title: "拍一張今天的天空", IsCompleted: false, TaskDate: todayStr},
		{UserID: primaryUser.ID, Title: "整理書桌 10 分鐘", IsCompleted: false, TaskDate: todayStr},
	}
	db.Create(&dailyTasks)

	// 3. Posts (Screen 2 & 3)
	posts := []models.Post{
		{
			UserID:       primaryUser.ID,
			AuthorName:   "bbb",
			AuthorAvatar: "",
			PostType:     "task",
			Title:        "拍一張今天的天空",
			Content:      "用一張照片記錄今天的狀態，完成後分享一句話。",
			ExpiresInfo:  "任務貼文 · 剩 18 小時",
			ResultCount:  3,
			LikesCount:   720,
		},
		{
			UserID:       primaryUser.ID,
			AuthorName:   "dafuu0000",
			AuthorAvatar: "",
			PostType:     "inspiration",
			Title:        "把通勤路上的顏色做成色票",
			Content:      "看到有趣的配色先收藏，之後可以轉成自己的創作。",
			ExpiresInfo:  "靈感貼文 · 收藏保留 30 天",
			ImageURL:     "placeholder-palette",
			LikesCount:   86,
			CommentsCount: 4,
			SavesCount:   12,
		},
		{
			UserID:       primaryUser.ID,
			AuthorName:   "aaa",
			AuthorAvatar: "",
			PostType:     "task",
			Title:        "整理書桌10分鐘",
			Content:      "去整理啦",
			ExpiresInfo:  "任務貼文 · 剩 18 小時",
			ResultCount:  0,
			LikesCount:   720,
		},
	}
	db.Create(&posts)

	// PostJoin for post 1 & post 3 (matching seed MyTasks)
	if len(posts) > 2 {
		db.Create(&models.PostJoin{
			PostID: posts[0].ID,
			UserID: primaryUser.ID,
		})
		db.Create(&models.PostJoin{
			PostID: posts[2].ID,
			UserID: primaryUser.ID,
		})

		// Submissions for Post 1 (Screen 3 timeline)
		submissions := []models.TaskSubmission{
			{
				PostID:     posts[0].ID,
				AuthorName: "kb_q_p",
				TimeAgo:    "4小時",
				Content:    "應該長這樣",
				ImageURL:   "placeholder-sky-1",
			},
			{
				PostID:     posts[0].ID,
				AuthorName: "同學A",
				TimeAgo:    "2小時",
				Content:    "今天的天空是藍灰色，很適合當背景。",
			},
			{
				PostID:     posts[0].ID,
				AuthorName: "同學B",
				TimeAgo:    "2小時",
				Content:    "我看倒像綠豆糕。",
			},
		}
		db.Create(&submissions)
	}

	// 4. My Tasks (Screen 5)
	post1ID := posts[0].ID
	post3ID := posts[2].ID
	myTasks := []models.MyTask{
		{
			UserID:       primaryUser.ID,
			TaskType:     "private",
			Title:        "完成企劃草稿",
			Description:  "私人任務，永久留在自己的清單",
			IsCompleted:  false,
			ProgressNote: "",
		},
		{
			UserID:       primaryUser.ID,
			PostID:       &post1ID,
			TaskType:     "public",
			Title:        "公共任務：拍一張今天的天空",
			Description:  "來自任務貼文，完成後需分享成果",
			IsCompleted:  false,
		},
		{
			UserID:       primaryUser.ID,
			PostID:       &post3ID,
			TaskType:     "public",
			Title:        "公共任務：整理書桌 10 分鐘",
			Description:  "來自任務貼文，完成後需分享成果",
			IsCompleted:  false,
		},
	}
	db.Create(&myTasks)

	// 5. Inspirations (Screen 6)
	inspirations := []models.Inspiration{
		{
			UserID:         primaryUser.ID,
			IndexNum:       1,
			Title:          "去租書店把柯南漫畫裡的壞人都圈出來",
			DateGroup:      "2026/06/24",
			IsPinned:       false,
		},
		{
			UserID:         primaryUser.ID,
			IndexNum:       2,
			Title:          "蛇的尿道跟肛門是同一個，尿液結晶有時會堵塞",
			PreviewContent: "貼文連結 / 照片預覽",
			ImageURL:       "placeholder-notes",
			DateGroup:      "2026/06/24",
			IsPinned:       false,
		},
	}
	db.Create(&inspirations)

	// 6. Portfolio Items (Screen 7 - 9 items grid)
	portfolioItems := []models.PortfolioItem{
		{UserID: primaryUser.ID, Title: "天空紀錄", Category: "成果", Tag: "任務成果", HasImage: true},
		{UserID: primaryUser.ID, Title: "書桌整理", Category: "成果", Tag: "任務成果", HasImage: false},
		{UserID: primaryUser.ID, Title: "色票靈感", Category: "成果", Tag: "靈感轉化", HasImage: true},
		{UserID: primaryUser.ID, Title: "租書店任務", Category: "成果", Tag: "發起任務", HasImage: true},
		{UserID: primaryUser.ID, Title: "專題草圖", Category: "成果", Tag: "靈感轉化", HasImage: false},
		{UserID: primaryUser.ID, Title: "晚霞照片", Category: "成果", Tag: "任務成果", HasImage: true},
		{UserID: primaryUser.ID, Title: "手寫觀察", Category: "成果", Tag: "私人靈感", HasImage: false},
		{UserID: primaryUser.ID, Title: "桌面改造", Category: "成果", Tag: "任務成果", HasImage: true},
		{UserID: primaryUser.ID, Title: "素材備忘", Category: "成果", Tag: "靈感整理", HasImage: false},
	}
	db.Create(&portfolioItems)

	log.Println("Seed data successfully populated!")
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
