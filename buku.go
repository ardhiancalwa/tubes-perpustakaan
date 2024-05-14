package main

import "fmt"

const MAXBUKU int = 1000

type buku struct {
	isbn          int
	judul         string
	penulis       string
	penerbit      string
	tahunTerbit   string
	jumlahHalaman string
}

type pinjamBuku struct {
	isbn        int
	tgl_pinjam  int
	tgl_kembali int
	denda       float64
}

type dataBuku [MAXBUKU]buku
type dataPinjamBuku [MAXBUKU]pinjamBuku

func menuUtama(data *dataBuku, n int) {
	pilihan := 0
	keluar := false

	for !keluar {
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
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			menuBuku(data, n)
		case 2:
			menuPinjam(data, n)
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

func menuPinjam(data *dataBuku, n int) {
	fmt.Println("Belum dibuat brohh")
}

func menuBuku(data *dataBuku, n int) {
	var isbn int
	var judul, penulis, penerbit, tahunTerbit, jumlahHalaman string
	pilihan := 0
	keluar := false // Boolean variable to control the loop

	for !keluar { // Continue looping until keluar is true
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
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			TambahBuku(data, &n)
		case 2:
			CetakBuku(data, &n)
		case 3:
			UbahBuku(data, isbn, judul, penulis, penerbit, tahunTerbit, jumlahHalaman, n)
		case 4:
			HapusBuku(*data, isbn, &n)
		case 5:
			menuUtama(data, n)
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

// Create Data Buku
func TambahBuku(data *dataBuku, n *int) {
	var isbn int
	var bukuBaru buku
	var isBerhasil, isISBN bool
	var jumlahBukuBaru int

	// Membersihkan layar sebelum menampilkan form tambah data
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul form tambah data
	fmt.Println("===============================")
	fmt.Println("|    TAMBAH DATA BUKU BARU    |")
	fmt.Println("===============================")
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
		isISBN = isUniqueISBN(data, isbn, *n)
		for !isISBN {
			fmt.Println("ISBN yang telah anda masukkan sudah dipakai. Silahkan masukkan ISBN yang lain.")
			fmt.Print("Masukkan ISBN ulang: ")
			fmt.Scan(&bukuBaru.isbn)
			isUniqueISBN(data, bukuBaru.isbn, *n)
		}
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
	fmt.Println("\nTekan Enter untuk kembali ke menu buku")
	fmt.Scanln()
}

// Function to check if the ISBN is unique
func isUniqueISBN(data *dataBuku, isbn int, n int) bool {
	for i := 0; i < n; i++ {
		if (*data)[i].isbn == isbn {
			return false // ISBN is not unique
		}
	}
	return true // ISBN is unique
}

// Update Data Buku
func UbahBuku(data *dataBuku, isbn int, judulBaru string, penulisBaru string, penerbitBaru string, tahunTerbitBaru string, jumlahHalamanBaru string, n int) {
	var index int

	fmt.Println("UBAH DATA BUKU")
	fmt.Println("Masukkan ISBN yang ingin diubah: ")
	fmt.Scan(&isbn)

	// update data
	index = findDataWithSequentialSearch(*data, isbn)
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

func findDataWithSequentialSearch(data dataBuku, isbn int) int {
	var index int = -1
	var i int = 0
	for i < len(data) && index == -1 {
		if data[i].isbn == isbn {
			index = i
		}
		i++
	}
	return index
}

func HapusBuku(data dataBuku, isbn int, n *int) {
	var index int
	fmt.Println("Hapus Data Buku")
	fmt.Print("Masukkan ISBN buku yang ingin dihapus: ")
	fmt.Scan(&isbn)

	index = findDataWithBinarySearch(data, isbn, *n)
	for i := index; i <= *n-1; i++ {
		data[i] = data[i+1]
	}
	*n = *n - 1
	fmt.Println("Data buku berhasil dihapus")
}

func findDataWithBinarySearch(data dataBuku, isbn int, n int) int {
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

func CetakBuku(data *dataBuku, n *int) {
	fmt.Print("\033[H\033[2J")

	// Menampilkan judul daftar buku
	fmt.Println("===============================")
	fmt.Println("|         DAFTAR BUKU         |")
	fmt.Println("===============================")
	selectionSort(data, n)
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
				fmt.Printf("ISBN: %d\n", buku.isbn)
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

func selectionSort(data *dataBuku, n *int) {
    for i := 0; i < *n-1; i++ {
        minIndex := i
        for j := i + 1; j < *n; j++ {
            if (*data)[j].isbn < (*data)[minIndex].isbn {
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
	// var isbn, tahunTerbit, jumlahHalaman int
	// var judul, penulis, penerbit string
	// create data
	menuUtama(&data, nData)
}
