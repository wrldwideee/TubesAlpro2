package main
import "fmt"

//variable global
const MAX int = 100
const ADMIN int = 0
const USER  int = 1

var daftarProyek [MAX]Proyek
var jumlahProyek int = 0
var daftarDonatur [MAX]Donatur
var jumlahDonatur int = 0

var daftarUser [MAX]User
var jumlahUser int = 0
var currentUser User
var nextUserID int = 1 

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

type User struct {
	id int
	username string
	password string
	userType int
}

func main() {
	initAdmin()
	var pilihan int
	loggedIn := false
	aplikasiAktif := true
	
	for aplikasiAktif {
		if !loggedIn {
			fmt.Println("\n====================== SimpleFund Login Menu ======================")
			fmt.Println("1. Login")
			fmt.Println("2. Register")
			fmt.Println("3. Keluar Aplikasi")
			fmt.Print("Pilih menu: ")
			fmt.Scan(&pilihan)
			
			if pilihan == 1 {
				loggedIn = login()
			} else if pilihan == 2 {
				register()
			} else if pilihan == 3 {
				fmt.Println("Terima kasih telah menggunakan aplikasi.")
				aplikasiAktif = false
			} else {
				fmt.Println("Pilihan tidak valid.")
			}
		} else {
			if currentUser.userType == ADMIN {
				adminMenu()
			} else {
				userMenu()
			}
			//logout
			loggedIn = false
		}
	}
}

// inisialisasi admin
func initAdmin() {
	adminUser := User{
		id:       0,   
		username: "admin",
		password: "admin",
		userType: ADMIN,
	}
	daftarUser[jumlahUser] = adminUser
	jumlahUser++
}

func login() bool {
	var username, password string
	fmt.Println("\n========================== Login SimpleFund ==========================")
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	for i := 0; i < jumlahUser; i++ {
		if daftarUser[i].username == username && daftarUser[i].password == password {
			currentUser = daftarUser[i]
			fmt.Printf("Login berhasil! Selamat datang %s (ID: %d)\n", currentUser.username, currentUser.id)
			return true
		}
	}
	fmt.Println("Username atau password salah!")
	return false
}

// Register user baru
func register() {
	if jumlahUser >= MAX {
		fmt.Println("Kapasitas pengguna penuh!")
		return
	}

	var newUser User
	fmt.Println("\n======================== Register SimpleFund ========================")
	
	isUnique := false
	for !isUnique {
		fmt.Print("Username baru: ")
		fmt.Scan(&newUser.username)
		
		isUnique = true
		for i := 0; i < jumlahUser && isUnique; i++ {
			if daftarUser[i].username == newUser.username {
				fmt.Println("Username sudah digunakan. Silakan pilih username lain.")
				isUnique = false
			}
		}
	}
	
	fmt.Print("Password: ")
	fmt.Scan(&newUser.password)
	
	// ID user selalu mulai dari 1
	newUser.userType = USER
	
	// assign ID otomatis (dimulai dari 1, bertambah secara berurutan)
	newUser.id = nextUserID
	nextUserID++
	
	daftarUser[jumlahUser] = newUser
	jumlahUser++
	
	fmt.Printf("Registrasi berhasil! Anda telah terdaftar dengan ID: %d\n", newUser.id)
	fmt.Println("Silakan login dengan username dan password Anda.")
}

// menu user biasa
func userMenu() {
	var pilihan int
	var pilihanSorting int
	selesai := false
	urutkanDefault := true
	urutkanNamaProyek := false
	urutkanDanaTerkumpul := false
	urutkanDanaDibutuhkan := false
	ProyekDicari := false
	
	
	for !selesai {
		if !ProyekDicari{
			fmt.Println("\n======================= Menu User SimpleFund =======================")
		}
		if urutkanDefault == false && urutkanNamaProyek == true && urutkanDanaTerkumpul == false && urutkanDanaDibutuhkan == false && ProyekDicari == false{
			tampilkanProyekUrutkanNama(daftarProyek, jumlahProyek)
		} else if urutkanDefault == false && urutkanNamaProyek == false && urutkanDanaTerkumpul == true && urutkanDanaDibutuhkan == false && ProyekDicari == false{
			tampilkanProyekUrutkanDanaTerkumpul(daftarProyek, jumlahProyek)
		} else if urutkanDefault == false && urutkanNamaProyek == false && urutkanDanaTerkumpul == false && urutkanDanaDibutuhkan == true && ProyekDicari == false{
			tampilkanProyekUrutkanDanaDibutuhkan(daftarProyek, jumlahProyek)
		} else if urutkanDefault == true && urutkanNamaProyek == false && urutkanDanaTerkumpul == false && urutkanDanaDibutuhkan == false && ProyekDicari == false{
			tampilkanProyekDefault(daftarProyek, jumlahProyek)
		} else {
			//tidak menampilkan
		}
		fmt.Println("\n============================ Pilih Menu ============================")
		fmt.Println("1. Tambah Proyek")
		fmt.Println("2. Donasi ke Proyek")
		fmt.Println("3. Urutkan Proyek")
		fmt.Println("4. Cari Proyek")
		fmt.Println("5. Logout")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahProyek(&daftarProyek, &jumlahProyek)
		} else if pilihan == 2 {
			donasiUser(&daftarProyek, jumlahProyek, &daftarDonatur, &jumlahDonatur)
		} else if pilihan == 3 {
			fmt.Println("====== Pilih Jenis Sorting ======")
			fmt.Println("1. Urutkan berdasarkan nama (Ascending)")
			fmt.Println("2. Urutkan berdasarkan dana terkumpul (Descending)")
			fmt.Println("3. Urutkan berdasarkan dana dibutuhkan (Descending)")
			fmt.Println("4. Kembali ke Urutan Default")
			fmt.Print("Pilih menu: ")
			fmt.Scan(&pilihanSorting)
			if pilihanSorting == 1{
				urutkanDefault = false
				urutkanNamaProyek = true
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = false
				ProyekDicari = false
			} else if pilihanSorting == 2{
				urutkanDefault = false
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = true
				urutkanDanaDibutuhkan = false
				ProyekDicari = false
			}  else if pilihanSorting == 3{
				urutkanDefault = false
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = true
				ProyekDicari = false
			}  else if pilihanSorting == 4{
				urutkanDefault = true
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = false
				ProyekDicari = false
			} else {
				fmt.Println("Pilihan tidak valid.")
			}
		} else if pilihan == 4{
			fmt.Println("== Cari Proyek ==")
			fmt.Println("1. Cari berdasarkan ID")
			fmt.Println("2. Cari berdasarkan Nama")
			fmt.Print("Pilihan: ")
			var metodeCari int
			fmt.Scan(&metodeCari)

			if metodeCari == 1 {
				var dicariID int
				urutkanDefault = false
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = false
				ProyekDicari = true
				fmt.Print("Masukkan ID Proyek yang ingin dicari: ")
				fmt.Scan(&dicariID)

				indexDitemukan := linearSearchIDProyek(daftarProyek, jumlahProyek, dicariID)
				fmt.Println("\n======================= Menu User SimpleFund =======================")
				tampilkanProyekDicari(daftarProyek, indexDitemukan)
				
			} else if metodeCari == 2 {
				var dicariNama string
				urutkanDefault = false
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = false
				ProyekDicari = true
				fmt.Print("Masukkan Nama Proyek yang ingin dicari: ")
				fmt.Scan(&dicariNama)

				indexDitemukan := linearSearchNamaProyek(daftarProyek, jumlahProyek, dicariNama)
				fmt.Println("\n======================= Menu User SimpleFund =======================")
				tampilkanProyekDicari(daftarProyek, indexDitemukan)
			} else {
				fmt.Println("Pilihan tidak valid.")
				urutkanDefault = true
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = false
				ProyekDicari = false
			}
		} else if pilihan == 5 {
			selesai = true
			fmt.Println("Logout berhasil.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// menu admin
func adminMenu() {
	var pilihan int
	var pilihanSorting int
	urutkanDefault := true
	urutkanNamaProyek := false
	urutkanDanaTerkumpul := false
	urutkanDanaDibutuhkan := false
	urutkanDefaultDonatur := false
	urutkanNamaDonatur := false
	urutkanTotalDonasiDonatur := false
	ProyekDicari := false
	DonaturDicari := false

	selesai := false
	i := 1

	for !selesai {
		if i == 1 {
			fmt.Println("\n================== Selamat Datang Admin SimpleFund ==================")
		}
		if !ProyekDicari && !DonaturDicari{
			fmt.Println("============================ Menu Admin ============================")
		}
		if !urutkanDefault && !urutkanNamaProyek && !urutkanDanaTerkumpul && !urutkanDanaDibutuhkan && urutkanDefaultDonatur && !urutkanNamaDonatur && !urutkanTotalDonasiDonatur && !ProyekDicari && !DonaturDicari {
			tampilkanDonaturDefault(daftarDonatur, jumlahDonatur)
		} else if !urutkanDefault && !urutkanNamaProyek && !urutkanDanaTerkumpul && !urutkanDanaDibutuhkan && !urutkanDefaultDonatur && urutkanNamaDonatur && !urutkanTotalDonasiDonatur && !ProyekDicari && !DonaturDicari {
			tampilkanDonaturUrutkanNama(daftarDonatur, jumlahDonatur)
		} else if !urutkanDefault && !urutkanNamaProyek && !urutkanDanaTerkumpul && !urutkanDanaDibutuhkan && !urutkanDefaultDonatur && !urutkanNamaDonatur && urutkanTotalDonasiDonatur && !ProyekDicari && !DonaturDicari {
			tampilkanDonaturUrutkanTotalDonasi(daftarDonatur, jumlahDonatur)
		} else if !urutkanDefault && urutkanNamaProyek && !urutkanDanaTerkumpul && !urutkanDanaDibutuhkan && !urutkanDefaultDonatur && !urutkanNamaDonatur && !urutkanTotalDonasiDonatur && !ProyekDicari && !DonaturDicari {
			tampilkanProyekUrutkanNama(daftarProyek, jumlahProyek)
		} else if !urutkanDefault && !urutkanNamaProyek && urutkanDanaTerkumpul && !urutkanDanaDibutuhkan && !urutkanDefaultDonatur && !urutkanNamaDonatur && !urutkanTotalDonasiDonatur && !ProyekDicari && !DonaturDicari {
			tampilkanProyekUrutkanDanaTerkumpul(daftarProyek, jumlahProyek)
		} else if !urutkanDefault && !urutkanNamaProyek && !urutkanDanaTerkumpul && urutkanDanaDibutuhkan && !urutkanDefaultDonatur && !urutkanNamaDonatur && !urutkanTotalDonasiDonatur && !ProyekDicari && !DonaturDicari {
			tampilkanProyekUrutkanDanaDibutuhkan(daftarProyek, jumlahProyek)
		} else if urutkanDefault && !urutkanNamaProyek && !urutkanDanaTerkumpul && !urutkanDanaDibutuhkan && !urutkanDefaultDonatur && !urutkanNamaDonatur && !urutkanTotalDonasiDonatur && !ProyekDicari && !DonaturDicari {
			tampilkanProyekDefault(daftarProyek, jumlahProyek)
		} else {
			//tidak menampilkan
		}
		
		fmt.Println("\n============================ Pilih Menu ============================")
		fmt.Println("1. Tambah Proyek")
		fmt.Println("2. Tambah Donatur")
		fmt.Println("3. Donasi ke Proyek")
		if (urutkanDefault || urutkanNamaProyek || urutkanDanaTerkumpul || urutkanDanaDibutuhkan) {
			fmt.Println("4. Tampilkan Donatur")
		} else {
			fmt.Println("4. Tampilkan Proyek")
		}
		if (urutkanDefault || urutkanNamaProyek || urutkanDanaTerkumpul || urutkanDanaDibutuhkan) {
			fmt.Println("5. Urutkan Proyek")
		} else {
			fmt.Println("5. Urutkan Donatur")
		}
		fmt.Println("6. Cari Proyek")
		fmt.Println("7. Cari Donatur")
		fmt.Println("8. Edit Proyek")
		fmt.Println("9. Edit Donatur")
		fmt.Println("10. Hapus Proyek")
		fmt.Println("11. Hapus Donatur")
		fmt.Println("12. Logout")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		i++

		if pilihan == 1 {
			tambahProyek(&daftarProyek, &jumlahProyek)
		} else if pilihan == 2 {
			tambahDonatur(&daftarDonatur, &jumlahDonatur)
		} else if pilihan == 3 {
			donasi(&daftarProyek, jumlahProyek, &daftarDonatur, jumlahDonatur)
		//menampilkan proyek jika sebelumnya menampilkan donatur dan sebaliknya
		} else if pilihan == 4 && (urutkanDefault || urutkanNamaProyek || urutkanDanaTerkumpul || urutkanDanaDibutuhkan) {
			urutkanDefault = false
			urutkanNamaProyek = false
			urutkanDanaTerkumpul = false
			urutkanDanaDibutuhkan = false
			urutkanDefaultDonatur = true
			urutkanNamaDonatur = false
			urutkanTotalDonasiDonatur = false
		////menampilkan proyek jika sebelumnya menampilkan donatur dan sebaliknya
		} else if pilihan == 4 && !(urutkanDefault || urutkanNamaProyek || urutkanDanaTerkumpul || urutkanDanaDibutuhkan) {
			urutkanDefault = true
			urutkanNamaProyek = false
			urutkanDanaTerkumpul = false
			urutkanDanaDibutuhkan = false
			urutkanDefaultDonatur = false
			urutkanNamaDonatur = false
			urutkanTotalDonasiDonatur = false
		//mensorting proyek jika sedang menampilkan proyek
		} else if pilihan == 5 && (urutkanDefault || urutkanNamaProyek || urutkanDanaTerkumpul || urutkanDanaDibutuhkan) {
			fmt.Println("\n========== Pilih Jenis Sorting ==========")
			fmt.Println("1. Urutkan berdasarkan nama (Ascending)")
			fmt.Println("2. Urutkan berdasarkan dana terkumpul (Descending)")
			fmt.Println("3. Urutkan berdasarkan dana dibutuhkan (Descending)")
			fmt.Println("4. Kembali ke Urutan Default")
			fmt.Print("Pilih menu: ")
			fmt.Scan(&pilihanSorting)
			if pilihanSorting == 1{
				urutkanDefault = false
				urutkanNamaProyek = true 
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = false
				urutkanDefaultDonatur = false
				urutkanNamaDonatur = false
				urutkanTotalDonasiDonatur = false
				ProyekDicari = false
				DonaturDicari = false
			} else if pilihanSorting == 2{
				urutkanDefault = false
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = true 
				urutkanDanaDibutuhkan = false
				urutkanDefaultDonatur = false
				urutkanNamaDonatur = false
				urutkanTotalDonasiDonatur = false
				ProyekDicari = false
				DonaturDicari = false
			}  else if pilihanSorting == 3{
				urutkanDefault = false
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = true 
				urutkanDefaultDonatur = false
				urutkanNamaDonatur = false
				urutkanTotalDonasiDonatur = false
				ProyekDicari = false
				DonaturDicari = false
			}  else if pilihanSorting == 4{
				urutkanDefault = true 
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = false
				urutkanDefaultDonatur = false
				urutkanNamaDonatur = false
				urutkanTotalDonasiDonatur = false
				ProyekDicari = false
				DonaturDicari = false
			} else {
				fmt.Println("Pilihan tidak valid.")
			}
		//mensorting donatur jika sedang menampilkan donatur
		} else if pilihan == 5 && !(urutkanDefault || urutkanNamaProyek || urutkanDanaTerkumpul || urutkanDanaDibutuhkan) {
			fmt.Println("\n========== Pilih Jenis Sorting ==========")
			fmt.Println("1. Urutkan berdasarkan nama (Ascending)")
			fmt.Println("2. Urutkan berdasarkan total donasi (Descending)")
			fmt.Println("3. Kembali ke Urutan Default")
			fmt.Print("Pilih menu: ")
			fmt.Scan(&pilihanSorting)
			if pilihanSorting == 1 {
				urutkanDefault = false
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = false
				urutkanDefaultDonatur = false
				urutkanNamaDonatur = true
				urutkanTotalDonasiDonatur = false
				ProyekDicari = false
				DonaturDicari = false
			} else if pilihanSorting == 2 {
				urutkanDefault = false
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = false
				urutkanDefaultDonatur = false
				urutkanNamaDonatur = false
				urutkanTotalDonasiDonatur = true
				ProyekDicari = false
				DonaturDicari = false
			}  else if pilihanSorting == 3 {
				urutkanDefault = false
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = false
				urutkanDefaultDonatur = true
				urutkanNamaDonatur = false
				urutkanTotalDonasiDonatur = false
				ProyekDicari = false
				DonaturDicari = false
			}  else {
				fmt.Println("Pilihan tidak valid.")
			}
		} else if pilihan == 6 {
			fmt.Println("== Cari Proyek ==")
			fmt.Println("1. Cari berdasarkan ID")
			fmt.Println("2. Cari berdasarkan Nama")
			fmt.Print("Pilihan: ")
			var metodeCari int
			fmt.Scan(&metodeCari)

			if metodeCari == 1 {
				var dicariID int
				urutkanDefault = false
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = false
				urutkanDefaultDonatur = false
				urutkanNamaDonatur = false
				urutkanTotalDonasiDonatur = false
				ProyekDicari = true
				DonaturDicari = false
				fmt.Print("Masukkan ID Proyek yang ingin dicari: ")
				fmt.Scan(&dicariID)

				indexDitemukan := linearSearchIDProyek(daftarProyek, jumlahProyek, dicariID)
				fmt.Println("============================ Menu Admin ============================")
				tampilkanProyekDicari(daftarProyek, indexDitemukan)
				
			} else if metodeCari == 2 {
				var dicariNama string
				urutkanDefault = false
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = false
				urutkanDefaultDonatur = false
				urutkanNamaDonatur = false
				urutkanTotalDonasiDonatur = false
				ProyekDicari = true
				DonaturDicari = false
				fmt.Print("Masukkan Nama Proyek yang ingin dicari: ")
				fmt.Scan(&dicariNama)

				indexDitemukan := linearSearchNamaProyek(daftarProyek, jumlahProyek, dicariNama)
				fmt.Println("============================ Menu Admin ============================")
				tampilkanProyekDicari(daftarProyek, indexDitemukan)
			} else {
				fmt.Println("Pilihan tidak valid.")
				urutkanDefault = true
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = false
				urutkanDefaultDonatur = false
				urutkanNamaDonatur = false
				urutkanTotalDonasiDonatur = false
				ProyekDicari = false
				DonaturDicari = false
			}
		} else if pilihan == 7 {
			fmt.Println("== Cari Donatur ==")
			fmt.Println("1. Cari berdasarkan ID")
			fmt.Println("2. Cari berdasarkan Nama")
			fmt.Print("Pilihan: ")

			var metodeCari int
			fmt.Scan(&metodeCari)

			if metodeCari == 1 {
				var dicariID int
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = false
				urutkanDefaultDonatur = false
				urutkanNamaDonatur = false
				urutkanTotalDonasiDonatur = false
				ProyekDicari = false
				DonaturDicari = true
				fmt.Print("Masukkan ID Donatur yang ingin dicari: ")
				fmt.Scan(&dicariID)

				indexDitemukan := binarySearchIDDonatur(daftarDonatur, jumlahDonatur, dicariID)
				fmt.Println("============================ Menu Admin ============================")
				tampilkanDonaturDicari(daftarDonatur, indexDitemukan)

			} else if metodeCari == 2 {
				var dicariNama string
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = false
				urutkanDefaultDonatur = false
				urutkanNamaDonatur = false
				urutkanTotalDonasiDonatur = false
				ProyekDicari = false
				DonaturDicari = true
				fmt.Print("Masukkan Nama Donatur yang ingin dicari: ")
				fmt.Scan(&dicariNama)

				indexDitemukan := binarySearchNamaDonatur(daftarDonatur, jumlahDonatur, dicariNama)
				fmt.Println("============================ Menu Admin ============================")
				tampilkanDonaturDicari(daftarDonatur, indexDitemukan)
			} else {
				fmt.Println("Pilihan tidak valid.")
				urutkanNamaProyek = false
				urutkanDanaTerkumpul = false
				urutkanDanaDibutuhkan = false
				urutkanDefaultDonatur = true
				urutkanNamaDonatur = false
				urutkanTotalDonasiDonatur = false
				ProyekDicari = false
				DonaturDicari = false
			}
		} else if pilihan == 8 {
			editProyek(&daftarProyek, jumlahProyek)
		} else if pilihan == 9 {
			editDonatur(&daftarDonatur, jumlahDonatur)
		} else if pilihan == 10 {
			hapusProyek(&daftarProyek, &jumlahProyek)
		} else if pilihan == 11 {
			hapusDonatur(&daftarDonatur, &jumlahDonatur)
		} else if pilihan == 12 {
			selesai = true
			fmt.Println("Logout berhasil.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

/*list 
	linearsearch : linearSearchNamaProyek, linearSearchIDProyek
	binarysearch : binarySearchNamaDonatur, binarySearchIDDonatur
	insertionSort : insertionSortNamaProyekASC, insertionSortDanaTerkumpulProyekDSC, insertionSortDanaDibutuhkanProyekDSC
	selectionSort : selectionSortNamaDonaturASC, selectionSortTotalDonasiDonaturDSC, 
*/

func linearSearchNamaProyek(data [MAX]Proyek, n int, dicari string) int {
    for i := 0; i < n; i++ {
        if data[i].nama == dicari{
            return i
        }
    }
    return -1
}

func linearSearchIDProyek(data [MAX]Proyek, n int, dicari int) int {
    for i := 0; i < n; i++ {
        if data[i].id == dicari{
            return i
        }
    }
    return -1
}

func binarySearchNamaDonatur(data [MAX]Donatur, n int, dicari string) int {
    kiri := 0
    kanan := n - 1

    for kiri <= kanan {
        tengah := (kiri + kanan) / 2

        if data[tengah].nama == dicari {
            return tengah
        } else {
            if data[tengah].nama < dicari {
                kiri = tengah + 1
            } else {
                kanan = tengah - 1
            }
        }
    }

    return -1
}

func binarySearchIDDonatur(data [MAX]Donatur, n int, dicari int) int {
    kiri := 0
    kanan := n - 1

    for kiri <= kanan {
        tengah := (kiri + kanan) / 2

        if data[tengah].id == dicari {
            return tengah
        } else {
            if data[tengah].id < dicari {
                kiri = tengah + 1
            } else {
                kanan = tengah - 1
            }
        }
    }

    return -1
}


// Fungsi insertion sort untuk proyek berdasarkan nama (ascending)
func insertionSortNamaProyekASC(data *[MAX]Proyek, jumlah int) {
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

// Fungsi insertion sort untuk Dana terkumpul proyek (descending)
func insertionSortDanaTerkumpulProyekDSC(data *[MAX]Proyek, jumlah int) {
	for i := 1; i < jumlah; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].danaTerkumpul < key.danaTerkumpul {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}



// Fungsi insertion sort untuk Dana dibutuhkan proyek (descending)
func insertionSortDanaDibutuhkanProyekDSC(data *[MAX]Proyek, jumlah int) {
	for i := 1; i < jumlah; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].danaDibutuhkan < key.danaDibutuhkan {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

// Fungsi selection sort untuk donatur berdasarkan totalDonasi (descending)
func selectionSortTotalDonasiDonaturDSC(data *[MAX]Donatur, jumlah int) {
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

// Fungsi selection sort untuk donatur berdasarkan nama (ascending)
func selectionSortNamaDonaturASC(data *[MAX]Donatur, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlah; j++ {
			if data[j].nama < data[minIdx].nama {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
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

//tampilkan proyek dicari
func tampilkanProyekDicari(data [MAX]Proyek, indexDitemukan int) {
	fmt.Println("\n🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦")
	if indexDitemukan == -1 {
		fmt.Println("============================= Proyek Yang DiCari Tidak Ditemukan ============================= ")
	} else {
		fmt.Println("============================= Proyek Ditemukan ============================= ")
	}
	fmt.Printf("%-4s | %-10s | %-20s | %-15s | %-15s | %-30s\n", "No", "ID Proyek", "Nama Proyek", "Dana Terkumpul", "Dana Dibutuhkan", "Status Proyek")
	fmt.Println("----------------------------------------------------------------------------------------------------------------")

	if indexDitemukan == -1 {
		fmt.Printf("%-4s | %-10s | %-20s | %-15s | %-15s | %-30s\n", "-", "-", "-", "-", "-", "-")
	} else {
		status := ""
		if data[indexDitemukan].danaTerkumpul >= data[indexDitemukan].danaDibutuhkan {
			status = "Dana Sudah Mencukupi"
			fmt.Printf("%-4d | %-10d | %-20s | Rp%-13d | Rp%-13d | %s\n", indexDitemukan+1, data[indexDitemukan].id,data[indexDitemukan].nama, data[indexDitemukan]. danaTerkumpul, data[indexDitemukan].danaDibutuhkan, status)
		} else {
			status = "Kurang Rp"
			sisaKebutuhan := data[indexDitemukan].danaDibutuhkan - data[indexDitemukan].danaTerkumpul
			fmt.Printf("%-4d | %-10d | %-20s | Rp%-13d | Rp%-13d | %s%d\n", indexDitemukan+1, data[indexDitemukan].id,data[indexDitemukan].nama, data[indexDitemukan]. danaTerkumpul, data[indexDitemukan].danaDibutuhkan, status, sisaKebutuhan)
		}
	}
	fmt.Println("🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦")
}

//tampilkan proyek default
func tampilkanProyekDefault(data [MAX]Proyek, jumlah int) {
	fmt.Println("\n🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦")
	fmt.Println("============================= Daftar Proyek ============================= ")
	fmt.Printf("%-4s | %-10s | %-20s | %-15s | %-15s | %-30s\n", "No", "ID Proyek", "Nama Proyek", "Dana Terkumpul", "Dana Dibutuhkan", "Status Proyek")
	fmt.Println("----------------------------------------------------------------------------------------------------------------")

	if jumlah == 0 {
		fmt.Printf("%-4s | %-10s | %-20s | %-15s | %-15s | %-30s\n", "-", "-", "-", "-", "-", "-")
	}

	for i := 0; i < jumlah; i++ {
		status := ""
		if data[i].danaTerkumpul >= data[i].danaDibutuhkan {
			status = "Dana Sudah Mencukupi"
			fmt.Printf("%-4d | %-10d | %-20s | Rp%-13d | Rp%-13d | %s\n", i+1, data[i].id,data[i].nama, data[i]. danaTerkumpul, data[i].danaDibutuhkan, status)
		} else {
			status = "Kurang Rp"
			sisaKebutuhan := data[i].danaDibutuhkan - data[i].danaTerkumpul
			fmt.Printf("%-4d | %-10d | %-20s | Rp%-13d | Rp%-13d | %s%d\n", i+1, data[i].id,data[i].nama, data[i]. danaTerkumpul, data[i].danaDibutuhkan, status, sisaKebutuhan)
		}
	}
	fmt.Println("🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦")
}

//tampilkan proyek udah diurut nama
func tampilkanProyekUrutkanNama(data [MAX]Proyek, jumlah int) {
	insertionSortNamaProyekASC(&data, jumlah)
	fmt.Println("\n🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦")
	fmt.Println("========================= Daftar Proyek (Ascending Nama) ===========================")
	fmt.Printf("%-4s | %-10s | %-20s | %-15s | %-15s | %-30s\n", "No", "ID Proyek", "Nama Proyek", "Dana Terkumpul", "Dana Dibutuhkan", "Status Proyek")
	fmt.Println("----------------------------------------------------------------------------------------------------------------")

	if jumlah == 0 {
		fmt.Printf("%-4s | %-10s | %-20s | %-15s | %-15s | %-30s\n", "-", "-", "-", "-", "-", "-")
	}

	for i := 0; i < jumlah; i++ {
		status := ""
		if data[i].danaTerkumpul >= data[i].danaDibutuhkan {
			status = "Dana Sudah Mencukupi"
			fmt.Printf("%-4d | %-10d | %-20s | Rp%-13d | Rp%-13d | %s\n", i+1, data[i].id,data[i].nama, data[i]. danaTerkumpul, data[i].danaDibutuhkan, status)
		} else {
			status = "Kurang Rp"
			sisaKebutuhan := data[i].danaDibutuhkan - data[i].danaTerkumpul
			fmt.Printf("%-4d | %-10d | %-20s | Rp%-13d | Rp%-13d | %s%d\n", i+1, data[i].id,data[i].nama, data[i]. danaTerkumpul, data[i].danaDibutuhkan, status, sisaKebutuhan)
		}
	}
	fmt.Println("🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦")
}

//tampilkan proyek udah urut dana terkumpul
func tampilkanProyekUrutkanDanaTerkumpul(data [MAX]Proyek, jumlah int) {
	insertionSortDanaTerkumpulProyekDSC(&data, jumlah)
	fmt.Println("\n🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦")
	fmt.Println("========================= Daftar Proyek (Descending Dana Terkumpul) ===========================")
	fmt.Printf("%-4s | %-10s | %-20s | %-15s | %-15s | %-30s\n", "No", "ID Proyek", "Nama Proyek", "Dana Terkumpul", "Dana Dibutuhkan", "Status Proyek")
	fmt.Println("----------------------------------------------------------------------------------------------------------------")

	if jumlah == 0 {
		fmt.Printf("%-4s | %-10s | %-20s | %-15s | %-15s | %-30s\n", "-", "-", "-", "-", "-", "-")
	}

	for i := 0; i < jumlah; i++ {
		status := ""
		if data[i].danaTerkumpul >= data[i].danaDibutuhkan {
			status = "Dana Sudah Mencukupi"
			fmt.Printf("%-4d | %-10d | %-20s | Rp%-13d | Rp%-13d | %s\n", i+1, data[i].id,data[i].nama, data[i]. danaTerkumpul, data[i].danaDibutuhkan, status)
		} else {
			status = "Kurang Rp"
			sisaKebutuhan := data[i].danaDibutuhkan - data[i].danaTerkumpul
			fmt.Printf("%-4d | %-10d | %-20s | Rp%-13d | Rp%-13d | %s%d\n", i+1, data[i].id,data[i].nama, data[i]. danaTerkumpul, data[i].danaDibutuhkan, status, sisaKebutuhan)
		}
	}
	fmt.Println("🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦")
}

//tampilkan proyek udah urut dana dibutuhkan
func tampilkanProyekUrutkanDanaDibutuhkan(data [MAX]Proyek, jumlah int) {
	insertionSortDanaDibutuhkanProyekDSC(&data, jumlah)
	fmt.Println("\n🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦")
	fmt.Println("========================= Daftar Proyek (Descending Dana Dibutuhkan) ===========================")
	fmt.Printf("%-4s | %-10s | %-20s | %-15s | %-15s | %-30s\n", "No", "ID Proyek", "Nama Proyek", "Dana Terkumpul", "Dana Dibutuhkan", "Status Proyek")
	fmt.Println("----------------------------------------------------------------------------------------------------------------")

	if jumlah == 0 {
		fmt.Printf("%-4s | %-10s | %-20s | %-15s | %-15s | %-30s\n", "-", "-", "-", "-", "-", "-")
	}

	for i := 0; i < jumlah; i++ {
		status := ""
		if data[i].danaTerkumpul >= data[i].danaDibutuhkan {
			status = "Dana Sudah Mencukupi"
			fmt.Printf("%-4d | %-10d | %-20s | Rp%-13d | Rp%-13d | %s\n", i+1, data[i].id,data[i].nama, data[i]. danaTerkumpul, data[i].danaDibutuhkan, status)
		} else {
			status = "Kurang Rp"
			sisaKebutuhan := data[i].danaDibutuhkan - data[i].danaTerkumpul
			fmt.Printf("%-4d | %-10d | %-20s | Rp%-13d | Rp%-13d | %s%d\n", i+1, data[i].id,data[i].nama, data[i]. danaTerkumpul, data[i].danaDibutuhkan, status, sisaKebutuhan)
		}
	}
	fmt.Println("🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦")
}

//tampilkan donatur yang dicari
func tampilkanDonaturDicari(data [MAX]Donatur, indexDitemukan int) {
	fmt.Println("\nMenampilkan Donatur")
	fmt.Println("🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩")
	if indexDitemukan == -1{
		fmt.Println("========================= Donatur Yang Dicari Tidak Ditemukan ========================= ")
	} else {
		fmt.Println("========================= Donatur Ditemukan ========================= ")
	}
	fmt.Printf("%-10s | %-20s | %-25s\n", "ID Donatur", "Nama Donatur", "Total Donasi")
	fmt.Println("------------------------------------------------------------------------------")

	if indexDitemukan == -1 {
		fmt.Printf("%-10s | %-20s | %-25s\n", "-", "-", "-")
	} else {
		fmt.Printf("%-10d | %-20s | %-25d\n", data[indexDitemukan].id, data[indexDitemukan].nama,  data[indexDitemukan].totalDonasi)
	}
	fmt.Println("🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩")
}

func tampilkanDonaturDefault(data [MAX]Donatur, jumlah int) {
	fmt.Println("\nMenampilkan Donatur")
	fmt.Println("🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩")
	fmt.Println("========================= Daftar Donatur (Descending Total Donasi) ========================= ")
	fmt.Printf("%-10s | %-20s | %-25s\n", "ID Donatur", "Nama Donatur", "Total Donasi")
	fmt.Println("------------------------------------------------------------------------------")

	if jumlah == 0 {
		fmt.Printf("%-10s | %-20s | %-25s\n", "-", "-", "-")
	}

	for i := 0; i < jumlah; i++ {
		fmt.Printf("%-10d | %-20s | %-25d\n", data[i].id, data[i].nama,  data[i].totalDonasi)
	}
	fmt.Println("🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩")
}

func tampilkanDonaturUrutkanNama(data [MAX]Donatur, jumlah int) {
	fmt.Println("\nMenampilkan Donatur Berurutkan Nama")
	fmt.Println("🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩")
	fmt.Println("========================= Daftar Donatur (Descending Total Donasi) ========================= ")
	fmt.Printf("%-10s | %-20s | %-25s\n", "ID Donatur", "Nama Donatur", "Total Donasi")
	fmt.Println("------------------------------------------------------------------------------")

	if jumlah == 0 {
		fmt.Printf("%-10s | %-20s | %-25s\n", "-", "-", "-")
	}

	for i := 0; i < jumlah; i++ {
		fmt.Printf("%-10d | %-20s | %-25d\n", data[i].id, data[i].nama,  data[i].totalDonasi)
	}
	fmt.Println("🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩")
}

func tampilkanDonaturUrutkanTotalDonasi(data [MAX]Donatur, jumlah int) {
	fmt.Println("\nMenampilkan Donatur Berurutkan Total Donasi")
	fmt.Println("🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩")
	fmt.Println("========================= Daftar Donatur (Descending Total Donasi) ========================= ")
	fmt.Printf("%-10s | %-20s | %-25s\n", "ID Donatur", "Nama Donatur", "Total Donasi")
	fmt.Println("------------------------------------------------------------------------------")

	if jumlah == 0 {
		fmt.Printf("%-10s | %-20s | %-25s\n", "-", "-", "-")
	}

	for i := 0; i < jumlah; i++ {
		fmt.Printf("%-10d | %-20s | %-25d\n", data[i].id, data[i].nama,  data[i].totalDonasi)
	}
	fmt.Println("🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩")
}

func tambahProyek(data *[MAX]Proyek, jumlah *int){
    var n int
    fmt.Print("Berapa proyek yang ingin ditambahkan? ")
    fmt.Scan(&n)
    
    for i := 0; i < n && *jumlah < MAX; i++ {
        var p Proyek
        fmt.Println("\n--- Tambah Proyek ---")

        // Input ID
        ulang := true
        for ulang {
            fmt.Print("Masukkan ID Proyek: ")
            fmt.Scan(&p.id)
            duplikat := false
            
            for j := 0; j < *jumlah && !duplikat; j++ {
                if data[j].id == p.id {
                    fmt.Println("ID proyek sudah digunakan. Silakan masukkan ID lain.")
                    duplikat = true
                }
            }
            
            ulang = duplikat
        }

        fmt.Print("Nama Proyek: ")
        fmt.Scan(&p.nama)
        fmt.Print("Dana Dibutuhkan: ")
        fmt.Scan(&p.danaDibutuhkan)

        p.danaTerkumpul = 0
        data[*jumlah] = p
        *jumlah++
        fmt.Println("\n========================= Proyek berhasil ditambahkan =========================")
    }
    
    if *jumlah >= MAX {
        fmt.Println("\nKapasitas proyek penuh.")
    }
}

func tambahDonatur(data *[MAX]Donatur, jumlah *int){
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
			
			for j := 0; j < *jumlah && !duplikat; j++ {
				if data[j].id == d.id {
					fmt.Println("ID donatur sudah digunakan. Silakan masukkan ID lain.")
					duplikat = true
				}
			}
			
			ulang = duplikat
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

// fungsi donasi khusus untuk user biasa
func donasiUser(proyek *[MAX]Proyek, jumlahProyek int, donatur *[MAX]Donatur, jumlahDonatur *int) {
    var idProyek, nominal int
    fmt.Println("\n--- Donasi ke Proyek ---")
    fmt.Print("Masukkan ID proyek: ")
    fmt.Scan(&idProyek)
    fmt.Print("Masukkan jumlah donasi: ")
    fmt.Scan(&nominal)

    if nominal <= 0 {
        fmt.Println("Nominal donasi harus lebih dari 0.")
        return
    }

    indeksProyek := cariProyekByID(idProyek, *proyek, jumlahProyek)
    if indeksProyek == -1 {
        fmt.Println("ID proyek tidak ditemukan.")
        return
    }
    
    // cari donatur berdasarkan username user yang sedang login
    indeksDonatur := -1
    i := 0
    ditemukanDonatur := false
    
    for i < *jumlahDonatur && !ditemukanDonatur {
        // cari donatur dengan nama yang sama dengan username user
        if donatur[i].nama == currentUser.username {
            indeksDonatur = i
            ditemukanDonatur = true
        }
        i++
    }
    
    // jika donatur belum ada, buat donatur baru secara otomatis
    if indeksDonatur == -1 {
        if *jumlahDonatur >= MAX {
            fmt.Println("Kapasitas donatur penuh. Hubungi admin untuk bantuan.")
            return
        }
        
        // ID donatur pake ID user
        var newID int = currentUser.id
        
        // buat donatur baru dengan nama sama dengan username
        var d Donatur
        d.id = newID
        d.nama = currentUser.username
        d.totalDonasi = 0
        
        // menambahkan ke daftar donatur
        donatur[*jumlahDonatur] = d
        indeksDonatur = *jumlahDonatur
        *jumlahDonatur++
    }

    sisaKebutuhan := proyek[indeksProyek].danaDibutuhkan - proyek[indeksProyek].danaTerkumpul
    if nominal > sisaKebutuhan {
        fmt.Println("Donasi melebihi kebutuhan proyek. Maksimal yang bisa didonasikan:", sisaKebutuhan)
        return
    }

    proyek[indeksProyek].danaTerkumpul += nominal
    donatur[indeksDonatur].totalDonasi += nominal
    fmt.Println("Donasi berhasil ditambahkan.")
    fmt.Printf("Anda (%s) dengan ID %d telah mendonasikan Rp%d ke proyek %s.\n", 
               currentUser.username, currentUser.id, nominal, proyek[indeksProyek].nama)
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