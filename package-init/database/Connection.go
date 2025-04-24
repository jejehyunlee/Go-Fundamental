package database

var connection = "MYSQL"

func InitDB() string {

	return "Database" + connection + "Connected"
}

func GetDB() string {
	return InitDB()
}
