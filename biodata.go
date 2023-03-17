package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	printMember()
}

type Member struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func printMember() {

	allMembers := []Member{
		{Nama: "Andi", Alamat: "Sumatera Utara", Pekerjaan: "Polisi", Alasan: "Belajar"},
		{Nama: "Budi", Alamat: "Sumatera Selatan", Pekerjaan: "Dosen", Alasan: "Belajar"},
		{Nama: "Carol", Alamat: "Papua", Pekerjaan: "Mahasiswa", Alasan: "Belajar"},
		{Nama: "Daniels", Alamat: "Kalimatan Timur", Pekerjaan: "Guru", Alasan: "Belajar"},
		{Nama: "Efraim", Alamat: "Kalimantan Barat", Pekerjaan: "Dokter", Alasan: "Belajar"},
		{Nama: "Firdaus", Alamat: "Kalimantan Selatan", Pekerjaan: "Instruktur", Alasan: "Belajar"},
		{Nama: "Gerry", Alamat: "Sulawesi Utara", Pekerjaan: "Pelajar", Alasan: "Belajar"},
		{Nama: "Hardy", Alamat: "Jawab Barat", Pekerjaan: "Koki", Alasan: "Belajar"},
		{Nama: "Imam", Alamat: "Jawa Timur", Pekerjaan: "Guru Bhs Inggris", Alasan: "Belajar"},
		{Nama: "Januzaj", Alamat: "Jawa Tengah", Pekerjaan: "Gamers", Alasan: "Belajar"},
		{Nama: "Kelvin", Alamat: "Jakarta", Pekerjaan: "Conten Creator", Alasan: "Belajar"},
		{Nama: "Ramos", Alamat: "Nusa Tenggara Timur", Pekerjaan: "Editor Video", Alasan: "Belajar"},
	}

	args := os.Args
	var show int
	for show = 1; show < len(args); show++ {
		switch {
		case args[show] == strconv.Itoa(1):
			fmt.Printf("Nama : %v\nAlamat : %v\nPekerjaan : %v\nAlasan : %v\n", allMembers[0].Nama,
				allMembers[0].Alamat, allMembers[0].Pekerjaan, allMembers[0].Alasan)

		case args[show] == strconv.Itoa(2):
			fmt.Printf("Nama : %v\nAlamat : %v\nPekerjaan : %v\nAlasan : %v\n", allMembers[1].Nama,
				allMembers[1].Alamat, allMembers[1].Pekerjaan, allMembers[1].Alasan)

		case args[show] == strconv.Itoa(3):
			fmt.Printf("Nama : %v\nAlamat : %v\nPekerjaan : %v\nAlasan : %v\n", allMembers[2].Nama,
				allMembers[2].Alamat, allMembers[2].Pekerjaan, allMembers[2].Alasan)

		case args[show] == strconv.Itoa(4):
			fmt.Printf("Nama : %v\nAlamat : %v\nPekerjaan : %v\nAlasan : %v\n", allMembers[3].Nama,
				allMembers[3].Alamat, allMembers[3].Pekerjaan, allMembers[3].Alasan)

		case args[show] == strconv.Itoa(5):
			fmt.Printf("Nama : %v\nAlamat : %v\nPekerjaan : %v\nAlasan : %v\n", allMembers[4].Nama,
				allMembers[4].Alamat, allMembers[4].Pekerjaan, allMembers[1].Alasan)

		case args[show] == strconv.Itoa(6):
			fmt.Printf("Nama : %v\nAlamat : %v\nPekerjaan : %v\nAlasan : %v\n", allMembers[5].Nama,
				allMembers[5].Alamat, allMembers[5].Pekerjaan, allMembers[5].Alasan)

		case args[show] == strconv.Itoa(7):
			fmt.Printf("Nama : %v\nAlamat : %v\nPekerjaan : %v\nAlasan : %v\n", allMembers[6].Nama,
				allMembers[6].Alamat, allMembers[6].Pekerjaan, allMembers[6].Alasan)

		case args[show] == strconv.Itoa(8):
			fmt.Printf("Nama : %v\nAlamat : %v\nPekerjaan : %v\nAlasan : %v\n", allMembers[7].Nama,
				allMembers[7].Alamat, allMembers[7].Pekerjaan, allMembers[7].Alasan)

		case args[show] == strconv.Itoa(9):
			fmt.Printf("Nama : %v\nAlamat : %v\nPekerjaan : %v\nAlasan : %v\n", allMembers[8].Nama,
				allMembers[8].Alamat, allMembers[8].Pekerjaan, allMembers[8].Alasan)

		case args[show] == strconv.Itoa(10):
			fmt.Printf("Nama : %v\nAlamat : %v\nPekerjaan : %v\nAlasan : %v\n", allMembers[9].Nama,
				allMembers[9].Alamat, allMembers[9].Pekerjaan, allMembers[9].Alasan)

		case args[show] == strconv.Itoa(11):
			fmt.Printf("Nama : %v\nAlamat : %v\nPekerjaan : %v\nAlasan : %v\n", allMembers[10].Nama,
				allMembers[10].Alamat, allMembers[10].Pekerjaan, allMembers[10].Alasan)

		case args[show] == strconv.Itoa(12):
			fmt.Printf("Nama : %v\nAlamat : %v\nPekerjaan : %v\nAlasan : %v\n", allMembers[11].Nama,
				allMembers[11].Alamat, allMembers[11].Pekerjaan, allMembers[11].Alasan)
		default:
			fmt.Println("Data belum tersedia, Silahkan tambah ada")
		}
	}
}
