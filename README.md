# 🌙 DarkTimeDiary

A modern, responsive photo diary web application built with Go backend and vanilla HTML/CSS/JavaScript frontend. DarkTimeDiary allows users to upload, organize, and browse their photo memories with a sleek dark theme interface.

![DarkTimeDiary](https://img.shields.io/badge/Go-1.24.3-blue?logo=go) ![MySQL](https://img.shields.io/badge/MySQL-Database-orange?logo=mysql) ![License](https://img.shields.io/badge/License-MIT-green)

## ✨ Features

- 📷 **Photo Upload**: Upload multiple images with custom titles and descriptions
- 🖼️ **Gallery View**: Browse photos in a responsive grid layout
- 🔍 **Search Functionality**: Search through photos by title or description
- 📱 **Mobile Responsive**: Optimized for both desktop and mobile devices
- 🌙 **Dark Theme**: Beautiful gradient dark theme UI
- 🗃️ **Database Storage**: MySQL backend for reliable data persistence
- ⚡ **Real-time Updates**: Dynamic photo loading with pagination
- 🎨 **Modern UI**: Built with Tailwind CSS for a clean, modern look

## 🛠️ Tech Stack

**Backend:**
- Go 1.24.3
- MySQL Database
- go-sql-driver/mysql

**Frontend:**
- HTML5
- CSS3 (Tailwind CSS)
- Vanilla JavaScript
- Responsive Design

## 🚀 Quick Start

### Prerequisites

Before running DarkTimeDiary, make sure you have:

- [Go](https://golang.org/dl/) (version 1.24.3 or later)
- [MySQL](https://dev.mysql.com/downloads/) server running
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/mdhsaikats/DarkTimeDiary.git
   cd DarkTimeDiary
   ```

2. **Install Go dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up MySQL database**
   ```bash
   # Create database
   mysql -u root -p -e "CREATE DATABASE darktimediary;"
   
   # Import the database schema
   mysql -u root -p darktimediary < darktimediary.sql
   ```

4. **Configure database connection**
   
   Edit the database connection string in `main.go` (line 28):
   ```go
   dsn := "root:YOUR_PASSWORD@tcp(127.0.0.1:3306)/darktimediary?parseTime=true&charset=utf8mb4"
   ```
   Replace `YOUR_PASSWORD` with your MySQL root password.

5. **Run the application**
   ```bash
   go run main.go
   ```

6. **Access the application**
   
   Open your browser and navigate to: `http://localhost:8080`

## 📁 Project Structure

```
DarkTimeDiary/
├── main.go              # Go backend server
├── index.html           # Frontend HTML
├── style.css           # Custom styles
├── darktimediary.sql   # Database schema and sample data
├── go.mod              # Go module dependencies
├── go.sum              # Go module checksums
├── README.md           # Project documentation
├── asset/              # Static assets
│   └── instagram.png   # App logo
└── uploads/            # User uploaded images
    ├── *.jpg
    ├── *.png
    └── ...
```

## 🔧 Configuration

### Database Configuration

The application uses MySQL as its database. You can modify the database connection in `main.go`:

```go
// Update these credentials to match your MySQL setup
dsn := "username:password@tcp(host:port)/database_name?parseTime=true&charset=utf8mb4"
```

### Server Configuration

By default, the server runs on port 8080. You can modify this in the `main()` function:

```go
log.Fatal(http.ListenAndServe(":8080", nil))
```

## 📡 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Serve the main application |
| GET | `/api/photos?page=1&limit=12` | Get paginated photos |
| POST | `/api/upload` | Upload new photos |
| GET | `/uploads/*` | Serve uploaded images |

### Upload API Example

```bash
curl -X POST http://localhost:8080/api/upload \
  -F "images=@photo1.jpg" \
  -F "images=@photo2.jpg" \
  -F "titles=My Photo 1" \
  -F "titles=My Photo 2" \
  -F "descriptions=Description 1" \
  -F "descriptions=Description 2"
```

## 🎨 Features Overview

### Photo Upload
- Support for multiple image formats (JPEG, PNG, etc.)
- Custom titles and descriptions for each photo
- Automatic file naming with timestamps
- Real-time upload progress

### Gallery
- Responsive grid layout
- Lazy loading for better performance
- Pagination support
- Modal view for full-size images

### Search
- Real-time search functionality
- Search by photo title or description
- Instant results filtering

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🐛 Known Issues

- Large image files may take longer to upload
- Search is case-sensitive

## 🔮 Future Enhancements

- [ ] Image compression and optimization
- [ ] User authentication and authorization
- [ ] Photo categories and tags
- [ ] Export functionality
- [ ] Image editing capabilities
- [ ] Social sharing features

## 📞 Support

If you have any questions or need help, please:

1. Check the [Issues](https://github.com/mdhsaikats/DarkTimeDiary/issues) page
2. Create a new issue if your problem isn't already reported
3. Contact the maintainer: [mdhsaikats](https://github.com/mdhsaikats)

## 🙏 Acknowledgments

- Built with ❤️ using Go and modern web technologies
- UI components styled with Tailwind CSS
- Icons and assets from various open-source projects

---

**Happy photo journaling! 📸✨**
