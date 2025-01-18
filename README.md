# Nama Proyek

sistem ticketing untuk final projek bootcamp


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
setting db di main.go
