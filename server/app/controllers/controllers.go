package controllers

var (
	User               = new(UserController)
	Record             = new(RecordController)
	RecordBestSingle   = new(RecordBestSingleController)
	RecordBestAverage  = new(RecordBestAverageController)
	RecordBestStep     = new(RecordBestStepController)
	Scramble           = new(ScrambleController)
	Notification       = new(NotificationController)
	AdminAuthorization = new(AdminAuthorizationController)
	Admin              = new(AdminController)
)
