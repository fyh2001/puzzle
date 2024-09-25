package services

import (
	"errors"
	"github.com/spf13/cast"
	"puzzle/app/models"
	"puzzle/app/services/options"
	"puzzle/app/types"
	"puzzle/database"
)

type RecordService struct{}

func (RecordService) Create(req []*models.Record, reqUid int64) error {

	for _, r := range req {
		if err := CheckReq(r); err != nil {
			return err
		}

		if r.UserID != reqUid {
			return errors.New("权限不足")
		}
	}

	if err := database.GetMySQL().Create(req).Error; err != nil {
		return err
	}

	return nil
}

func (RecordService) List(req *models.RecordReq, reqUid int64) (resp *models.RecordListResp, err error) {

	options.ApplyStringOptions(req,
		options.WithStringFieldStr("ID", req.IDStr),
		options.WithStringFieldStr("UserID", req.UserIDStr),
		options.WithStringFieldStr("IDx", req.IDxStr),
	)

	if len(req.IDsStr) > 0 {
		req.IDs = make([]int64, len(req.IDsStr))

		for i, idStr := range req.IDsStr {
			req.IDs[i] = cast.ToInt64(idStr)
		}
	}

	if reqUid != types.ReqAdminUid {
		// 普通用户
		if req.UserID != 0 && req.UserID != reqUid {
			return &models.RecordListResp{}, errors.New("权限不足")
		}

		req.UserID = reqUid
	}

	db := database.GetMySQL().Model(&models.Record{})

	filters := []options.QueryOption{
		options.WithQueryID64(req.ID),
		options.WithQueryIDs64(req.IDs),
		options.WithQueryUserID(req.UserID),
		options.WithQueryDimension(req.Dimension),
		options.WithQueryType(int(req.Type)),
		options.WithQueryIdx(req.IDx),
		options.WithQueryStatus(int(req.Status)),
		options.WithQueryDurationRange(req.DurationRange),
		options.WithQueryStepRange(req.StepRange),
		options.WithQueryDateRange(req.DateRange),
		options.WithQueryPagination(req.Pagination),
		options.WithQueryOrderBy(req.OrderBy, req.Sorted),
		options.WithQueryUserData(req.NeedUserData),
	}

	options.ApplyQueryFilters(db, filters...)

	if err := db.Count(&resp.Total).Error; err != nil {
		return &models.RecordListResp{}, errors.New("查询失败")
	}

	if err := db.Find(&resp.Records).Error; err != nil {
		return &models.RecordListResp{}, errors.New("查询失败")
	}

	// 校验是否符合权限
	if reqUid != types.ReqAdminUid && len(req.IDs) > 0 {
		for _, v := range resp.Records {
			if cast.ToInt64(v.UserID) != reqUid {
				return &models.RecordListResp{}, errors.New("权限不足")
			}
		}
	}

	return resp, nil
}

func (RecordService) Update(req []*models.Record, reqUid int64) (err error) {

	if reqUid != types.ReqAdminUid {
		// 普通用户
		if len(req) != 0 || (req[0].UserID != 0 && req[0].UserID != reqUid) {
			return errors.New("权限不足")
		}
	}

	if err = database.GetMySQL().Model(&models.Record{}).Updates(req).Error; err != nil {
		return errors.New("更新失败")
	}

	return nil
}
