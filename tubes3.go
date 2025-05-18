package main

import "fmt"

const MAX = 100

type Proyek struct {
	id             int
	nama           string
	deskripsi      string
	danaDibutuhkan int
	danaTerkumpul  int
}

type Donatur struct {
	id    int
	nama  string
	email string
}

var daftarProyek [MAX]Proyek
var jumlahProyek int = 0

var daftarDonatur [MAX]Donatur
var jumlahDonatur int = 0

func tambahProyek(data *[MAX]Proyek, jumlah *int) {
	if *jumlah < MAX {
		var p Proyek
		fmt.Println("\n--- Tambah Proyek ---")
		fmt.Print("ID Proyek: ")
		fmt.Scan(&p.id)
		fmt.Print("Nama Proyek: ")
		fmt.Scan(&p.nama)
		fmt.Print("Deskripsi: ")
		fmt.Scan(&p.deskripsi)
		fmt.Print("Dana Dibutuhkan: ")
		fmt.Scan(&p.danaDibutuhkan)

		p.danaTerkumpul = 0
		data[*jumlah] = p
		*jumlah = *jumlah + 1
		fmt.Println("Proyek berhasil ditambahkan.")
	} else {
		fmt.Println("Kapasitas proyek penuh.")
	}
}

func tambahDonatur(data *[MAX]Donatur, jumlah *int) {
	if *jumlah < MAX {
		var d Donatur
		fmt.Println("\n--- Tambah Donatur ---")
		fmt.Print("ID Donatur: ")
		fmt.Scan(&d.id)
		fmt.Print("Nama Donatur: ")
		fmt.Scan(&d.nama)
		fmt.Print("Email Donatur: ")
		fmt.Scan(&d.email)

		data[*jumlah] = d
		*jumlah = *jumlah + 1
		fmt.Println("Donatur berhasil ditambahkan.")
	} else {
		fmt.Println("Kapasitas donatur penuh.")
	}
}

func donasi(proyek *[MAX]Proyek, jumlah int) {
	var id int
	var nominal int

	fmt.Println("\n--- Donasi ke Proyek ---")
	fmt.Print("Masukkan ID proyek yang ingin didanai: ")
	fmt.Scan(&id)

	index := cariProyekByID(id, *proyek, jumlah)
	if index != -1 {
		fmt.Print("Masukkan nominal donasi: ")
		fmt.Scan(&nominal)
		proyek[index].danaTerkumpul = proyek[index].danaTerkumpul + nominal
		fmt.Println("Donasi berhasil ditambahkan.")
	} else {
		fmt.Println("Proyek tidak ditemukan.")
	}
}

func ubahProyek(data *[MAX]Proyek, jumlah int) {
	var id int
	fmt.Println("\n--- Ubah Data Proyek ---")
	fmt.Print("Masukkan ID proyek: ")
	fmt.Scan(&id)

	index := cariProyekByID(id, *data, jumlah)
	if index != -1 {
		fmt.Println("Data ditemukan. Silakan ubah:")
		fmt.Print("Nama Proyek: ")
		fmt.Scan(&data[index].nama)
		fmt.Print("Deskripsi: ")
		fmt.Scan(&data[index].deskripsi)
		fmt.Print("Dana Dibutuhkan: ")
		fmt.Scan(&data[index].danaDibutuhkan)
		fmt.Println("Data proyek berhasil diubah.")
	} else {
		fmt.Println("Proyek tidak ditemukan.")
	}
}

func hapusProyek(data *[MAX]Proyek, jumlah *int) {
	var id int
	fmt.Println("\n--- Hapus Proyek ---")
	fmt.Print("Masukkan ID proyek: ")
	fmt.Scan(&id)

	index := cariProyekByID(id, *data, *jumlah)
	if index != -1 {
		for i := index; i < *jumlah-1; i++ {
			data[i] = data[i+1]
		}
		*jumlah = *jumlah - 1
		fmt.Println("Proyek berhasil dihapus.")
	} else {
		fmt.Println("Proyek tidak ditemukan.")
	}
}

func cariProyekByNama(nama string, data [MAX]Proyek, jumlah int) int {
	var i int = 0
	var ketemu bool = false

	for i < jumlah && !ketemu {
		if data[i].nama == nama {
			ketemu = true
		} else {
			i = i + 1
		}
	}

	if ketemu {
		return i
	} else {
		return -1
	}
}

func cariProyekByID(id int, data [MAX]Proyek, jumlah int) int {
	var kiri, kanan, tengah int
	kiri = 0
	kanan = jumlah - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if data[tengah].id == id {
			return tengah
		} else if id < data[tengah].id {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	return -1
}

func selectionSortNama(data *[MAX]Proyek, jumlah int, ascending bool) {
	var i, j, idx int
	for i = 0; i < jumlah-1; i++ {
		idx = i
		for j = i + 1; j < jumlah; j++ {
			if ascending {
				if data[j].nama < data[idx].nama {
					idx = j
				}
			} else {
				if data[j].nama > data[idx].nama {
					idx = j
				}
			}
		}
		temp := data[i]
		data[i] = data[idx]
		data[idx] = temp
	}
}

func insertionSortDana(data *[MAX]Proyek, jumlah int, ascending bool) {
	var i, j int
	var key Proyek

	for i = 1; i < jumlah; i++ {
		key = data[i]
		j = i - 1

		if ascending {
			for j >= 0 && data[j].danaTerkumpul > key.danaTerkumpul {
				data[j+1] = data[j]
				j = j - 1
			}
		} else {
			for j >= 0 && data[j].danaTerkumpul < key.danaTerkumpul {
				data[j+1] = data[j]
				j = j - 1
			}
		}

		data[j+1] = key
	}
}

func tampilkanMenu() {
	fmt.Println("\n=== Menu Utama ===")
	fmt.Println("1. Tambah Proyek")
	fmt.Println("2. Tambah Donatur")
	fmt.Println("3. Donasi ke Proyek")
	fmt.Println("4. Ubah Proyek")
	fmt.Println("5. Hapus Proyek")
	fmt.Println("6. Cari Proyek")
	fmt.Println("7. Tampilkan Proyek")
	fmt.Println("8. Keluar")
	fmt.Print("Pilih menu (1-8): ")
}

func main() {
	var pilihan int
	var selesai bool = false

	for !selesai {
		tampilkanMenu()
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahProyek(&daftarProyek, &jumlahProyek)
		} else if pilihan == 2 {
			tambahDonatur(&daftarDonatur, &jumlahDonatur)
		} else if pilihan == 3 {
			donasi(&daftarProyek, jumlahProyek)
		} else if pilihan == 4 {
			ubahProyek(&daftarProyek, jumlahProyek)
		} else if pilihan == 5 {
			hapusProyek(&daftarProyek, &jumlahProyek)
		} else if pilihan == 6 {
			var nama string
			fmt.Print("Masukkan nama proyek yang dicari: ")
			fmt.Scan(&nama)
			index := cariProyekByNama(nama, daftarProyek, jumlahProyek)
			if index != -1 {
				fmt.Println("Proyek ditemukan:")
				fmt.Println("ID:", daftarProyek[index].id)
				fmt.Println("Nama:", daftarProyek[index].nama)
				fmt.Println("Deskripsi:", daftarProyek[index].deskripsi)
				fmt.Println("Dana Terkumpul:", daftarProyek[index].danaTerkumpul)
				fmt.Println("Target Dana:", daftarProyek[index].danaDibutuhkan)
			} else {
				fmt.Println("Proyek tidak ditemukan.")
			}
		} else if pilihan == 7 {
			var kategori int
			var urutan int
			fmt.Println("Kategori pengurutan:")
			fmt.Println("1. Nama Proyek (Selection Sort)")
			fmt.Println("2. Dana Terkumpul (Insertion Sort)")
			fmt.Print("Pilih kategori: ")
			fmt.Scan(&kategori)
			fmt.Print("Urutan (1 = Ascending, 2 = Descending): ")
			fmt.Scan(&urutan)

			if kategori == 1 {
				selectionSortNama(&daftarProyek, jumlahProyek, urutan == 1)
			} else if kategori == 2 {
				insertionSortDana(&daftarProyek, jumlahProyek, urutan == 1)
			}

			fmt.Println("\n=== Daftar Proyek ===")
			var i int = 0
			for i < jumlahProyek {
				fmt.Println("ID:", daftarProyek[i].id)
				fmt.Println("Nama:", daftarProyek[i].nama)
				fmt.Println("Deskripsi:", daftarProyek[i].deskripsi)
				fmt.Println("Terkumpul:", daftarProyek[i].danaTerkumpul)
				fmt.Println("Target:", daftarProyek[i].danaDibutuhkan)
				fmt.Println("----------------------")
				i = i + 1
			}
		} else if pilihan == 8 {
			selesai = true
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
