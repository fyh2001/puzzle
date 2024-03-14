package controllers

var (
	User                   = new(UserController)
	Record                 = new(RecordController)
	RecordBestSingle       = new(RecordBestSingleController)
	RecordBestAverage      = new(RecordBestAverageController)
	RecordBestStep         = new(RecordBestStepController)
	Scramble               = new(ScrambleController)
	Notification           = new(NotificationController)
	NotificationUserStatus = new(NotificationUserStatusController)
	AdminAuthorization     = new(AdminAuthorizationController)
	Admin                  = new(AdminController)
)
