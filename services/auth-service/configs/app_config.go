package config

// AppConfig holds all the configuration for the application
//type AppConfig struct {
//	Port             string // Server port
//	DatabaseDSN      string // Database connection string
//	JWTSecret        string // JWT secret key
//	ResetTokenSecret string // Password reset secret key
//	TokenExpiry      int    // Access token expiry in seconds
//	RefreshExpiry    int    // Refresh token expiry in seconds
//	PasswordResetURL string
//	MigrationPath    string
//}
//
//// configInstance is a singleton instance of AppConfig
//var configInstance *AppConfig
//var once sync.Once
//
//// LoadConfig initializes the application configuration
//func LoadConfig() {
//	once.Do(func() {
//		log.Println("Loading application configuration...")
//		env.LoadEnv() // Load environment variables
//
//		// Construct the absolute path
//		cwd, _ := os.Getwd()
//		absoluteMigrationPath := filepath.Join(cwd, "db", "migrations")
//
//		configInstance = &AppConfig{
//			Port:             env.GetEnv("APP_PORT", "8081"),
//			DatabaseDSN:      env.GetEnv("DATABASE_DSN", "postgres://postgres:admin@localhost:5432/devdojo?sslmode=disable"),
//			JWTSecret:        env.GetEnv("JWT_SECRET", "202ed20f8188b90391022c1df7f789cba1af91fa30b6d86a145edcdd73d65b2e685f519d43152f403480b318e9934e43c9cf5d31a2f45a66bf5159ff88cc416e6349c6af58efc10814aa36780682e5ea9f37d964d5ec64d8a054f9eb519b35a852de9a0874d4279181a35e97c7b31041f313c788f808243c137e9b6739199aa44c46bd9ec786cc2c6faf3fe88744ba7fe1499996f2ceb87aafc6e39b9011b36b01d2cf108f731acf443069a23362d5c5161b350f0c1a0807ccf5727292a20717d6cb787f1a9a0cb793469dd245a728fd5c2c376562932e5b10327559cbbb7511628ed4f4411f6e0dd88827ce4212a93ab78be69adf9ad2e5dd92c38235c1743f"),
//			ResetTokenSecret: env.GetEnv("RESET_SECRET", "auAs0tV7dS"),
//			TokenExpiry:      env.GetEnvAsInt("TOKEN_EXPIRY", 604800),   // Default: 24 hour
//			RefreshExpiry:    env.GetEnvAsInt("REFRESH_EXPIRY", 604800), // Default: 1 week
//			PasswordResetURL: env.GetEnv("PASSWORD_URL", "http://localhost:8081"),
//			MigrationPath:    env.GetEnv("MIGRATION_PATH", absoluteMigrationPath),
//		}
//
//		log.Println("Configuration loaded successfully")
//	})
//}

//type Config struct {
//	Server   ServerConfig   `yaml:"server"`
//	JWT      JWTConfig      `yaml:"jwt"`
//	Database DatabaseConfig `yaml:"database"`
//	//Redis    RedisConfig    `yaml:"redis"`
//}
//
//type ServerConfig struct {
//	Port string `yaml:"port"`
//}
//
//type JWTConfig struct {
//	Secret string `yaml:"secret"`
//	Expiry string `yaml:"expiry"`
//}
//
//type DatabaseConfig struct {
//	Host     string `yaml:"host"`
//	Port     int    `yaml:"port"`
//	User     string `yaml:"user"`
//	Password string `yaml:"password"`
//	DBName   string `yaml:"dbname"`
//}
//
////type RedisConfig struct {
////	Host     string `yaml:"host"`
////	Port     int    `yaml:"port"`
////	Password string `yaml:"password"`
////}
//
//var AppConfig Config
//var configInstance *Config

//func LoadConfig() {
//	// Load the configuration file
//	configPath := "./internal/configs/config.yaml"
//	if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
//		configPath = envPath
//	}
//	file, err := os.Open(configPath)
//	if err != nil {
//		log.Fatalf("Failed to open config file: %v", err)
//	}
//	defer file.Close()
//
//	decoder := yaml.NewDecoder(file)
//	if err := decoder.Decode(&AppConfig); err != nil {
//		log.Fatalf("Failed to decode config file: %v", err)
//	}
//}

// GetConfig returns the singleton instance of AppConfig
//func GetConfig() *AppConfig {
//	if configInstance == nil {
//		LoadConfig()
//	}
//	return configInstance
//}
