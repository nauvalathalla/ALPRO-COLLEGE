package main

import "fmt"

type buku struct {
	Kode                int
	Judul               string
	Pengarang           string
	Penerbit            string
	TahunTerbit         int
	TanggalPinjam       int
	TanggalPengembalian int
}

const NMAX int = 100

type Tabbuku [NMAX]buku

func main() {

	jumlahBuku = 0
	jumlahPeminjaman = 0

	for {
		mainMenu()
	}
}

func menu() {
	fmt.Println("\n=== Aplikasi Perpustakaan ===")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Edit Buku")
	fmt.Println("3. Hapus Buku")
	fmt.Println("4. Tampilkan Buku")
	fmt.Println("5. Cari Buku")
	fmt.Println("6. Tambah Peminjaman")
	fmt.Println("7. Edit Peminjaman")
	fmt.Println("8. Hapus Peminjaman")
	fmt.Println("9. Tampilkan Peminjaman")
	fmt.Println("10. Cari Peminjaman")
	fmt.Println("11. Buku Terpinjam")
	fmt.Println("12. Buku Sering Dipinjam")
	fmt.Println("13. Buku Jarang Dipinjam")
	fmt.Println("14. Keluar")
	fmt.Print("Pilih menu: ")

}

func inputBuku() {
	var b buku

	fmt.Print("Masukkan Kode Buku: ")
	fmt.Scanln(&b.Kode)
	fmt.Print("Masukkan Judul Buku: ")
	fmt.Scanln(&b.Judul)
	fmt.Print("Masukkan Pengarang Buku: ")
	fmt.Scanln(&b.Pengarang)
	fmt.Print("Masukkan Penerbit Buku: ")
	fmt.Scanln(&b.Pengarang)
	fmt.Print("Masukkan Tahun Terbit: ")
	fmt.Scanln(&b.TahunTerbit)
	fmt.Print("Masukkan Tanggal Pinjam: ")
	fmt.Scanln(&b.TanggalPinjam)
	fmt.Print("Masukkan Tanggal Pengembalian: ")
	fmt.Scanln(&b.TanggalPengembalian)

	buku[jumlahBuku] = Buku{
		Kode:        kode,
		Judul:       judul,
		Pengarang:   pengarang,
		Penerbit:    penerbit,
		TahunTerbit: tahunTerbit,
	}
	jumlahBuku++
	fmt.Println("Data buku berhasil ditambahkan!")
}
