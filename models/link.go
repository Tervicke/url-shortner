package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)
type link struct{
	ID int;
	ShortUrl string;
	LongUrl string;
}

func GetAllLinks(username string) []link {
	db , err := sql.Open("sqlite3","./db/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var links []link
	rows , err := db.Query("SELECT ID,shortcut,url from links where username = ?",username)
	defer rows.Close()
	for rows.Next(){
		var l link
		err :=rows.Scan(&l.ID,&l.ShortUrl,&l.LongUrl)
		if err != nil {
				log.Fatal(err)
		}
		links = append(links,l)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return links
}
func GetLink(shortcut string,username string) (string,error) {
	db , err := sql.Open("sqlite3","./db/database.db")
	if err != nil {
		return "",err
	}
	defer db.Close()	
	var link string
	err = db.QueryRow("SELECT url FROM links WHERE username = ? AND shortcut = ?",username , shortcut).Scan(&link)
	if err != nil {
		return "",err
	}
	return link , nil
}

func Addurl(shorturl string , longurl string , username string) (int , error ){
	db , err := sql.Open("sqlite3","./db/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	res,err := db.Exec("INSERT INTO links(username,shortcut,url) values (?,?,?)",username,shorturl,longurl)
	if err != nil {
		return 0,err
	}
	id , err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id)  , nil
}
func Deleteurl(ID string) (int , error){
	db , err := sql.Open("sqlite3","./db/database.db")
	if err != nil {
		log.Fatal(err)
		return 0 , err
	}
	defer db.Close()
	res ,err := db.Exec("DELETE FROM links where id =?",ID)
	if err != nil {
		log.Fatal(err)
		return 0 , err
	}
	fmt.Println(res)
	return 1 , nil	
}
