package handlers

import (
	"errors"
	"puzzle/database"
)

type RankUpdate struct {
	Dimension int
	Type      int
}

// UpdateRecordBestSingleRank 更新记录最佳单次排名
func UpdateRecordBestSingleRank(rankUpdate any) error {

	// 将rankUpdate转换为RankUpdate类型
	rankUpdateData := rankUpdate.(RankUpdate)

	db := database.GetMySQL()

	// 开启事务
	tx := db.Begin()

	// 创建临时表
	err := tx.Exec("CREATE TEMPORARY TABLE temp_rank SELECT id FROM record_best_single WHERE dimension = ? ORDER BY record_duration", rankUpdateData.Dimension).Error
	if err != nil {
		tx.Rollback() // 回滚事务
		return errors.New("创建临时表失败")
	}

	// 设置变量
	err = tx.Exec("SET @ranked = 0").Error
	if err != nil {
		tx.Rollback() // 回滚事务
		return errors.New("设置变量失败")
	}

	// 更新记录
	err = tx.Exec("UPDATE record_best_single AS r JOIN temp_rank AS tr ON r.id = tr.id SET r.ranked = (@ranked := @ranked + 1)").Error
	if err != nil {
		tx.Rollback() // 回滚事务
		return errors.New("更新记录失败")
	}

	// 删除临时表
	err = tx.Exec("DROP TEMPORARY TABLE IF EXISTS temp_rank").Error
	if err != nil {
		tx.Rollback() // 回滚事务
		return errors.New("删除临时表失败")
	}

	// 提交事务
	tx.Commit()

	return nil
}

// UpdateRecordBestAverageRank 更新记录最佳平均排名
func UpdateRecordBestAverageRank(rankUpdate any) error {
	// 将rankUpdate转换为RankUpdate类型
	rankUpdateData := rankUpdate.(RankUpdate)

	db := database.GetMySQL()

	// 开启事务
	tx := db.Begin()

	// 创建临时表
	err := tx.Exec("CREATE TEMPORARY TABLE temp_rank SELECT id FROM record_best_average WHERE dimension = ? AND type = ? ORDER BY record_average_duration", rankUpdateData.Dimension, rankUpdateData.Type).Error
	if err != nil {
		tx.Rollback() // 回滚事务
		return errors.New("创建临时表失败")
	}

	// 设置变量
	err = tx.Exec("SET @ranked = 0").Error
	if err != nil {
		tx.Rollback() // 回滚事务
		return errors.New("设置变量失败")
	}

	// 更新记录
	err = tx.Exec("UPDATE record_best_average AS r JOIN temp_rank AS tr ON r.id = tr.id SET r.ranked = (@ranked := @ranked + 1)").Error
	if err != nil {
		tx.Rollback() // 回滚事务
		return errors.New("更新记录失败")
	}

	// 删除临时表
	err = tx.Exec("DROP TEMPORARY TABLE IF EXISTS temp_rank").Error
	if err != nil {
		tx.Rollback() // 回滚事务
		return errors.New("删除临时表失败")
	}

	// 提交事务
	tx.Commit()

	return nil
}

// UpdateRecordBestSingleRank 更新记录最佳单次排名
func UpdateRecordBestStepRank(rankUpdate any) error {

	// 将rankUpdate转换为RankUpdate类型
	rankUpdateData := rankUpdate.(RankUpdate)

	db := database.GetMySQL()

	// 开启事务
	tx := db.Begin()

	// 创建临时表
	err := tx.Exec("CREATE TEMPORARY TABLE temp_rank SELECT id FROM record_best_step WHERE dimension = ? ORDER BY record_step", rankUpdateData.Dimension).Error
	if err != nil {
		tx.Rollback() // 回滚事务
		return errors.New("创建临时表失败")
	}

	// 设置变量
	err = tx.Exec("SET @ranked = 0").Error
	if err != nil {
		tx.Rollback() // 回滚事务
		return errors.New("设置变量失败")
	}

	// 更新记录
	err = tx.Exec("UPDATE record_best_step AS r JOIN temp_rank AS tr ON r.id = tr.id SET r.ranked = (@ranked := @ranked + 1)").Error
	if err != nil {
		tx.Rollback() // 回滚事务
		return errors.New("更新记录失败")
	}

	// 删除临时表
	err = tx.Exec("DROP TEMPORARY TABLE IF EXISTS temp_rank").Error
	if err != nil {
		tx.Rollback() // 回滚事务
		return errors.New("删除临时表失败")
	}

	// 提交事务
	tx.Commit()

	return nil
}
