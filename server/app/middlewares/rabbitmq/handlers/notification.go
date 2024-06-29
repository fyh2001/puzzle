package handlers

import (
	"errors"
	"puzzle/app/models"
	"puzzle/database"
	"puzzle/utils"
)

type NotificationMsg struct {
	TypeId  int     `json:"type_id"`
	Content string  `json:"content"`
	UserIds []int64 `json:"user_ids"`
}

func SendNotification(notification any) error {
	msg := notification.(NotificationMsg)

	var notifications []models.Notification

	snowflake := utils.Snowflake{}

	// 发送通知
	for _, userId := range msg.UserIds {
		notification := models.Notification{
			Id:         snowflake.NextVal(),
			UserId:     userId,
			TypeId:     msg.TypeId,
			Content:    msg.Content,
			ReadStatus: 1,
			Status:     1,
		}

		notifications = append(notifications, notification)
	}

	// 插入通知
	db := database.GetMySQL()

	tx := db.Begin()

	for _, notification := range notifications {
		err := tx.Create(&notification).Error
		if err != nil {
			tx.Rollback()
			return errors.New("[rabbitmq]插入通知失败")
		}
	}

	tx.Commit()

	return nil
}
