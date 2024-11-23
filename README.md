# Final Project Sanbercode Golang Batch 61
## Cinema Booking API
Cinema Booking API adalah sebuah aplikasi yang dirancang untuk memanage data di dalam ekosistem bioskop. API ini memberikan akses untuk membuat daftar cinema hall (studio), daftar film, daftar jadwal film , dan transaksi booking.

Dengan fitur-fitur yang telah disediakan ini, agar membantu dalam memanajemen data yang ada di dalam transaksi bioskop.

## Kegunaan
Aplikasi ini memungkinkan pengguna untuk melakukan operasi berikut:

* CRUD untuk Customers: Menambah, mengubah, melihat, dan menghapus daftar customer.
* CRUD untuk Cinema Halls: Menambah, mengubah, melihat, dan menghapus daftar cinema hall(studio).
* CRUD untuk Movies: Menambah, mengubah, melihat, dan menghapus daftar movie(film).
* CRUD untuk Showtimes: Menambah, mengubah, melihat, dan menghapus daftar showtime(jadwal film).
* Transaksi Bookings dari jadwal film yang ada.
* JWT Authentication: Menggunakan JWT Auth untuk memastikan hanya pengguna yang terautentikasi yang dapat mengakses API.
## Persyaratan
* Go: Pastikan Go terpasang di sistem Anda.
* PostgreSQL: Database yang digunakan adalah PostgreSQL. Anda harus menyiapkan koneksi database terlebih dahulu.
* GORM : ORM untuk transaksi repository ke database
* JWT Auth : Proses autentikasi token untuk keperluan akses API
* Gin Framework : Framweork GIN untuk handlers API
## Instalasi
### Clone Repository

Clone repositori ini ke mesin lokal Anda:

* git clone https://github.com/robbyars/cinema.git
* cd cinema

### Install Dependencies

Install dependencies menggunakan go get:
* https://github.com/gin-gonic/gin
* https://github.com/lib/pq
* https://github.com/rubenv/sql-migrate
* gorm.io/gorm
* gorm.io/driver/postgres
* github.com/dgrijalva/jwt-go
* github.com/joho/godotenv


### Konfigurasi Database

* Buat database di PostgreSQL dan sesuaikan konfigurasi di file configs/config.json untuk menyesuaikan kredensial dan URL koneksi Anda.

## Menjalankan Aplikasi

Jalankan server menggunakan perintah berikut:

* go run main.go
Server akan berjalan di port 8080. (jika router.Run(":8080") aktif)

Cara Menggunakan
API ini menggunakan Basic Authentication untuk mengakses endpoint yang membutuhkan otentikasi. Anda dapat mengirimkan header Authorization dengan format berikut:

* Authorization: Bearer <token>
* Gunakan Postman untuk menguji endpoint API.

### List Path API yang Tersedia
#### Customer (customers)
* POST /signup
 
Singup user awal untuk keperluan login

Contoh request:

```json
{
  "username": "user1",
  "password": "user12345",
  "re_type_password":"user12345",
  "fullname": "User Coba",
  "email": "user@gmail.com",
  "phone":"081234232122"
}
```
* POST /login

Proses login untuk akses API

Contoh Request:

```json
{
  "username": "user1",
  "password": "user12345"
}
```
* GET /api/customers

Menampilkan seluruh data customers yang sudah signup.

Contoh Response:

```json
{
    "success": true,
    "message": "successfully get all customer data",
    "data": [
        {
            "id": 1,
            "username": "admin",
            "password": "$2a$10$45ABRNnVTJ.DdV5UC/pKWeUgFI1pDixo0DVudQykv9npx3ut/svsq",
            "fullname": "Admin Coba",
            "email": "admin@gmail.com",
            "phone": "081234232122",
            "CreatedAt": "2024-11-23T10:40:29.685827Z",
            "ModifiedAt": "2024-11-23T10:40:29.685827Z"
        },
        {
            "id": 2,
            "username": "user1",
            "password": "$2a$10$XtAhLEdefpdSyQuqZZbseeCTyZyT5yl7q9WdYPfOKnh3yAJ.mBxsi",
            "fullname": "User Coba",
            "email": "user@gmail.com",
            "phone": "081234232122",
            "CreatedAt": "2024-11-23T11:02:33.131793Z",
            "ModifiedAt": "2024-11-23T11:02:33.131793Z"
        }
    ],
    "total_data": 2,
    "trace_id": ""
}
```
* PUT /api/customers/:username

Mengubah data customers berdasarkan username.

Contoh Request:

```json
{
    "fullname": "Admin",
    "email":"admin1@gmail.com",
    "phone":"081234232122"
}
```
* DELETE /api/customers/:id

Menghapus customers berdasarkan ID.

#### Cinema Hall (cinema_halls)
* POST /api/cinema_halls

Menambahkan cinema hall (studio).

Contoh Request:

```json
{
  "Name": "Cinema Hall A4",
  "Capacity": 150,
  "Location": "Bali"
}
```
* GET /api/cinema_halls

Menampilkan semua data cinema hall (studio).

Contoh Response:

```json
{
    "success": true,
    "message": "successfully get all cinema_hall data",
    "data": [
        {
            "id": 1,
            "name": "Cinema Hall A1",
            "location": "Bali",
            "capacity": 150,
            "CreatedAt": "2024-11-23T10:46:55.133798Z",
            "ModifiedAt": "2024-11-23T10:46:55.133798Z",
            "showtimes": [
                {
                    "id": 1,
                    "cinema_hall_id": 1,
                    "movie_id": 1,
                    "showtime_date": "2024-11-23T14:00:00Z",
                    "price": 45000,
                    "CreatedAt": "2024-11-23T10:57:07.146969Z",
                    "ModifiedAt": "2024-11-23T10:57:07.146969Z"
                },
                {
                    "id": 2,
                    "cinema_hall_id": 1,
                    "movie_id": 2,
                    "showtime_date": "2024-11-23T18:00:00Z",
                    "price": 45000,
                    "CreatedAt": "2024-11-23T10:57:19.937152Z",
                    "ModifiedAt": "2024-11-23T10:57:19.937152Z"
                },
                {
                    "id": 3,
                    "cinema_hall_id": 1,
                    "movie_id": 3,
                    "showtime_date": "2024-11-23T21:00:00Z",
                    "price": 45000,
                    "CreatedAt": "2024-11-23T10:57:31.346438Z",
                    "ModifiedAt": "2024-11-23T10:57:31.346438Z"
                }
            ]
        },
        {
            "id": 2,
            "name": "Cinema Hall A2",
            "location": "Bali",
            "capacity": 150,
            "CreatedAt": "2024-11-23T10:47:01.788973Z",
            "ModifiedAt": "2024-11-23T10:47:01.788973Z",
            "showtimes": []
        },
        {
            "id": 3,
            "name": "Cinema Hall A3",
            "location": "Bali",
            "capacity": 150,
            "CreatedAt": "2024-11-23T10:47:07.535107Z",
            "ModifiedAt": "2024-11-23T10:47:07.535107Z",
            "showtimes": []
        },
        {
            "id": 4,
            "name": "Cinema Hall A4",
            "location": "Bali",
            "capacity": 150,
            "CreatedAt": "2024-11-23T10:47:12.438158Z",
            "ModifiedAt": "2024-11-23T10:47:12.438158Z",
            "showtimes": []
        }
    ],
    "total_data": 4,
    "trace_id": ""
}
```
* GET /api/cinema_halls/:id

Menampilkan data cinema hall berdasarkan ID.

Contoh Response:

```json
{
    "success": true,
    "message": "successfully get cinema_hall data",
    "data": {
        "id": 1,
        "name": "Cinema Hall A1",
        "location": "Bali",
        "capacity": 150,
        "CreatedAt": "2024-11-23T10:46:55.133798Z",
        "ModifiedAt": "2024-11-23T10:46:55.133798Z",
        "showtimes": [
            {
                "id": 1,
                "cinema_hall_id": 1,
                "movie_id": 1,
                "showtime_date": "2024-11-23T14:00:00Z",
                "price": 45000,
                "CreatedAt": "2024-11-23T10:57:07.146969Z",
                "ModifiedAt": "2024-11-23T10:57:07.146969Z"
            },
            {
                "id": 2,
                "cinema_hall_id": 1,
                "movie_id": 2,
                "showtime_date": "2024-11-23T18:00:00Z",
                "price": 45000,
                "CreatedAt": "2024-11-23T10:57:19.937152Z",
                "ModifiedAt": "2024-11-23T10:57:19.937152Z"
            },
            {
                "id": 3,
                "cinema_hall_id": 1,
                "movie_id": 3,
                "showtime_date": "2024-11-23T21:00:00Z",
                "price": 45000,
                "CreatedAt": "2024-11-23T10:57:31.346438Z",
                "ModifiedAt": "2024-11-23T10:57:31.346438Z"
            }
        ]
    },
    "total_data": 0,
    "trace_id": ""
}
```

* PUT /api/cinema_halls/:id

Mengupdate data cinema hall berdasarkan ID.

Contoh Request:

```json
{
  "Name": "Cinema Hall A2",
  "Capacity": 150,
  "Location":"Bali"
}
```

* DELETE /api/cinema_halls/:id

Menghapus data cinema hall berdasarkan ID.

#### Movie (movies)
* POST /api/movies

Menambahkan movies (film).

Contoh Request:

```json
{
  "Title": "Batman VS Superman",
  "Genre":"Superhero",
  "Duration": "152 minutes",
  "Rating": "PG-13",
  "Release_date": "2024-07-18T00:00:00Z",
  "Description": "Super hero bertarung disini."
}
```
* GET /api/movies

Menampilkan semua data movies (film).

Contoh Response:

```json
{
    "success": true,
    "message": "successfully get all movie data",
    "data": [
        {
            "id": 1,
            "title": "Sangkuriang",
            "genre": "History",
            "duration": "152 minutes",
            "rating": "PG-13",
            "release_date": "2008-07-18T00:00:00Z",
            "description": "Film sejarah indonesia.",
            "CreatedAt": "2024-11-23T10:49:29.464769Z",
            "ModifiedAt": "2024-11-23T10:49:29.464769Z",
            "showtimes": [
                {
                    "id": 1,
                    "cinema_hall_id": 1,
                    "movie_id": 1,
                    "showtime_date": "2024-11-23T14:00:00Z",
                    "price": 45000,
                    "CreatedAt": "2024-11-23T10:57:07.146969Z",
                    "ModifiedAt": "2024-11-23T10:57:07.146969Z"
                }
            ]
        },
        {
            "id": 2,
            "title": "Spiderman",
            "genre": "Superhero",
            "duration": "152 minutes",
            "rating": "PG-13",
            "release_date": "2008-07-18T00:00:00Z",
            "description": "Super hero laba laba dari amerika.",
            "CreatedAt": "2024-11-23T10:49:50.923661Z",
            "ModifiedAt": "2024-11-23T10:49:50.923661Z",
            "showtimes": [
                {
                    "id": 2,
                    "cinema_hall_id": 1,
                    "movie_id": 2,
                    "showtime_date": "2024-11-23T18:00:00Z",
                    "price": 45000,
                    "CreatedAt": "2024-11-23T10:57:19.937152Z",
                    "ModifiedAt": "2024-11-23T10:57:19.937152Z"
                }
            ]
        },
        {
            "id": 3,
            "title": "Batman VS Superman",
            "genre": "Superhero",
            "duration": "152 minutes",
            "rating": "PG-13",
            "release_date": "2024-07-18T00:00:00Z",
            "description": "Super hero bertarung disini.",
            "CreatedAt": "2024-11-23T10:53:34.907159Z",
            "ModifiedAt": "2024-11-23T10:53:34.907159Z",
            "showtimes": [
                {
                    "id": 3,
                    "cinema_hall_id": 1,
                    "movie_id": 3,
                    "showtime_date": "2024-11-23T21:00:00Z",
                    "price": 45000,
                    "CreatedAt": "2024-11-23T10:57:31.346438Z",
                    "ModifiedAt": "2024-11-23T10:57:31.346438Z"
                }
            ]
        }
    ],
    "total_data": 3,
    "trace_id": ""
}
```
* GET /api/movies/:title

Menampilkan data movies berdasarkan title film.

Contoh Response:

```json
{
    "success": true,
    "message": "successfully get cinema_hall data",
    "data": {
        "id": 1,
        "name": "Cinema Hall A1",
        "location": "Bali",
        "capacity": 150,
        "CreatedAt": "2024-11-23T10:46:55.133798Z",
        "ModifiedAt": "2024-11-23T10:46:55.133798Z",
        "showtimes": [
            {
                "id": 1,
                "cinema_hall_id": 1,
                "movie_id": 1,
                "showtime_date": "2024-11-23T14:00:00Z",
                "price": 45000,
                "CreatedAt": "2024-11-23T10:57:07.146969Z",
                "ModifiedAt": "2024-11-23T10:57:07.146969Z"
            },
            {
                "id": 2,
                "cinema_hall_id": 1,
                "movie_id": 2,
                "showtime_date": "2024-11-23T18:00:00Z",
                "price": 45000,
                "CreatedAt": "2024-11-23T10:57:19.937152Z",
                "ModifiedAt": "2024-11-23T10:57:19.937152Z"
            },
            {
                "id": 3,
                "cinema_hall_id": 1,
                "movie_id": 3,
                "showtime_date": "2024-11-23T21:00:00Z",
                "price": 45000,
                "CreatedAt": "2024-11-23T10:57:31.346438Z",
                "ModifiedAt": "2024-11-23T10:57:31.346438Z"
            }
        ]
    },
    "total_data": 0,
    "trace_id": ""
}
```

* PUT /api/movies/:id

Mengupdate data movies berdasarkan ID.

Contoh Request:

```json
{
            "title": "The Dark Knight",
            "genre": "Action, Crime, Drama",
            "duration": "152 minutes",
            "rating": "PG-18",
            "release_date": "2008-07-18T00:00:00Z",
            "description": "Batman faces off against the Joker, a criminal mastermind who seeks to create chaos in Gotham City."
}
```

* DELETE /api/movie/:id

Menghapus data movie berdasarkan ID.

#### Show Time (showtimes)
* POST /api/showtimes

Menambahkan showtime (jadwal film).

Contoh Request:

```json
{
  "movie_id": 3,
  "cinema_hall_id": 1,
  "showtime_date": "2024-11-23T21:00:00Z",
  "price":45000
}
```
* GET /api/showtimes

Menampilkan semua data showtime (jadwal tayang).

Contoh Response:

```json
{
    "success": true,
    "message": "successfully get all showtime data",
    "data": [
        {
            "id": 1,
            "cinema_hall_id": 1,
            "movie_id": 1,
            "showtime_date": "2024-11-23T14:00:00Z",
            "price": 45000,
            "CreatedAt": "2024-11-23T10:57:07.146969Z",
            "ModifiedAt": "2024-11-23T10:57:07.146969Z"
        },
        {
            "id": 2,
            "cinema_hall_id": 1,
            "movie_id": 2,
            "showtime_date": "2024-11-23T18:00:00Z",
            "price": 45000,
            "CreatedAt": "2024-11-23T10:57:19.937152Z",
            "ModifiedAt": "2024-11-23T10:57:19.937152Z"
        },
        {
            "id": 3,
            "cinema_hall_id": 1,
            "movie_id": 3,
            "showtime_date": "2024-11-23T21:00:00Z",
            "price": 45000,
            "CreatedAt": "2024-11-23T10:57:31.346438Z",
            "ModifiedAt": "2024-11-23T10:57:31.346438Z"
        }
    ],
    "total_data": 3,
    "trace_id": ""
}
```
* GET /api/showtimes/:cinema_id

Menampilkan data showtime berdasarkan cinema id.

Contoh Response:

```json
{
    "success": true,
    "message": "successfully get showtime data",
    "data": [
        {
            "id": 1,
            "cinema_hall_id": 1,
            "movie_id": 1,
            "showtime_date": "2024-11-23T14:00:00Z",
            "price": 45000,
            "CreatedAt": "2024-11-23T10:57:07.146969Z",
            "ModifiedAt": "2024-11-23T10:57:07.146969Z"
        },
        {
            "id": 2,
            "cinema_hall_id": 1,
            "movie_id": 2,
            "showtime_date": "2024-11-23T18:00:00Z",
            "price": 45000,
            "CreatedAt": "2024-11-23T10:57:19.937152Z",
            "ModifiedAt": "2024-11-23T10:57:19.937152Z"
        },
        {
            "id": 3,
            "cinema_hall_id": 1,
            "movie_id": 3,
            "showtime_date": "2024-11-23T21:00:00Z",
            "price": 45000,
            "CreatedAt": "2024-11-23T10:57:31.346438Z",
            "ModifiedAt": "2024-11-23T10:57:31.346438Z"
        }
    ],
    "total_data": 3,
    "trace_id": ""
}
```

* PUT /api/showtimes/:id

Mengupdate data showtime berdasarkan ID.

Contoh Request:

```json
{
        "cinema_hall_id": 1,
        "movie_id": 1,
        "showtime_date": "2024-11-22T10:00:00Z",
        "price": 45000
}
```

* DELETE /api/showtimes/:id

Menghapus data showtime berdasarkan ID.

#### Booking (bookings)
* POST /api/bookings

Menambahkan booking (tiket). Untuk Request hanya showtime_id untuk pemilihan jadwal tayang dan seat number yang di dapat dari front end. Untuk CustomerID didapat dari context header userid yang login dan melakukan bookings.

Contoh Request:

```json
{
  "showtime_id": 2,
  "seat_number":1
}
```
* GET /api/bookings

Menampilkan semua data bookings (tiket). Untuk parameter userID didapatkan dari user yang login didapat dari context header userid yang login dan melakukan bookings.

Contoh Response:

```json
{
    "success": true,
    "message": "successfully get booking data",
    "data": [
        {
            "id": 1,
            "customer_id": 2,
            "showtime_id": 2,
            "booking_date": "2024-11-23T11:04:08.762895Z",
            "seat_number": 1,
            "status": "Paid",
            "CreatedAt": "2024-11-23T11:04:08.771833Z",
            "ModifiedAt": "2024-11-23T11:06:23.408862Z",
            "customer": {
                "id": 2,
                "fullname": "User Coba",
                "email": "user@gmail.com",
                "phone": "081234232122"
            },
            "showtime": {
                "id": 2,
                "cinema_hall_id": 1,
                "movie_id": 2,
                "showtime_date": "2024-11-23T18:00:00Z",
                "price": 45000,
                "CreatedAt": "2024-11-23T10:57:19.937152Z",
                "ModifiedAt": "2024-11-23T10:57:19.937152Z"
            }
        }
    ],
    "total_data": 1,
    "trace_id": ""
}
```

* PUT /api/bookings

Mengupdate data booking status menjadi "Paid". Request data hanya mengirimkan showtime_id dan status untuk user booking yang status nya "Not Paid".

Contoh Request:
```json
{   
    "showtime_id":2,
    "status":"Paid"
}
```


### Autentikasi
API ini menggunakan JWT Authentication untuk memastikan bahwa hanya pengguna yang terautentikasi yang dapat mengakses endpoint tertentu. Autentikasi ini dilakukan dengan mengirimkan username dan password dalam header 

Authorization menggunakan format:
```bash
Authorization: Bearer <token>
```

Error Handling
API ini mengembalikan kode status HTTP berikut dalam berbagai situasi:

* 200 OK: Sesuai dengan error message berhasil dari setiap fungsi.
* 400 Bad Request: Input yang diberikan tidak valid.
* 401 Unauthorized: Pengguna tidak terautentikasi atau kredensial tidak valid.
* 404 Not Found: Data yang diminta tidak ditemukan.
* 500 Internal Server Error: Terjadi kesalahan pada server.

## Contributing
Jika Anda ingin berkontribusi pada proyek ini, silakan fork repositori ini, buat perubahan, dan kirimkan pull request.