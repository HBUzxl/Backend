package config

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

var AppConfig Config

func InitConfig() {
	AppConfig = Config{
		DBHost:     "localhost",
		DBPort:     "3306",
		DBUser:     "root",
		DBPassword: "root",
		DBName:     "online_diagnosis_system",
		JWTSecret:  "your-secret-key", // 在生产环境中应该从配置文件中读取

	}
	InitDB()
}
