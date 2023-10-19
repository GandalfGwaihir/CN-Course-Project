package main

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "student.db")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM student_activity_data")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "PRN\tName\tCGPA\tActivity\tDate\tAadhar\tPhone\tEmail")
	fmt.Fprintln(w, "--------------------------------------------------------------")
	for rows.Next() {
		var prn, name, cgpa, activity, date, aadhar, phone, email string
		if err := rows.Scan(&prn, &name, &cgpa, &activity, &date, &aadhar, &phone, &email); err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%-8s\t%-20s\t%-4s\t%-15s\t%-12s\t%-12s\t%-12s\t%-40s\n", prn, name, cgpa, activity, date, aadhar, phone, email)
	}
}

func handleGetRequest2(w http.ResponseWriter, r *http.Request) {
	
	name := r.URL.Query().Get("name")

	db, err := sql.Open("sqlite3", "student.db")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM student_activity_data WHERE name=?", name)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "PRN\tName\tCGPA\tActivity\tDate\tAadhar\tPhone\tEmail")
	fmt.Fprintln(w, "--------------------------------------------------------------")
	for rows.Next() {
		var prn, name, cgpa, activity, date, aadhar, phone, email string
		if err := rows.Scan(&prn, &name, &cgpa, &activity, &date, &aadhar, &phone, &email); err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%-8s\t%-20s\t%-4s\t%-15s\t%-12s\t%-12s\t%-12s\t%-40s\n", prn, name, cgpa, activity, date, aadhar, phone, email)
	}
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "student.db")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err := r.ParseForm(); err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	prn := r.PostFormValue("prn")
	name := r.PostFormValue("name")
	cgpa := r.PostFormValue("cgpa")
	activity := r.PostFormValue("activity")
	date := r.PostFormValue("date")
	aadhar := r.PostFormValue("aadhar")
	phone := r.PostFormValue("phone")
	email := r.PostFormValue("email")

	_, err = db.Exec("INSERT INTO student_activity_data (prn, name, cgpa, activity, date, aadhar, phone, email) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		prn, name, cgpa, activity, date, aadhar, phone, email)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Data created")
}

func handleDeleteRequest(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	db, err := sql.Open("sqlite3", "student.db")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM student_activity_data WHERE name=?", name)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Data deleted")
}

func handlePutRequest(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	activity := r.URL.Query().Get("activity")

	db, err := sql.Open("sqlite3", "student.db")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE student_activity_data SET activity=? WHERE name=?", activity, name)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Data updated")
}

func main() {
	http.HandleFunc("/get", handleGetRequest)
	http.HandleFunc("/get2", handleGetRequest2)
	http.HandleFunc("/post", handlePostRequest)
	http.HandleFunc("/delete", handleDeleteRequest)
	http.HandleFunc("/put", handlePutRequest)

	fmt.Println("Server is listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
