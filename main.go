// program to get all payroll

package main

import (
	"encoding/json"
	"go_crud_payroll/data"
	"log"
	"net/http"
	"strconv"
)

func main() {
	data.InitDB()
	defer data.CloseDB()

	// get Payroll by empid
	http.HandleFunc("/payroll/", func(w http.ResponseWriter, r *http.Request) {
		// get the empid from the path parameter after /payroll/
		empid, err := strconv.Atoi(r.URL.Path[len("/payroll/"):])
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(400), 400)
			return
		}

		// print empid
		log.Default().Println(empid)

		log.Default().Println("Inside /payroll/empid")

		payroll, err := data.GetPayrollByEmpID(empid)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payroll)
	})

	http.HandleFunc("/payroll", func(w http.ResponseWriter, r *http.Request) {
		payroll, err := data.GetAllPayroll()
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payroll)
	})

	http.ListenAndServe(":8080", nil)
}
