package config

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var AppConfig Config

func InitConfig() {
	AppConfig = Config{
		DBHost:     "localhost",
		DBPort:     "3306",
		DBUser:     "root",
		DBPassword: "root",
		DBName:     "online_diagnosis_system",
	}
	InitDB()
}
