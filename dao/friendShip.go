package dao

import (
	"github.com/pkc918/chat/db"
	"github.com/pkc918/chat/dto"
)

func GetFriendList(id string) ([]dto.FriendItemDTO, error) {
	var friendList []dto.FriendItemDTO
	if err := db.DB.Raw("select a.uid, a.created_at, a.avatar, a.email, f.status from account a left join friend_ship f on a.uid = f.friend_id where f.user_id is not null and f.user_id = ?", id).Scan(&friendList).Error; err != nil {
		return nil, err
	}
	return friendList, nil
}
