#### Roadmap.sh 
https://roadmap.sh/projects/task-tracker

# Task Tracker

Task Tracker adalah aplikasi command line sederhana untuk mengelola dan melacak tugas. Aplikasi ini memungkinkan pengguna untuk menambahkan, memperbarui, menghapus, dan memantau status tugas.

## Requirements

Aplikasi ini harus berjalan dari command line, menerima aksi dan input pengguna sebagai argumen, dan menyimpan tugas dalam file JSON. Pengguna harus dapat:

- Menambahkan, Memperbarui, dan Menghapus tugas
- Menandai tugas sebagai "in progress" atau "done"
- Mendaftar semua tugas
- Mendaftar semua tugas yang selesai
- Mendaftar semua tugas yang belum selesai
- Mendaftar semua tugas yang sedang dikerjakan

### Constraints

- Anda dapat menggunakan bahasa pemrograman apa pun untuk membangun proyek ini.
- Gunakan argumen posisi di command line untuk menerima input pengguna.
- Gunakan file JSON untuk menyimpan tugas di direktori saat ini.
- File JSON harus dibuat jika tidak ada.
- Gunakan modul sistem file bawaan dari bahasa pemrograman Anda untuk berinteraksi dengan file JSON.
- Jangan menggunakan pustaka atau framework eksternal untuk membangun proyek ini.
- Pastikan untuk menangani kesalahan dan kasus tepi dengan baik.

## Task Properties

Setiap tugas harus memiliki properti berikut:

- **id**: Pengidentifikasi unik untuk tugas
- **description**: Deskripsi singkat tentang tugas
- **status**: Status tugas (todo, in-progress, done)
- **createdAt**: Tanggal dan waktu saat tugas dibuat
- **updatedAt**: Tanggal dan waktu saat tugas terakhir diperbarui

Pastikan untuk menambahkan properti ini ke file JSON saat menambahkan tugas dan memperbarui mereka saat melakukan pembaruan.

## Example

Berikut adalah daftar perintah dan penggunaannya:

### Menambahkan Tugas Baru

```bash
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
