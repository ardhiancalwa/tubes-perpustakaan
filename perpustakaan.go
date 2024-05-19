package main

import (
	"fmt"
)

const MAXBUKU int = 1000

type admin struct {
	username string
	email    string
	password string
}

type buku struct {
	isbn               string
	judul              string
	penulis            string
	penerbit           string
	tahunTerbit        int
	jumlahHalaman      int
	status             bool
	jumlahBukuDipinjam int
}

type pinjamBuku struct {
	isbn         string
	idPinjamBuku int
	idAnggota    int
	tgl_pinjam   int
	tgl_kembali  int
	jumlahPinjam int
	denda        float64
	status       bool
}

var jumlahDataAdmin, jumlahDataBuku, jumlahDataPinjam int

type dataAdmin [MAXBUKU]admin
type dataBuku [MAXBUKU]buku
type dataPinjamBuku [MAXBUKU]pinjamBuku

func ViewBukuMenu() {
	// Membersihkan layar sebelum menampilkan menu
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul menu dengan garis pembatas
	fmt.Println("===============================")
	fmt.Println("|          Menu Buku          |")
	fmt.Println("===============================")
	fmt.Println("\n1. Tampilkan Daftar Buku")
	fmt.Println("2. Tambah Data Buku")
	fmt.Println("3. Ubah Data Buku")
	fmt.Println("4. Hapus Data Buku")
	fmt.Println("5. Pencarian Data Buku")
	fmt.Println("6. Kembali")
	fmt.Print("Masukkan pilihan (1/2/3/4/5/6): ")
}

func ViewDeleteDataBukuMenu() {
	// Membersihkan layar sebelum menampilkan menu
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul menu dengan garis pembatas
	fmt.Println("================================")
	fmt.Println("|     Menu Hapus Data Buku     |")
	fmt.Println("================================")
	fmt.Println("\n1. Hapus Data Buku Berdasarkan ISBN")
	fmt.Println("2. Hapus Data Buku Berdasarkan Judul")
	fmt.Println("3. Hapus Data Buku Berdasarkan Penulis")
	fmt.Println("4. Hapus Data Buku Berdasarkan Tahun Terbit")
	fmt.Println("5. Hapus Semua Data Buku")
	fmt.Println("6. Kembali ke Menu Utama")
	fmt.Print("Masukkan Pilihan (1/2/3/4/5/6): ")
}

func ViewShowDataBukuMenu() {
	// Membersihkan layar sebelum menampilkan menu
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul menu dengan garis pembatas
	fmt.Println("================================")
	fmt.Println("|     Menu Cetak Data Buku     |")
	fmt.Println("================================")
	fmt.Println("\n1. Tampilkan Semua Data Buku")
	fmt.Println("2. Tampilkan Data Buku Berdasarkan Kategori")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Masukkan Pilihan (1/2/3): ")
}

func ViewShowDataBukuByKategoriMenu() {
	// Membersihkan layar sebelum menampilkan menu
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul menu dengan garis pembatas
	fmt.Println("=====================================================")
	fmt.Println("|     Menu Cetak Data Buku Berdasarkan Kategori     |")
	fmt.Println("=====================================================")
	fmt.Println("\n1. Tampilkan Data Buku Berdasarkan Judul")
	fmt.Println("2. Tampilkan Data Buku Berdasarkan Penulis")
	fmt.Println("3. Tampilkan Data Buku Berdasarkan Penerbit")
	fmt.Println("4. Tampilkan Data Buku Berdasarkan Tahun Terbit")
	fmt.Println("5. Kembali ke Menu Utama")
	fmt.Print("Masukkan Pilihan (1/2/3/4/5): ")
}

func ViewCreateDataBuku() {
	// Membersihkan layar sebelum menampilkan form tambah data
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul form tambah data
	fmt.Println("===============================")
	fmt.Println("|    TAMBAH DATA BUKU BARU    |")
	fmt.Println("===============================")
}

func ViewUpdateDataBuku() {
	// Membersihkan layar sebelum menampilkan form tambah data
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul form tambah data
	fmt.Println("===============================")
	fmt.Println("|       UBAH DATA BUKU        |")
	fmt.Println("===============================")
}

func ViewDeleteDataBuku() {
	// Membersihkan layar sebelum menampilkan form tambah data
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul form tambah data
	fmt.Println("===============================")
	fmt.Println("|        HAPUS DATA BUKU      |")
	fmt.Println("===============================")
}

func ViewSearchDataBuku() {
	// Membersihkan layar sebelum menampilkan form tambah data
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul form tambah data
	fmt.Println("===============================")
	fmt.Println("|       CARI DATA BUKU        |")
	fmt.Println("===============================")
}

func ViewShowDataBuku() {
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul daftar buku
	fmt.Println("===============================")
	fmt.Println("|         DAFTAR BUKU         |")
	fmt.Println("===============================")
}

func ViewAdminMenu() {
	// Membersihkan layar sebelum menampilkan menu
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul menu dengan garis pembatas
	fmt.Println("===============================")
	fmt.Println("|         Menu Admin          |")
	fmt.Println("===============================")

	// Menampilkan pilihan menu
	fmt.Println("\n1. Login")
	fmt.Println("2. Register")
	fmt.Println("3. Keluar")

	// Meminta input user
	fmt.Print("Masukkan pilihan (1/2/3): ")
}

func ViewLoginAdmin() {
	// Membersihkan layar sebelum menampilkan menu
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul menu dengan garis pembatas
	fmt.Println("===============================")
	fmt.Println("|         Login Admin         |")
	fmt.Println("===============================")
}

func ViewRegisterAdmin() {
	// Membersihkan layar sebelum menampilkan menu
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul menu dengan garis pembatas
	fmt.Println("===============================")
	fmt.Println("|         Register Admin      |")
	fmt.Println("===============================")
}

func ViewMainMenu() {
	// Membersihkan layar sebelum menampilkan menu
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul menu dengan garis pembatas
	fmt.Println("===============================")
	fmt.Println("|         Menu Utama          |")
	fmt.Println("===============================")

	// Menampilkan pilihan menu
	fmt.Println("\n1. Data Buku")
	fmt.Println("2. Pinjam Buku")
	fmt.Println("3. Kembali")

	// Meminta input user
	fmt.Print("Masukkan pilihan (1/2/3): ")
}

func login(dataAdmin *dataAdmin) {
	var dataBuku dataBuku
	var dataPinjam dataPinjamBuku
	var username, password string
	var isLogin bool

	isLogin = false

	if jumlahDataAdmin == 0 {
		ViewLoginAdmin()
		fmt.Println("\nData admin masih kosong!")
		fmt.Println("Silahkan register terlebih dahulu")
		fmt.Scanln()
	} else {
		ViewLoginAdmin()
		fmt.Print("\nMasukkan username: ")
		fmt.Scan(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&password)

		// Periksa username dan password di database
		// Contoh: menggunakan slice dataAdmin
		for i := 0; i < jumlahDataAdmin; i++ {
			if dataAdmin[i].username == username && dataAdmin[i].password == password {
				isLogin = true
			}
		}

		if isLogin {
			fmt.Println("\nLogin berhasil!")
			fmt.Scanln()
			fmt.Println("\nTekan Enter untuk ke menu buku")
			fmt.Scanln()
			MainMenu(&dataBuku, &dataPinjam)
		} else {
			fmt.Println("Username atau password salah!")
			fmt.Scanln()
		}
	}
	fmt.Println("\nTekan Enter untuk kembali ke menu buku")
	fmt.Scanln()
}

func registration(dataAdmin *dataAdmin) {
	var adminBaru admin

	ViewRegisterAdmin()

	fmt.Print("\nMasukkan username: ")
	fmt.Scan(&adminBaru.username)
	fmt.Print("Masukkan email: ")
	fmt.Scan(&adminBaru.email)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&adminBaru.password)

	// Validasi username
	for i := 0; i < jumlahDataAdmin; i++ {
		if dataAdmin[i].username == adminBaru.username {
			fmt.Println("Username sudah ada. Silakan coba username lain.")
			registration(dataAdmin)
		}
	}

	// Simpan data admin baru ke slice dataAdmin
	dataAdmin[jumlahDataAdmin] = adminBaru
	jumlahDataAdmin++

	fmt.Println("Registrasi berhasil!")
	fmt.Println("\nTekan Enter untuk kembali ke menu buku")
	fmt.Scanln()
}

func AdminMenu(dataAdmin *dataAdmin) {
	pilihan := 0
	keluar := false

	for !keluar {
		ViewAdminMenu()
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			login(dataAdmin)
		case 2:
			registration(dataAdmin)
		case 3:
			keluar = true
			fmt.Println("\n======================================")
			fmt.Println("|         KELUAR DARI PROGAM         |")
			fmt.Println("======================================\n")
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func MainMenu(dataBuku *dataBuku, dataPinjam *dataPinjamBuku) {
	pilihan := 0
	keluar := false

	for !keluar {
		ViewMainMenu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			BukuMenu(dataBuku, dataPinjam)
		case 2:
			PinjamMenu(dataPinjam, dataBuku)
		case 3:
			keluar = true
			fmt.Println("\n======================================")
			fmt.Println("|         KELUAR DARI PROGAM         |")
			fmt.Println("======================================\n")
		default:
			fmt.Println("Pilihan tidak valid!")
		}

		// Menunggu user menekan Enter untuk melanjutkan
		if !keluar {
			fmt.Println("\nTekan Enter untuk kembali ke menu utama")
			fmt.Scanln()
		}
	}
}

func BukuMenu(dataBuku *dataBuku, dataPinjam *dataPinjamBuku) {
	pilihan := 0
	keluar := false // Boolean variable to control the loop

	for !keluar { // Continue looping until keluar is true
		ViewBukuMenu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			ShowDataBukuMenu(dataBuku, dataPinjam)
		case 2:
			AddDataBuku(dataBuku)
		case 3:
			UpdateDataBuku(dataBuku)
		case 4:
			DeleteDataBukuMenu(dataBuku, *dataPinjam)
		case 5:
			SearchDataBukuByKeyword(dataBuku)
		case 6:
			MainMenu(dataBuku, dataPinjam)
			keluar = true // Set keluar to true to exit the loop
		default:
			fmt.Println("Pilihan tidak valid!")
			fmt.Print("Tolong masukkan pilihan kembali: ")
			fmt.Scan(&pilihan)
		}

		if !keluar { // Check if we should continue showing the menu
			// Tampilkan menu kembali setelah setiap operasi CRUD
			fmt.Scanln() // Menunggu user menekan Enter
		}
	}
}

func ShowDataBukuMenu(dataBuku *dataBuku, dataPinjam *dataPinjamBuku) {
	pilihan := 0
	keluar := false

	for !keluar {
		ViewShowDataBukuMenu()
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			ShowDataBuku(dataBuku)
		case 2:
			ShowDataBukuByKategoriMenu(dataBuku, dataPinjam)
		case 3:
			BukuMenu(dataBuku, dataPinjam)
			keluar = true
		default:
			fmt.Println("Pilihan tidak valid!")
			fmt.Print("Tolong masukkan pilihan kembali: ")
			fmt.Scan(&pilihan)
		}
	}

	if !keluar { // Check if we should continue showing the menu
		// Tampilkan menu kembali setelah setiap operasi CRUD
		fmt.Scanln() // Menunggu user menekan Enter
	}
}

func ShowDataBukuByKategoriMenu(dataBuku *dataBuku, dataPinjam *dataPinjamBuku) {
	pilihan := 0
	keluar := false

	for !keluar {
		ViewShowDataBukuByKategoriMenu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			ShowDataBukuByJudul(dataBuku)
		case 2:
			ShowDataBukuByPenulis(dataBuku)
		case 3:
			ShowDataBukuByPenerbit(dataBuku)
		case 4:
			ShowDataBukuByTahunTerbit(dataBuku)
		case 5:
			BukuMenu(dataBuku, dataPinjam)
			keluar = true
		default:
			fmt.Println("Pilihan tidak valid!")
			fmt.Print("Tolong masukkan pilihan kembali: ")
			fmt.Scan(&pilihan)
		}

		if !keluar { // Check if we should continue showing the menu
			// Tampilkan menu kembali setelah setiap operasi CRUD
			fmt.Scanln() // Menunggu user menekan Enter
		}
	}
}

func DeleteDataBukuMenu(dataBuku *dataBuku, dataPinjam dataPinjamBuku) {
	pilihan := 0
	keluar := false

	for !keluar {
		ViewDeleteDataBukuMenu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			DeleteDataBukuByISBN(dataBuku)
		case 2:
			DeleteDataBukuByJudul(dataBuku)
		case 3:
			DeleteDataBukuByPenulis(dataBuku)
		case 4:
			DeleteDataBukuByTahunTerbit(dataBuku)
		case 5:
			DeleteAllDataBuku(dataBuku, &dataPinjam)
		case 6:
			BukuMenu(dataBuku, &dataPinjam)
			keluar = true
		default:
			fmt.Println("Pilihan tidak valid!")
			fmt.Print("Tolong masukkan pilihan kembali: ")
			fmt.Scan(&pilihan)
		}

		if !keluar { // Check if we should continue showing the menu
			// Tampilkan menu kembali setelah setiap operasi CRUD
			fmt.Scanln() // Menunggu user menekan Enter
		}
	}
}

func AddDataBuku(data *dataBuku) {
	var bukuBaru buku
	var inputJumlahDataBuku int
	var isBerhasil bool

	ViewCreateDataBuku()
	// Menjelaskan format input
	fmt.Println("\nMasukkan data buku dengan format yang benar.")

	// Meminta input jumlah buku baru
	fmt.Print("Masukkan jumlah buku yang akan ditambahkan: ")
	fmt.Scan(&inputJumlahDataBuku)
	isBerhasil = false
	for i := 0; i < inputJumlahDataBuku; i++ {
		isDuplicate := false
		// Meminta input ISBN
		fmt.Print("\nISBN		: ")
		fmt.Scan(&bukuBaru.isbn)
		for j := 0; j < jumlahDataBuku; j++ {
			if data[j].isbn == bukuBaru.isbn {
				isDuplicate = true
				// Jika ISBN sudah ada, minta ISBN lain
				fmt.Println("\nBuku dengan ISBN yang sama sudah ada. Silakan masukkan ISBN lain.")
				i-- // Mengurangi indeks untuk mengulang input buku ini
			}
		}
		if !isDuplicate {
			// Meminta input judul
			fmt.Print("Judul		: ")
			fmt.Scan(&bukuBaru.judul)
			// Meminta input penulis
			fmt.Print("Penulis		: ")
			fmt.Scan(&bukuBaru.penulis)
			// Meminta input penerbit
			fmt.Print("Penerbit	: ")
			fmt.Scan(&bukuBaru.penerbit)
			// Meminta input tahun terbit
			fmt.Print("Tahun Terbit	: ")
			fmt.Scan(&bukuBaru.tahunTerbit)
			// Meminta input jumlah halaman
			fmt.Print("Jumlah Halaman	: ")
			fmt.Scan(&bukuBaru.jumlahHalaman)
			bukuBaru.status = true

			if !isBerhasil {
				data[jumlahDataBuku] = bukuBaru
				fmt.Println("\nBuku berhasil ditambahkan")
				isBerhasil = false
				jumlahDataBuku++
			} else {
				fmt.Println("\nBuku gagal ditambahkan")
				isBerhasil = true
			}
		}
	}
	inputJumlahDataBuku++
	fmt.Println("\nTekan Enter untuk kembali ke menu buku")
	fmt.Scanln()
}

func ShowDataBuku(data *dataBuku) {
	// SortingDataBukuByTahunTerbit(data)
	if jumlahDataBuku == 0 {
		ViewShowDataBuku()
		// Menampilkan pesan jika tidak ada data buku
		fmt.Println("\nTidak ada data buku")
		fmt.Println("\nTekan Enter untuk kembali ke menu buku")
		fmt.Scanln()
	} else {
		ViewShowDataBuku()
		// Menampilkan daftar buku
		fmt.Println("\nSemua Daftar Buku:")
		for i := 0; i < jumlahDataBuku; i++ {
			fmt.Println()
			buku := (*data)[i]
			if buku.isbn != "" { // Periksa apakah buku memiliki judul
				fmt.Printf("ISBN		: %s\n", buku.isbn)
				fmt.Printf("Judul		: %s\n", buku.judul)
				fmt.Printf("Penulis		: %s\n", buku.penulis)
				fmt.Printf("Penerbit	: %s\n", buku.penerbit)
				fmt.Printf("Tahun Terbit	: %d\n", buku.tahunTerbit)
				fmt.Printf("Jumlah Halaman	: %d\n", buku.jumlahHalaman)
				if buku.status == true {
					fmt.Printf("Status Buku	: %s\n", "Buku tersedia")
				} else {
					fmt.Printf("Status Buku	: %s\n", "Buku sedang dipinjam")
				}
				fmt.Println("===================================")
			}
		}
		fmt.Println("\nTekan Enter untuk kembali ke menu utama")
		fmt.Scanln()
	}
	fmt.Scanln()
}

func ShowDataBukuByJudul(data *dataBuku) {
	if jumlahDataBuku == 0 {
		ViewShowDataBuku()
		// Menampilkan pesan jika tidak ada data buku
		fmt.Println("\nTidak ada data buku")
		fmt.Println("\nTekan Enter untuk kembali ke menu buku")
		fmt.Scanln()
	} else {
		ViewShowDataBuku()
		// Bubble sort algorithm to sort books by title
		SortingDataBukuByJudul(data)
		// Print sorted book data
		fmt.Println("\nDaftar Buku Berdasarkan Judul:")
		for i := 0; i < jumlahDataBuku; i++ {
			buku := data[i]
			if buku.judul != "" { // Check if the book has a title
				fmt.Printf("ISBN            : %s\n", buku.isbn)
				fmt.Printf("Judul           : %s\n", buku.judul)
				fmt.Printf("Penulis         : %s\n", buku.penulis)
				fmt.Printf("Penerbit        : %s\n", buku.penerbit)
				fmt.Printf("Tahun Terbit    : %d\n", buku.tahunTerbit)
				fmt.Printf("Jumlah Halaman  : %d\n", buku.jumlahHalaman)
				if buku.status == true {
					fmt.Printf("Status Buku     : %s\n", "Buku tersedia")
				} else {
					fmt.Printf("Status Buku     : %s\n", "Buku sedang dipinjam")
				}
				fmt.Println("===================================")
			}
		}
		fmt.Println("\nTekan Enter untuk kembali ke menu utama")
		fmt.Scanln()
	}
}

func ShowDataBukuByPenulis(data *dataBuku) {
	if jumlahDataBuku == 0 {
		ViewShowDataBuku()
		// Menampilkan pesan jika tidak ada data buku
		fmt.Println("\nTidak ada data buku")
		fmt.Println("\nTekan Enter untuk kembali ke menu buku")
		fmt.Scanln()
	} else {
		ViewShowDataBuku()
		// Bubble sort algorithm to sort books by title
		SortingDataBukuByPenulis(data)
		// Print sorted book data
		fmt.Println("\nDaftar Buku Berdasarkan Penulis:")
		for i := 0; i < jumlahDataBuku; i++ {
			buku := data[i]
			if buku.penulis != "" { // Check if the book has a title
				fmt.Printf("ISBN            : %s\n", buku.isbn)
				fmt.Printf("Judul           : %s\n", buku.judul)
				fmt.Printf("Penulis         : %s\n", buku.penulis)
				fmt.Printf("Penerbit        : %s\n", buku.penerbit)
				fmt.Printf("Tahun Terbit    : %d\n", buku.tahunTerbit)
				fmt.Printf("Jumlah Halaman  : %d\n", buku.jumlahHalaman)
				if buku.status == true {
					fmt.Printf("Status Buku     : %s\n", "Buku tersedia")
				} else {
					fmt.Printf("Status Buku     : %s\n", "Buku sedang dipinjam")
				}
				fmt.Println("===================================")
			}
		}
		fmt.Println("\nTekan Enter untuk kembali ke menu utama")
		fmt.Scanln()
	}
}

func ShowDataBukuByPenerbit(data *dataBuku) {
	if jumlahDataBuku == 0 {
		ViewShowDataBuku()
		// Menampilkan pesan jika tidak ada data buku
		fmt.Println("\nTidak ada data buku")
		fmt.Println("\nTekan Enter untuk kembali ke menu buku")
		fmt.Scanln()
	} else {
		ViewShowDataBuku()
		// Bubble sort algorithm to sort books by title
		SortingDataBukuByPenerbit(data)
		// Print sorted book data
		fmt.Println("\nDaftar Buku Berdasarkan Penerbit:")
		for i := 0; i < jumlahDataBuku; i++ {
			buku := data[i]
			if buku.penerbit != "" { // Check if the book has a title
				fmt.Printf("ISBN            : %s\n", buku.isbn)
				fmt.Printf("Judul           : %s\n", buku.judul)
				fmt.Printf("Penulis         : %s\n", buku.penulis)
				fmt.Printf("Penerbit        : %s\n", buku.penerbit)
				fmt.Printf("Tahun Terbit    : %d\n", buku.tahunTerbit)
				fmt.Printf("Jumlah Halaman  : %d\n", buku.jumlahHalaman)
				if buku.status == true {
					fmt.Printf("Status Buku     : %s\n", "Buku tersedia")
				} else {
					fmt.Printf("Status Buku     : %s\n", "Buku sedang dipinjam")
				}
				fmt.Println("===================================")
			}
		}
		fmt.Println("\nTekan Enter untuk kembali ke menu utama")
		fmt.Scanln()
	}
}

func ShowDataBukuByTahunTerbit(data *dataBuku) {
	if jumlahDataBuku == 0 {
		ViewShowDataBuku()
		// Menampilkan pesan jika tidak ada data buku
		fmt.Println("\nTidak ada data buku")
		fmt.Println("\nTekan Enter untuk kembali ke menu buku")
		fmt.Scanln()
	} else {
		ViewShowDataBuku()
		// Bubble sort algorithm to sort books by title
		SortingDataBukuByTahunTerbit(data)
		// Print sorted book data
		fmt.Println("\nDaftar Buku Berdasarkan Tahun Terbit:")
		for i := 0; i < jumlahDataBuku; i++ {
			buku := data[i]
			if buku.tahunTerbit != 0 { // Check if the book has a title
				fmt.Printf("ISBN            : %s\n", buku.isbn)
				fmt.Printf("Judul           : %s\n", buku.judul)
				fmt.Printf("Penulis         : %s\n", buku.penulis)
				fmt.Printf("Penerbit        : %s\n", buku.penerbit)
				fmt.Printf("Tahun Terbit    : %d\n", buku.tahunTerbit)
				fmt.Printf("Jumlah Halaman  : %d\n", buku.jumlahHalaman)
				if buku.status == true {
					fmt.Printf("Status Buku     : %s\n", "Buku tersedia")
				} else {
					fmt.Printf("Status Buku     : %s\n", "Buku sedang dipinjam")
				}
				fmt.Println("===================================")
			}
		}
		fmt.Println("\nTekan Enter untuk kembali ke menu utama")
		fmt.Scanln()
	}
}

// mengubah data buku
func UpdateDataBuku(data *dataBuku) {
	var indexFound int
	var isbnBuku string
	var judulBaru, penulisBaru, penerbitBaru string
	var tahunTerbitBaru, jumlahHalamanBaru int
	ViewUpdateDataBuku()
	fmt.Print("\nMasukkan ISBN yang ingin diubah	: ")
	fmt.Scan(&isbnBuku)
	fmt.Println()

	// update data
	indexFound = FindDataBukuByISBNWithSequentialSearch(*data, isbnBuku)
	if indexFound != -1 {
		fmt.Print("Masukkan Judul Baru		: ")
		fmt.Scan(&judulBaru)
		fmt.Print("Masukkan Penulis Baru		: ")
		fmt.Scan(&penulisBaru)
		fmt.Print("Masukkan Penerbit Baru		: ")
		fmt.Scan(&penerbitBaru)
		fmt.Print("Masukkan Tahun Terbit Baru	: ")
		fmt.Scan(&tahunTerbitBaru)
		fmt.Print("Masukkan Jumlah Halaman Baru	: ")
		fmt.Scan(&jumlahHalamanBaru)

		data[indexFound].judul = judulBaru
		data[indexFound].penulis = penulisBaru
		data[indexFound].penerbit = penerbitBaru
		data[indexFound].tahunTerbit = tahunTerbitBaru
		data[indexFound].jumlahHalaman = jumlahHalamanBaru
		fmt.Println("Buku berhasil diubah")
	} else {
		fmt.Println("Maaf, buku dengan ISBN tersebut tidak dapat ditemukan")
	}
}

func DeleteDataBukuByKeyword(data *dataBuku) {
	var kataKunci string
	var index int
	var found bool

	ViewDeleteDataBuku()
	fmt.Print("Masukkan kata kunci buku yang ingin dihapus: ")
	fmt.Scan(&kataKunci)

	found = true
	for found {
		index = FindDataBukuWithSequentialSearch(*data, kataKunci)
		if index != -1 {
			// Menggeser semua elemen setelah indeks yang ditemukan
			for i := index; i < jumlahDataBuku-1; i++ {
				data[i] = data[i+1]
			}
			jumlahDataBuku--
		} else {
			found = false
		}
	}

	if jumlahDataBuku < MAXBUKU {
		fmt.Printf("Semua data buku dengan kata kunci \"%s\" berhasil dihapus\n", kataKunci)
	} else {
		fmt.Println("\nMaaf, buku dengan judul tersebut tidak dapat ditemukan")
	}
	fmt.Println("\nTekan Enter untuk kembali ke menu utama")
	fmt.Scanln()
}

func DeleteAllDataBuku(dataBuku *dataBuku, dataPinjam *dataPinjamBuku) {
	// Mengatur jumlahDataBuku menjadi 0 akan menghapus semua data
	keluar := false
	jumlahDataBuku = 0

	if !keluar {
		fmt.Println("Semua data buku berhasil dihapus.")
		keluar = true
		fmt.Println("\nTekan Enter untuk kembali ke menu utama")
		fmt.Scanln()
	} else {
		fmt.Println("\nTekan Enter untuk kembali ke menu utama")
		fmt.Scanln()
	}
}

// menghapus data buku by isbn
func DeleteDataBukuByISBN(data *dataBuku) {
	var index int
	var isbn string

	ViewDeleteDataBuku()
	fmt.Print("Masukkan ISBN buku yang ingin dihapus: ")
	fmt.Scan(&isbn)

	index = FindDataBukuByISBNWithSequentialSearch(*data, isbn)
	if index != -1 {
		for i := index; i <= jumlahDataBuku-1; i++ {
			data[i] = data[i+1]
		}
		jumlahDataBuku = jumlahDataBuku - 1
		fmt.Printf("Data buku dengan ISBN %s berhasil dihapus", isbn)
	} else {
		fmt.Println("\nMaaf, buku dengan ISBN tersebut tidak dapat ditemukan")
		fmt.Println("\nTekan Enter untuk kembali ke menu utama")
		fmt.Scanln()
	}
}

// delete data buku by judul
func DeleteDataBukuByJudul(data *dataBuku) {
	var judul string
	var index int
	var found bool

	ViewDeleteDataBuku()
	fmt.Print("Masukkan judul buku yang ingin dihapus: ")
	fmt.Scan(&judul)

	found = true
	for found {
		index = FindDataBukuByJudulWithSequentialSearch(*data, judul)
		if index != -1 {
			// Menggeser semua elemen setelah indeks yang ditemukan
			for i := index; i < jumlahDataBuku-1; i++ {
				data[i] = data[i+1]
			}
			jumlahDataBuku--
		} else {
			found = false
		}
	}

	if jumlahDataBuku < MAXBUKU {
		fmt.Printf("Semua data buku dengan judul \"%s\" berhasil dihapus\n", judul)
	} else {
		fmt.Println("\nMaaf, buku dengan judul tersebut tidak dapat ditemukan")
	}
	fmt.Println("\nTekan Enter untuk kembali ke menu utama")
	fmt.Scanln()
}

// delete data buku by penulis
func DeleteDataBukuByPenulis(data *dataBuku) {
	var penulis string
	var index int
	var found bool

	ViewDeleteDataBuku()
	fmt.Print("Masukkan penulis buku yang ingin dihapus: ")
	fmt.Scan(&penulis)

	found = true
	for found {
		index = FindDataBukuByPenulisWithSequentialSearch(*data, penulis)
		if index != -1 {
			// Menggeser semua elemen setelah indeks yang ditemukan
			for i := index; i < jumlahDataBuku-1; i++ {
				data[i] = data[i+1]
			}
			jumlahDataBuku--
		} else {
			found = false
		}
	}

	if jumlahDataBuku < MAXBUKU {
		fmt.Printf("Semua data buku dengan penulis %s berhasil dihapus\n", penulis)
	} else {
		fmt.Println("\nMaaf, buku dengan penulis tersebut tidak dapat ditemukan")
	}
	fmt.Println("\nTekan Enter untuk kembali ke menu hapus")
	fmt.Scanln()
}

// delete data buku by tahun terbit
func DeleteDataBukuByTahunTerbit(data *dataBuku) {
	var tahunTerbit int
	var index int
	var found bool

	ViewDeleteDataBuku()
	fmt.Print("Masukkan tahun terbit buku yang ingin dihapus: ")
	fmt.Scan(&tahunTerbit)

	found = true
	for found {
		index = FindDataBukuByTahunTerbitWithSequentialSearch(*data, tahunTerbit)
		if index != -1 {
			// Menggeser semua elemen setelah indeks yang ditemukan
			for i := index; i < jumlahDataBuku-1; i++ {
				data[i] = data[i+1]
			}
			jumlahDataBuku--
		} else {
			found = false
		}
	}

	if jumlahDataBuku < MAXBUKU {
		fmt.Printf("Semua data buku dengan tahun terbit %d berhasil dihapus\n", tahunTerbit)
	} else {
		fmt.Println("\nMaaf, buku dengan tahun terbit tersebut tidak dapat ditemukan")
	}
	fmt.Println("\nTekan Enter untuk kembali ke menu hapus")
	fmt.Scanln()
}

func SearchDataBukuByKeyword(data *dataBuku) {
	var kataKunci string
	var found bool

	ViewSearchDataBuku()
	fmt.Print("Masukkan kata kunci buku yang ingin dicari: ")
	fmt.Scan(&kataKunci)

	found = true
	for i := 0; i < jumlahDataBuku; i++ {
		tahunTerbitString := fmt.Sprintf("%d", data[i].tahunTerbit)
		if data[i].isbn == kataKunci || data[i].judul == kataKunci || data[i].penulis == kataKunci || data[i].penerbit == kataKunci || tahunTerbitString == kataKunci {
			found = false
			fmt.Printf("\nData buku dengan kata kunci \"%s\":\n", kataKunci)
			fmt.Printf("ISBN        	: %s\n", data[i].isbn)
			fmt.Printf("Judul           : %s\n", data[i].judul)
			fmt.Printf("Penulis       	: %s\n", data[i].penulis)
			fmt.Printf("Penerbit       	: %s\n", data[i].penerbit)
			fmt.Printf("Tahun Terbit   	: %d\n", data[i].tahunTerbit)
			fmt.Printf("Jumlah Halaman 	: %d\n", data[i].jumlahHalaman)
		}
	}

	if found {
		fmt.Printf("\nMaaf, buku dengan kata kunci \"%s\" tidak dapat ditemukan.\n", kataKunci)
	}

	fmt.Println("\nTekan Enter untuk kembali ke menu utama.")
	fmt.Scanln()
}

// mencari data buku menggunakan sequential search
func FindDataWithSequentialSearch(data dataBuku, isbn string) int {
	var found int = -1
	var i int
	for i < jumlahDataBuku && found == -1 {
		if isbn == data[i].isbn {
			found = i
		}
		i++
	}
	return found
}

func FindDataBukuWithSequentialSearch(data dataBuku, kataKunci string) int {
	var found int = -1
	var i int
	for i < jumlahDataBuku && found == -1 {
		tahunTerbitString := fmt.Sprintf("%d", data[i].tahunTerbit)
		if data[i].isbn == kataKunci || data[i].judul == kataKunci || data[i].penulis == kataKunci || data[i].penerbit == kataKunci || tahunTerbitString == kataKunci {
			found = i
		}
		i++
	}
	return found
}

func FindDataBukuByISBNWithSequentialSearch(data dataBuku, isbn string) int {
	var found int = -1
	var i int
	for i < jumlahDataBuku && found == -1 {
		if isbn == data[i].isbn {
			found = i
		}
		i++
	}
	return found
}

func FindDataBukuByJudulWithSequentialSearch(data dataBuku, judul string) int {
	var found int = -1
	var i int
	for i < jumlahDataBuku && found == -1 {
		if judul == data[i].judul {
			found = i
		}
		i++
	}
	return found
}

func FindDataBukuByPenulisWithSequentialSearch(data dataBuku, penulis string) int {
	var found int = -1
	var i int
	for i < jumlahDataBuku && found == -1 {
		if penulis == data[i].penulis {
			found = i
		}
		i++
	}
	return found
}

func FindDataBukuByTahunTerbitWithSequentialSearch(data dataBuku, tahunTerbit int) int {
	var found int = -1
	var i int
	for i < jumlahDataBuku && found == -1 {
		if tahunTerbit == data[i].tahunTerbit {
			found = i
		}
		i++
	}
	return found
}

// mencari data buku menggunakan binary search
func FindDataWithBinarySearch(data dataBuku, isbn string, n int) int {
	var left, right, mid, index int
	left = 0
	right = n - 1
	index = -1
	for left <= right && index == -1 {
		mid = (left + right) / 2
		if isbn < data[mid].isbn {
			right = mid - 1
		} else if isbn > data[mid].isbn {
			left = mid + 1
		} else {
			index = mid
		}
	}
	return index
}

// mengurutkan data buku menggunakan selection sort
func SortingDataBukuByTahunTerbit(data *dataBuku) {

	for i := 0; i < jumlahDataBuku-1; i++ {
		maxIndex := i
		for j := i + 1; j < jumlahDataBuku; j++ {
			if data[j].tahunTerbit > data[maxIndex].tahunTerbit {
				maxIndex = j
			}
		}

		if i != maxIndex {
			data[i], data[maxIndex] = data[maxIndex], data[i]
		}
	}
}

func SortingDataBukuByJudul(data *dataBuku) {
	// Temporary variable for swapping
	var temp buku

	for i := 0; i < jumlahDataBuku-1; i++ {
		for j := 0; j < jumlahDataBuku-i-1; j++ {
			if data[j].judul > data[j+1].judul {
				// Swap the books
				temp = (*data)[j]
				(*data)[j] = (*data)[j+1]
				(*data)[j+1] = temp
			}
		}
	}
}

func SortingDataBukuByPenulis(data *dataBuku) {
	// Temporary variable for swapping
	var temp buku

	for i := 0; i < jumlahDataBuku-1; i++ {
		for j := 0; j < jumlahDataBuku-i-1; j++ {
			if data[j].penulis > data[j+1].penulis {
				// Swap the books
				temp = data[j]
				data[j] = data[j+1]
				data[j+1] = temp
			}
		}
	}
}

func SortingDataBukuByPenerbit(data *dataBuku) {
	// Temporary variable for swapping
	var temp buku

	for i := 0; i < jumlahDataBuku-1; i++ {
		for j := 0; j < jumlahDataBuku-i-1; j++ {
			if data[j].penerbit > data[j+1].penerbit {
				// Swap the books
				temp = data[j]
				data[j] = data[j+1]
				data[j+1] = temp
			}
		}
	}
}

// tampilan menu pinjam
func ViewPinjamMenu() {
	// Membersihkan layar sebelum menampilkan menu
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul menu dengan garis pembatas
	fmt.Println("================================")
	fmt.Println("|        Menu Pinjam Buku      |")
	fmt.Println("================================")
	fmt.Println("\n1. Tampilkan Daftar Buku yang Dipinjam")
	fmt.Println("2. Tambah Data Peminjaman")
	fmt.Println("3. Ubah Data Peminjaman")
	fmt.Println("4. Hapus Data Peminjaman")
	fmt.Println("5. Hitung Tarif Pinjaman dan Denda")
	fmt.Println("6. Kembali ke Menu Utama")
	fmt.Print("Masukkan pilihan (1/2/3/4/5/6): ")
}

// tampilan menu create data pinjam buku
func ViewCreateDataPinjamBuku() {
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul menu dengan garis pembatas
	fmt.Println("==========================================")
	fmt.Println("|        Tambah Data Peminjaman  Buku    |")
	fmt.Println("==========================================")
}

func ViewShowDataPinjamBuku() {
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul menu dengan garis pembatas
	fmt.Println("=========================================")
	fmt.Println("|      Daftar Data Peminjaman  Buku     |")
	fmt.Println("=========================================")
}

func ViewUpdateDataTanggalPinjamBuku() {
	// Membersihkan layar sebelum menampilkan form tambah data
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul form tambah data
	fmt.Println("==============================================")
	fmt.Println("|        UBAH DATA TANGGAL PINJAM BUKU       |")
	fmt.Println("==============================================")
}

// menu pinjam
func PinjamMenu(dataPinjam *dataPinjamBuku, dataBuku *dataBuku) {
	pilihan := 0
	keluar := false
	for !keluar {
		ViewPinjamMenu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			ShowDataPinjamBuku(dataPinjam, dataBuku)
		case 2:
			CreateDataPinjamBuku(dataPinjam, dataBuku)
		case 3:
			UpdateDataTanggalPinjamBuku(dataPinjam, dataBuku)
		case 4:
			DeleteDataPinjamBuku()
		case 5:
			CalculateDataPinjamBuku()
		case 6:
			MainMenu(dataBuku, dataPinjam)
			keluar = true // Set keluar to true to exit the loop
		default:
			fmt.Println("Pilihan tidak valid!")
			fmt.Print("Tolong masukkan pilihan kembali: ")
			fmt.Scan(&pilihan)
		}

		if !keluar { // Check if we should continue showing the menu
			// Tampilkan menu kembali setelah setiap operasi CRUD
			fmt.Scanln() // Menunggu user menekan Enter
		}
	}
}

// create data pinjam buku
func CreateDataPinjamBuku(dataPinjam *dataPinjamBuku, dataBuku *dataBuku) {
	var isbn string
	var pinjamBukuBaru pinjamBuku
	var inputJumlahDataPinjamBuku int
	var isBerhasil bool

	isBerhasil = false

	ViewCreateDataPinjamBuku()

	ShowDataBuku(dataBuku)

	fmt.Println("\n!!! Ubah spasi dengan _ !!!")
	fmt.Print("Masukkan ID Anggota			: ")
	fmt.Scan(&pinjamBukuBaru.idAnggota)
	fmt.Print("Masukkan jumlah buku yang akan dipinjam	: ")
	fmt.Scan(&inputJumlahDataPinjamBuku)
	for i := 0; i < inputJumlahDataPinjamBuku; i++ {
		fmt.Print("Masukkan ISBN				: ")
		fmt.Scan(&isbn)
		foundISBN := FindDataBukuByISBNWithSequentialSearch(*dataBuku, isbn)
		if foundISBN != -1 {
			fmt.Print("Masukkan Tanggal Peminjaman		: ")
			fmt.Scan(&pinjamBukuBaru.tgl_pinjam)
			pinjamBukuBaru.tgl_kembali = 0

			//MENGECEK ISBN BUKU TELAH DIPINJAM, JADI TIDAK BISA DIPINJAM LAGI

			// if !dataBuku[i].status {
			// 	fmt.Println("Maaf, buku dengan ISBN tersebut sudah dipinjam.")
			// } else {
			// 	if (*dataPinjam)[jumlahDataPinjam] == (pinjamBuku{}) && !isBerhasil {
			// 		(*dataPinjam)[jumlahDataPinjam] = pinjamBukuBaru
			// 		fmt.Println("\nBuku berhasil dipinjam")
			// 		isBerhasil = false
			// 		dataPinjam[i].jumlahBukuDipinjam++
			// 		jumlahDataPinjam++
			// 		dataPinjam[i].idPinjamBuku++
			// 		dataBuku[foundISBN].status = false
			// 	} else {
			// 		fmt.Println("\nBuku gagal dipinjam")
			// 		isBerhasil = true
			// 	}
			// }

			if !isBerhasil {
				pinjamBukuBaru.isbn = dataBuku[foundISBN].isbn
				dataPinjam[jumlahDataPinjam] = pinjamBukuBaru
				fmt.Println("\nBuku berhasil dipinjam")

				jumlahDataPinjam++
				dataPinjam[foundISBN].idPinjamBuku++
				isBerhasil = false

				dataBuku[foundISBN].jumlahBukuDipinjam++
				dataBuku[foundISBN].status = false
			} else {
				fmt.Println("\nBuku gagal dipinjam")
				isBerhasil = true
			}
		} else {
			fmt.Println("Maaf, buku dengan ISBN tersebut tidak dapat ditemukan")
		}
	}
	inputJumlahDataPinjamBuku++
	fmt.Println("\nTekan Enter untuk kembali ke menu buku")
	fmt.Scanln()
}

// mencetak data buku yang di pinjam
func ShowDataPinjamBuku(dataPinjam *dataPinjamBuku, dataBuku *dataBuku) {
	// var isbn string
	if jumlahDataPinjam == 0 {
		ViewShowDataPinjamBuku()
		fmt.Println("\nTidak ada data buku yang sedang dipinjam")
		fmt.Println("\nTekan Enter untuk kembali ke menu buku")
		fmt.Scanln()
	} else {
		ViewShowDataPinjamBuku()
		for i := 0; i < jumlahDataPinjam; i++ {
			if dataPinjam[i].isbn == dataBuku[i].isbn && dataPinjam[i].status == false {
				pinjam := dataPinjam[i]
				buku := dataBuku[i]
				fmt.Println("ID Pinjam Buku		: ", pinjam.idPinjamBuku)
				fmt.Println("ID Anggota		: ", pinjam.idAnggota)
				fmt.Println("ISBN			: ", buku.isbn)
				fmt.Println("Judul			: ", buku.judul)
				fmt.Println("Penulis			: ", buku.penulis)
				fmt.Println("Penerbit		: ", buku.penerbit)
				fmt.Println("Tahun Terbit		: ", buku.tahunTerbit)
				fmt.Println("Tanggal dipinjam	: ", pinjam.tgl_pinjam)

				if pinjam.tgl_kembali == 0 {
					fmt.Println("Tanggal dikembalikan	:  Buku masih dipinjam")
				} else {
					fmt.Println("Tanggal dikembalikan	: ", pinjam.tgl_kembali)
				}

				fmt.Println("Jumlah Buku Dipinjam	: ", buku.jumlahBukuDipinjam)
				fmt.Println("=====================================")
			} else {
				fmt.Println("Maaf, buku dengan ISBN tersebut tidak dapat ditemukan")
				fmt.Println("\nTekan Enter untuk kembali ke menu pinjam")
			}
		}
	}
	fmt.Scanln()

}

// update data pinjam buku
func UpdateDataTanggalPinjamBuku(dataPinjam *dataPinjamBuku, dataBuku *dataBuku) {
	var indexFound int
	var idPinjamBuku int
	var tanggalKembali int
	ViewUpdateDataTanggalPinjamBuku()
	fmt.Print("Masukkan ID Pinjam Buku yang ingin diubah: ")
	fmt.Scan(&idPinjamBuku)
	indexFound = FindDataPinjamBukuByIdPinjamBukuWithSequentialSearch(*dataPinjam, idPinjamBuku)
	if indexFound != -1 {
		fmt.Print("Masukkan tanggal kembali: ")
		fmt.Scan(&tanggalKembali)
		(*dataPinjam)[indexFound].tgl_kembali = tanggalKembali
		fmt.Println("\nTanggal kembali berhasil diubah")
		dataBuku[indexFound].status = true
	} else {
		fmt.Println("Maaf, ID Pinjam Buku tersebut tidak dapat ditemukan")
	}
	fmt.Println("\nTekan Enter untuk kembali ke menu pinjam")
	fmt.Scanln()
}

func DeleteDataPinjamBuku() {
}

func CalculateDataPinjamBuku() {
}

func FindDataPinjamBukuByISBNWithSequentialSearch(data dataPinjamBuku, isbn string) int {
	var found int = -1
	var i int
	for i < jumlahDataBuku && found == -1 {
		if isbn == data[i].isbn {
			found = i
		}
		i++
	}
	return found
}

func FindDataPinjamBukuByIdPinjamBukuWithSequentialSearch(data dataPinjamBuku, idPinjamBuku int) int {
	var found int = -1
	var i int
	for i < jumlahDataBuku && found == -1 {
		if idPinjamBuku == data[i].idPinjamBuku {
			found = i
		}
		i++
	}
	return found
}

func main() {
	var dataAdmin dataAdmin
	AdminMenu(&dataAdmin)
}