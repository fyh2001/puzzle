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

var tableName = "notification"

func check(notification *models.Notification) error {
	return nil
}

func Insert(notificationReq *models.NotificationReq) error {

	if notificationReq.UserIdStr != "" {
		notificationReq.UserId, _ = strconv.ParseInt(notificationReq.UserIdStr, 10, 64)
	}

	snowflake := utils.Snowflake{}

	notification := &models.Notification{
		Id:        snowflake.NextVal(),
		UserId:    notificationReq.UserId,
		TypeId:    notificationReq.TypeId,
		Content:   notificationReq.Content,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := check(notification)
	if err != nil {
		return err
	}

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

	db := database.GetMySQL().Table("notification")

	if notificationReq.Sorted != "" {
		db.Order(notificationReq.OrderBy + " " + notificationReq.Sorted)
	}

	if notificationReq.Id != 0 {
		db.Where(tableName+"."+"id = ?", notificationReq.Id)
	}

	if len(notificationReq.Ids) > 0 {
		db.Where(tableName+"."+"id in ?", notificationReq.Ids)
	}

	if notificationReq.UserId != 0 {
		db.Where(tableName+"."+"user_id = 0").Or(tableName+"."+"user_id = ?", notificationReq.UserId)
	}

	if notificationReq.TypeId != 0 {
		db.Where(tableName+"."+"type_id = ?", notificationReq.TypeId)
	}

	if notificationReq.Content != "" {
		db.Where(tableName+"."+"content Like ?", "%"+notificationReq.Content+"%")
	}

	if notificationReq.Status != 0 {
		db.Where(tableName+"."+"status = ?", notificationReq.Status)
	}

	if len(notificationReq.DateRange) == 2 {
		if notificationReq.DateRange[0] != (time.Time{}) {
			db.Where(tableName+"."+"created_at >= ?", notificationReq.DateRange[0])
		}
		if notificationReq.DateRange[1] != (time.Time{}) {
			db.Where(tableName+"."+"created_at <= ?", notificationReq.DateRange[1])
		}
	}

	// 查询总数
	err := db.Count(&notificationListResp.Total).Error
	if err != nil {
		return notificationListResp, errors.New("查询通知总数失败")
	}

	db.Preload("NotificationTypeInfo")
	db.Preload("NotificationUserStatusInfo", "notification_user_status.user_id = ?", notificationReq.UserId)

	db.Joins("LEFT JOIN notification_user_status ON notification.id = notification_user_status.notification_id AND notification_user_status.user_id = ?", notificationReq.UserId).Order("notification_user_status.status asc").Order("notification.updated_at desc")

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
