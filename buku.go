package main

import "fmt"

const MAXBUKU int = 1000

type buku struct {
	isbn          string
	judul         string
	penulis       string
	penerbit      string
	tahunTerbit   int
	jumlahHalaman int
	status        bool
}

type pinjamBuku struct {
	isbn               string
	idPinjamBuku       int
	idAnggota          int
	tgl_pinjam         int
	tgl_kembali        int
	jumlahPinjam       int
	denda              float64
	status             bool
	jumlahBukuDipinjam int
}

var jumlahDataBuku, jumlahDataPinjam int

type dataBuku [MAXBUKU]buku
type dataPinjamBuku [MAXBUKU]pinjamBuku

// menu utama
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
			fmt.Println("Keluar dari program")
			keluar = true
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

// tampilan menu utama
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
	fmt.Println("3. Keluar")

	// Meminta input user
	fmt.Print("Masukkan pilihan (1/2/3): ")
}

// tampilan menu buku
func ViewBukuMenu() {
	// Membersihkan layar sebelum menampilkan menu
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul menu dengan garis pembatas
	fmt.Println("===============================")
	fmt.Println("|          Menu Buku          |")
	fmt.Println("===============================")
	fmt.Println("\n1. Tambah Data Buku")
	fmt.Println("2. Tampilkan Daftar Buku")
	fmt.Println("3. Ubah Data Buku")
	fmt.Println("4. Hapus Data Buku")
	fmt.Println("5. Kembali")
	fmt.Print("Masukkan pilihan (1/2/3/4/5): ")
}  

func ViewDeleteDataBukuMenu() {
	// Membersihkan layar sebelum menampilkan menu
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul menu dengan garis pembatas
	fmt.Println("================================")
    fmt.Println("|     Menu Hapus Data Buku    |")
    fmt.Println("================================")
    fmt.Println("\n1. Hapus Data Buku Berdasarkan ISBN")
    fmt.Println("2. Hapus Data Buku Berdasarkan Judul")
    fmt.Println("3. Hapus Data Buku Berdasarkan Penulis")
    fmt.Println("4. Hapus Data Buku Berdasarkan Tahun Terbit")
    fmt.Println("5. Kembali ke Menu Utama")
    fmt.Print("Masukkan Pilihan (1/2/3/4/5): ")
}

// tampilan prosedur menambahkan data buku
func ViewCreateDataBuku() {
	// Membersihkan layar sebelum menampilkan form tambah data
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul form tambah data
	fmt.Println("===============================")
	fmt.Println("|    TAMBAH DATA BUKU BARU    |")
	fmt.Println("===============================")
}

// tampilan prosedur mengubah data buku
func ViewUpdateDataBuku() {
	// Membersihkan layar sebelum menampilkan form tambah data
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul form tambah data
	fmt.Println("===============================")
	fmt.Println("|       UBAH DATA BUKU        |")
	fmt.Println("===============================")
}

// tampilan prosedur menghapus data buku
func ViewDeleteDataBuku() {
	// Membersihkan layar sebelum menampilkan form tambah data
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul form tambah data
	fmt.Println("===============================")
	fmt.Println("|        HAPUS DATA BUKU      |")
	fmt.Println("===============================")
	fmt.Print("Masukkan ISBN buku yang ingin dihapus: ")
}

func ViewReadDataBuku() {
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul daftar buku
	fmt.Println("===============================")
	fmt.Println("|         DAFTAR BUKU         |")
	fmt.Println("===============================")
}

func BukuMenu(dataBuku *dataBuku, dataPinjam *dataPinjamBuku) {
	pilihan := 0
	keluar := false // Boolean variable to control the loop

	for !keluar { // Continue looping until keluar is true
		ViewBukuMenu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			CreateDataBuku(dataBuku)
		case 2:
			ReadDataBuku(dataBuku)
		case 3:
			UpdateDataBuku(dataBuku)
		case 4:
			DeleteDataBukuMenu(dataBuku, *dataPinjam)
		case 5:
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

// menambahkan data buku
func CreateDataBuku(data *dataBuku) {
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
		fmt.Print("\nISBN: ")
		fmt.Scan(&bukuBaru.isbn)
		for j := 0; j < jumlahDataBuku; j++ {
			if (*data)[j].isbn == bukuBaru.isbn {
				isDuplicate = true
				// Jika ISBN sudah ada, minta ISBN lain
				fmt.Println("\nBuku dengan ISBN yang sama sudah ada. Silakan masukkan ISBN lain.")
				i-- // Mengurangi indeks untuk mengulang input buku ini
			}
		}
		if !isDuplicate {
			// Meminta input judul
			fmt.Print("Judul: ")
			fmt.Scan(&bukuBaru.judul)
			// Meminta input penulis
			fmt.Print("Penulis: ")
			fmt.Scan(&bukuBaru.penulis)
			// Meminta input penerbit
			fmt.Print("Penerbit: ")
			fmt.Scan(&bukuBaru.penerbit)
			// Meminta input tahun terbit
			fmt.Print("Tahun Terbit: ")
			fmt.Scan(&bukuBaru.tahunTerbit)
			// Meminta input jumlah halaman
			fmt.Print("Jumlah Halaman: ")
			fmt.Scan(&bukuBaru.jumlahHalaman)
			bukuBaru.status = false

			if (*data)[jumlahDataBuku] == (buku{}) && !isBerhasil {
				(*data)[jumlahDataBuku] = bukuBaru
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

// mengubah data buku
func UpdateDataBuku(data *dataBuku) {
	var indexFound int
	var isbnBuku string
	var judulBaru, penulisBaru, penerbitBaru string
	var tahunTerbitBaru, jumlahHalamanBaru int
	ViewUpdateDataBuku()
	fmt.Println("Masukkan ISBN yang ingin diubah: ")
	fmt.Scan(&isbnBuku)

	// update data
	indexFound = FindDataWithSequentialSearch(*data, isbnBuku)
	if indexFound != -1 {
		fmt.Print("Masukkan Judul Baru: ")
		fmt.Scan(&judulBaru)
		fmt.Print("Masukkan Penulis Baru: ")
		fmt.Scan(&penulisBaru)
		fmt.Print("Masukkan Penerbit Baru: ")
		fmt.Scan(&penerbitBaru)
		fmt.Print("Masukkan Tahun Terbit Baru: ")
		fmt.Scan(&tahunTerbitBaru)
		fmt.Print("Masukkan Jumlah Halaman Baru: ")
		fmt.Scan(&jumlahHalamanBaru)

		(*data)[indexFound].judul = judulBaru
		(*data)[indexFound].penulis = penulisBaru
		(*data)[indexFound].penerbit = penerbitBaru
		(*data)[indexFound].tahunTerbit = tahunTerbitBaru
		(*data)[indexFound].jumlahHalaman = jumlahHalamanBaru
		fmt.Println("Buku berhasil diubah")
	} else {
		fmt.Println("Maaf, buku dengan ISBN tersebut tidak dapat ditemukan")
	}
}

// menghapus data buku by isbn
func DeleteDataBukuByISBN(data *dataBuku) {
	var index int
	var isbn string

	ViewDeleteDataBuku()
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

// delete data buku by tahun terbit
func DeleteDataBukuByJudul(data *dataBuku) {
	var index int
	var judul string

	ViewDeleteDataBuku()
	fmt.Scan(&judul)

	index = FindDataBukuByJudulWithSequentialSearch(*data, judul)
	if index != -1 {
		for i := index; i <= jumlahDataBuku-1; i++ {
			data[i] = data[i+1]
		}
		jumlahDataBuku = jumlahDataBuku - 1
		fmt.Printf("Data buku dengan judul %s berhasil dihapus \n", judul)
	} else {
		fmt.Println("\nMaaf, buku dengan judul tersebut tidak dapat ditemukan")
		fmt.Println("\nTekan Enter untuk kembali ke menu utama")
		fmt.Scanln()
	}
}

// delete data buku by tahun terbit
func DeleteDataBukuByPenulis(data *dataBuku) {
	var index int
	var penulis string

	ViewDeleteDataBuku()
	fmt.Scan(&penulis)

	index = FindDataBukuByPenulisWithSequentialSearch(*data, penulis)
	if index != -1 {
		for i := index; i <= jumlahDataBuku-1; i++ {
			data[i] = data[i+1]
		}
		jumlahDataBuku = jumlahDataBuku - 1
		fmt.Printf("Data buku dengan penulis %s berhasil dihapus \n", penulis)
	} else {
		fmt.Println("\nMaaf, buku dengan penulis tersebut tidak dapat ditemukan")
		fmt.Println("\nTekan Enter untuk kembali ke menu utama")
		fmt.Scanln()
	}	
}

// delete data buku by tahun terbit
func DeleteDataBukuByTahunTerbit(data *dataBuku) {
	var index int
	var tahunTerbit int

	ViewDeleteDataBuku()
	fmt.Scan(&tahunTerbit)

	index = FindDataBukuByTahunTerbitWithSequentialSearch(*data, tahunTerbit)
	if index != -1 {
		for i := index; i <= jumlahDataBuku-1; i++ {
			data[i] = data[i+1]
		}
		jumlahDataBuku = jumlahDataBuku - 1
		fmt.Printf("Data buku dengan tahun terbit %d berhasil dihapus", tahunTerbit)
	} else {
		fmt.Println("\nMaaf, buku dengan tahun terbit tersebut tidak dapat ditemukan")
		fmt.Println("\nTekan Enter untuk kembali ke menu utama")
		fmt.Scanln()
	}
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

// mencetak data buku
func ReadDataBuku(data *dataBuku) {
	SortingDataBukuByTahunTerbit(data)
	if jumlahDataBuku == 0 {
		ViewReadDataBuku()
		// Menampilkan pesan jika tidak ada data buku
		fmt.Println("\nTidak ada data buku")
		fmt.Println("\nTekan Enter untuk kembali ke menu buku")
		fmt.Scanln()
	} else {
		ViewReadDataBuku()
		// Menampilkan daftar buku
		for i := 0; i < jumlahDataBuku; i++ {
			fmt.Println()
			buku := (*data)[i]
			if buku.judul != "" { // Periksa apakah buku memiliki judul
				fmt.Printf("ISBN		: %s\n", buku.isbn)
				fmt.Printf("Judul		: %s\n", buku.judul)
				fmt.Printf("Penulis		: %s\n", buku.penulis)
				fmt.Printf("Penerbit	: %s\n", buku.penerbit)
				fmt.Printf("Tahun Terbit	: %d\n", buku.tahunTerbit)
				fmt.Printf("Jumlah Halaman	: %d\n", buku.jumlahHalaman)
				if buku.status == true {
					fmt.Printf("Status Buku	: %s\n", "Buku sedang dipinjam")
				} else {
					fmt.Printf("Status Buku	: %s\n", "Buku tersedia")
				}
				fmt.Println("===================================")
			}
		}
		fmt.Println("\nTekan Enter untuk kembali ke menu utama")
		fmt.Scanln()
	}
}

// mengurutkan data buku menggunakan selection sort
func SortingDataBukuByTahunTerbit(data *dataBuku) {

	for i := 0; i < jumlahDataBuku-1; i++ {
		maxIndex := i
		for j := i + 1; j < jumlahDataBuku; j++ {
			if (*data)[j].tahunTerbit > (*data)[maxIndex].tahunTerbit {
				maxIndex = j
			}
		}

		if i != maxIndex {
			(*data)[i], (*data)[maxIndex] = (*data)[maxIndex], (*data)[i]
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

func ViewCreateDataPinjamBuku() {
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul menu dengan garis pembatas
	fmt.Println("==========================================")
	fmt.Println("|        Tambah Data Peminjaman  Buku    |")
	fmt.Println("==========================================")
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
			ReadDataPinjamBuku(dataPinjam, dataBuku)
		case 2:
			CreateDataPinjamBuku(dataPinjam, dataBuku)
		case 3:
			UpdateDataPinjamBuku()
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

func CreateDataPinjamBuku(dataPinjam *dataPinjamBuku, dataBuku *dataBuku) {
	var isbn string
	var pinjamBukuBaru pinjamBuku
	var inputJumlahDataPinjamBuku int
	var isBerhasil bool

	isBerhasil = false

	ViewCreateDataPinjamBuku()
	ReadDataBuku(dataBuku)

	fmt.Println("\n!!! Ubah spasi dengan _ !!!")
	fmt.Print("Masukkan ID Anggota: ")
	fmt.Scan(&pinjamBukuBaru.idAnggota)
	fmt.Print("Masukkan jumlah buku yang akan dipinjam: ")
	fmt.Scan(&inputJumlahDataPinjamBuku)
	for i := 0; i < inputJumlahDataPinjamBuku; i++ {
		fmt.Print("Masukkan ISBN: ")
		fmt.Scan(&isbn)
		foundISBN := FindDataWithSequentialSearch(*dataBuku, isbn)
		if foundISBN != -1 {
			fmt.Print("Masukkan Tanggal Peminjaman: ")
			fmt.Scan(&pinjamBukuBaru.tgl_pinjam)
			fmt.Print("Masukkan Tanggal Pengembalian: ")
			fmt.Scan(&pinjamBukuBaru.tgl_kembali)
			dataBuku[i].status = true

			if (*dataPinjam)[jumlahDataPinjam] == (pinjamBuku{}) && !isBerhasil {
				(*dataPinjam)[jumlahDataPinjam] = pinjamBukuBaru
				fmt.Println("\nBuku berhasil ditambahkan")
				isBerhasil = false
				dataPinjam[i].jumlahBukuDipinjam++
				jumlahDataPinjam++
			} else {
				fmt.Println("\nBuku gagal ditambahkan")
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

func ReadDataPinjamBuku(dataPinjam *dataPinjamBuku, dataBuku *dataBuku) {
	if jumlahDataPinjam == 0 {
		fmt.Println("\nTidak ada data buku yang sedang dipinjam")
		fmt.Println("\nTekan Enter untuk kembali ke menu buku")
		fmt.Scanln()
	} else {
		for i := 0; i < jumlahDataPinjam; i++ {
			pinjam := (*dataPinjam)[i]
			bukuIndex := FindDataPinjamBukuByISBNWithSequentialSearch(*dataPinjam, pinjam.isbn)
			if bukuIndex != -1 {
				fmt.Println("ID Anggota		: ", pinjam.idAnggota)
				fmt.Println("ISBN			: ", (*dataBuku)[bukuIndex].isbn)
				fmt.Println("Judul			: ", (*dataBuku)[bukuIndex].judul)
				fmt.Println("Penulis			: ", (*dataBuku)[bukuIndex].penulis)
				fmt.Println("Penerbit		: ", (*dataBuku)[bukuIndex].penerbit)
				fmt.Println("Tahun Terbit		: ", (*dataBuku)[bukuIndex].tahunTerbit)
				fmt.Println("Tanggal dipinjam	: ", pinjam.tgl_pinjam)
				fmt.Println("Tanggal dikembalikan	: ", pinjam.tgl_kembali)
				fmt.Println("Jumlah Buku Dipinjam	: ", pinjam.jumlahBukuDipinjam)
				fmt.Println("=====================================")
			} else {
				fmt.Println("Maaf, buku dengan ISBN tersebut tidak dapat ditemukan")
			}
			fmt.Println("\nTekan Enter untuk kembali ke menu pinjam")
			fmt.Scanln()
		}
	}
}

func UpdateDataPinjamBuku() {
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

// func FindDataBukuByISBNWithSequentialSearch(dataBuku *dataBuku, isbn string) bool {
// 	var n int
// 	var found bool = false
// 	var j int = 0

// 	jumlahDataPinjam = n
// 	for j < n && !found {
// 		if dataBuku[j].isbn == isbn {
// 			found = true
// 		} else {
// 			j++
// 		}
// 	}
// 	return found
// }

func main() {
	var dataBuku dataBuku
	var dataPinjam dataPinjamBuku
	MainMenu(&dataBuku, &dataPinjam)

}
