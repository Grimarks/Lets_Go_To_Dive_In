# Lets_Go_To_Dive_In

Nama Proyek: Lets Go to Dive In!  
Deskripsi:  
Proyek ini merupakan pengembangan dari backend_started dengan fitur tambahan berupa otentikasi JWT dan validasi input. Setiap endpoint telah dilengkapi dengan dokumentasi API menggunakan format JSON.  

Struktur Direktori:  
- controllers: Berisi fungsi untuk menangani request endpoint dan otentikasi.  
- middlewares: Berisi middleware untuk mengamankan API.  
- routes: Berisi konfigurasi rute untuk endpoint API yang aman.  
- models: Berisi struktur data untuk validasi input.  

Cara Penggunaan:  
1. Pastikan Anda telah mengikuti langkah pada folder backend_started untuk setup database dan server.  
2. Jalankan server backend menggunakan perintah:  
   ```
   go run main.go
   ```  
3. Gunakan Postman untuk mengakses endpoint API sesuai dokumentasi.  
   - Login menggunakan endpoint /login untuk mendapatkan JWT token.  
   - Gunakan token yang diterima untuk mengakses endpoint lain dengan menambahkan header Authorization: Bearer {token}.  

Dokumentasi API:  
1. POST /login  
   - Login untuk mendapatkan JWT token.  
2. GET /products  
   - Mengambil daftar produk. Membutuhkan header Authorization.  
3. POST /products  
   - Menambahkan produk baru. Membutuhkan header Authorization.  

Fitur Tambahan:  
- Proteksi endpoint menggunakan JWT.  
- Validasi input pada setiap endpoint untuk memastikan data yang dimasukkan sesuai format.  
- Dokumentasi API untuk memudahkan integrasi.  

Catatan:  
Pastikan Anda menyimpan JWT token dengan aman selama sesi berlangsung. Token memiliki masa berlaku 72 jam. 
