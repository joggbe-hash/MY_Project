package models

import (
	"time"
)

type User struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Username   string    `gorm:"size:64;uniqueIndex" json:"username"`
	Password   string    `gorm:"size:255" json:"-"`
	Nickname   string    `gorm:"size:64" json:"nickname"`
	AvatarURL  string    `gorm:"size:255" json:"avatar_url"`
	Bio        string    `gorm:"size:255" json:"bio"`
	WeeklyInfo string    `gorm:"size:255" json:"weekly_info"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type DailyTask struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"user_id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	IsCompleted bool      `gorm:"default:false" json:"is_completed"`
	TaskDate    string    `gorm:"size:32" json:"task_date"` // YYYY-MM-DD
	CreatedAt   time.Time `json:"created_at"`
}

type Post struct {
	ID          uint             `gorm:"primaryKey" json:"id"`
	UserID      uint             `json:"user_id"`
	AuthorName  string           `gorm:"size:64;not null" json:"author_name"`
	AuthorAvatar string          `gorm:"size:255" json:"author_avatar"`
	PostType    string           `gorm:"size:32;not null" json:"post_type"` // 'task' or 'inspiration'
	Title       string           `gorm:"size:255;not null" json:"title"`
	Content     string           `gorm:"type:text" json:"content"`
	ImageURL    string           `gorm:"size:255" json:"image_url"`
	ExpiresInfo string           `gorm:"size:128" json:"expires_info"` // e.g. "任務貼文 · 剩 18 小時"
	ResultCount int              `gorm:"default:0" json:"result_count"`
	LikesCount  int              `gorm:"default:0" json:"likes_count"`
	CommentsCount int            `gorm:"default:0" json:"comments_count"`
	SavesCount  int              `gorm:"default:0" json:"saves_count"`
	IsJoined    bool             `gorm:"-" json:"is_joined"` // dynamic field computed per user
	Submissions []TaskSubmission `gorm:"foreignKey:PostID" json:"submissions,omitempty"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

type PostJoin struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PostID    uint      `gorm:"index;not null" json:"post_id"`
	UserID    uint      `gorm:"index;not null" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type TaskSubmission struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	PostID      uint      `gorm:"index;not null" json:"post_id"`
	AuthorName  string    `gorm:"size:64;not null" json:"author_name"`
	AuthorAvatar string   `gorm:"size:255" json:"author_avatar"`
	TimeAgo     string    `gorm:"size:64" json:"time_ago"` // e.g. "4小時", "2小時"
	Content     string    `gorm:"type:text" json:"content"`
	ImageURL    string    `gorm:"size:255" json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type MyTask struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `json:"user_id"`
	PostID       *uint     `json:"post_id"` // null if private task
	TaskType     string    `gorm:"size:32;not null" json:"task_type"` // 'private' or 'public'
	Title        string    `gorm:"size:255;not null" json:"title"`
	Description  string    `gorm:"size:255" json:"description"`
	IsCompleted  bool      `gorm:"default:false" json:"is_completed"`
	ProgressNote string    `gorm:"type:text" json:"progress_note"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Inspiration struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	UserID         uint      `json:"user_id"`
	IndexNum       int       `json:"index_num"`
	Title          string    `gorm:"type:text;not null" json:"title"`
	PreviewContent string    `gorm:"size:255" json:"preview_content"`
	ImageURL       string    `gorm:"size:255" json:"image_url"`
	DateGroup      string    `gorm:"size:32" json:"date_group"` // e.g. "2026/06/24"
	IsPinned       bool      `gorm:"default:false" json:"is_pinned"`
	CreatedAt      time.Time `json:"created_at"`
}

type PortfolioItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	Title     string    `gorm:"size:128;not null" json:"title"`
	Category  string    `gorm:"size:64;not null" json:"category"` // '成果', '任務', '靈感', '發起'
	Tag       string    `gorm:"size:64" json:"tag"`               // e.g. '任務成果', '靈感轉化', '發起任務', '私人靈感'
	HasImage  bool      `gorm:"default:true" json:"has_image"`
	ImageURL  string    `gorm:"size:255" json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}
