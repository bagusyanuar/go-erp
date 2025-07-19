# Go ERP Project

Ini adalah proyek ERP (Enterprise Resource Planning) berbasis bahasa Go menggunakan PostgreSQL sebagai database dan `golang-migrate` untuk migrasi database.

## Prasyarat

Sebelum menjalankan proyek ini, pastikan kamu sudah menginstal:

- Go 1.21+
- PostgreSQL
- [Scoop](https://scoop.sh/) (untuk pengguna Windows)

## Instalasi Scoop (Windows)

Jika belum memiliki `scoop`, jalankan di PowerShell:

```powershell
Set-ExecutionPolicy RemoteSigned -scope CurrentUser
irm get.scoop.sh | iex

scoop install migrate
scoop install make

```