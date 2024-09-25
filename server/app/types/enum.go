package types

const (
	UserStatusNormal UserStatus = 1 + iota
	UserStatusFrozen
)

const ReqAdminUid int64 = 0

const (
	RecordStatusNormal RecordStatus = 1 + iota
	RecordStatusFrozen
)

const (
	RecordTypePractice = 1 + iota
	RecordTypeRank
	RecordTypeBattle
)
