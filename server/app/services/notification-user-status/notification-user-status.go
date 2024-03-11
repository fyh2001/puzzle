package notificationuserstatus

import (
	"errors"
	"puzzle/app/models"
	"puzzle/database"
	"puzzle/utils"
	"strconv"
)

func check(notificationUserStatus *models.NotificationUserStatus) error {
	if notificationUserStatus.NotificationId == 0 {
		return errors.New("通知ID不能为空")
	}

	if notificationUserStatus.UserId == 0 {
		return errors.New("用户ID不能为空")
	}

	return nil
}

func Insert(notificationUserStatusReq *models.NotificationUserStatusReq) error {
	notificationId, _ := strconv.ParseInt(notificationUserStatusReq.NotificationIdStr, 10, 64)
	notificationUserStatus := &models.NotificationUserStatus{
		NotificationId: notificationId,
		UserId:         notificationUserStatusReq.UserId,
	}

	err := check(notificationUserStatus)
	if err != nil {
		return err
	}

	snowflake := utils.Snowflake{}

	notificationUserStatus.Id = snowflake.NextVal()
	notificationUserStatus.Status = 1

	err = database.GetMySQL().Create(notificationUserStatus).Error
	if err != nil {
		return err
	}

	return nil
}
