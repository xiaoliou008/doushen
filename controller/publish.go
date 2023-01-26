package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/simple-demo/common"
	"github.com/simple-demo/service"
	"net/http"
	"strconv"
)

type VideoListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")

	if _, exist := service.UsersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, service.Publish(token, title, data))
}

// PublishList 获取登录用户的发布列表
func PublishList(c *gin.Context) {
	userID := c.Query("user_id")
	ID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: common.Response{
				StatusCode: 0,
			},
			VideoList: []common.Video{},
		})
		return
	}
	token := c.Query("token")
	user, exist := service.UsersLoginInfo[token]
	if !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	if user.Id != ID {
		c.JSON(http.StatusOK, common.Response{StatusCode: 4, StatusMsg: "User ID doesn't match token"})
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: service.GetPublishList(ID),
	})
}

//// Publish check token then save upload file to public directory
//func PublishDemo(c *gin.Context) {
//	token := c.PostForm("token")
//
//	if _, exist := service.UsersLoginInfo[token]; !exist {
//		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
//		return
//	}
//
//	data, err := c.FormFile("data")
//	if err != nil {
//		c.JSON(http.StatusOK, common.Response{
//			StatusCode: 1,
//			StatusMsg:  err.Error(),
//		})
//		return
//	}
//
//	filename := filepath.Base(data.Filename)
//	user := service.UsersLoginInfo[token]
//	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
//	saveFile := filepath.Join("./public/", finalName)
//	if err := c.SaveUploadedFile(data, saveFile); err != nil {
//		c.JSON(http.StatusOK, common.Response{
//			StatusCode: 1,
//			StatusMsg:  err.Error(),
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, common.Response{
//		StatusCode: 0,
//		StatusMsg:  finalName + " uploaded successfully",
//	})
//}
//
//// PublishList all users have same publish video list
//func PublishListDemo(c *gin.Context) {
//	c.JSON(http.StatusOK, VideoListResponse{
//		Response: common.Response{
//			StatusCode: 0,
//		},
//		VideoList: DemoVideos,
//	})
//}
