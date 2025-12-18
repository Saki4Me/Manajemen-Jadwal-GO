package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Jadwal struct {
	Nama string
	Hari string
}

var ListHari = []string{
	"Senin",
	"Selasa",
	"Rabu",
	"Kamis",
	"Jumat",
	"Sabtu",
	"Minggu",
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	// Func menu utama

	var daftarJadwal []Jadwal
	// Deklarasi variabel di dalam []Jadwal

	for {
		// Looping  pada func menu

		fmt.Println("\n|___MENU JADWAL___|")
		fmt.Println("1. TAMBAHKAN JADWAL")
		fmt.Println("2. LIST JADWAL")
		fmt.Println("3. HAPUS JADWAL")
		fmt.Println("4. KELUAR")
		// Opsi pada func menu

		var pilihan int
		fmt.Println("PILIH MENU (1-4) : ")
		fmt.Scan(&pilihan)

		switch pilihan {
		// Pilihan pada menu yang akan mengarah ke func lain

		case 1:
			TambahJadwal(&daftarJadwal)
			// membawa ke func TambahJadwal

		case 2:
			ListJadwal(daftarJadwal)
			// membawa ke func ListJadwal

		case 3:
			HapusJadwal(&daftarJadwal)
			// membawa ke func HapusJadwal

		case 4:
			// mengakhiri program

			fmt.Println("*PROGRAM SELESAI*")
			return

		default:
			fmt.Println("pilihan tidak valid")
		}
	}
}

func TambahJadwal(daftar *[]Jadwal) {
	// func tambah jadwal

	var pilihan int
	// untuk memilih hari pada ListHari

	fmt.Println("PILIH HARI : ")
	for i, h := range ListHari {
		fmt.Printf("%d. %s\n", i+1, h)
	}

	fmt.Println("PILIHAN :")
	fmt.Scan(&pilihan)
	reader.ReadString('\n')

	if pilihan < 1 || pilihan > 7 {
		fmt.Println("PILIHAN HARI TIDAK VALID")
		return
	}
	// bila user memilih diluar opsi yang diberikan

	hari := ListHari[pilihan-1]
	// pilihan hari sesuai dengan opsi

	for {
		fmt.Print("MASUKKAN KEGIATAN : ")
		kegiatan, _ := reader.ReadString('\n')
		kegiatan = strings.TrimSpace(kegiatan)
		// menambahkan kegiatan yang dimasukkan user kedalam jadwal

		if kegiatan == "" {
			fmt.Println("Kegiatan tidak boleh kosong")
			continue
		}

		*daftar = append(*daftar, Jadwal{
			Hari: hari,
			Nama: kegiatan,
			/* membuat output menjadi :
			senin :
			- abc

			selasa :
			- (tidak ada jadwal)

			dst*/
		})

		fmt.Print("Tambah kegiatan dihari yang sama? (ya/tdk) : ")
		lanjut, _ := reader.ReadString('\n')
		lanjut = strings.TrimSpace(strings.ToLower(lanjut))

		if lanjut != "ya" {
			break
		}
	}
}

func ListJadwal(daftar []Jadwal) {
	fmt.Println("===== LIST JADWAL =====")

	for _, hari := range ListHari {
		fmt.Println(hari, ":")

		ada := false

		for _, j := range daftar {
			if j.Hari == hari {
				fmt.Println("-", j.Nama)
				ada = true
			}
		}

		if !ada {
			fmt.Println("(tidak ada jadwal)")
		}

		fmt.Println()
	}
}

func HapusJadwal(daftar *[]Jadwal) {
	// tampilkan jadwal terlebih dahulu
	ListJadwal(*daftar)

	// pilih hari
	fmt.Println("PILIH HARI YANG AKAN DIHAPUS JADWALNYA :")
	for i, h := range ListHari {
		fmt.Printf("%d. %s\n", i+1, h)
	}

	fmt.Print("PILIHAN : ")
	var pilihan int
	fmt.Scan(&pilihan)
	reader.ReadString('\n')

	if pilihan < 1 || pilihan > len(ListHari) {
		fmt.Println("Pilihan hari tidak valid")
		return
	}

	hariDipilih := ListHari[pilihan-1]

	// kumpulkan indeks jadwal di hari tersebut
	var indeks []int
	for i, j := range *daftar {
		if j.Hari == hariDipilih {
			indeks = append(indeks, i)
		}
	}

	if len(indeks) == 0 {
		fmt.Println("Tidak ada jadwal pada hari tersebut")
		return
	}

	// tampilkan kegiatan
	fmt.Println("KEGIATAN PADA", hariDipilih, ":")
	for i, idx := range indeks {
		fmt.Printf("%d. %s\n", i+1, (*daftar)[idx].Nama)
	}

	// pilih kegiatan
	fmt.Print("PILIH KEGIATAN YANG AKAN DIHAPUS : ")
	var pilihKegiatan int
	fmt.Scan(&pilihKegiatan)
	reader.ReadString('\n')

	if pilihKegiatan < 1 || pilihKegiatan > len(indeks) {
		fmt.Println("Pilihan kegiatan tidak valid")
		return
	}

	// hapus dari slice utama
	hapusIndex := indeks[pilihKegiatan-1]
	*daftar = append((*daftar)[:hapusIndex], (*daftar)[hapusIndex+1:]...)

	fmt.Println("Jadwal berhasil dihapus")
}
