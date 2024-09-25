package services

import (
	"errors"
	"puzzle/app/models"
	"puzzle/app/services/options"
	"puzzle/app/types"
	"puzzle/database"
)

type RecordBestSingleService struct{}

func (RecordBestSingleService) Create(req *models.RecordBestSingle, reqUid int64) (err error) {

	if err := CheckReq(req); err != nil {
		return err
	}

	if req.UserID != reqUid {
		return errors.New("权限不足")
	}

	if err := database.GetMySQL().Create(req).Error; err != nil {
		return err
	}

	return nil
}

func (RecordBestSingleService) List(req *models.RecordBestSingleReq, reqUid int64) (resp *models.RecordBestSingleListResp, err error) {

	options.ApplyStringOptions(req,
		options.WithStringFieldStr("ID", req.IDStr),
		options.WithStringFieldStr("UserID", req.UserIDStr),
		options.WithStringFieldStr("RecordID", req.RecordIDStr),
	)

	if reqUid != types.ReqAdminUid {
		// 普通用户
		if req.UserID != 0 && req.UserID != reqUid {
			return &models.RecordBestSingleListResp{}, errors.New("权限不足")
		}

		req.UserID = reqUid
	}

	db := database.GetMySQL().Model(&models.RecordBestSingle{})

	filter := []options.QueryOption{
		options.WithQueryID64(req.ID),
		options.WithQueryUserID(req.UserID),
		options.WithQueryDimension(req.Dimension),
		options.WithQueryRecordID(req.RecordID),
		options.WithQueryRecordBreakCount(req.RecordBreakCount),
		options.WithQueryDurationRange(req.DurationRange),
		options.WithQueryStepRange(req.StepRange),
		options.WithQueryDateRange(req.DateRange),
		options.WithQueryRankRange(req.RankRange),
		options.WithQueryBreakCountRange(req.BreakCountRange),
		options.WithQueryPagination(req.Pagination),
		options.WithQueryOrderBy(req.OrderBy, req.Sorted),
		options.WithQueryUserData(req.NeedUserData),
		options.WithQueryRecordData(req.NeedRecordData),
	}

	options.ApplyQueryFilters(db, filter...)

	if err := db.Count(&resp.Total).Error; err != nil {
		return &models.RecordBestSingleListResp{}, errors.New("查询失败")
	}

	if err := db.Find(&resp.Records).Error; err != nil {
		return &models.RecordBestSingleListResp{}, errors.New("查询失败")
	}

	return resp, nil
}

func (RecordBestSingleService) Update(req *models.RecordBestSingle, reqUid int64) (err error) {

	if reqUid != types.ReqAdminUid {
		// 普通用户
		if req.UserID != 0 && req.UserID != reqUid {
			return errors.New("权限不足")
		}
	}

	if err := database.GetMySQL().Model(&models.RecordBestSingle{}).Updates(req).Error; err != nil {
		return err
	}

	return nil
}
