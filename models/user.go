package models

import (
	"database/sql"
	"log"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/mattn/go-sqlite3"
)

func UserAuthenticate(username string , Password string) bool {
	db , err := sql.Open("sqlite3","./db/database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var PasswordHash string;
	err = db.QueryRow("SELECT password_hash FROM users WHERE username = ?",username).Scan(&PasswordHash)
	fmt.Println(PasswordHash ," - " ,PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows{
			return false
		}
		log.Fatal(err)
	}
	if checkPasswordHash( Password , PasswordHash ) {
		return true
	} else {
		return false
	}
}

func UserExist(username string) bool {
	db , err := sql.Open("sqlite3","./db/database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var ID int
	err = db.QueryRow("SELECT ID FROM users WHERE username = ?",username).Scan(&ID)
	if ID != 0{
		return true 
	}
	return false
}

func UserExistEmail(email string) bool {
	db , err := sql.Open("sqlite3","./db/database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var ID int
	err = db.QueryRow("SELECT ID FROM users WHERE email = ?",email).Scan(&ID)
	if ID != 0{
		return true 
	}
	return false
}
//returns the hashed password of the given string 
func hashPassword(password string) (string, error) {
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
		fmt.Println(hashedBytes)
    return string(hashedBytes), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}

/*
func test_hash(){
	 password := "mySecurePassword"
    hashedPassword, err := hashPassword(password)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Hashed Password:", hashedPassword)

    match := checkPasswordHash(password, hashedPassword)
    fmt.Println("Password Match:", match)
}
*/
