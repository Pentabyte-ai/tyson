package model


func MigrateDB(){
	db.AutoMigrate(
		Bike{},
		User{},
		Provider{},
		Plan{},
	)
}