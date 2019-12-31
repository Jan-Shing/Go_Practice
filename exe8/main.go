package main

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

const(
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "goisgood"
	dbname = "exe8_db"
	table = "phone_numbers"
)

type phone struct {
	id int
	number string
}



func main(){
	psqlInfo := fmt.Sprintf("host=%s port =%d user =%s" +
							" password=%s dbname=%s sslmode=disable",
						    host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	must(err)
	//err = createDB(db, dbname)
	//must(err)
	defer db.Close()

	must(createPhoneNumbersTable(db))
	
	err = seed(db)
	must(err)

	number, err := getphone(db, 1)
	must(err)
	fmt.Println("Number is....", number)

	phones, err := allPhones(db)
	must(err)
	for _, p := range phones{
		fmt.Printf("Working on... %+v\n",p)
		number := normalize(p.number)
		if number != p.number{
			fmt.Println("Updating or removing....", number)
			exsiting, err := findPhone(db, number)
			must(err)
			if exsiting != nil{
				must(deletePhone(db, p.id))
			}else{
				p.number = number
				must(updatePhone(db, p))
			}
		}else{
			fmt.Println("No changes required")
		}
	}

}

func seed(db *sql.DB)error{
	data := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
	for _, number := range data{
		if _, err := insertPhone(db, number); err != nil{
			return err
		}
	}
	return nil
}

func allPhones(db *sql.DB)([]phone, error){
	rows, err := db.Query("SELECT id, value FROM phone_numbers")
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var ret []phone
	for rows.Next(){
		var p phone
		if err := rows.Scan(&p.id, &p.number); err != nil{
			return nil, err
		}
		ret = append(ret, p)
	}
	if err := rows.Err(); err != nil{
		return nil, err
	}
	return ret, nil
}

/* INSERT INTO table_name(column1, column2...)*/
/* VALUES(value1, value2,...)*/
func insertPhone(db *sql.DB, phone string)(int, error){
	//statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil{
		return -1, err
	}
	return id, nil
}

func getphone(db *sql.DB, id int)(string , error){
	var number string
	row := db.QueryRow("SELECT * FROM phone_numbers WHERE id=$1", id)
	err := row.Scan(&id, &number)
	if err != nil{
		return "",err
	}	
	return number, nil
}

func findPhone(db *sql.DB, number string)(*phone, error){
	var p phone
	row := db.QueryRow("SELECT * FROM phone_numbers WHERE value=$1", number)
	err := row.Scan(&p.id, &p.number)
	if err != nil{
		if err == sql.ErrNoRows{
			return nil, nil
		}else{
			return nil, err
		}
	}
	return &p, nil
}

func updatePhone(db *sql.DB, p phone) error{
	statement := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.Exec(statement, p.id, p.number)
	return err
}

func deletePhone(db *sql.DB, id int) error{
	statement := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.Exec(statement, id)
	return err
}

func createPhoneNumbersTable(db *sql.DB)error{
	statement :=` CREATE TABLE IF NOT EXISTS phone_numbers (
		id SERIAL, value VARCHAR(255))`
	_, err := db.Exec(statement)
	return err
}

func createDB(db *sql.DB, name string)error{
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
}

func must(err error){
	if err != nil{
		panic(err)
	}
}

func normalize(phone string) string{
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phone, "")
}


