package main

import (
	"fmt"
)

type Mahasiswa struct {
	NIM          string
	Nama         string
	NilaiUTS     float64
	NilaiUAS     float64
	NilaiLainnya float64
	NilaiAkhir   float64
}

func hitungNilaiAkhir(mahasiswa *Mahasiswa) {
	mahasiswa.NilaiAkhir = 0.3*mahasiswa.NilaiUTS + 0.45*mahasiswa.NilaiUAS + 0.25*mahasiswa.NilaiLainnya
}

func selectionSortASC(mahasiswas []Mahasiswa) {
	for i := 0; i < len(mahasiswas)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(mahasiswas); j++ {
			if mahasiswas[j].NilaiAkhir < mahasiswas[minIndex].NilaiAkhir {
				minIndex = j
			}
		}
		mahasiswas[i], mahasiswas[minIndex] = mahasiswas[minIndex], mahasiswas[i]
	}
}

func insertionSortDESC(mahasiswas []Mahasiswa) {
	for i := 1; i < len(mahasiswas); i++ {
		key := mahasiswas[i]
		j := i - 1
		for j >= 0 && mahasiswas[j].NilaiAkhir < key.NilaiAkhir {
			mahasiswas[j+1] = mahasiswas[j]
			j--
		}
		mahasiswas[j+1] = key
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	var mahasiswas = make([]Mahasiswa, n)

	// Hitung nilai akhir untuk setiap mahasiswa
	for i := 0; i < len(mahasiswas); i++ {
		hitungNilaiAkhir(&mahasiswas[i])
	}

	// Urutkan nilai akhir secara ASC menggunakan Selection Sort
	selectionSortASC(mahasiswas)
	fmt.Println("Urutan Nilai Akhir ASC:")
	for i := 0; i < len(mahasiswas); i++ {
		fmt.Printf("%d. %s - Nilai Akhir: %.2f\n", i+1, mahasiswas[i].Nama, mahasiswas[i].NilaiAkhir)
	}

	// Urutkan nilai akhir secara DESC menggunakan Insertion Sort
	insertionSortDESC(mahasiswas)
	fmt.Println("Urutan Nilai Akhir DESC:")
	for i := 0; i < len(mahasiswas); i++ {
		fmt.Printf("%d. %s - Nilai Akhir: %.2f\n", i+1, mahasiswas[i].Nama, mahasiswas[i].NilaiAkhir)
	}
}
