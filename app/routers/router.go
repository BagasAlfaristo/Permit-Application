package routers

func InitRouter(e *echo.Echo, db *gorm.DB, s3 *s3.S3, s3Bucket string) {
	hashService := encrypts.NewHashService()
	helperUserService := helperuser.NewHelperService()
	helperService := helper.NewHelperService(s3, s3Bucket, _adminData.New(db), _userData.New(db, helperUserService))

	userData := _userData.New(db, helperUserService)
	userService := _userService.New(userData, hashService, helperService)
	userHandlerAPI := _userHandler.New(userService, hashService)
}
