package main
import (
"fmt"
"os"
"os/exec"
)

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

var systemMessage string

// otentikasi tipe user
const (
	ADMIN = 0
	USER  = 1
)

type User struct {
	id int
	username string
	password string
	userType int // admin = 0, user = 1
}

var daftarUser [MAX]User
var jumlahUser int = 0
var currentUser User
var nextUserID int = 1 // variabel untuk menyimpan ID berikutnya yang tersedia

// inisialisasi admin
func initAdmin() {
	adminUser := User{
		id:       0,     // admin selalu memiliki ID 0
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
			systemMessage = fmt.Sprintf("Login berhasil! Selamat datang %s (ID: %d)\n", currentUser.username, currentUser.id)
			return true
		}
	}
	systemMessage = "Username atau password salah!"
	return false
}

// register user baru
func register() {
	if jumlahUser >= MAX {
		systemMessage = "Kapasitas pengguna penuh!"
		return
	}

	var newUser User
	fmt.Println("\n======================== Register SimpleFund ========================")
	
	// cek keunikan username
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
	
	systemMessage = fmt.Sprintf("Registrasi berhasil! Anda telah terdaftar dengan ID: %d\nSilahkan login dengan username dan password Anda.", newUser.id)
}

// menu user biasa
func userMenu() {
	var pilihan int
	selesai := false

	for !selesai {
		clear()
		fmt.Println("\n======================= Menu User SimpleFund =======================")
		tampilkanProyek(daftarProyek, jumlahProyek)
		fmt.Println("\n============================ Pilih Menu ============================")
		fmt.Println("1. Tambah Proyek")
		fmt.Println("2. Donasi ke Proyek")
		fmt.Println("3. Logout")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahProyek(&daftarProyek, &jumlahProyek)
		} else if pilihan == 2 {
			donasiUser(&daftarProyek, jumlahProyek, &daftarDonatur, &jumlahDonatur)
		} else if pilihan == 3 {
			selesai = true
			systemMessage = "Logout berhasil."
		} else {
			systemMessage = "Pilihan tidak valid."
		}
	}
}

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
	fmt.Println("\n🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦")
	fmt.Println("========================= Daftar Proyek (Ascending Nama) ===========================")
	fmt.Printf("%-4s | %-10s | %-20s | %-15s | %-12s | %-30s\n", "No", "ID Proyek", "Nama Proyek", "Dana Terkumpul", "Target Dana", "Status Proyek")
	fmt.Println("----------------------------------------------------------------------------------------------------------------")

	if jumlah == 0 {
		fmt.Printf("%-4s | %-10s | %-20s | %-15s | %-12s | %-25s\n", "-", "-", "-", "-", "-", "-")
	}

	for i := 0; i < jumlah; i++ {
		status := ""
		if data[i].danaTerkumpul >= data[i].danaDibutuhkan {
			status = "Dana Sudah Mencukupi"
			fmt.Printf("%-4d | %-10d | %-20s | Rp%-13d | Rp%-10d | %s\n", i+1, data[i].id,data[i].nama, data[i]. danaTerkumpul, data[i].danaDibutuhkan, status)
		} else {
			status = "Kurang Rp"
			sisaKebutuhan := data[i].danaDibutuhkan - data[i].danaTerkumpul
			fmt.Printf("%-4d | %-10d | %-20s | Rp%-13d | Rp%-10d | %s%d\n", i+1, data[i].id,data[i].nama, data[i]. danaTerkumpul, data[i].danaDibutuhkan, status, sisaKebutuhan)
		}
	}
	fmt.Println("🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦🟦")
}


func tampilkanDonatur(data [MAX]Donatur, jumlah int) {
	var pilihan int
	selectionSortDonatur(&data, jumlah)
	ulang := true
	for ulang{
		clear()
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
		fmt.Println("\n==================================== Pilih Menu ====================================")
		fmt.Println("1. Tambah Proyek")
		fmt.Println("2. Tambah Donatur")
		fmt.Println("3. Donasi ke Proyek")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Println("5. Edit Proyek")
		fmt.Println("6. Hapus Proyek")
		fmt.Println("7. Edit Donatur")
		fmt.Println("8. Hapus Donatur")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			ulang = false
			tambahProyek(&daftarProyek, &jumlahProyek)
		} else if pilihan == 2 {
			ulang = false
			tambahDonatur(&daftarDonatur, &jumlahDonatur)
		} else if pilihan == 3 {
			ulang = false
			donasi(&daftarProyek, jumlahProyek, &daftarDonatur, jumlahDonatur)
		} else if pilihan == 4 {
			ulang = false
		} else if pilihan == 5 {
			ulang = false
			editProyek(&daftarProyek, jumlahProyek)
		} else if pilihan == 6 {
			ulang = false
			hapusProyek(&daftarProyek, &jumlahProyek)
		} else if pilihan == 7 {
			ulang = false
			editDonatur(&daftarDonatur, jumlahDonatur)
		} else if pilihan == 8 {
			ulang = false
			hapusDonatur(&daftarDonatur, &jumlahDonatur)
		}  else {
			systemMessage = "Pilihan tidak valid."
		}		
	}
	fmt.Println("")
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
        systemMessage = "\n========================= Proyek berhasil ditambahkan ========================="
    }
    
    if *jumlah >= MAX {
        systemMessage = "\nKapasitas proyek penuh."
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
		systemMessage = "Donatur berhasil ditambahkan."
	}
	
	if *jumlah >= MAX {
		systemMessage = "Kapasitas donatur penuh."
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
		systemMessage = "Nominal donasi harus lebih dari 0."
		return
	}

	indeksProyek := cariProyekByID(idProyek, *proyek, jumlahProyek)
	indeksDonatur := cariDonaturByID(idDonatur, *donatur, jumlahDonatur)

	if indeksProyek == -1 || indeksDonatur == -1 {
		systemMessage = "ID proyek atau donatur tidak ditemukan."
		return
	}

	sisaKebutuhan := proyek[indeksProyek].danaDibutuhkan - proyek[indeksProyek].danaTerkumpul
	if nominal > sisaKebutuhan {
		systemMessage = fmt.Sprintf("Donasi melebihi kebutuhan proyek. Maksimal yang bisa didonasikan:", sisaKebutuhan)
		return
	}

	proyek[indeksProyek].danaTerkumpul += nominal
	donatur[indeksDonatur].totalDonasi += nominal
	systemMessage = "Donasi berhasil ditambahkan."
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
        systemMessage = "Nominal donasi harus lebih dari 0."
        return
    }

    indeksProyek := cariProyekByID(idProyek, *proyek, jumlahProyek)
    if indeksProyek == -1 {
        systemMessage = "ID proyek tidak ditemukan."
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
            systemMessage = "Kapasitas donatur penuh. Hubungi admin untuk bantuan."
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
        systemMessage = fmt.Sprintf("Donasi melebihi kebutuhan proyek. Maksimal yang bisa didonasikan:", sisaKebutuhan)
        return
    }

    proyek[indeksProyek].danaTerkumpul += nominal
    donatur[indeksDonatur].totalDonasi += nominal
    systemMessage = fmt.Sprintf("Donasi berhasil ditambahkan.\nAnda (%s) dengan ID %d telah mendonasikan Rp%d ke proyek %s.\n", 
               currentUser.username, currentUser.id, nominal, proyek[indeksProyek].nama)
}

func editProyek(data *[MAX]Proyek, jumlah int) {
	var id int
	fmt.Print("Masukkan ID proyek yang ingin diedit: ")
	fmt.Scan(&id)
	indeks := cariProyekByID(id, *data, jumlah)

	if indeks == -1 {
		systemMessage = "Proyek dengan ID tersebut tidak ditemukan."
		return
	}

	fmt.Println("Proyek ditemukan. Masukkan data baru:")
	fmt.Print("Nama Proyek: ")
	fmt.Scan(&data[indeks].nama)
	fmt.Print("Dana Dibutuhkan: ")
	fmt.Scan(&data[indeks].danaDibutuhkan)
	systemMessage = "Data proyek berhasil diperbarui."
}

func hapusProyek(data *[MAX]Proyek, jumlah *int) {
	var id int
	fmt.Print("Masukkan ID proyek yang ingin dihapus: ")
	fmt.Scan(&id)
	indeks := cariProyekByID(id, *data, *jumlah)

	if indeks == -1 {
		systemMessage = "Proyek dengan ID tersebut tidak ditemukan."
		return
	}

	for i := indeks; i < *jumlah-1; i++ {
		data[i] = data[i+1]
	}
	*jumlah--
	systemMessage = "Proyek berhasil dihapus."
}

func editDonatur(data *[MAX]Donatur, jumlah int) {
	var id int
	fmt.Print("Masukkan ID donatur yang ingin diedit: ")
	fmt.Scan(&id)
	indeks := cariDonaturByID(id, *data, jumlah)

	if indeks == -1 {
		systemMessage = "Donatur dengan ID tersebut tidak ditemukan."
		return
	}

	fmt.Println("Donatur ditemukan. Masukkan data baru:")
	fmt.Print("Nama Donatur: ")
	fmt.Scan(&data[indeks].nama)
	systemMessage = "Data donatur berhasil diperbarui."
}

func hapusDonatur(data *[MAX]Donatur, jumlah *int) {
	var id int
	fmt.Print("Masukkan ID donatur yang ingin dihapus: ")
	fmt.Scan(&id)
	indeks := cariDonaturByID(id, *data, *jumlah)

	if indeks == -1 {
		systemMessage = "Donatur dengan ID tersebut tidak ditemukan."
		return
	}

	for i := indeks; i < *jumlah-1; i++ {
		data[i] = data[i+1]
	}
	*jumlah--
	systemMessage = "Donatur berhasil dihapus."
}

func main() {
	//login admin
	initAdmin()
	
	var pilihan int
	loggedIn := false
	aplikasiAktif := true
	
	for aplikasiAktif {
		if !loggedIn {
			clear()
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
				systemMessage = "Pilihan tidak valid."
			}
		} else {
			//biar masuk ke menu admin
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

// menu admin
func adminMenu() {
	var pilihan int
	selesai := false
	i := 1

	for !selesai {
		clear()
		if i == 1 {
			fmt.Println("\n================== Selamat Datang Admin SimpleFund ==================")
		}
		fmt.Println("============================ Menu Admin ============================")
		tampilkanProyek(daftarProyek, jumlahProyek)
		fmt.Println("\n============================ Pilih Menu ============================")
		fmt.Println("1. Tambah Proyek")
		fmt.Println("2. Tambah Donatur")
		fmt.Println("3. Donasi ke Proyek")
		fmt.Println("4. Tampilkan Donatur")
		fmt.Println("5. Edit Proyek")
		fmt.Println("6. Hapus Proyek")
		fmt.Println("7. Edit Donatur")
		fmt.Println("8. Hapus Donatur")
		fmt.Println("9. Logout")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		i++

		if pilihan == 1 {
			tambahProyek(&daftarProyek, &jumlahProyek)
		} else if pilihan == 2 {
			tambahDonatur(&daftarDonatur, &jumlahDonatur)
		} else if pilihan == 3 {
			donasi(&daftarProyek, jumlahProyek, &daftarDonatur, jumlahDonatur)
		} else if pilihan == 4 {
			tampilkanDonatur(daftarDonatur, jumlahDonatur)
		} else if pilihan == 5 {
			editProyek(&daftarProyek, jumlahProyek)
		} else if pilihan == 6 {
			hapusProyek(&daftarProyek, &jumlahProyek)
		} else if pilihan == 7 {
			editDonatur(&daftarDonatur, jumlahDonatur)
		} else if pilihan == 8 {
			hapusDonatur(&daftarDonatur, &jumlahDonatur)
		} else if pilihan == 9 {
			selesai = true
			systemMessage = "Logout berhasil."
		} else {
			systemMessage = "Pilihan tidak valid."
		}
	}
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	
	if systemMessage != ""{
		fmt.Println(systemMessage)
		systemMessage = ""
	}
}