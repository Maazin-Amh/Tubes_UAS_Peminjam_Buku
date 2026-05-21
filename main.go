package main

import "fmt"

const NMAX int = 1000

type Buku struct {
	judul_buku string
	jumlah     int
}

type Peminjam struct {
	nama       string
	judul_buku string
	jumlah     int
	status     bool
}

type arrbook [NMAX]Buku
type arrpeminjam [NMAX]Peminjam

func masukanbuku(A *arrbook, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Println("masukan judul buku :")
		fmt.Scan(&A[i].judul_buku)

		fmt.Println("masukan jumlah buku :")
		fmt.Scan(&A[i].jumlah)
	}
}

func masukanpeminjam(B *arrpeminjam, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Println("masukan nama peminjam :")
		fmt.Scan(&B[i].nama)
		fmt.Println("masukan judul buku yang ingin dipinjam :")
		fmt.Scan(&B[i].judul_buku)
		fmt.Println("masukan jumlah yang ingin dipinjam :")
		fmt.Scan(&B[i].jumlah)
	}
}

func logiclist(A arrbook)  {
	
}

func logicpeminjaman(A *arrbook, B *arrpeminjam, n int) {
	var i int

	for i = 0; i < n; i++ {
		if B[i].judul_buku == A[i].judul_buku {
			if A[i].jumlah >= B[i].jumlah {
				A[i].jumlah = A[i].jumlah - B[i].jumlah
				B[i].status = true
			} else {
				B[i].status = false
			}
		}
	}
}

func cetak(A arrbook, B arrpeminjam, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Println("nama", B[i].nama)

		if B[i].status == true {
			fmt.Println("Meminjam")
		} else {
			fmt.Println("Tidak meminjam")
		}

		fmt.Println("Judul buku :", A[i].judul_buku)
		fmt.Println("sisa jumlah buku :", A[i].jumlah)
	}

}

func main() {

	var buku arrbook
	var peminjam arrpeminjam
	var n int

	fmt.Println("Masukan nilai jumlah:")
	fmt.Scan(&n)

	masukanbuku(&buku, n)
	masukanpeminjam(&peminjam, n)
	logicpeminjaman(&buku, &peminjam, n)
	cetak(buku, peminjam, n)
}
