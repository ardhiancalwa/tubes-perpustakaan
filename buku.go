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
}

type pinjamBuku struct {
	isbn        int
	tgl_pinjam  int
	tgl_kembali int
	denda       float64
}

type dataBuku [MAXBUKU]buku
type dataPinjamBuku [MAXBUKU]pinjamBuku

func MainMenu(data *dataBuku, n int) {
	pilihan := 0
	keluar := false

	for !keluar {
		ViewMainMenu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			BukuMenu(data, n)
		case 2:
			PinjamMenu(data, n)
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

func PinjamMenu(data *dataBuku, n int) {
	fmt.Println("Belum dibuat brohh")
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
	fmt.Println("3. Keluar")

	// Meminta input user
	fmt.Print("Masukkan pilihan (1/2/3): ")
}

func ViewBukuMenu() {
	// Membersihkan layar sebelum menampilkan menu
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul menu dengan garis pembatas
	fmt.Println("===============================")
	fmt.Println("|          Menu Buku          |")
	fmt.Println("===============================")
	fmt.Println("\n1. Tambah Buku")
	fmt.Println("2. Tampilkan Daftar Buku")
	fmt.Println("3. Ubah Buku")
	fmt.Println("4. Hapus Buku")
	fmt.Println("5. Kembali")
	fmt.Print("Masukkan pilihan (1/2/3/4/5): ")
}

func ViewPinjamMenu() {
	// Membersihkan layar sebelum menampilkan menu
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul menu dengan garis pembatas
	fmt.Println("===============================")
	fmt.Println("|         Menu Pinjam         |")
	fmt.Println("===============================")
	fmt.Println("\n1. Tambah Pinjam")
	fmt.Println("2. Tampilkan Daftar Pinjam")
	fmt.Println("3. Ubah Pinjam")
	fmt.Println("4. Hapus Pinjam")
	fmt.Println("5. Kembali")
	fmt.Print("Masukkan pilihan: ")
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
	fmt.Println("Masukkan ISBN yang ingin diubah: ")
}

func ViewDeleteDataBuku() {
	// Membersihkan layar sebelum menampilkan form tambah data
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul form tambah data
	fmt.Println("===============================")
	fmt.Println("|        HAPUS DATA BUKU      |")
	fmt.Println("===============================")
	fmt.Print("Masukkan ISBN buku yang ingin dihapus: ")
}

func BukuMenu(data *dataBuku, n int) {
	var tahunTerbit, jumlahHalaman int
	var judul, penulis, penerbit, isbn string
	pilihan := 0
	keluar := false // Boolean variable to control the loop

	for !keluar { // Continue looping until keluar is true
		ViewBukuMenu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			CreateDataBuku(data, &n)
		case 2:
			ReadDataBuku(data, &n)
		case 3:
			UpdateDataBuku(data, isbn, judul, penulis, penerbit, tahunTerbit, jumlahHalaman, n)
		case 4:
			DeleteDataBuku(data, isbn, &n)
		case 5:
			MainMenu(data, n)
			keluar = true // Set keluar to true to exit the loop
		default:
			fmt.Println("Pilihan tidak valid!")
		}

		if !keluar { // Check if we should continue showing the menu
			// Tampilkan menu kembali setelah setiap operasi CRUD
			fmt.Scanln() // Menunggu user menekan Enter
		}
	}
}

func CreateDataBuku(data *dataBuku, n *int) {
	var bukuBaru buku
	var isBerhasil bool
	var jumlahBukuBaru int

	ViewCreateDataBuku()
	// Menjelaskan format input
	fmt.Println("\n!!! Ubah spasi dengan _ !!!")

	// Meminta input jumlah buku baru
	fmt.Print("Masukkan jumlah buku yang akan ditambahkan: ")
	fmt.Scan(&jumlahBukuBaru)
	isBerhasil = false
	for i := 0; i < jumlahBukuBaru; i++ {
		// Meminta input ISBN
		fmt.Print("\nISBN: ")
		fmt.Scan(&bukuBaru.isbn)
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

		if (*data)[*n] == (buku{}) && !isBerhasil {
			(*data)[*n] = bukuBaru
			fmt.Println("\nBuku berhasil ditambahkan")
			isBerhasil = false
			*n++ // Update the global count of books
		} else {
			fmt.Println("\nBuku gagal ditambahkan")
			isBerhasil = true
		}
	}
	jumlahBukuBaru++
	fmt.Println("\nTekan Enter untuk kembali ke menu buku")
	fmt.Scanln()
}

func UpdateDataBuku(data *dataBuku, isbn string, judulBaru string, penulisBaru string, penerbitBaru string, tahunTerbitBaru int, jumlahHalamanBaru int, n int) {
	var index int
	ViewUpdateDataBuku()
	fmt.Scan(&isbn)

	// update data
	index = FindDataWithSequentialSearch(*data, isbn, n)
	if index != -1 {
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

		(*data)[index].judul = judulBaru
		(*data)[index].penulis = penulisBaru
		(*data)[index].penerbit = penerbitBaru
		(*data)[index].tahunTerbit = tahunTerbitBaru
		(*data)[index].jumlahHalaman = jumlahHalamanBaru
		fmt.Println("Buku berhasil diubah")
	} else {
		fmt.Println("Error, Buku dengan ISBN tersebut tidak ditemukan")
	}
}

func DeleteDataBuku(data *dataBuku, isbn string, n *int) {
	var index int
	ViewDeleteDataBuku()
	fmt.Scan(&isbn)

	index = FindDataWithSequentialSearch(*data, isbn, *n)
	for i := index; i <= *n-1; i++ {
		data[i] = data[i+1]
	}
	*n = *n - 1
	fmt.Println("Data buku berhasil dihapus")
}

func FindDataWithSequentialSearch(data dataBuku, isbn string, n int) int {
	var found int = -1
	var i int
	for i < n && found == -1 {
		if isbn == data[i].isbn {
			found = i
		}
		i++
	}
	return found
}

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

func ReadDataBuku(data *dataBuku, n *int) {
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul daftar buku
	fmt.Println("===============================")
	fmt.Println("|         DAFTAR BUKU         |")
	fmt.Println("===============================")
	SortingDataWithSelectionSort(data, n)
	if *n == 0 {
		// Menampilkan pesan jika tidak ada data buku
		fmt.Println("\nTidak ada data buku")
		fmt.Println("\nTekan Enter untuk kembali ke menu utama")
		fmt.Scanln()
	} else {
		// Menampilkan daftar buku
		for i := 0; i < *n; i++ {
			fmt.Println()
			buku := (*data)[i]
			if buku.judul != "" { // Periksa apakah buku memiliki judul
				fmt.Printf("ISBN: %s\n", buku.isbn)
				fmt.Printf("Judul: %s\n", buku.judul)
				fmt.Printf("Penulis: %s\n", buku.penulis)
				fmt.Printf("Penerbit: %s\n", buku.penerbit)
				fmt.Printf("Tahun Terbit: %d\n", buku.tahunTerbit)
				fmt.Printf("Jumlah Halaman: %d\n", buku.jumlahHalaman)
				fmt.Println("=======================")
			}
		}
		fmt.Println("\nTekan Enter untuk kembali ke menu utama")
		fmt.Scanln()
	}
}

func SortingDataWithSelectionSort(data *dataBuku, n *int) {
	for i := 0; i < *n-1; i++ {
		minIndex := i
		for j := i + 1; j < *n; j++ {
			if (*data)[j].tahunTerbit < (*data)[minIndex].tahunTerbit {
				minIndex = j
			}
		}

		if i != minIndex {
			(*data)[i], (*data)[minIndex] = (*data)[minIndex], (*data)[i]
		}
	}
}

func main() {
	var data dataBuku
	var nData int
	MainMenu(&data, nData)
}
