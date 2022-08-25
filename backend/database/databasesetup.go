package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func createConnection() *sql.DB {
	// Comment below to use env variables coming from a docker-compose.yml
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	var dataSourceName string

	if password == "" {
		dataSourceName = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=%s",
			os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USERNAME"),
			os.Getenv("POSTGRES_DBNAME"), os.Getenv("POSTGRES_SSL"))
	} else {
		dataSourceName = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USERNAME"),
			password, os.Getenv("POSTGRES_DBNAME"), os.Getenv("POSTGRES_SSL"))
	}

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.(any))
	}

	fmt.Println("Successfully connected to postgres")

	return db
}

/*func DBSet() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to connect to mongodb")
		return nil
	}

	log.Println("successfully connected to mongodb")
	return client
}

func UserData(client *mongo.Client, collectionName string) (userCollection *mongo.Collection) {
	userCollection = client.Database("Ecommerce").Collection(collectionName)
	return userCollection
}

func ProductData(client *mongo.Client, collectionName string) (productCollection *mongo.Collection) {
	productCollection = client.Database("Ecommerce").Collection(collectionName)
	return productCollection
}
*/
