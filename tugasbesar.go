package main

import "fmt"

const MAX = 100

type Proyek struct {
	id             int
	nama           string
	danaDibutuhkan int
	danaTerkumpul  int
}

type Donatur struct {
	id          int
	nama        string
	totalDonasi int
}

var daftarProyek [MAX]Proyek
var jumlahProyek int = 0

var daftarDonatur [MAX]Donatur
var jumlahDonatur int = 0

// Fungsi insertion sort untuk proyek berdasarkan nama (ascending)
func insertionSortProyek(data *[MAX]Proyek, jumlah int) {
	for i := 1; i < jumlah; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].nama > key.nama {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

// Fungsi selection sort untuk donatur berdasarkan totalDonasi (descending)
func selectionSortDonatur(data *[MAX]Donatur, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		maxIdx := i
		for j := i + 1; j < jumlah; j++ {
			if data[j].totalDonasi > data[maxIdx].totalDonasi {
				maxIdx = j
			}
		}
		data[i], data[maxIdx] = data[maxIdx], data[i]
	}
}

func cariProyekByID(id int, data [MAX]Proyek, jumlah int) int {
	for i := 0; i < jumlah; i++ {
		if data[i].id == id {
			return i
		}
	}
	return -1
}

func cariDonaturByID(id int, data [MAX]Donatur, jumlah int) int {
	for i := 0; i < jumlah; i++ {
		if data[i].id == id {
			return i
		}
	}
	return -1
}

func tampilkanProyek(data [MAX]Proyek, jumlah int) {
	insertionSortProyek(&data, jumlah)
	fmt.Println("\n=== Daftar Proyek (Ascending Nama) ===")
	for i := 0; i < jumlah; i++ {
		fmt.Println("ID:", data[i].id)
		fmt.Println("Nama:", data[i].nama)
		fmt.Println("Dana Terkumpul:", data[i].danaTerkumpul)
		fmt.Println("Target Dana:", data[i].danaDibutuhkan)
		fmt.Println("------------------------")
	}
}

func tampilkanDonatur(data [MAX]Donatur, jumlah int) {
	selectionSortDonatur(&data, jumlah)
	fmt.Println("\n=== Daftar Donatur (Descending Total Donasi) ===")
	for i := 0; i < jumlah; i++ {
		fmt.Println("ID:", data[i].id)
		fmt.Println("Nama:", data[i].nama)
		fmt.Println("Total Donasi:", data[i].totalDonasi)
		fmt.Println("------------------------")
	}
}

func tambahProyek(data *[MAX]Proyek, jumlah *int) {
	var n int
	fmt.Print("Berapa proyek yang ingin ditambahkan? ")
	fmt.Scan(&n)
	for i := 0; i < n && *jumlah < MAX; i++ {
		var p Proyek
		fmt.Println("\n--- Tambah Proyek ---")

		ulang := true
		for ulang {
			fmt.Print("ID Proyek: ")
			fmt.Scan(&p.id)
			duplikat := false
			for j := 0; j < *jumlah; j++ {
				if data[j].id == p.id {
					fmt.Println("ID proyek sudah digunakan. Silakan masukkan ID lain.")
					duplikat = true
					break
				}
			}
			if !duplikat {
				ulang = false
			}
		}

		fmt.Print("Nama Proyek: ")
		fmt.Scan(&p.nama)
		fmt.Print("Dana Dibutuhkan: ")
		fmt.Scan(&p.danaDibutuhkan)

		p.danaTerkumpul = 0
		data[*jumlah] = p
		*jumlah++
		fmt.Println("Proyek berhasil ditambahkan.")
	}
	if *jumlah >= MAX {
		fmt.Println("Kapasitas proyek penuh.")
	}
}

func tambahDonatur(data *[MAX]Donatur, jumlah *int) {
	var n int
	fmt.Print("Berapa donatur yang ingin ditambahkan? ")
	fmt.Scan(&n)
	for i := 0; i < n && *jumlah < MAX; i++ {
		var d Donatur
		fmt.Println("\n--- Tambah Donatur ---")

		ulang := true
		for ulang {
			fmt.Print("ID Donatur: ")
			fmt.Scan(&d.id)
			duplikat := false
			for j := 0; j < *jumlah; j++ {
				if data[j].id == d.id {
					fmt.Println("ID donatur sudah digunakan. Silakan masukkan ID lain.")
					duplikat = true
					break
				}
			}
			if !duplikat {
				ulang = false
			}
		}

		fmt.Print("Nama Donatur: ")
		fmt.Scan(&d.nama)

		d.totalDonasi = 0
		data[*jumlah] = d
		*jumlah++
		fmt.Println("Donatur berhasil ditambahkan.")
	}
	if *jumlah >= MAX {
		fmt.Println("Kapasitas donatur penuh.")
	}
}

func donasi(proyek *[MAX]Proyek, jumlahProyek int, donatur *[MAX]Donatur, jumlahDonatur int) {
	var idProyek, idDonatur, nominal int
	fmt.Println("\n--- Donasi ke Proyek ---")
	fmt.Print("Masukkan ID proyek: ")
	fmt.Scan(&idProyek)
	fmt.Print("Masukkan ID donatur: ")
	fmt.Scan(&idDonatur)
	fmt.Print("Masukkan jumlah donasi: ")
	fmt.Scan(&nominal)

	if nominal <= 0 {
		fmt.Println("Nominal donasi harus lebih dari 0.")
		return
	}

	indeksProyek := cariProyekByID(idProyek, *proyek, jumlahProyek)
	indeksDonatur := cariDonaturByID(idDonatur, *donatur, jumlahDonatur)

	if indeksProyek == -1 || indeksDonatur == -1 {
		fmt.Println("ID proyek atau donatur tidak ditemukan.")
		return
	}

	sisaKebutuhan := proyek[indeksProyek].danaDibutuhkan - proyek[indeksProyek].danaTerkumpul
	if nominal > sisaKebutuhan {
		fmt.Println("Donasi melebihi kebutuhan proyek. Maksimal yang bisa didonasikan:", sisaKebutuhan)
		return
	}

	proyek[indeksProyek].danaTerkumpul += nominal
	donatur[indeksDonatur].totalDonasi += nominal
	fmt.Println("Donasi berhasil ditambahkan.")
}

func editProyek(data *[MAX]Proyek, jumlah int) {
	var id int
	fmt.Print("Masukkan ID proyek yang ingin diedit: ")
	fmt.Scan(&id)
	indeks := cariProyekByID(id, *data, jumlah)

	if indeks == -1 {
		fmt.Println("Proyek dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Println("Proyek ditemukan. Masukkan data baru:")
	fmt.Print("Nama Proyek: ")
	fmt.Scan(&data[indeks].nama)
	fmt.Print("Dana Dibutuhkan: ")
	fmt.Scan(&data[indeks].danaDibutuhkan)
	fmt.Println("Data proyek berhasil diperbarui.")
}

func hapusProyek(data *[MAX]Proyek, jumlah *int) {
	var id int
	fmt.Print("Masukkan ID proyek yang ingin dihapus: ")
	fmt.Scan(&id)
	indeks := cariProyekByID(id, *data, *jumlah)

	if indeks == -1 {
		fmt.Println("Proyek dengan ID tersebut tidak ditemukan.")
		return
	}

	for i := indeks; i < *jumlah-1; i++ {
		data[i] = data[i+1]
	}
	*jumlah--
	fmt.Println("Proyek berhasil dihapus.")
}

func editDonatur(data *[MAX]Donatur, jumlah int) {
	var id int
	fmt.Print("Masukkan ID donatur yang ingin diedit: ")
	fmt.Scan(&id)
	indeks := cariDonaturByID(id, *data, jumlah)

	if indeks == -1 {
		fmt.Println("Donatur dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Println("Donatur ditemukan. Masukkan data baru:")
	fmt.Print("Nama Donatur: ")
	fmt.Scan(&data[indeks].nama)
	fmt.Println("Data donatur berhasil diperbarui.")
}

func hapusDonatur(data *[MAX]Donatur, jumlah *int) {
	var id int
	fmt.Print("Masukkan ID donatur yang ingin dihapus: ")
	fmt.Scan(&id)
	indeks := cariDonaturByID(id, *data, *jumlah)

	if indeks == -1 {
		fmt.Println("Donatur dengan ID tersebut tidak ditemukan.")
		return
	}

	for i := indeks; i < *jumlah-1; i++ {
		data[i] = data[i+1]
	}
	*jumlah--
	fmt.Println("Donatur berhasil dihapus.")
}

func main() {
	var pilihan int
	selesai := false

	for !selesai {
		fmt.Println("\n=== Menu Utama ===")
		fmt.Println("1. Tambah Proyek")
		fmt.Println("2. Tambah Donatur")
		fmt.Println("3. Donasi ke Proyek")
		fmt.Println("4. Tampilkan Proyek")
		fmt.Println("5. Tampilkan Donatur")
		fmt.Println("6. Edit Proyek")
		fmt.Println("7. Hapus Proyek")
		fmt.Println("8. Edit Donatur")
		fmt.Println("9. Hapus Donatur")
		fmt.Println("10. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahProyek(&daftarProyek, &jumlahProyek)
		} else if pilihan == 2 {
			tambahDonatur(&daftarDonatur, &jumlahDonatur)
		} else if pilihan == 3 {
			donasi(&daftarProyek, jumlahProyek, &daftarDonatur, jumlahDonatur)
		} else if pilihan == 4 {
			tampilkanProyek(daftarProyek, jumlahProyek)
		} else if pilihan == 5 {
			tampilkanDonatur(daftarDonatur, jumlahDonatur)
		} else if pilihan == 6 {
			editProyek(&daftarProyek, jumlahProyek)
		} else if pilihan == 7 {
			hapusProyek(&daftarProyek, &jumlahProyek)
		} else if pilihan == 8 {
			editDonatur(&daftarDonatur, jumlahDonatur)
		} else if pilihan == 9 {
			hapusDonatur(&daftarDonatur, &jumlahDonatur)
		} else if pilihan == 10 {
			selesai = true
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
