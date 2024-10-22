package main

import (
	"fmt"
	"time"
)

type buku struct {
	Kode                              int
	Judul, Jenis, Pengarang, Penerbit string
	TahunTerbit                       int
}

const NMAX int = 100

type Tabbuku [NMAX]buku
type Tabpinjam [NMAX]buku

func main() {
	var buku Tabbuku
	var pilih int
	var jumlahBuku int
	// var jumlahPeminjaman int

	jumlahBuku = 0
	// jumlahPeminjaman = 0
	for {
		menu()
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahBuku(&buku, &jumlahBuku)
		} else if pilih == 2 {
			editBuku(&buku, jumlahBuku)
		} else if pilih == 3 {
			hapusBuku(&buku, &jumlahBuku)
		} else if pilih == 4 {
			cetakBuku(buku, jumlahBuku)

		} else if pilih == 5 {
			var Judul string
			fmt.Println("Masukkan Judul Buku yang ingin dicari: ")
			fmt.Scan(&Judul)
			idx := cariBuku(buku, jumlahBuku, Judul) // Mencari buku dan menyimpan indeksnya

			if idx == -1 {
				fmt.Println("Buku tidak ditemukan.")
			} else {
				fmt.Println("Buku ditemukan:", buku[idx].Judul) // Menampilkan data buku
			}

		} else if pilih == 15 {
			Denda()
		} else if pilih == 6 {
			tambahBukuPinjam(&buku, &p, &n, judul, judulbaru)
		} else if pilih == 7 {
			editBukuPinjam(&b, n)
		} else if pilih == 8 {
			hapusBukuPinjam(&b, &n)
		} else if pilih == 9 {

		}
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
	fmt.Println("15. Denda")
	fmt.Print("Pilih menu: ")

}

func tambahBuku(buku *Tabbuku, jumlahBuku *int) {

	fmt.Println("Masukkan Kode Buku: ")
	fmt.Scan(&buku[*jumlahBuku].Kode)
	fmt.Println("Masukkan Judul Buku: ")
	fmt.Scan(&buku[*jumlahBuku].Judul)
	fmt.Println("Masukkan Pengarang Buku: ")
	fmt.Scan(&buku[*jumlahBuku].Pengarang)
	fmt.Println("Masukkan Penerbit Buku: ")
	fmt.Scan(&buku[*jumlahBuku].Penerbit)
	fmt.Println("Masukkan Tahun Terbit: ")
	fmt.Scan(&buku[*jumlahBuku].TahunTerbit)

	*jumlahBuku++
	fmt.Println("Data buku berhasil ditambahkan!")
}

func editBuku(b *Tabbuku, n int) {
	var judul string
	var idx int
	fmt.Scan(&judul)
	idx = cariBuku(*b, n, judul)
	if idx == -1 {
		fmt.Println("buku tidak ditemukan")
	} else {
		fmt.Scan(&b[idx].Judul)
		fmt.Println("Data buku berhasil diedit!")
	}

}
func hapusBuku(b *Tabbuku, n *int) {
	var judul string
	var idx int
	fmt.Scan(&judul)
	idx = cariBuku(*b, *n, judul)
	if idx == -1 {
		fmt.Println("buku tidak ditemukan")
	} else {
		for i := idx; i < *n-1; i++ {
			b[i] = b[i+1]
		}
		*n--
		fmt.Println("Data buku berhasil dihapus!")
	}

}
func urutBuku(b *Tabbuku, n int) {

	var i int
	var j int
	var idx_min int
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if b[idx_min].Judul > b[j].Judul {
				idx_min = j
			}
			j = j + 1
		}
		t := b[idx_min]
		b[idx_min] = b[i-1]
		b[i-1] = t
		i = i + 1
	}

}
func cetakBuku(b Tabbuku, n int) {
	urutBuku(&b, n)
	for i := 0; i < n; i++ {
		fmt.Println(b[i].Kode, b[i].Judul, b[i].Pengarang, b[i].Penerbit, b[i].TahunTerbit)
	}
}

func tambahBukuPinjam(b *Tabbuku, p *Tabpinjam, n *int, judul, judulbaru string) {
	var idx int = cariBuku(*b, *n, judul)
	if idx == -1 {
		fmt.Println("buku tidak ditemukan")
	} else {
		p[*n] = b[idx]
		*n++
		fmt.Println("Data buku berhasil ditambahkan!")
	}

}

func editBukuPinjam(b *Tabpinjam, n int) {
	var judul string
	var idx int
	fmt.Scan(&judul)
	idx = cariBukuPinjam(*b, n, judul)
	if idx == -1 {
		fmt.Println("buku tidak ditemukan")
	} else {
		fmt.Scan(&b[idx].Judul)
		fmt.Println("Data buku berhasil diedit!")
	}

}

func hapusBukuPinjam(b *Tabpinjam, n *int) {
	var judul string
	var idx int
	fmt.Scan(&judul)
	idx = cariBukuPinjam(*b, *n, judul)
	if idx == -1 {
		fmt.Println("buku tidak ditemukan")
	} else {
		for i := idx; i < *n-1; i++ {
			b[i] = b[i+1]
		}
		*n--
		fmt.Println("Data buku berhasil dihapus!")
	}

}

func cetakBuku(b Tabpinjam, n int) {
	urutBuku(&b, n)
	for i := 0; i < n; i++ {
		fmt.Println(b[i].Kode, b[i].Judul, b[i].Pengarang, b[i].Penerbit, b[i].TahunTerbit)
	}
}

func cariBuku(b Tabbuku, n int, judul string) int {
	/* mengembalikan indeks dari X apabila X ditemukan di dalam array T yang berisi
	n buah teks, atau -1 apabila X tidak ditemukan */
	var found int = -1
	var j int = 0
	for j < n && found == -1 {
		if b[j].Judul == judul {
			found = j
		}
		j = j + 1
	}
	return found
}

//

func cariBukuPinjam(b Tabpinjam, n int, judul string) int {
	/* mengembalikan indeks dari X apabila X ditemukan di dalam array T yang berisi
	n buah teks, atau -1 apabila X tidak ditemukan */
	var found int = -1
	var j int = 0
	for j < n && found == -1 {
		if b[j].Judul == judul {
			found = j
		}
		j = j + 1
	}
	return found
}

// Fungsi untuk menghitung denda keterlambatan
func hitungDenda(pinjam time.Time, kembali time.Time) int {
	// Hitung selisih hari antara tanggal pinjam dan tanggal kembali
	selisihHari := int(kembali.Sub(pinjam).Hours() / 24)

	// Hitung denda
	denda := selisihHari * 1000

	// Batasi denda maksimal
	if denda > 10000 {
		denda = 10000
	}

	return denda
}

func Denda() {
	// Deklarasi variabel
	var idBuku int
	var tglPinjam, tglKembali string

	// Meminta input dari pengguna
	fmt.Print("Masukkan ID buku: ")
	fmt.Scan(&idBuku)

	fmt.Print("Masukkan tanggal pinjam (YYYY-MM-DD): ")
	fmt.Scan(&tglPinjam)

	fmt.Print("Masukkan tanggal kembali (YYYY-MM-DD): ")
	fmt.Scan(&tglKembali)

	t, _ := time.Parse("2024-08-18", tglPinjam)
	t2, _ := time.Parse("2024-08-18", tglKembali)
	// Menghitung denda
	denda := hitungDenda(t, t2)

	// Menampilkan hasil perhitungan
	fmt.Println("Informasi peminjaman:")
	fmt.Println("ID buku:", idBuku)
	fmt.Println("Judul buku:" /* Judul buku */) // Ambil judul buku berdasarkan ID buku
	fmt.Println("Tanggal pinjam:", tglPinjam)
	fmt.Println("Tanggal kembali:", tglKembali)
	fmt.Println("Denda:", denda)
}
