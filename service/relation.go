package service

import (
	"github.com/gin-gonic/gin"
	"github.com/simple-demo/common"
	"github.com/simple-demo/dao"
)

// RelationAction 首先尝试更新数据库，如果没有数据，则插入新的数据
func RelationAction(fan common.User, toUserId int64, actionType int8) error {
	if affectRows := dao.UpdateRelation(fan.Id, toUserId, actionType); affectRows < 1 {
		err := dao.InsertFavorite(fan.Id, toUserId, actionType)
		return err
	}
	return nil
}

// FollowList 查询关注列表
func FollowList(user common.User) []common.User {
	follows := dao.FindRelationsByFanID(user.Id)
	var IDList []int64
	for _, follow := range follows {
		IDList = append(IDList, follow.UserId)
	}
	res := dao.FindUsersByIDList(IDList)
	return convertUsers(res)
}

// FollowerList 查询粉丝列表
func FollowerList(user common.User) []common.User {
	fans := dao.FindRelationsByUserID(user.Id)
	var IDList []int64
	for _, fan := range fans {
		IDList = append(IDList, fan.UserId)
	}
	res := dao.FindUsersByIDList(IDList)
	return convertUsers(res)
}

// FriendList TODO
func FriendList(c *gin.Context) []common.User {
	return []common.User{}
}
