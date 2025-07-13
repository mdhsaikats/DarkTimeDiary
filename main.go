package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Photo struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

var db *sql.DB

func initDB() {
	var err error
	// Update these credentials to match your MySQL setup
	dsn := "root:29112003@tcp(127.0.0.1:3306)/darktimediary?parseTime=true&charset=utf8mb4"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	os.Mkdir("./uploads", 0755)
	r.ParseMultipartForm(32 << 20)

	files := r.MultipartForm.File["images"]
	titles := r.MultipartForm.Value["titles"]
	var uploaded []string

	for i, fh := range files {
		file, _ := fh.Open()
		defer file.Close()

		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), fh.Filename)
		dst, _ := os.Create(filepath.Join("./uploads", filename))
		defer dst.Close()

		io.Copy(dst, file)

		// Get title for this image
		title := fh.Filename // Default to filename
		if i < len(titles) && titles[i] != "" {
			title = titles[i]
		}

		db.Exec("INSERT INTO photos (url, title) VALUES (?, ?)", "/uploads/"+filename, title)
		uploaded = append(uploaded, filename)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "files": uploaded})
}

func photosHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 12
	}
	offset := (page - 1) * limit

	rows, err := db.Query("SELECT url, title FROM photos ORDER BY id DESC LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var photos []Photo
	for rows.Next() {
		var p Photo
		if err := rows.Scan(&p.URL, &p.Title); err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		photos = append(photos, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(photos)
}

// Modal functions should be implemented in your frontend JavaScript, not in this Go file.

func main() {
	initDB()
	defer db.Close()

	// Serve static files
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// Serve uploaded images
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))

	// API endpoints
	http.HandleFunc("/api/photos", photosHandler)
	http.HandleFunc("/api/upload", uploadHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
