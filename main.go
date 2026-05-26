package main

import "fmt"

const NMAX int = 1000

type Komik struct {
	id         int
	judul_buku string
	jumlah     int
}

type member struct {
	id     int
	nama   string
	status bool
}

type peminjam struct {
	id            int
	idkomik       int
	idmember      int
	jumlah        int
	keterlambatan int
	denda         int
}

type arrbook [NMAX]Komik
type arrmember [NMAX]member
type arrpeminjam [NMAX]peminjam

func masukankomik(A *arrbook, n int) {
	var i int

	for i = 0; i < n; i++ {

		A[i].id = i + 1

		fmt.Print("Masukan judul buku : ")
		fmt.Scan(&A[i].judul_buku)

		fmt.Print("Masukan jumlah buku : ")
		fmt.Scan(&A[i].jumlah)
	}
}

func masukanmember(B *arrmember, n int) {
	var i int

	for i = 0; i < n; i++ {

		B[i].id = i + 1

		fmt.Print("Masukan nama member : ")
		fmt.Scan(&B[i].nama)

		B[i].status = true
	}
}

func masukanpeminjam(P *arrpeminjam, n int) {
	var i int

	for i = 0; i < n; i++ {

		P[i].id = i + 1

		fmt.Print("Masukan ID Komik : ")
		fmt.Scan(&P[i].idkomik)

		fmt.Print("Masukan ID Member : ")
		fmt.Scan(&P[i].idmember)

		fmt.Print("Masukan jumlah pinjam : ")
		fmt.Scan(&P[i].jumlah)

		fmt.Print("Masukan hari keterlambatan : ")
		fmt.Scan(&P[i].keterlambatan)

		P[i].denda = P[i].keterlambatan * 5000
	}
}

func listkomik(A arrbook, n int) {
	var i int

	for i = 0; i < n; i++ {

		fmt.Println("ID :", A[i].id)
		fmt.Println("Judul :", A[i].judul_buku)
		fmt.Println("Jumlah :", A[i].jumlah)
		fmt.Println()
	}
}

func listmember(B arrmember, n int) {
	var i int

	for i = 0; i < n; i++ {

		fmt.Println("ID :", B[i].id)
		fmt.Println("Nama :", B[i].nama)

		if B[i].status {
			fmt.Println("Status : aktif")
		} else {
			fmt.Println("Status : tidak aktif")
		}

		fmt.Println()
	}
}

func listpeminjam(P arrpeminjam, n int) {
	var i int

	for i = 0; i < n; i++ {

		fmt.Println("ID :", P[i].id)
		fmt.Println("ID Komik :", P[i].idkomik)
		fmt.Println("ID Member :", P[i].idmember)
		fmt.Println("Jumlah :", P[i].jumlah)
		fmt.Println("Terlambat :", P[i].keterlambatan)
		fmt.Println("Denda :", P[i].denda)
		fmt.Println()
	}
}

func carikomik(A arrbook, n, id int) int {
	var i int

	for i = 0; i < n; i++ {
		if A[i].id == id {
			return i
		}
	}

	return -1
}

func carimember(B arrmember, n, id int) int {
	var i int

	for i = 0; i < n; i++ {
		if B[i].id == id {
			return i
		}
	}

	return -1
}

func caripeminjam(P arrpeminjam, n, id int) int {
	var i int

	for i = 0; i < n; i++ {
		if P[i].id == id {
			return i
		}
	}

	return -1
}

func tampilCariKomik(A arrbook, n int) {
	var id, idx int

	fmt.Print("Masukan ID komik : ")
	fmt.Scan(&id)

	idx = carikomik(A, n, id)

	if idx != -1 {

		fmt.Println("Data ditemukan")
		fmt.Println("ID :", A[idx].id)
		fmt.Println("Judul :", A[idx].judul_buku)
		fmt.Println("Jumlah :", A[idx].jumlah)

	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func tampilCariMember(B arrmember, n int) {
	var id, idx int

	fmt.Print("Masukan ID member : ")
	fmt.Scan(&id)

	idx = carimember(B, n, id)

	if idx != -1 {

		fmt.Println("Data ditemukan")
		fmt.Println("ID :", B[idx].id)
		fmt.Println("Nama :", B[idx].nama)

	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func editkomik(A *arrbook, n int) {
	var id, idx int

	fmt.Print("Masukan ID komik : ")
	fmt.Scan(&id)

	idx = carikomik(*A, n, id)

	if idx != -1 {

		fmt.Print("Masukan judul baru : ")
		fmt.Scan(&A[idx].judul_buku)

		fmt.Print("Masukan jumlah baru : ")
		fmt.Scan(&A[idx].jumlah)

		fmt.Println("Data berhasil diubah")

	} else {
		fmt.Println("Data tidak ada")
	}
}

func editmember(B *arrmember, n int) {
	var id, idx int

	fmt.Print("Masukan ID member : ")
	fmt.Scan(&id)

	idx = carimember(*B, n, id)

	if idx != -1 {

		fmt.Print("Masukan nama baru : ")
		fmt.Scan(&B[idx].nama)

		fmt.Println("Data berhasil diubah")

	} else {
		fmt.Println("Data tidak ada")
	}
}

func editpeminjam(P *arrpeminjam, n int) {
	var id, idx int

	fmt.Print("Masukan ID peminjam : ")
	fmt.Scan(&id)

	idx = caripeminjam(*P, n, id)

	if idx != -1 {

		fmt.Print("Masukan keterlambatan : ")
		fmt.Scan(&P[idx].keterlambatan)

		P[idx].denda = P[idx].keterlambatan * 5000

		fmt.Println("Data berhasil diubah")

	} else {
		fmt.Println("Data tidak ada")
	}
}

func hapuskomik(A *arrbook, n *int) {
	var id, idx, i int

	fmt.Print("Masukan ID : ")
	fmt.Scan(&id)

	idx = carikomik(*A, *n, id)

	if idx != -1 {

		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}

		*n--

		fmt.Println("Data berhasil dihapus")

	} else {
		fmt.Println("Data tidak ada")
	}
}

func hapusmember(B *arrmember, n *int) {
	var id, idx, i int

	fmt.Print("Masukan ID : ")
	fmt.Scan(&id)

	idx = carimember(*B, *n, id)

	if idx != -1 {

		for i = idx; i < *n-1; i++ {
			B[i] = B[i+1]
		}

		*n--

		fmt.Println("Data berhasil dihapus")

	} else {
		fmt.Println("Data tidak ada")
	}
}

func hapuspeminjam(P *arrpeminjam, n *int) {
	var id, idx, i int

	fmt.Print("Masukan ID : ")
	fmt.Scan(&id)

	idx = caripeminjam(*P, *n, id)

	if idx != -1 {

		for i = idx; i < *n-1; i++ {
			P[i] = P[i+1]
		}

		*n--

		fmt.Println("Data berhasil dihapus")

	} else {
		fmt.Println("Data tidak ada")
	}
}

func hitungTotalDenda(P arrpeminjam, n int) int {
	var i, total int

	for i = 0; i < n; i++ {
		total += P[i].denda
	}

	return total
}

func hitungTotalUang(nMember int, totalDenda int) int {
	const biayaDaftar = 10000

	return (nMember * biayaDaftar) + totalDenda
}

func main() {
	var komik arrbook
	var member arrmember
	var peminjam arrpeminjam
	var nKomik, nMember, nPeminjam, menu, totalDenda, totalUang int
	var ulang string

	ulang = "y"

	for ulang == "iya" || ulang == "y" {

		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Tambah Komik")
		fmt.Println("2. List Komik")
		fmt.Println("3. Edit Komik")
		fmt.Println("4. Hapus Komik")
		fmt.Println("5. Tambah Member")
		fmt.Println("6. List Member")
		fmt.Println("7. Edit Member")
		fmt.Println("8. Hapus Member")
		fmt.Println("9. Tambah Peminjam")
		fmt.Println("10. List Peminjam")
		fmt.Println("11. Edit Peminjam")
		fmt.Println("12. Hapus Peminjam")
		fmt.Println("13. Cari Komik")
		fmt.Println("14. Cari Member")
		fmt.Println("15. Hitung Total Denda")
		fmt.Println("16. Hitung Total Uang")
		fmt.Println("0. Keluar")

		fmt.Print("Masukan pilihan : ")
		fmt.Scan(&menu)

		if menu == 1 {
			fmt.Print("Jumlah komik : ")
			fmt.Scan(&nKomik)
			masukankomik(&komik, nKomik)
		} else if menu == 2 {
			listkomik(komik, nKomik)
		} else if menu == 3 {
			editkomik(&komik, nKomik)
		} else if menu == 4 {
			hapuskomik(&komik, &nKomik)
		} else if menu == 5 {
			fmt.Print("Jumlah member : ")
			fmt.Scan(&nMember)
			masukanmember(&member, nMember)
		} else if menu == 6 {
			listmember(member, nMember)
		} else if menu == 7 {
			editmember(&member, nMember)
		} else if menu == 8 {
			hapusmember(&member, &nMember)
		} else if menu == 9 {
			fmt.Print("Jumlah peminjam : ")
			fmt.Scan(&nPeminjam)
			masukanpeminjam(&peminjam, nPeminjam)
		} else if menu == 10 {
			listpeminjam(peminjam, nPeminjam)
		} else if menu == 11 {
			editpeminjam(&peminjam, nPeminjam)
		} else if menu == 12 {
			hapuspeminjam(&peminjam, &nPeminjam)
		} else if menu == 13 {
			tampilCariKomik(komik, nKomik)
		} else if menu == 14 {
			tampilCariMember(member, nMember)
		} else if menu == 15 {
			totalDenda = hitungTotalDenda(peminjam, nPeminjam)
			fmt.Println("Total Denda :", totalDenda)
		} else if menu == 16 {
			totalDenda = hitungTotalDenda(peminjam, nPeminjam)
			totalUang = hitungTotalUang(nMember, totalDenda)
			fmt.Println("Total Uang :", totalUang)

		} else if menu == 0 {
			fmt.Println("Program selesai")
			break
		} else {
			fmt.Println("Menu tidak tersedia")
		}

		fmt.Print("\nIngin input lagi? (y/n) : ")
		fmt.Scan(&ulang)
	}
}
