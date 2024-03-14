package services

import (
	"errors"
	"fmt"
	"puzzle/app/models"
	"puzzle/database"
	"puzzle/utils"
	"strconv"
	"strings"
	"time"
)

type ScrambleSerivce interface {
	check(scramble *models.Scramble) error
	Insert(scramble *models.Scramble) error
	List(scrambleReq *models.ScrambleReq) (models.ScrambleListResp, error)
	GetNewScamble(getNewScrambleReq *models.GetNewScambleReq) (models.ScrambleResp, error)
	GetUserScramble(getNewScrambleReq *models.GetNewScambleReq) (models.ScrambleResp, error)
}

type ScrambleImpl struct{}

// check 检查数据
func (ScrambleImpl) check(scramble *models.Scramble) error {
	if scramble.Dimension == 0 {
		return errors.New("阶数不能为空")
	}

	if scramble.Scramble == "" {
		return errors.New("打乱公式不能为空")
	}

	if scramble.Idx == 0 {
		return errors.New("随机数索引不能为空")
	}

	return nil
}

// Insert 插入打乱公式
func (ScrambleImpl) Insert(scramble *models.Scramble) error {
	if err := Scramble.check(scramble); err != nil {
		return err
	}

	snowflake := utils.Snowflake{}
	scramble.Id = snowflake.NextVal()
	scramble.Status = 1

	return database.GetMySQL().Create(scramble).Error
}

// List 获取打乱公式列表
func (ScrambleImpl) List(scrambleReq *models.ScrambleReq) (models.ScrambleListResp, error) {
	var scrambleListResp models.ScrambleListResp
	db := database.GetMySQL().Table("scramble")

	if scrambleReq.Id != 0 {
		db.Where("id = ?", scrambleReq.Id)
	}

	if scrambleReq.Dimension != 0 {
		db.Where("dimension = ?", scrambleReq.Dimension)
	}

	if scrambleReq.Idx != 0 {
		db.Where("idx = ?", scrambleReq.Idx)
	}

	if scrambleReq.Scramble != "" {
		db.Where("scramble = ?", scrambleReq.Scramble)
	}

	if scrambleReq.Status != 0 {
		db.Where("status = ?", scrambleReq.Status)
	}

	if len(scrambleReq.DateRange) == 2 {
		db.Where("created_at BETWEEN ? AND ?", scrambleReq.DateRange[0], scrambleReq.DateRange[1])
	}

	if scrambleReq.Sorted != "" {
		db.Order("created_at " + scrambleReq.Sorted)
	}

	// 获取总数
	err := db.Count(&scrambleListResp.Total).Error
	if err != nil {
		return scrambleListResp, errors.New("查询失败")
	}

	// 分页
	if scrambleReq.Pagination.PageSize > 0 && scrambleReq.Pagination.Page > 0 {
		db.Scopes(utils.Paginate(&scrambleReq.Pagination))
	}

	// 查询记录
	err = db.Find(&scrambleListResp.Records).Error
	if err != nil {
		return models.ScrambleListResp{}, errors.New("查询失败")
	}

	return scrambleListResp, nil
}

// GetNewScamble 获取新的打乱公式
func (ScrambleImpl) GetNewScamble(getNewScrambleReq *models.GetNewScambleReq) (models.ScrambleResp, error) {
	var scrambleResp models.ScrambleResp

	// 获取用户当前的完成状态
	scrambledUserStatusReq := models.ScrambledUserStatusReq{
		UserId:    getNewScrambleReq.UserId,
		Dimension: getNewScrambleReq.Dimension,
		Pagination: utils.Pagination{
			PageSize: 1,
			Page:     1,
		},
	}

	scrambledUserStatusResp, err := ScrambledUserStatus.List(&scrambledUserStatusReq)
	if err != nil {
		return models.ScrambleResp{}, errors.New("查询用户打乱公式状态失败")
	}

	// 如果没有找到用户的完成状态，或是用户的完成状态为已完成，则生成新的打乱公式
	if scrambledUserStatusResp.Total == 0 || scrambledUserStatusResp.Records[0].Status == 2 {
		idx := time.Now().UnixMilli()
		scramble := utils.Shuffle(getNewScrambleReq.Dimension, int(idx))
		scrambleStr := strings.Trim(strings.Replace(fmt.Sprint(scramble), " ", ",", -1), "[]")

		snowflake := utils.Snowflake{}

		scrambleModel := &models.Scramble{
			Id:        snowflake.NextVal(),
			Dimension: getNewScrambleReq.Dimension,
			Idx:       idx,
			Scramble:  scrambleStr,
			Status:    1,
		}

		err = Scramble.Insert(scrambleModel)
		if err != nil {
			return models.ScrambleResp{}, err
		}

		// 查询用户的打乱公式状态
		if scrambledUserStatusResp.Total == 0 {
			// 新增用户的打乱公式状态
			scrambledUserStatus := &models.ScrambledUserStatus{
				UserId:     getNewScrambleReq.UserId,
				Dimension:  getNewScrambleReq.Dimension,
				ScrambleId: scrambleModel.Id,
				Status:     1,
			}

			err = ScrambledUserStatus.Insert(scrambledUserStatus)
			if err != nil {
				return models.ScrambleResp{}, err
			}
		} else {
			// 更新用户的打乱公式状态
			id, _ := strconv.ParseInt(scrambledUserStatusResp.Records[0].Id, 10, 64)

			scrambledUserStatus := &models.ScrambledUserStatus{
				Id:         id,
				ScrambleId: scrambleModel.Id,
				Status:     1,
			}

			err = ScrambledUserStatus.Update(scrambledUserStatus)
			if err != nil {
				return models.ScrambleResp{}, err
			}
		}

		// 返回打乱公式
		scrambleResp = models.ScrambleResp{
			Id:        scrambleModel.Id,
			Dimension: scrambleModel.Dimension,
			Idx:       scrambleModel.Idx,
			Scramble:  scrambleModel.Scramble,
			Status:    scrambleModel.Status,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		return scrambleResp, nil
	}

	return models.ScrambleResp{}, errors.New("当前的打乱公式未完成")
}

// GetUserScramble 获取用户的打乱公式
func (ScrambleImpl) GetUserScramble(getNewScrambleReq *models.GetNewScambleReq) (models.ScrambleResp, error) {
	var scrambleResp models.ScrambleResp

	// 获取用户当前的打乱信息
	scrambledUserStatusReq := &models.ScrambledUserStatusReq{
		UserId:    getNewScrambleReq.UserId,
		Dimension: getNewScrambleReq.Dimension,
		Pagination: utils.Pagination{
			PageSize: 1,
			Page:     1,
		},
	}

	scrambledUserStatusResp, err := ScrambledUserStatus.List(scrambledUserStatusReq)
	if err != nil {
		return scrambleResp, errors.New("查询用户打乱公式状态失败")
	}

	// 如果没有找到用户的完成状态，或是用户的完成状态为已完成，则不生成新的打乱公式
	if scrambledUserStatusResp.Total == 0 || scrambledUserStatusResp.Records[0].Status == 2 {
		return scrambleResp, nil
	}

	// 查询打乱信息
	scrambleReq := &models.ScrambleReq{
		Id: scrambledUserStatusResp.Records[0].ScrambleId,
	}

	scrambleListResp, err := Scramble.List(scrambleReq)
	if err != nil {
		return scrambleResp, errors.New("查询打乱公式失败")
	}

	if scrambleListResp.Total == 0 {
		return scrambleResp, errors.New("未找到打乱公式")
	}

	return scrambleListResp.Records[0], nil
}
