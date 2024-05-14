package main

import "fmt"

const MAX_PINJAMAN int = 1000

type pinjamBuku1 struct {
	isbn int
	tgl_pinjam int
	tgl_kembali int
	denda float64
}

type dataPinjamanBuku [MAX_PINJAMAN]pinjamBuku

func TambahPinjamanBuku(data *dataPinjamanBuku, n *int) {
	// var isbn int
	var pinjam pinjamBuku
	var isBerhasil bool
	var jumlahPinjam int

	fmt.Println()
	fmt.Println("TAMBAH PINJAMAN BUKU")
	fmt.Println("!!! Ubah spasi dengan _ !!!")
	fmt.Print("Masukkan jumlah buku yang ingin dipinjam: ")
	fmt.Scan(&jumlahPinjam)
	isBerhasil = false
	for i := 0; i< jumlahPinjam; i++ {
		fmt.Print("ISBN: ")
		fmt.Scan(&pinjam.isbn)
		fmt.Print("Tanggal Pinjam: ")
		fmt.Scan(&pinjam.tgl_pinjam)
		fmt.Print("Tanggal Kembali: ")
		fmt.Scan(&pinjam.tgl_kembali)
		fmt.Print("Denda: ")

		if (*data)[*n] == (pinjamBuku{}) && !isBerhasil {
			(*data)[*n] = pinjam
			fmt.Println("Data pinjam buku berhasil ditambahkan")
			isBerhasil = true
			*n++
		} else {
			fmt.Println("Data pinjam buku gagal ditambahkan")
			isBerhasil = false
		}
	}
	fmt.Println()
}

func denda() {

}

func UbahPinjamanBuku() {

}

func HapusPinjamanBuku() {

}

func CetakPinjamanBuku() {

}

func main() {
	var data dataPinjamanBuku
	var nData int
	TambahPinjamanBuku(&data, &nData)
}