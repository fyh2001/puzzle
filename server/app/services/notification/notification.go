package notification

import (
	"errors"
	"puzzle/app/models"
	commonService "puzzle/app/services"
	"puzzle/database"
	"puzzle/utils"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func check(notification *models.Notification) error {
	return nil
}

func Insert(notification *models.Notification) error {

	err := check(notification)
	if err != nil {
		return err
	}

	snowflake := utils.Snowflake{}
	notification.Id = snowflake.NextVal()
	notification.Status = 1

	err = database.GetMySQL().Create(notification).Error
	if err != nil {
		return err
	}

	return nil
}

func List(notificationReq *models.NotificationReq) (models.NotificationListResp, error) {
	var notificationListResp models.NotificationListResp

	if notificationReq.Username != "" || notificationReq.Nickname != "" {
		userInfo, err := commonService.GetUserInfoByUsernameOrNickname(notificationReq.Username, notificationReq.Nickname)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return notificationListResp, errors.New("查询用户信息失败")
		}

		if userInfo.Id == "" {
			notificationReq.UserIdStr = "-1"
		} else {
			notificationReq.UserIdStr = userInfo.Id
		}
	}

	if notificationReq.IdStr != "" {
		notificationReq.Id, _ = strconv.ParseInt(notificationReq.IdStr, 10, 64)
	}

	if notificationReq.UserIdStr != "" {
		notificationReq.UserId, _ = strconv.ParseInt(notificationReq.UserIdStr, 10, 64)
	}

	if len(notificationReq.IdsStr) > 0 {
		notificationReq.Ids = make([]int64, len(notificationReq.IdsStr))
		for i, v := range notificationReq.IdsStr {
			notificationReq.Ids[i], _ = strconv.ParseInt(v, 10, 64)
		}
	}

	if notificationReq.OrderBy == "" {
		notificationReq.OrderBy = "updated_at"
	}

	db := database.GetMySQL().Table("notification").Order(notificationReq.OrderBy + " " + notificationReq.Sorted)

	if notificationReq.Id != 0 {
		db.Where("id = ?", notificationReq.Id)
	}

	if len(notificationReq.Ids) > 0 {
		db.Where("id in ?", notificationReq.Ids)
	}

	if notificationReq.UserId != 0 {
		db.Where("user_id = 0 OR user_id = ?", notificationReq.UserId)
	}

	if notificationReq.TypeId == 0 {
		db.Where("type_id = ?", notificationReq.TypeId)
	}

	if notificationReq.Content != "" {
		db.Where("content Like ?", "%"+notificationReq.Content+"%")
	}

	if notificationReq.Status != 0 {
		db.Where("status = ?", notificationReq.Status)
	}

	if len(notificationReq.DateRange) == 2 {
		if notificationReq.DateRange[0] != (time.Time{}) {
			db.Where("created_at >= ?", notificationReq.DateRange[0])
		}
		if notificationReq.DateRange[1] != (time.Time{}) {
			db.Where("created_at <= ?", notificationReq.DateRange[1])
		}
	}

	// 查询总数
	err := db.Count(&notificationListResp.Total).Error
	if err != nil {
		return notificationListResp, errors.New("查询通知总数失败")
	}

	// db.Left

	// 分页
	if notificationReq.Pagination.Page > 0 && notificationReq.Pagination.PageSize > 0 {
		db.Scopes(utils.Paginate(&notificationReq.Pagination))
	}

	// 查询列表
	err = db.Find(&notificationListResp.Records).Error
	if err != nil {
		return notificationListResp, errors.New(("查询通知列表失败"))
	}

	return notificationListResp, nil
}
