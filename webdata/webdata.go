package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func search(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("search.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		var input = r.Form["input"]
		fmt.Println("input:", input)
		getData(w, input[0])

	}
}

func main() {
	http.HandleFunc("/search", search)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getData(w http.ResponseWriter, input string) {

	db, err := sql.Open("mysql", "login:password@tcp(8.8.8.8:3306)/search?charset=utf8") // CHANGEME
	checkErr(err)

	rows, err := db.Query("SELECT * FROM mytable WHERE RONumber=?", input)
	checkErr(err)

	fmt.Fprintf(w, "<html><head><title>HTML Tables</title></head>")

	fmt.Fprintf(w, "<form action='/search' method='post'>")
	fmt.Fprintf(w, "Search:<input type='text' name='input'>")
	fmt.Fprintf(w, "<input type='submit' value='Submit'></form><br>")

	fmt.Fprintf(w, "<br>Input: "+input)
	fmt.Fprintf(w, "<br><br>")

	fmt.Fprintf(w, "<body><table border=%q1%q><tr>")
	fmt.Fprintf(w, "<tr>")
	fmt.Fprintf(w, "<td>Part Number</td>")
	fmt.Fprintf(w, "<td>RONumber</td>")
	fmt.Fprintf(w, "<td>PartRequestNumber</td>")
	fmt.Fprintf(w, "<td>FulfillLocation</td>")
	fmt.Fprintf(w, "<td>RequestLocation</td>")
	fmt.Fprintf(w, "<td>TrackingNumber</td>")
	fmt.Fprintf(w, "</tr>")

	for rows.Next() {
		var PartNumber string
		var RONumber string
		var PartRequestNumber string
		var FulfillLocation string
		var RequestLocation string
		var TrackingNumber string
		err = rows.Scan(&PartRequestNumber, &PartNumber, &FulfillLocation, &RequestLocation, &TrackingNumber, &RONumber)
		checkErr(err)
		fmt.Fprintf(w, "<tr>")
		fmt.Fprintf(w, "<td>"+PartNumber+"</td>")
		fmt.Fprintf(w, "<td>"+RONumber+"</td>")
		fmt.Fprintf(w, "<td>"+PartRequestNumber+"</td>")
		fmt.Fprintf(w, "<td>"+FulfillLocation+"</td>")
		fmt.Fprintf(w, "<td>"+RequestLocation+"</td>")
		fmt.Fprintf(w, "<td>"+TrackingNumber+"</td>")
		fmt.Fprintf(w, "</tr>")
	}

	fmt.Fprintf(w, "</tr></table></body></html>")

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
