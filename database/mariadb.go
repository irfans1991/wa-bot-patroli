package database


func MariadbConnect(ctx context.Context) *gorm.DB {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Connection string
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Makassar", dbHost, dbUser, dbPassword, dbName, dbPort)
	
	// Open connection to the database
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
		// 	log.Fatalf("Error connecting to the database: %v", err)
		// }
		
		// fmt.Println("Successfully connected to the database")
		
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Makassar", dbHost, dbUser, dbPassword, dbName, dbPort)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Database connection successful")

	// Example of using GORM
	// var users []Users
	// db.Find(&users)
	// // db.Raw("SELECT * FROM users").Scan(&result)

	// for _, user := range users {
	// 	fmt.Println(user)
	// }
	// fmt.Println("Result of 'SELECT 1':", users)

	return db
}
