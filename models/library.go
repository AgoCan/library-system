package model

import "github.com/jinzhu/gorm"

// Blog 博客
type Blog struct {
	gorm.Model
	BlogID uint64 `json:"blog_id"`
	Title string  `json:"title"`
	SiteName string `json:"site_name"`
	Theme string    `json:"theme"`
}

// Category 章节
type Category struct {
	gorm.Model
	CategoryID uint64  `json:"category_id"`
	Title string		`json:"title"`
	Blog Blog `gorm:"foreignkey:BlogID"`
}

// Tag 标签
type Tag struct {
	gorm.Model
	TagID uint64 	`json:"tag_id"`
	Title string	`json:"title"`
	Blog Blog `gorm:"foreignkey:BlogID"`
	Articles []*Article `gorm:"many2many:article_tags;"`
}

// Article 文章列表
type Article struct {
	gorm.Model
	ArticleID uint64 	`json:"article_id"`
	Title string		`json:"title"`
	Desc string			`json:"desc"`
	Content string		`json:"content"`
	// 评论和点赞
	CommentCount int	`json:"comment_count"`
	UpCount int			`json:"up_count"`
	DownCount int		`json:"down_count"`
	// 外键没有使用，主要是靠代码实现
	UserID UserInfo `gorm:"foreignkey:UserID"`
	CategoryID UserInfo `gorm:"foreignkey:CategoryID"`

	Tags []*Tag `gorm:"many2many:article_tags;"`

}

type Comment struct {
	gorm.Model
	CommentID uint64
	Article Article
	User UserInfo
	Content string
}

func blogMigrate()  {
	DB.AutoMigrate(&Blog{})
	DB.AutoMigrate(&Category{})
	DB.AutoMigrate(&Tag{})
	DB.AutoMigrate(&Article{})
}