package services

var (
	User                   = new(UserImpl)
	Record                 = new(RecordImpl)
	RecordBestSingle       = new(RecordBestSingleImpl)
	RecordBestAverage      = new(RecordBestAverageImpl)
	RecordBestStep         = new(RecordBestStepImpl)
	Scramble               = new(ScrambleImpl)
	ScrambledUserStatus    = new(ScrambledUserStatusImpl)
	Notification           = new(NotificationImpl)
	NotificationUserStatus = new(NotificationUserStatusImpl)
	Cos                    = new(CosImpl)
	AdminAuthorization     = new(AdminAuthorizationImpl)
)
