package service

import (
	"fmt"
	"github.com/simple-demo/common"
	"github.com/simple-demo/dao"
)

// Feed same demo video list for every request
func Feed() []common.Video {
	videos := dao.FindAllVideos()
	return convertVideos(videos)
}

func convertVideos(videos []dao.Video) []common.Video {
	var res []common.Video
	for _, video := range videos {
		if name, err := dao.FindUserByID(video.Author); err == nil {
			res = append(res, common.Video{
				Id: video.ID,
				Author: common.User{
					Id:            video.Author,
					Name:          name,
					FollowCount:   0,
					FollowerCount: 0,
					IsFollow:      false,
				},
				PlayUrl:       video.PlayUrl,
				CoverUrl:      video.CoverUrl,
				FavoriteCount: 0,
				CommentCount:  0,
				IsFavorite:    false,
			})
		} else {
			fmt.Printf("GetPublishList Error: %v", err)
		}
	}
	return res
}
