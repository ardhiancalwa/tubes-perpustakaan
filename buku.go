package main

import "fmt"

const MAXBUKU int = 1000

type buku struct {
	isbn          int
	judul         string
	penulis       string
	penerbit      string
	tahunTerbit   int
	jumlahHalaman int
}

type dataBuku [MAXBUKU]buku

// Function Create Data Buku
func TambahBuku(data *dataBuku) {
	var judul, penulis, penerbit string
	var isbn, tahunTerbit, jumlahHalaman int
	var bukuBaru buku
	var berhasil bool

	fmt.Println("TAMBAH BUKU BARU")
	fmt.Println("!!! Ubah spasi dengan _ !!!")
	fmt.Print("ISBN: ")
	fmt.Scan(&isbn)
	fmt.Print("Judul: ")
	fmt.Scan(&judul)
	fmt.Print("Penulis: ")
	fmt.Scan(&penulis)
	fmt.Print("Penerbit: ")
	fmt.Scan(&penerbit)
	fmt.Print("Tahun Terbit: ")
	fmt.Scan(&tahunTerbit)
	fmt.Print("Jumlah Halaman: ")
	fmt.Scan(&jumlahHalaman)

	bukuBaru = buku{
		isbn:          isbn,
		judul:         judul,
		penulis:       penulis,
		penerbit:      penerbit,
		tahunTerbit:   tahunTerbit,
		jumlahHalaman: jumlahHalaman,
	}

	berhasil = false
	for i := 0; i < len(data); i++ {
		/*
			data[i] == buku {} digunakan untuk memeriksa apakah
			data index ke i kosong.
			buku{} digunakan untuk mewakili elemen yang kosong.
			jika benar data[i] kosong, maka percabangan akan dijalankan.
		*/
		if (data)[i] == (buku{}) && !berhasil {
			(data)[i] = bukuBaru
			fmt.Println("Buku berhasil ditambahkan")
			berhasil = true
		}
	}

	if !berhasil {
		fmt.Println("Error, Array sudah penuh")
	}
}

// Function Update Data Buku
func UbahBuku(data *dataBuku, isbn int) {
	var judulBaru, penulisBaru, penerbitBaru string
	var tahunTerbitBaru, jumlahHalamanBaru int
	var bukuDitemukan bool

	fmt.Println("UBAH DATA BUKU")
	fmt.Println("Masukkan ISBN yang ingin diubah: ")
	fmt.Scan(&isbn)

	bukuIndex := -1
	for i := 0; i < len(*data); i++ {
		if (*data)[i].isbn == isbn {
			bukuIndex = i
		}
	}

	bukuDitemukan = bukuIndex != -1

	if bukuDitemukan {
		fmt.Println("Masukkan Judul Baru: ")
		fmt.Scan(&judulBaru)
		fmt.Println("Masukkan Penulis Baru: ")
		fmt.Scan(&penulisBaru)
		fmt.Println("Masukkan Penerbit Baru: ")
		fmt.Scan(&penerbitBaru)
		fmt.Println("Masukkan Tahun Terbit Baru: ")
		fmt.Scan(&tahunTerbitBaru)
		fmt.Println("Masukkan Jumlah Halaman Baru: ")
		fmt.Scan(&jumlahHalamanBaru)

		(*data)[bukuIndex].judul = judulBaru
		(*data)[bukuIndex].penulis = penulisBaru
		(*data)[bukuIndex].penerbit = penerbitBaru
		(*data)[bukuIndex].tahunTerbit = tahunTerbitBaru
		(*data)[bukuIndex].jumlahHalaman = jumlahHalamanBaru
		fmt.Println("Buku berhasil diubah")
	} else {
		fmt.Println("Error, Buku dengan ISBN tersebut tidak ditemukan")
	}
}

func HapusBuku(data dataBuku, isbn string) dataBuku {
	fmt.Println("Hapus Data Buku")
	fmt.Print("Masukkan ISBN buku yang ingin dihapus: ")
	fmt.Scanln(&isbn)

	bukuIndex := -1
	for i := bukuIndex + 1; i < len(data); i++ {
		data[i-1] = data[i]
	}
	data[len(data)-1] = buku{}
	fmt.Println("Data buku berhasil dihapus")
	return data
}

func CetakBuku(data *dataBuku) {
	fmt.Println("DAFTAR BUKU")
	fmt.Println("=======================")
	for i := 0; i < len(data); i++ {
		buku := (*data)[i]
		if buku.judul != "" { // Periksa apakah buku memiliki judul
			fmt.Printf("ISBN: %d\nJudul: %s\nPenulis: %s\nPenerbit: %s\nTahun Terbit: %d\nJumlah Halaman: %d\n", buku.isbn, buku.judul, buku.penulis, buku.penerbit, buku.tahunTerbit, buku.jumlahHalaman)
			fmt.Println("=======================")
		}
	}
}

func main() {
	var data dataBuku
	var isbn int
	TambahBuku(&data)
	CetakBuku(&data)
	UbahBuku(&data, isbn)
	CetakBuku(&data)
	// HapusBuku(data, isbn)
}
