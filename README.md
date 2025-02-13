# Nama Proyek

sistem ticketing untuk final projek bootcamp
ticket pengaduan client terkait issue / kerusakan

## ALUR 
1. client membuat ticket pengaduan
2. admin memberikan ticket ke engineer untuk di kerjakan
3. engineer mengerjakan ticket yang di berikan

## Prasyarat

- Go 
- MySQL 
- GORM (untuk ORM di Golang)
- GIN

## Instalasi


### Langkah Langkah Instalasi

1. Clone repositori ini ke mesin lokal Anda.
   git clone https://github.com/iyan1994/final_projek_ticketing.git 
2. Membuat database pada mysql 
3. Ubah koneksi database pada file main.go 
4. jalankan go run main.go "table otomatis di buat dengan menjalan kan migrate pada folder models/migrate_table"
5. analisa system ada di file "analisa ticketing system.pptx"

### MENJALANKAN DOCKER FILE
Docker build image docker file
docker build -t final_projek_ticketing . 

Docker create container
docker container create --name final_projek_ticketing -e PORT=8083 -e INSTANCE_ID="my first instance" -p 8083:8083 final_projek_ticketing

### CONFIG DATABASE 
setting db di .env (jika menggunakan container gunakan host "host.docker.internal") 

### 
ROLE 
- 1 = admin
- 2 = client
- 3 = enginner

CATEGORIE 
- 1 = Request
- 2 = Problem
- 3 = Corective

### ENDPOINT 

-- login --
"/login" method POST  // semua role
Request body 
{
    "username" : "test",
    "password" : "test"
}

-- create user --
"/user" method POST // semua role
AUTHORIZATION = barier token
Request boy
{
     "username" : "client",
	  "email" : "client@gmail.com",
	  "name" : "client",
	  "password" : "Cicicuit@12345",
	  "id_role" : 2,
     "title" : "pt ini",
	  "no_telepon" : 872723774,
	  "address" : "jln. gatot subroto"
}

-- pembuatan ticket pengaduan --
"/ticket" method POST // hanya role client
AUTHORIZATION = barier token
Request body
{
	"id_category" : 1,
	"subjek"    :"kendala tidak bisa connetc internet",
	"deksripsi" :"kendala tidak bisa akses internet"
}

-- pembuatan myticket client --
"/ticket/myticket?page=1&page_size=20&status=Closed&start_date=2025-01-01&end_date=2025-01-20" method GET // hanya role client
AUTHORIZATION = barier token


-- create feedback -- 
"/ticket/feedback/id_ticket" method POST // hanya role client
AUTHORIZATION = barier token
Request body
{
    "satisfaction" : "Puas", //sangat puas. tidak.
    "deksripsi": "perkerjaan bagugs"
}

-- assign ticket untuk di kerjakan ke engineer -- 
"/assignticket" method POST // hanya role admin
AUTHORIZATION = barier token
Request body
{
   "id_ticket" : id_ticket,
	"id_priority" : 1, 
	"id_teknisi" : id_teknisi
}

-- start ticket -- 
"/assignticket/start/id_assign_ticket" method PUT // hanya role engineer
AUTHORIZATION = barier token

-- closed ticket -- 
"/assignticket/closed/id_assign_ticket" method PUT // hanya role engineer
AUTHORIZATION = barier token

-- upload solution -- 
"/assignticket/solution/id_assign_ticket" method POST // hanya role engineer
AUTHORIZATION = barier token
Request body FORM-DATA
{
   "image" : img,
	"name_image" : "nama image" 
	"description" : "update deksripsi solution"
}

-- view assign ticket -- 
"/assignticket/myassignticket?page=1&page_size=10" method GET // hanya role engineer
AUTHORIZATION = barier token

-- view all ticket -- 
"/ticket/allticket?page=1&page_size=30&status=Closed&start_date=2025-01-01&end_date=2025-01-20" method GET// hanya role admin
AUTHORIZATION = barier token

-- delete user -- 
"user/delete/id_user" method DELETE // hanya role admin
AUTHORIZATION = barier token

