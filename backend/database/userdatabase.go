package database

import (
	"ecommerce/models"
	"fmt"
	"log"
)

// CreateUser creates a user for the first time.
// email and password must be provided
func CreateUser(user models.User) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `
					INSERT INTO users (email, password, first_name, last_name)
					VALUES ($1, $2, $3, $4) RETURNING id
					`
	var id int64

	err := db.QueryRow(sqlStatement, user.Email, user.Password, user.FirstName, user.LastName).Scan(&id)
	if err != nil {
		log.Fatalf("unable to create user in database. %v", err)
	}

	fmt.Printf("inserted a single record %v", id)
	return id
}

func GetUserByEmail(email string) (models.User, error) {
	db := createConnection()
	defer db.Close()

	var user models.User

	sqlStatement := `
					SELECT id,
						   created_at,
						   updated_at,
						   email,
						   password,
						   COALESCE(first_name, '') as first_name,
						   COALESCE(last_name, '') as last_name,
						   COALESCE(token, '') as token,
						   COALESCE(refresh_token, '') as refresh_token,
						   COALESCE(csrf, '') as csrf					    
					FROM users
					WHERE email = $1;
					`

	err := db.QueryRow(sqlStatement, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Token,
		&user.RefreshToken,
		&user.CSRF,
	)

	if err != nil {
		fmt.Println("unable get user by email:", err)
		return models.User{}, err
	}

	return user, nil
}

func GetUserById(userId string) (models.User, error) {
	db := createConnection()
	defer db.Close()

	var user models.User

	sqlStatement := `
					SELECT id,
						   created_at,
						   updated_at,
						   email,
						   password,
						   COALESCE(first_name, '') as first_name,
						   COALESCE(last_name, '') as last_name,
						   COALESCE(token, '') as token,
						   COALESCE(refresh_token, '') as refresh_token,
						   COALESCE(csrf, '') as csrf					    
					FROM users
					WHERE id = $1;
					`

	err := db.QueryRow(sqlStatement, userId).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Token,
		&user.RefreshToken,
		&user.CSRF,
	)

	if err != nil {
		fmt.Println("unable get user by id:", err)
		return models.User{}, err
	}

	return user, nil
}

func UpdateUser(user models.User) error {
	db := createConnection()
	defer db.Close()

	sqlStatement := `
					UPDATE users
					first_name = $2
					last_name = $3
					password = $4
					email = $5
					phone = $6
					token = $7
					refresh_token = $8
					csrf = $9
					created_at = $10
					updated_at = $11
					user_cart = $12
					address_details = $13
					order_status = $14
					WHERE id = $1
					RETURNING token;
					`

	_ = db.QueryRow(sqlStatement,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Password,
		user.Email,
		user.Phone,
		user.Token,
		user.RefreshToken,
		user.CSRF,
		user.CreatedAt,
		user.UpdatedAt,
		user.UserCart,
		user.AddressDetails,
		user.OrderStatus,
	)

	fmt.Println("Updated user successfully")
	return nil
}

func StoreUserTokens(id int64, token, refreshToken, csrf string) error {
	db := createConnection()
	defer db.Close()

	sqlStatement := `
					UPDATE users
					SET token = $2,
						refresh_token = $3,
						csrf = $4
					
					WHERE id = $1
					RETURNING id;
					`

	err := db.QueryRow(sqlStatement, id, token, refreshToken, csrf).Scan(&id)

	if err != nil {
		fmt.Println("unable to store user tokens in database:", err)
		return err
	}

	fmt.Println("tokens stored for the user:", id)
	return nil
}
