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
	var i, pass, idx, pilihan int
	var temp Komik
	var status string

	fmt.Println("Apakah anda ingin diurutkan? y/n")
	fmt.Scan(&status)

	if status == "y" || status == "iya" {

		fmt.Println("1. Jumlah terbesar ke terkecil")
		fmt.Println("2. Jumlah terkecil ke terbesar")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			for pass = 1; pass <= n-1; pass++ {
				idx = pass - 1
				i = pass

				for i < n {
					if A[idx].jumlah < A[i].jumlah {
						idx = i
					}
					i++
				}

				temp = A[pass-1]
				A[pass-1] = A[idx]
				A[idx] = temp
			}

		} else if pilihan == 2 {
			for pass = 1; pass <= n-1; pass++ {
				idx = pass - 1
				i = pass

				for i < n {
					if A[idx].jumlah > A[i].jumlah {
						idx = i
					}
					i++
				}

				temp = A[pass-1]
				A[pass-1] = A[idx]
				A[idx] = temp
			}
		}
	}

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
			fmt.Println("Status : terdaftar")
		} else {
			fmt.Println("Status : tidak terdaftar")
		}

		fmt.Println()
	}
}

func listpeminjam(P arrpeminjam, n int) {
	var pass, i, pilihan int
	var status string
	var temp peminjam

	fmt.Println("Apakah anda ingin diurutkan? y/n")
	fmt.Scan(&status)

	if status == "y" || status == "iya" {

		fmt.Println("1. denda terbesar ke terkecil")
		fmt.Println("2. denda terkecil ke terbesar")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			for pass = 1; pass > n; pass++ {
				i = pass
				temp = P[pass]
				for i > 0 && temp.denda < P[i-1].denda {
					P[i] = P[i-1]
					i--
				}
				P[i] = temp
			}
		} else if pilihan == 2 {
			for pass = 1; pass < n; pass++ {
				i = pass
				temp = P[pass]
				for i > 0 && temp.denda < P[i-1].denda {
					P[i] = P[i-1]
					i--
				}
				P[i] = temp
			}
		} else {
			println("angka yang dipilih tidak tersedia ")
		}
	}

	for i = 0; i < n; i++ {
		fmt.Println("ID :", P[i].id)
		fmt.Println("ID Komik :", P[i].idkomik)
		fmt.Println("ID Member :", P[i].idmember)
		fmt.Println("Jumlah :", P[i].jumlah)
		fmt.Println("Terlambat :", P[i].keterlambatan)
		fmt.Println("Denda :", P[i].denda)
	}
}

func carikomik(A arrbook, n, id int) int {
	var left, right, mid int

	left = 0
	right = n - 1

	for left <= right {
		mid = (left + right) / 2
		if A[mid].id == id {
			return mid
		} else if A[mid].id > id {
			right = mid - 1
		} else if A[mid].id < id {
			left = mid + 1
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
	var nKomik, nMember, nPeminjam, menu, subMenu, totalDenda, totalUang int
	var ulang string

	ulang = "y"

	for ulang == "iya" || ulang == "y" {

		fmt.Println("\n------- MENU UTAMA -------")
		fmt.Println("1. Kelola Komik")
		fmt.Println("2. Kelola Member")
		fmt.Println("3. Kelola Peminjaman")
		fmt.Println("4. Keuangan & Pendapatan")
		fmt.Println("0. Keluar")

		fmt.Print("Masukan pilihan : ")
		fmt.Scan(&menu)

		if menu == 1 {
			fmt.Println("\n--- Kelola Komik ---")
			fmt.Println("1. Tambah Komik")
			fmt.Println("2. List Komik")
			fmt.Println("3. Edit Komik")
			fmt.Println("4. Hapus Komik")
			fmt.Println("5. Cari Komik")
			fmt.Print("Pilih aksi : ")
			fmt.Scan(&subMenu)

			if subMenu == 1 {
				fmt.Print("Jumlah komik : ")
				fmt.Scan(&nKomik)
				masukankomik(&komik, nKomik)
			} else if subMenu == 2 {
				listkomik(komik, nKomik)
			} else if subMenu == 3 {
				editkomik(&komik, nKomik)
			} else if subMenu == 4 {
				hapuskomik(&komik, &nKomik)
			} else if subMenu == 5 {
				tampilCariKomik(komik, nKomik)
			} else {
				fmt.Println("Aksi tidak valid")
			}

		} else if menu == 2 {
			fmt.Println("\n--- Kelola Member ---")
			fmt.Println("1. Tambah Member")
			fmt.Println("2. List Member")
			fmt.Println("3. Edit Member")
			fmt.Println("4. Hapus Member")
			fmt.Println("5. Cari Member")
			fmt.Print("Pilih aksi : ")
			fmt.Scan(&subMenu)

			if subMenu == 1 {
				fmt.Print("Jumlah member : ")
				fmt.Scan(&nMember)
				masukanmember(&member, nMember)
			} else if subMenu == 2 {
				listmember(member, nMember)
			} else if subMenu == 3 {
				editmember(&member, nMember)
			} else if subMenu == 4 {
				hapusmember(&member, &nMember)
			} else if subMenu == 5 {
				tampilCariMember(member, nMember)
			} else {
				fmt.Println("Aksi tidak valid")
			}

		} else if menu == 3 {
			fmt.Println("\n--- Kelola Peminjaman ---")
			fmt.Println("1. Tambah Peminjam")
			fmt.Println("2. List Peminjam")
			fmt.Println("3. Edit Peminjam")
			fmt.Println("4. Hapus Peminjam")
			fmt.Print("Pilih aksi : ")
			fmt.Scan(&subMenu)

			if subMenu == 1 {
				fmt.Print("Jumlah peminjam : ")
				fmt.Scan(&nPeminjam)
				masukanpeminjam(&peminjam, nPeminjam)
			} else if subMenu == 2 {
				listpeminjam(peminjam, nPeminjam)
			} else if subMenu == 3 {
				editpeminjam(&peminjam, nPeminjam)
			} else if subMenu == 4 {
				hapuspeminjam(&peminjam, &nPeminjam)
			} else {
				fmt.Println("Aksi tidak valid")
			}

		} else if menu == 4 {
			fmt.Println("\n--- Keuangan & Pendapatan ---")
			fmt.Println("1. Hitung Total Denda")
			fmt.Println("2. Hitung Total Uang")
			fmt.Print("Pilih aksi : ")
			fmt.Scan(&subMenu)

			if subMenu == 1 {
				totalDenda = hitungTotalDenda(peminjam, nPeminjam)
				fmt.Println("Total Denda :", totalDenda)
			} else if subMenu == 2 {
				totalDenda = hitungTotalDenda(peminjam, nPeminjam)
				totalUang = hitungTotalUang(nMember, totalDenda)
				fmt.Println("Total Uang :", totalUang)
			} else {
				fmt.Println("Aksi tidak valid")
			}

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
