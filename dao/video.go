package dao

type Video struct {
	//gorm.Model
	ID       int64
	Title    string
	Author   int64
	PlayUrl  string
	CoverUrl string
}

func FindAllVideos() []Video {
	var videos []Video
	// 获取全部记录
	db.Find(&videos)
	return videos
}

func FindVideosByAuthor(authorID int64) []Video {
	var videos []Video
	// 获取全部记录
	db.Where(&Video{Author: authorID}).Find(&videos)
	return videos
}

func InsertVideo(title string, author int64, playUrl string, coverUrl string) error {
	video := Video{
		Title:    title,
		Author:   author,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	}
	result := db.Create(&video) // 通过数据的指针来创建

	return result.Error // 返回 error
}
