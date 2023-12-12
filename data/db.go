package data

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Connect to the database to mysql
func InitDB() {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	// print the value of the all environment variables
	log.Println(user)
	log.Println(password)
	log.Println(hostname)
	log.Println(port)
	log.Println(dbname)

	var err error
	db, err = sql.Open("mysql", user+":"+password+"@tcp("+hostname+":"+port+")/"+dbname)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	// call get all payroll and print the value
	GetAllPayroll()
	log.Println("Connected to payroll database..updated")
}

// get All Payroll
func GetAllPayroll() ([]Payroll, error) {
	rows, err := db.Query("SELECT * FROM Payroll")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payroll := make([]Payroll, 0)
	for rows.Next() {
		p := Payroll{}
		//(empid, salary, pay_date, taxes, bonuses, net_pay)
		err := rows.Scan(&p.PayrollID, &p.EmployeeID, &p.Salary, &p.PayDate, &p.Taxes, &p.Bonuses, &p.NetPay)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		payroll = append(payroll, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// print the value of the all payroll
	log.Println(payroll)
	return payroll, nil
}

// GetPayroll get a single payroll given an employee id
func GetPayrollByEmpID(id int) (*Payroll, error) {
	p := &Payroll{}
	err := db.QueryRow("SELECT * FROM Payroll WHERE empid=?", id).Scan(&p.PayrollID, &p.EmployeeID, &p.Salary, &p.PayDate, &p.Taxes, &p.Bonuses, &p.NetPay)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// CloseDB close the database

func CloseDB() {
	db.Close()
}
