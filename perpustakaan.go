package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

const MAXBUKU int = 1000

type tanggal struct {
	hari  int
	bulan int
	tahun int
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
	isbn        string
	idAnggota   int
	tgl_pinjam  tanggal
	tgl_kembali tanggal
	denda       float64
}

var jumlahDataBuku, jumlahDataPinjam int

type dataBuku [MAXBUKU]buku
type dataPinjamBuku [MAXBUKU]pinjamBuku

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd = exec.Command("clear")
	} else {
		fmt.Println("Unsupported platform")
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ViewLoginMenu() {
	clearScreen()
	fmt.Println("===================================================")
	fmt.Println("|              Aplikasi Perpustakaan              |")
	fmt.Println("|       Dibuat oleh Ardhian Calwa Nugraha &       |")
	fmt.Println("|              Arya Raka Syahputra                |")
	fmt.Println("|           Algoritma Pemgrograman 2024           |")
	fmt.Println("===================================================")

	fmt.Println("\n1. Login")
	fmt.Println("2. Keluar")

	fmt.Print("Masukkan pilihan (1/2/3): ")
}

func ViewBukuMenu() {
	clearScreen()
	fmt.Println("=======================================")
	fmt.Println("|          Menu Perpustakaan          |")
	fmt.Println("=======================================")
	fmt.Println("\n1. Tampilkan Daftar Buku")
	fmt.Println("2. Tambah Data Buku")
	fmt.Println("3. Ubah Data Buku")
	fmt.Println("4. Hapus Data Buku")
	fmt.Println("5. Pencarian Data Buku")
	fmt.Println("6. Kembali")
	fmt.Print("Masukkan pilihan (1/2/3/4/5/6): ")
}

func ViewDeleteDataBukuMenu() {
	clearScreen()
	fmt.Println("================================")
	fmt.Println("|     Menu Hapus Data Buku     |")
	fmt.Println("================================")
	fmt.Println("\n1. Hapus Data Buku Berdasarkan ISBN")
	fmt.Println("2. Hapus Data Buku Berdasarkan Judul")
	fmt.Println("3. Hapus Data Buku Berdasarkan Penulis")
	fmt.Println("4. Hapus Data Buku Berdasarkan Penerbit")
	fmt.Println("5. Hapus Data Buku Berdasarkan Tahun Terbit")
	fmt.Println("6. Hapus Semua Data Buku")
	fmt.Println("7. Kembali")
	fmt.Print("Masukkan Pilihan (1/2/3/4/5/6): ")
}

func ViewShowDataBukuMenu() {
	clearScreen()
	fmt.Println("================================")
	fmt.Println("|     Menu Cetak Data Buku     |")
	fmt.Println("================================")
	fmt.Println("\n1. Tampilkan Semua Data Buku")
	fmt.Println("2. Tampilkan Data Buku Berdasarkan Kategori")
	fmt.Println("3. Kembali")
	fmt.Print("Masukkan Pilihan (1/2/3): ")
}

func ViewShowDataBukuByKategoriMenu() {
	clearScreen()
	fmt.Println("=====================================================")
	fmt.Println("|     Menu Cetak Data Buku Berdasarkan Kategori     |")
	fmt.Println("=====================================================")
	fmt.Println("\n1. Tampilkan Data Buku Berdasarkan Judul")
	fmt.Println("2. Tampilkan Data Buku Berdasarkan Penulis")
	fmt.Println("3. Tampilkan Data Buku Berdasarkan Penerbit")
	fmt.Println("4. Tampilkan Data Buku Berdasarkan Tahun Terbit")
	fmt.Println("5. Kembali")
	fmt.Print("Masukkan Pilihan (1/2/3/4/5): ")
}

func ViewCreateDataBuku() {
	clearScreen()
	fmt.Println("===============================")
	fmt.Println("|    TAMBAH DATA BUKU BARU    |")
	fmt.Println("===============================")
}

func ViewUpdateDataBuku() {
	clearScreen()
	fmt.Println("===============================")
	fmt.Println("|       UBAH DATA BUKU        |")
	fmt.Println("===============================")
}

func ViewDeleteDataBuku() {
	clearScreen()
	fmt.Println("===============================")
	fmt.Println("|        HAPUS DATA BUKU      |")
	fmt.Println("===============================")
}

func ViewSearchDataBuku() {
	clearScreen()
	fmt.Println("===============================")
	fmt.Println("|       CARI DATA BUKU        |")
	fmt.Println("===============================")
}

func ViewShowDataBuku() {
	clearScreen()
	fmt.Println("===============================")
	fmt.Println("|         DAFTAR BUKU         |")
	fmt.Println("===============================")
}

func ViewLoginAdmin() {
	clearScreen()

	fmt.Println("===============================")
	fmt.Println("|         Login Admin         |")
	fmt.Println("===============================")
}

func ViewTop5BukuTerfavorit() {
	clearScreen()

	fmt.Println("========================================")
	fmt.Println("|         Top 5 Buku Terfavorit        |")
	fmt.Println("========================================")
}

func ViewMainMenu() {
	clearScreen()

	fmt.Println("==========================================")
	fmt.Println("|         Menu Utama Perpustakaan        |")
	fmt.Println("==========================================")

	fmt.Println("\n1. Data Buku")
	fmt.Println("2. Pinjam Buku")
	fmt.Println("3. Rekomendasi 5 Buku Terfavorit")
	fmt.Println("4. Log out")

	fmt.Print("Masukkan pilihan (1/2/3/4): ")
}

func LoginMenu() {
	pilihan := 0
	keluar := false

	for !keluar {
		ViewLoginMenu()
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			clearScreen()
			login()
		case 2:
			keluar = true
			fmt.Println("\n======================================")
			fmt.Println("|         KELUAR DARI PROGAM         |")
			fmt.Println("======================================")
			// program akan berhenti
			os.Exit(0)
		default:
			clearScreen()
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func login() {
	var dataBuku dataBuku
	var dataPinjam dataPinjamBuku
	var username, password string
	var isLogin bool

	isLogin = false

	ViewLoginAdmin()
	fmt.Print("\nMasukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	if username == "admin" && password == "root" {
		isLogin = true
	}

	if isLogin {
		fmt.Println("\nLogin berhasil!")
		fmt.Scanln()
		fmt.Println("\nTekan Enter untuk lanjut ke menu perpustakaan")
		fmt.Scanln()
		MainMenu(&dataBuku, &dataPinjam)
		clearScreen()
	} else {
		fmt.Println("\nMasukkan \"admin\" untuk username dan \"root\" untuk password")
		fmt.Scanln()

	}
	fmt.Println("\nTekan Enter untuk kembali ke menu login")
	fmt.Scanln()
	clearScreen()
	LoginMenu()
}

func MainMenu(dataBuku *dataBuku, dataPinjam *dataPinjamBuku) {
	pilihan := 0
	keluar := false

	for !keluar {
		ViewMainMenu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			clearScreen()
			BukuMenu(dataBuku, dataPinjam)
		case 2:
			clearScreen()
			PinjamMenu(dataPinjam, dataBuku)
		case 3:
			clearScreen()
			Top5BukuTerfavorit(dataBuku)
		case 4:
			keluar = true
			fmt.Println("\n======================================")
			fmt.Println("|         KELUAR DARI PROGAM         |")
			fmt.Println("======================================")
			// program akan berhenti
			os.Exit(0)
		default:
			fmt.Println("\nPilihan tidak valid!")
			fmt.Scanln()
		}

		if !keluar {
			fmt.Println("\nTekan Enter untuk kembali")
			fmt.Scanln()
		}
	}
}

func BukuMenu(dataBuku *dataBuku, dataPinjam *dataPinjamBuku) {
	pilihan := 0
	keluar := false

	for !keluar {
		ViewBukuMenu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			clearScreen()
			ShowDataBukuMenu(dataBuku, dataPinjam)
		case 2:
			clearScreen()
			AddDataBuku(dataBuku)
		case 3:
			clearScreen()
			UpdateDataBuku(dataBuku)
		case 4:
			clearScreen()
			DeleteDataBukuMenu(dataBuku, *dataPinjam)
		case 5:
			clearScreen()
			SearchDataBukuByKeyword(dataBuku)
		case 6:
			MainMenu(dataBuku, dataPinjam)
			keluar = true
		default:
			clearScreen()
			fmt.Println("Pilihan tidak valid!")
		}

		if !keluar {
			fmt.Println("\nTekan Enter untuk kembali ke menu buku")
			fmt.Scanln()
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
			clearScreen()
			ShowDataBuku(dataBuku)
		case 2:
			clearScreen()
			ShowDataBukuByKategoriMenu(dataBuku, dataPinjam)
		case 3:
			BukuMenu(dataBuku, dataPinjam)
			keluar = true
		default:
			clearScreen()
			fmt.Println("Pilihan tidak valid!")
		}

		if !keluar {
			fmt.Println("\nTekan Enter untuk kembali ke menu buku")
			fmt.Scanln()
		}
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
			clearScreen()
			ShowDataBukuByJudul(dataBuku)
		case 2:
			clearScreen()
			ShowDataBukuByPenulis(dataBuku)
		case 3:
			clearScreen()
			ShowDataBukuByPenerbit(dataBuku)
		case 4:
			clearScreen()
			ShowDataBukuByTahunTerbit(dataBuku)
		case 5:
			clearScreen()
			BukuMenu(dataBuku, dataPinjam)
			keluar = true
		default:
			clearScreen()
			fmt.Println("Pilihan tidak valid!")
		}

		if !keluar {
			fmt.Println("\nTekan Enter untuk kembali ke menu buku")
			fmt.Scanln()
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
			clearScreen()
			DeleteDataBukuByISBN(dataBuku)
		case 2:
			clearScreen()
			DeleteDataBukuByJudul(dataBuku)
		case 3:
			clearScreen()
			DeleteDataBukuByPenulis(dataBuku)
		case 4:
			clearScreen()
			DeleteDataBukuByPenerbit(dataBuku)
		case 5:
			clearScreen()
			DeleteDataBukuByTahunTerbit(dataBuku)
		case 6:
			clearScreen()
			DeleteAllDataBuku(dataBuku, &dataPinjam)
		case 7:
			clearScreen()
			BukuMenu(dataBuku, &dataPinjam)
			keluar = true
		default:
			clearScreen()
			fmt.Println("Pilihan tidak valid!")
		}

		if !keluar {
			fmt.Println("\nTekan Enter untuk kembali ke menu hapus data buku")
			fmt.Scanln()
		}
	}
}

func AddDataBuku(data *dataBuku) {
	var bukuBaru buku
	var inputJumlahDataBuku int
	var isBerhasil bool
	var inputJumlah, inputTahunTerbit, inputHalamanBuku string

	ViewCreateDataBuku()

	fmt.Print("\nMasukkan jumlah buku yang akan ditambahkan: ")
	fmt.Scan(&inputJumlah)

	// konversi input ke integer
	inputJumlahDataBuku, err := strconv.Atoi(inputJumlah)
	for err != nil {
		fmt.Println("\nInput jumlah data buku yang akan ditambahkan harus berupa angka. Silakan coba lagi.")
		fmt.Scanln()
		return
	}

	isBerhasil = false

	for i := 0; i < inputJumlahDataBuku; i++ {
		isDuplicate := false
		fmt.Print("\nISBN		: ")
		fmt.Scan(&bukuBaru.isbn)

		for j := 0; j < jumlahDataBuku; j++ {
			if data[j].isbn == bukuBaru.isbn {
				isDuplicate = true
				fmt.Println("\nBuku dengan ISBN yang sama sudah ada. Silakan masukkan ISBN lain.")
				i--
			}
		}

		if !isDuplicate {
			fmt.Print("Judul		: ")
			fmt.Scan(&bukuBaru.judul)
			fmt.Print("Penulis		: ")
			fmt.Scan(&bukuBaru.penulis)
			fmt.Print("Penerbit	: ")
			fmt.Scan(&bukuBaru.penerbit)
			fmt.Print("Tahun Terbit	: ")
			fmt.Scan(&inputTahunTerbit)

			// konversi input ke integer
			bukuBaru.tahunTerbit, err = strconv.Atoi(inputTahunTerbit)
			for err != nil {
				fmt.Println("\nInput tahun terbit harus berupa angka. Silakan coba lagi.")
				fmt.Scanln()
				return
			}
			fmt.Print("Jumlah Halaman	: ")
			fmt.Scan(&inputHalamanBuku)

			// konversi input ke integer
			bukuBaru.jumlahHalaman, err = strconv.Atoi(inputHalamanBuku)
			for err != nil {
				fmt.Println("\nInput jumlah halaman buku harus berupa angka. Silakan coba lagi.")
				fmt.Scanln()
				return
			}

			if !isBerhasil {
				data[jumlahDataBuku] = bukuBaru
				fmt.Println("\nBuku berhasil ditambahkan")
				bukuBaru.status = false
				isBerhasil = false
				jumlahDataBuku++
			} else {
				fmt.Println("\nBuku gagal ditambahkan")
				isBerhasil = true
			}
		}
	}
	inputJumlahDataBuku++
	fmt.Scanln()
}

func ShowDataBuku(data *dataBuku) {
	if jumlahDataBuku == 0 {
		ViewShowDataBuku()
		fmt.Println("\nTidak ada data buku")
		fmt.Scanln()
	} else {
		clearScreen()
		ViewShowDataBuku()
		fmt.Println("\nSemua Daftar Buku:")
		for i := 0; i < jumlahDataBuku; i++ {
			fmt.Println()
			buku := data[i]
			if buku.isbn != "" {
				fmt.Printf("ISBN		: %s\n", buku.isbn)
				fmt.Printf("Judul		: %s\n", buku.judul)
				fmt.Printf("Penulis		: %s\n", buku.penulis)
				fmt.Printf("Penerbit	: %s\n", buku.penerbit)
				fmt.Printf("Tahun Terbit	: %d\n", buku.tahunTerbit)
				fmt.Printf("Jumlah Halaman	: %d\n", buku.jumlahHalaman)
				if !buku.status {
					fmt.Printf("Status Buku	: %s\n", "Buku tersedia")
				} else {
					fmt.Printf("Status Buku	: %s\n", "Buku sedang dipinjam")
				}
				fmt.Println("===================================")
			}
		}
		fmt.Scanln()
	}
}

func ShowDataBukuByJudul(data *dataBuku) {
	if jumlahDataBuku == 0 {
		ViewShowDataBuku()
		fmt.Println("\nTidak ada data buku")
		fmt.Scanln()
	} else {
		clearScreen()
		ViewShowDataBuku()
		SortingDataBukuByJudulBySelectionSort(data)
		fmt.Println("\nDaftar Buku Berdasarkan Judul:")
		for i := 0; i < jumlahDataBuku; i++ {
			buku := data[i]
			if buku.judul != "" {
				fmt.Printf("ISBN            : %s\n", buku.isbn)
				fmt.Printf("Judul           : %s\n", buku.judul)
				fmt.Printf("Penulis         : %s\n", buku.penulis)
				fmt.Printf("Penerbit        : %s\n", buku.penerbit)
				fmt.Printf("Tahun Terbit    : %d\n", buku.tahunTerbit)
				fmt.Printf("Jumlah Halaman  : %d\n", buku.jumlahHalaman)
				if !buku.status {
					fmt.Printf("Status Buku     : %s\n", "Buku tersedia")
				} else {
					fmt.Printf("Status Buku     : %s\n", "Buku sedang dipinjam")
				}
				fmt.Println("===================================")
			}
		}
		fmt.Scanln()
	}
}

func ShowDataBukuByPenulis(data *dataBuku) {
	if jumlahDataBuku == 0 {
		ViewShowDataBuku()
		fmt.Println("\nTidak ada data buku")
		fmt.Scanln()
	} else {
		clearScreen()
		ViewShowDataBuku()
		SortingDataBukuByPenulisBySelectionSort(data)
		fmt.Println("\nDaftar Buku Berdasarkan Penulis:")
		for i := 0; i < jumlahDataBuku; i++ {
			buku := data[i]
			if buku.penulis != "" {
				fmt.Printf("ISBN            : %s\n", buku.isbn)
				fmt.Printf("Judul           : %s\n", buku.judul)
				fmt.Printf("Penulis         : %s\n", buku.penulis)
				fmt.Printf("Penerbit        : %s\n", buku.penerbit)
				fmt.Printf("Tahun Terbit    : %d\n", buku.tahunTerbit)
				fmt.Printf("Jumlah Halaman  : %d\n", buku.jumlahHalaman)
				if !buku.status {
					fmt.Printf("Status Buku     : %s\n", "Buku tersedia")
				} else {
					fmt.Printf("Status Buku     : %s\n", "Buku sedang dipinjam")
				}
				fmt.Println("===================================")
			}
		}
		fmt.Scanln()
	}
}

func ShowDataBukuByPenerbit(data *dataBuku) {
	if jumlahDataBuku == 0 {
		ViewShowDataBuku()
		fmt.Println("\nTidak ada data buku")
		fmt.Scanln()
	} else {
		clearScreen()
		ViewShowDataBuku()
		SortingDataBukuByPenerbitBySelectionSort(data)
		fmt.Println("\nDaftar Buku Berdasarkan Penerbit:")
		for i := 0; i < jumlahDataBuku; i++ {
			buku := data[i]
			if buku.penerbit != "" {
				fmt.Printf("ISBN            : %s\n", buku.isbn)
				fmt.Printf("Judul           : %s\n", buku.judul)
				fmt.Printf("Penulis         : %s\n", buku.penulis)
				fmt.Printf("Penerbit        : %s\n", buku.penerbit)
				fmt.Printf("Tahun Terbit    : %d\n", buku.tahunTerbit)
				fmt.Printf("Jumlah Halaman  : %d\n", buku.jumlahHalaman)
				if !buku.status {
					fmt.Printf("Status Buku     : %s\n", "Buku tersedia")
				} else {
					fmt.Printf("Status Buku     : %s\n", "Buku sedang dipinjam")
				}
				fmt.Println("===================================")
			}
		}
		fmt.Scanln()
	}
}

func ShowDataBukuByTahunTerbit(data *dataBuku) {
	if jumlahDataBuku == 0 {
		ViewShowDataBuku()
		fmt.Println("\nTidak ada data buku")
		fmt.Scanln()
	} else {
		ViewShowDataBuku()
		SortingDataBukuByTahunTerbitByInsertionSort(data)
		fmt.Println("\nDaftar Buku Berdasarkan Tahun Terbit:")
		for i := 0; i < jumlahDataBuku; i++ {
			buku := data[i]
			if buku.tahunTerbit != 0 {
				fmt.Printf("ISBN            : %s\n", buku.isbn)
				fmt.Printf("Judul           : %s\n", buku.judul)
				fmt.Printf("Penulis         : %s\n", buku.penulis)
				fmt.Printf("Penerbit        : %s\n", buku.penerbit)
				fmt.Printf("Tahun Terbit    : %d\n", buku.tahunTerbit)
				fmt.Printf("Jumlah Halaman  : %d\n", buku.jumlahHalaman)
				if buku.status {
					fmt.Printf("Status Buku     : %s\n", "Buku sedang dipinjam")
				} else {
					fmt.Printf("Status Buku     : %s\n", "Buku tersedia")
				}
				fmt.Println("===================================")
			}
		}
		fmt.Scanln()
	}
}

func UpdateDataBuku(data *dataBuku) {
	var indexFound int
	var isbnBuku string
	var judulBaru, penulisBaru, penerbitBaru string
	var tahunTerbitBaru, jumlahHalamanBaru int
	ViewUpdateDataBuku()
	fmt.Print("\nMasukkan ISBN yang ingin diubah	: ")
	fmt.Scan(&isbnBuku)

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
		fmt.Printf("\nBuku dengan ISBN \"%s\" berhasil diubah\n", isbnBuku)
		fmt.Scanln()
	} else {
		fmt.Printf("\nMaaf, buku dengan ISBN \"%s\" tidak dapat ditemukan\n", isbnBuku)
		fmt.Scanln()
	}
}

func DeleteAllDataBuku(dataBuku *dataBuku, dataPinjam *dataPinjamBuku) {
	var input string
	ViewDeleteDataBuku()

	fmt.Print("Apakah anda yakin ingin menghapus semua data buku? (y/n) : ")
	fmt.Scan(&input)

	if input == "y" || input == "Y"{
		if jumlahDataBuku > 0 {
			jumlahDataBuku = 0
			fmt.Println("\nSemua data buku berhasil dihapus.")
			fmt.Scanln()
		} else {
			fmt.Println("\nTidak ada data buku yang dapat dihapus.")
			fmt.Scanln()
		}
	}
}

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
		fmt.Printf("\nData buku dengan ISBN \"%s\" berhasil dihapus\n", isbn)
		fmt.Scanln()
	} else {
		fmt.Printf("\nMaaf, buku dengan ISBN \"%s\" tidak dapat ditemukan\n", isbn)
		fmt.Scanln()
	}
}

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
			for i := index; i < jumlahDataBuku-1; i++ {
				data[i] = data[i+1]
			}
			jumlahDataBuku--
		} else {
			found = false
		}
	}

	if jumlahDataBuku < MAXBUKU {
		fmt.Printf("\nSemua data buku dengan judul \"%s\" berhasil dihapus\n", judul)
		fmt.Scanln()
	} else {
		fmt.Printf("\nMaaf, buku dengan judul \"%s\" tidak dapat ditemukan\n", judul)
		fmt.Scanln()
	}
}

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
			for i := index; i < jumlahDataBuku-1; i++ {
				data[i] = data[i+1]
			}
			jumlahDataBuku--
		} else {
			found = false
		}
	}

	if jumlahDataBuku < MAXBUKU {
		fmt.Printf("\nSemua data buku dengan penulis \"%s\" berhasil dihapus\n", penulis)
		fmt.Scanln()
	} else {
		fmt.Printf("\nMaaf, buku dengan penulis \"%s\" tidak dapat ditemukan\n", penulis)
		fmt.Scanln()
	}
}

func DeleteDataBukuByPenerbit(data *dataBuku) {
	var penerbit string
	var index int
	var found bool

	ViewDeleteDataBuku()
	fmt.Print("Masukkan penerbit buku yang ingin dihapus: ")
	fmt.Scan(&penerbit)

	found = true
	for found {
		index = FindDataBukuByPenerbitWithSequentialSearch(*data, penerbit)
		if index != -1 {
			for i := index; i < jumlahDataBuku-1; i++ {
				data[i] = data[i+1]
			}
			jumlahDataBuku--
		} else {
			found = false
		}
	}

	if jumlahDataBuku < MAXBUKU {
		fmt.Printf("\nSemua data buku dengan penerbit \"%s\" berhasil dihapus\n", penerbit)
		fmt.Scanln()
	} else {
		fmt.Printf("\nMaaf, buku dengan penerbit \"%s\" tidak dapat ditemukan\n", penerbit)
		fmt.Scanln()
	}
}

func DeleteDataBukuByTahunTerbit(data *dataBuku) {
	var tahunTerbit int
	var input string
	var index int
	var found bool

	ViewDeleteDataBuku()
	fmt.Print("Masukkan tahun terbit buku yang ingin dihapus: ")
	fmt.Scan(&input)

	tahunTerbit, err := strconv.Atoi(input)
	for err != nil {
		fmt.Println("\nInput tahun terbit harus berupa angka. Silakan coba lagi.")
		fmt.Scanln()
		return
	}

	// sorting data tahun terbit secara descending
	for i := 1; i < jumlahDataBuku; i++ {
		temp := data[i]
		j := i
		for j > 0 && (*data)[j-1].tahunTerbit > temp.tahunTerbit {
			data[j] = data[j-1]
			j = j - 1
		}
		data[j] = temp
	}

	found = true
	for found {
		index = FindDataBukuByTahunTerbitWithBinarySearch(*data, tahunTerbit)
		if index != -1 {
			for i := index; i < jumlahDataBuku-1; i++ {
				data[i] = data[i+1]
			}
			jumlahDataBuku--
		} else {
			found = false
		}
	}

	if jumlahDataBuku < MAXBUKU {
		fmt.Printf("\nSemua data buku dengan tahun terbit \"%d\" berhasil dihapus\n", tahunTerbit)
		fmt.Scanln()
	} else {
		fmt.Printf("\nMaaf, buku dengan tahun terbit \"%d\" tidak dapat ditemukan\n", tahunTerbit)
		fmt.Scanln()
	}
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
	fmt.Scanln()

	if found {
		fmt.Printf("\nMaaf, buku dengan kata kunci \"%s\" tidak dapat ditemukan.\n", kataKunci)
	}
}

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

func FindDataBukuByTahunTerbitWithBinarySearch(data dataBuku, tahunTerbit int) int {
	var found int = -1
	var kr int = 0
	var kn int = jumlahDataBuku - 1
	var med int

	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if tahunTerbit < data[med].tahunTerbit {
			kn = med - 1
		} else if tahunTerbit > data[med].tahunTerbit {
			kr = med + 1
		} else {
			found = med
		}
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

func FindDataBukuByPenerbitWithSequentialSearch(data dataBuku, penerbit string) int {
	var found int = -1
	var i int
	for i < jumlahDataBuku && found == -1 {
		if penerbit == data[i].penerbit {
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

func SortingDataBukuByTahunTerbitByInsertionSort(data *dataBuku) {
	for i := 1; i < jumlahDataBuku; i++ {
		temp := (*data)[i]
		j := i
		for j > 0 && (*data)[j-1].tahunTerbit < temp.tahunTerbit {
			(*data)[j] = (*data)[j-1]
			j = j - 1
		}
		data[j] = temp
	}
}

func SortingDataBukuByJudulBySelectionSort(data *dataBuku) {
	for i := 1; i < jumlahDataBuku; i++ {
		idxMin := i - 1
		for j := i; j < jumlahDataBuku; j++ {
			if data[idxMin].judul > data[j].judul {
				idxMin = j
			}
		}
		temp := data[idxMin]
		data[idxMin] = data[i-1]
		data[i-1] = temp
	}
}

func SortingDataBukuByPenulisBySelectionSort(data *dataBuku) {
	for i := 1; i < jumlahDataBuku; i++ {
		idxMin := i - 1
		for j := i; j < jumlahDataBuku; j++ {
			if (*data)[idxMin].penulis > (*data)[j].penulis {
				idxMin = j
			}
		}
		// Swap the books
		temp := data[idxMin]
		data[idxMin] = data[i-1]
		data[i-1] = temp
	}
}

func SortingDataBukuByPenerbitBySelectionSort(data *dataBuku) {
	for i := 1; i < jumlahDataBuku; i++ {
		idxMin := i - 1
		for j := i; j < jumlahDataBuku; j++ {
			if data[idxMin].penerbit > (*data)[j].penerbit {
				idxMin = j
			}
		}
		// Swap the books
		temp := data[idxMin]
		data[idxMin] = data[i-1]
		data[i-1] = temp
	}
}

func ViewPinjamMenu() {
	clearScreen()

	fmt.Println("=============================================")
	fmt.Println("|        Menu Pinjam Buku Perpustakaan      |")
	fmt.Println("=============================================")
	fmt.Println("\n1. Tampilkan Daftar Buku yang Dipinjam")
	fmt.Println("2. Tambah Data Peminjaman")
	fmt.Println("3. Ubah Data Peminjaman")
	fmt.Println("4. Hapus Data Peminjaman")
	fmt.Println("5. Kembali ke Menu Utama")
	fmt.Print("Masukkan pilihan (1/2/3/4/5/6): ")
}

func ViewCreateDataPinjamBuku() {
	clearScreen()

	fmt.Println("==========================================")
	fmt.Println("|        Tambah Data Peminjaman  Buku    |")
	fmt.Println("==========================================")
}

func ViewShowDataPinjamBuku() {
	clearScreen()

	fmt.Println("=========================================")
	fmt.Println("|      Daftar Data Peminjaman  Buku     |")
	fmt.Println("=========================================")
}

func ViewUpdateDataTanggalPinjamBuku() {
	clearScreen()

	fmt.Println("==============================================")
	fmt.Println("|        UBAH DATA TANGGAL PINJAM BUKU       |")
	fmt.Println("==============================================")
}

func ViewDeleteDataPinjamBuku() {
	clearScreen()

	fmt.Println("=========================================")
	fmt.Println("|      Hapus Data Peminjaman  Buku     |")
	fmt.Println("=========================================")
}

func PinjamMenu(dataPinjam *dataPinjamBuku, dataBuku *dataBuku) {
	var idPinjam pinjamBuku
	pilihan := 0
	keluar := false
	for !keluar {
		ViewPinjamMenu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			ShowDataPinjamBuku(dataPinjam, dataBuku, &idPinjam)
		case 2:
			AddDataPinjamBuku(dataPinjam, dataBuku)
		case 3:
			UpdateDataTanggalPinjamBuku(dataPinjam, dataBuku)
		case 4:
			DeleteDataPinjamBuku(dataPinjam, dataBuku)
		case 5:
			MainMenu(dataBuku, dataPinjam)
			keluar = true
		default:
			fmt.Println("Pilihan tidak valid!")
			fmt.Scanln()
		}

		if !keluar {
			fmt.Println("\nTekan Enter untuk kembali ke menu pinjam buku")
			fmt.Scanln()
		}
	}
}

func InputTanggal() tanggal {
	var tgl tanggal
	fmt.Print("\nMasukkan tanggal(1/2/3/../31)		: ")
	fmt.Scan(&tgl.hari)
	fmt.Print("Masukkan bulan(1/2/3/../12)		: ")
	fmt.Scan(&tgl.bulan)
	fmt.Print("Masukkan tahun				: ")
	fmt.Scan(&tgl.tahun)
	return tgl
}

func ValidasiTanggalPinjamDanKembali(tglPinjam, tglKembali tanggal) bool {
	if tglKembali.tahun > tglPinjam.tahun {
		return true
	} else if tglKembali.tahun == tglPinjam.tahun {
		if tglKembali.bulan > tglPinjam.bulan {
			return true
		} else if tglKembali.bulan == tglPinjam.bulan {
			if tglKembali.hari >= tglPinjam.hari {
				return true
			}
		}
	}
	return false
}

func AddDataPinjamBuku(dataPinjam *dataPinjamBuku, dataBuku *dataBuku) {
	var isbn string
	var pinjamBukuBaru pinjamBuku
	var inputJumlahDataPinjamBuku int
	var input string
	var foundISBN int

	ViewCreateDataPinjamBuku()

	if jumlahDataBuku == 0 {
		ViewShowDataBuku()
		fmt.Println("\nTidak ada data buku")
		fmt.Scanln()
		return
	} else {
		ViewShowDataBuku()
		fmt.Println("\nSemua Daftar Buku:")
		for i := 0; i < jumlahDataBuku; i++ {
			fmt.Println()
			buku := (*dataBuku)[i]
			if buku.isbn != "" {
				fmt.Printf("ISBN		: %s\n", buku.isbn)
				fmt.Printf("Judul		: %s\n", buku.judul)
				fmt.Printf("Penulis		: %s\n", buku.penulis)
				fmt.Printf("Penerbit	: %s\n", buku.penerbit)
				fmt.Printf("Tahun Terbit	: %d\n", buku.tahunTerbit)
				fmt.Printf("Jumlah Halaman	: %d\n", buku.jumlahHalaman)
				if !buku.status {
					fmt.Printf("Status Buku	: %s\n", "Buku tersedia")
				} else {
					fmt.Printf("Status Buku	: %s\n", "Buku sedang dipinjam")
				}
				fmt.Println("===================================")
			}
		}
		fmt.Scanln()
	}

	fmt.Println()

	fmt.Print("Masukkan ID Anggota			: ")
	fmt.Scan(&pinjamBukuBaru.idAnggota)
	fmt.Print("Masukkan jumlah buku yang akan dipinjam	: ")
	fmt.Scan(&input)

	inputJumlahDataPinjamBuku, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("\nInput jumlah data buku yang akan dipinjam harus berupa angka. Silakan coba lagi.")
		fmt.Scanln()
		return
	}

	for i := 0; i < inputJumlahDataPinjamBuku; i++ {
		validISBN := false
		for !validISBN {
			fmt.Print("Masukkan ISBN				: ")
			fmt.Scan(&isbn)
			foundISBN = FindDataBukuByISBNWithSequentialSearch(*dataBuku, isbn)
			if foundISBN != -1 {
				if dataBuku[foundISBN].status {
					fmt.Println("Maaf, buku dengan ISBN tersebut sudah dipinjam.")
				} else {
					validISBN = true
				}
			} else {
				fmt.Printf("\nMaaf, buku dengan ISBN \"%s\" tidak dapat ditemukan\n", isbn)
			}
		}

		if validISBN {
			pinjamBukuBaru.isbn = dataBuku[foundISBN].isbn
			dataPinjam[jumlahDataPinjam] = pinjamBukuBaru
			fmt.Println("\nBuku berhasil dipinjam")

			jumlahDataPinjam++
			dataBuku[foundISBN].jumlahBukuDipinjam++
			dataBuku[foundISBN].status = true
		}
	}
	pinjamBukuBaru.tgl_pinjam = InputTanggal()
	pinjamBukuBaru.tgl_kembali = tanggal{0, 0, 0}

	for i := 0; i < inputJumlahDataPinjamBuku; i++ {
		dataPinjam[jumlahDataPinjam-inputJumlahDataPinjamBuku+i].tgl_pinjam = pinjamBukuBaru.tgl_pinjam
		dataPinjam[jumlahDataPinjam-inputJumlahDataPinjamBuku+i].tgl_kembali = pinjamBukuBaru.tgl_kembali
	}

	fmt.Scanln()
}


func ShowDataPinjamBuku(dataPinjam *dataPinjamBuku, dataBuku *dataBuku, idPinjam *pinjamBuku) {
	if jumlahDataPinjam == 0 {
		ViewShowDataPinjamBuku()
		fmt.Println("\nTidak ada data buku yang sedang dipinjam")
	} else {
		ViewShowDataPinjamBuku()
		for i := 0; i < jumlahDataPinjam; i++ {
			pinjam := (*dataPinjam)[i]
			found := false

			for j := 0; j < jumlahDataBuku; j++ {
				if pinjam.isbn == (*dataBuku)[j].isbn {
					found = true
					buku := (*dataBuku)[j]

					fmt.Println("ID Anggota            : ", pinjam.idAnggota)
					fmt.Println("ISBN                  : ", buku.isbn)
					fmt.Println("Judul                 : ", buku.judul)
					fmt.Println("Penulis               : ", buku.penulis)
					fmt.Println("Penerbit              : ", buku.penerbit)
					fmt.Println("Tahun Terbit          : ", buku.tahunTerbit)
					fmt.Printf("Tanggal dipinjam      :  %02d-%02d-%04d\n", pinjam.tgl_pinjam.hari, pinjam.tgl_pinjam.bulan, pinjam.tgl_pinjam.tahun)

					if pinjam.tgl_kembali.hari == 0 && pinjam.tgl_kembali.bulan == 0 && pinjam.tgl_kembali.tahun == 0 {
						fmt.Println("Tanggal dikembalikan  :  Buku masih dipinjam")
					} else {
						fmt.Printf("Tanggal dikembalikan  :  %02d-%02d-%04d\n", pinjam.tgl_kembali.hari, pinjam.tgl_kembali.bulan, pinjam.tgl_kembali.tahun)
					}

					fmt.Println("Jumlah Buku Dipinjam  : ", buku.jumlahBukuDipinjam)
					fmt.Println("=====================================")
				}
			}

			if !found {
				fmt.Println("Maaf, buku dengan ISBN ", pinjam.isbn, " tidak dapat ditemukan")
			}
		}
	}
	fmt.Scanln()
}

func UpdateDataTanggalPinjamBuku(dataPinjam *dataPinjamBuku, dataBuku *dataBuku) {
	var indexFound, isbnFound int
	var isbn string

	ViewUpdateDataTanggalPinjamBuku()
	fmt.Print("Masukkan ISBN Buku yang ingin diubah		: ")
	fmt.Scan(&isbn)
	indexFound = FindDataPinjamBukuByISBNWithSequentialSearch(*dataPinjam, isbn)
	isbnFound = FindDataBukuByISBNWithSequentialSearch(*dataBuku, isbn)

	if indexFound != -1 && dataBuku[isbnFound].status {
		fmt.Print("Masukkan tanggal kembali ")
		tanggalKembali := InputTanggal()
		tanggalPinjam := (*dataPinjam)[indexFound].tgl_pinjam
		if ValidasiTanggalPinjamDanKembali(tanggalPinjam, tanggalKembali) {
			dataPinjam[indexFound].tgl_kembali = tanggalKembali
			// Hitung hari pinjam untuk tarif dan denda
			hariPinjam := HitungHariPinjam(tanggalPinjam, tanggalKembali)
			tarif, denda := HitungTarifDanDenda(hariPinjam)
			dataPinjam[indexFound].denda = denda

			fmt.Printf("\nTanggal kembali berhasil diubah\nTarif: %.2f, Denda: %.2f\n", tarif, denda)
			fmt.Println("Total Tarif : ", tarif+denda)
			dataBuku[isbnFound].status = false
			fmt.Scanln()
		} else {
			fmt.Println("Tanggal kembali tidak valid, lebih kecil dari tanggal pinjam")
			fmt.Scanln()
		}
	} else {
		fmt.Println("Maaf, ID Pinjam Buku tersebut tidak dapat ditemukan")
		fmt.Scanln()
	}
}

func DeleteDataPinjamBuku(dataPinjam *dataPinjamBuku, dataBuku *dataBuku) {
	ViewDeleteDataPinjamBuku()
	var indexFound, isbnFound int
	var isbn string
	fmt.Print("Masukkan ISBN Buku yang batal untuk dipinjam		: ")
	fmt.Scan(&isbn)
	indexFound = FindDataPinjamBukuByISBNWithSequentialSearch(*dataPinjam, isbn)
	isbnFound = FindDataBukuByISBNWithSequentialSearch(*dataBuku, isbn)
	if indexFound != -1 {
		for i := indexFound; i < jumlahDataPinjam-1; i++ {
			dataPinjam[i] = dataPinjam[i+1]
		}
		jumlahDataPinjam--
		dataBuku[isbnFound].jumlahBukuDipinjam--
		dataBuku[isbnFound].status = false
		fmt.Println("\nData berhasil dihapus")
		fmt.Scanln()
	} else {
		fmt.Println("\nMaaf, ISBN Buku tersebut tidak dapat ditemukan")
		fmt.Scanln()
	}
}

func HitungHariPinjam(tglPinjam, tglKembali tanggal) int {
	daysInMonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if tglKembali.tahun%4 == 0 && (tglKembali.tahun%100 != 0 || tglKembali.tahun%400 == 0) {
		daysInMonth[1] = 29
	}

	hariPinjam := 0

	if tglPinjam.bulan == tglKembali.bulan && tglPinjam.tahun == tglKembali.tahun {
		hariPinjam = tglKembali.hari - tglPinjam.hari
	} else {
		hariPinjam += daysInMonth[tglPinjam.bulan-1] - tglPinjam.hari
		month := tglPinjam.bulan
		year := tglPinjam.tahun
		for !(month == tglKembali.bulan && year == tglKembali.tahun) {
			month++
			if month > 12 {
				month = 1
				year++
			}
			hariPinjam += daysInMonth[month-1]
		}

		hariPinjam += tglKembali.hari
	}
	return hariPinjam
}

func HitungTarifDanDenda(hariPinjam int) (tarif, denda float64) {
	const batasHari = 7        // Misal batas pinjam 7 hari
	const tarifHarian = 1000.0 // Tarif per hari
	const dendaHarian = 500.0  // Denda per hari setelah batas hari

	if hariPinjam <= batasHari {
		tarif = float64(hariPinjam) * tarifHarian
		denda = 0
	} else {
		tarif = float64(batasHari) * tarifHarian
		denda = float64(hariPinjam-batasHari) * dendaHarian
	}
	return tarif, denda
}

func Top5BukuTerfavorit(dataBuku *dataBuku) {
	ViewTop5BukuTerfavorit()
	if jumlahDataBuku == 0 {
		fmt.Println("\nTidak ada data buku.")
		fmt.Scanln()
	} else {
		SortingTop5FavoritByInsertionSort(dataBuku)

		fmt.Println("\n5 Buku Terfavorit:")
		for i := 0; i < 5 && i < jumlahDataBuku; i++ {
			buku := dataBuku[i]
			fmt.Println("Judul		: ", buku.judul)
			fmt.Println("Penulis		: ", buku.penulis)
			fmt.Println("Penerbit	: ", buku.penerbit)
			fmt.Println("Jumlah Dipinjam	: ", buku.jumlahBukuDipinjam)
			fmt.Println("=====================================")
		}
		fmt.Scanln()
	}
}

func SortingTop5FavoritByInsertionSort(dataBuku *dataBuku) {
	var i, j int
	for i = 1; i < jumlahDataBuku; i++ {
		temp := dataBuku[i]
		j = i - 1
		for j >= 0 && temp.jumlahBukuDipinjam > dataBuku[j].jumlahBukuDipinjam {
			dataBuku[j+1] = dataBuku[j]
			j--
		}
		dataBuku[j+1] = temp
	}
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

func main() {
	LoginMenu()
}
