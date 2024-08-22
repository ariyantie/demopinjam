
# Kredit System

Proyek kredit sistem adalah sebuah aplikasi web yang menyediakan API untuk pengajuan pinjaman

## Pencegahan Keamanan

Langkah pencegahan
1. **Validasi Input**: melakukan validasi parameter

2. **Perlindungan Terhadap Serangan Injection**: menggunakan parameterized queries atau prepared statements saat berinteraksi dengan basis data untuk mencegah SQL Injection.
3. **Autentikasi**: menerapkan autentikasi yang kuat menggunakan token JWT (JSON Web Tokens)

## Documentation

[Penjelasan Flow Sistem](https://gist.github.com/arifkurniawan200/0b3c627a005cb52749193ba9adcc4fa4)

# Arsitektur Aplikasi
arsitektur aplikasi terdiri dari 3 layer yang digambarkan sebagai berikut

![alt text for screen readers](./docs/arcitecture_system.jpg "bank balance")

# Rancangan Database
rancangan dan relasi database digambarkan sebagai berikut

![alt text for screen readers](./docs/kredit_system.png "bank balance")
## Running System

change app.yaml.example in folder config to app.yaml and setting the configuration based on your machine

install dependencies

```bash
  go mod tidy
```

running database migration (create table and seed data into the table)

```bash
  go run main.go db:migrate up
```


reset database (delete database and existing data)

```bash
  go run main.go db:migrate reset
```

running api server

```bash
  go run main.go api
```




## Tech Stack

**Database:** MySQL

**Framework:** Echo golang

**Migration:** Goose
## API Reference

#### login as admin

```http
  GET /login
  
  param :
  email: admin@example.com
  password: secure_password
```

#### register user

```http
  GET /register
  
  param :
  username: {{username}}
  password: {{password}}
```

### login user
```http
  GET /login
  
  param :
  username: {{username}}
  password: {{password}}
```

### Authorization

To Access endpoint always using bearer Authorization

```
Bearer {{token from login}}
```


#### example operation

postman file already attached in repo


