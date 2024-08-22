package domain

type Config struct {
	DBURI    		string 			`mapstructure:"DB_URI"`
	DbName 			string 			`mapstructure:"DB_NAME"`
	Username    	string 			`mapstructure:"DB_USER"`
	Password  		string        	`mapstructure:"DB_PASSWORD"`
	SecretKey   	string        	`mapstructure:"JWT_SECRET"`
	EmailHost   	string			`mapstructure:"EMAIL_HOST"`
	EmailPort   	string 			`mapstructure:"EMAIL_PORT"`
	SenderEmail 	string 			`mapstructure:"SENDER_EMAIL"`
	SenderPassword	string			`mapstructure:"SENDER_PASSWORD"`
	CloudinaryUrl 	string			`mapstructure:"CLOUDINARY_URL"`
<<<<<<< Updated upstream
	GeminiApiKey	string			`mapstructure:"GEMINI_API_kEY"`
=======
	RedisAdr 		string			`mapstructure:"REDIS_ADDR"`
	RedisPass 		string			`mapstructure:"REDIS_PASS"`
>>>>>>> Stashed changes
}