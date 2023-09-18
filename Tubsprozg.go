package main

import (
	"fmt"
	"math"
)

type course struct {
	nim  int
	nama string
	m1   string
	m2   string
	m3   string
	m4   string
	m5   string
	m6   string
	m7   string
	m8   string
	sks  int
}
type coursevalue struct {
	namaku     string
	namats     string
	namaus     string
	namaot     string
	namagr     string
	kuis       string
	uts        string
	uas        string
	total      string
	grade      string
	alproku    float64
	matdisku   float64
	mavekku    float64
	bindoku    float64
	bingku     float64
	kalkulusku float64
	sisdigku   float64
	pbdku      float64
	alprots    float64
	matdists   float64
	mavekts    float64
	bindots    float64
	bingts     float64
	kalkulusts float64
	sisdigts   float64
	pbdts      float64
	alprous    float64
	matdisus   float64
	mavekus    float64
	bindous    float64
	bingus     float64
	kalkulusus float64
	sisdigus   float64
	pbdus      float64
	alproot    float64
	matdisot   float64
	mavekot    float64
	bindoot    float64
	bingot     float64
	kalkulusot float64
	sisdigot   float64
	pbdot      float64
	gralpro    string
	grmatdis   string
	grmavek    string
	grbindo    string
	grbing     string
	grkalkulus string
	grsisdig   string
	grpbd      string
	grAll      string
	totalAll   float64
	ipk        float64
}
type matkulMaha [2023]course
type nilaiMaha [2023]coursevalue

func inputDataMaha(tab *matkulMaha) {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Scan(&tab[i].nim, &tab[i].nama, &tab[i].m1, &tab[i].m2, &tab[i].m3, &tab[i].m4, &tab[i].m5, &tab[i].m6, &tab[i].m7, &tab[i].m8, &tab[i].sks)
	}
}
func inputNilaiMaha(tab *nilaiMaha) {
	var j int
	for j = 0; j < 10; j++ {
		fmt.Scan(&tab[j].namaku, &tab[j].kuis, &tab[j].alproku, &tab[j].matdisku, &tab[j].mavekku, &tab[j].bindoku, &tab[j].bingku, &tab[j].kalkulusku, &tab[j].sisdigku, &tab[j].pbdku)
		fmt.Scan(&tab[j].namats, &tab[j].uts, &tab[j].alprots, &tab[j].matdists, &tab[j].mavekts, &tab[j].bindots, &tab[j].bingts, &tab[j].kalkulusts, &tab[j].sisdigts, &tab[j].pbdts)
		fmt.Scan(&tab[j].namaus, &tab[j].uas, &tab[j].alprous, &tab[j].matdisus, &tab[j].mavekus, &tab[j].bindous, &tab[j].bingus, &tab[j].kalkulusus, &tab[j].sisdigus, &tab[j].pbdus)

	}
}

// meghitung total semua nilai total dari semua matkul
func totalSemua(tab *nilaiMaha, tab2 *matkulMaha, k int) {
	var j int
	for j = 0; j < k; j++ {
		if tab2[j].sks == 23 {
			tab[j].totalAll = math.Round(((tab[j].alproot+tab[j].matdisot+tab[j].mavekot+tab[j].bindoot+tab[j].bingot+tab[j].kalkulusot+tab[j].sisdigot+tab[j].pbdot)*100)/8) / 100
		} else {
			tab[j].totalAll = math.Round(((tab[j].alproot+tab[j].matdisot+tab[j].mavekot+tab[j].bindoot+tab[j].bingot+tab[j].kalkulusot+tab[j].sisdigot+tab[j].pbdot)*100)/7) / 100

		}
	}
}
func printIPK(tab *nilaiMaha, Tab *matkulMaha, idx int) {
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	fmt.Print("\t", "Total", "\t\t", Tab[idx].sks, " SKS\t", "IPK: ", tab[idx].ipk, "\t\t", "\n")

}
func OutputdataAwalnilai(tab *nilaiMaha) {
	var j int
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	for j = 0; j < 10; j++ {
		fmt.Println("lokasi:", j)
		fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
		fmt.Println(tab[j].namaku, "\t", tab[j].kuis, "\t", tab[j].alproku, "\t", tab[j].matdisku, "\t", tab[j].mavekku, "\t", tab[j].bindoku, "\t", tab[j].bingku, "\t", tab[j].kalkulusku, "\t", tab[j].sisdigku, "\t", tab[j].pbdku)
		fmt.Println(tab[j].namats, "\t", tab[j].uts, "\t", tab[j].alprots, "\t", tab[j].matdists, "\t", tab[j].mavekts, "\t", tab[j].bindots, "\t", tab[j].bingts, "\t", tab[j].kalkulusts, "\t", tab[j].sisdigts, "\t", tab[j].pbdts)
		fmt.Println(tab[j].namaus, "\t", tab[j].uas, "\t", tab[j].alprous, "\t", tab[j].matdisus, "\t", tab[j].mavekus, "\t", tab[j].bindous, "\t", tab[j].bingus, "\t", tab[j].kalkulusus, "\t", tab[j].sisdigus, "\t", tab[j].pbdus)
		fmt.Println("---------------------------------------------------------------------------------------------------------------------------")

	}

}
func OutputNilaiMaha(tab *nilaiMaha, tab2 *matkulMaha, k int) {
	var j int
	for j = 0; j < k; j++ {
		totalperMatkul(tab, k)   //meghitung total permatkul
		totalSemua(tab, tab2, k) //meghitung total semua nilai dar total dari semua matkul
		hitungGrade(tab, k)      //menghitungg grade per matkul dari nilai total
		fmt.Println("lokasi:", j)
		printtranskrip(*tab, *tab2, j)
		gradekeNilai(tab, k)     //grade ke nilai digunakan untuk hitung ipk saja
		totalSemua(tab, tab2, k) //meghitung total keseluruhan nilai total dari semua matkul
		gradeAll(tab, k)         //grade keseluruhan(total dari semua matkul) untuk membantu hitung ipk
		totalperMatkul(tab, k)   //menghitung  total permatkul
		totalSemua(tab, tab2, k) //meghitung total semua nilai dari total dari semua matkul
		hitungIPK(tab, k)        //hitung ipk berdasarkan grAll dan totalAll yang mana didapat dari fungsi totalSemua ,dan totalSemua perlu totalperMatkul
		printIPK(tab, tab2, j)
		fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	}
}
func OutputDataMaha(tab *matkulMaha) {
	var j int
	for j = 0; j < 10; j++ {
		fmt.Println("lokasi:", j)
		fmt.Println(tab[j].nim, "\t", tab[j].nama, "\t", tab[j].m1, "\t", tab[j].m2, "\t", tab[j].m3, "\t", tab[j].m4, "\t", tab[j].m5, "\t", tab[j].m6, "\t", tab[j].m7, "\t", tab[j].m8, "\t", tab[j].sks)

	}
}

func tampilanmenu() {
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("                                       Menu                                     ")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("1a. Menambahkan data Matkul mahasiswa\n1b. Menambahkan data Nilai Mahasiswa\n2a. Mengubah data Matkul mahasiswa\n2b. Mengubah data Nilai mahasiswa\n3a. Menghapus data Matkul mahasiswa\n3b.Menghapus data nilai mahasiswa\n3c. Menghapus Matkul beserta nilai mahasiswa\n4a. Menampilkan daftar mahasiswa yang mengambil suatu matkul\n4b. Menampilkan daftar matakuliah yang diambil seorang mahasiswa\n5. Menampilkan daftar mahasiswa terurut descending berdasarkan nilai dan jumlah sks yang diambil\n6. Menampilkan daftar mahasiswa terurut ascending berdasarkan nilai dan sks yang diambil\n7. Menampilkan nama seorang mahasiswa dari total nilainya\n8. Menampilkan transkrip nilai mahasiswa\n9. Akhiri pengolahan data")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
}

// proses
func mainloop(tabmatkul *matkulMaha, tabnilai *nilaiMaha, pilihan string) {
	var j, k int
	j = 10
	k = 10
	for pilihan != "9" {
		tampilanmenu()
		fmt.Println("Masukan pilihan anda:")
		fmt.Scan(&pilihan)
		switch pilihan {
		case "1a":
			caller(tabmatkul, tabnilai, pilihan, j, k)
			outputAfterProcess(tabmatkul, tabnilai, pilihan, j, k)
			j++

		case "1b":
			caller(tabmatkul, tabnilai, pilihan, j, k)
			outputAfterProcess(tabmatkul, tabnilai, pilihan, j, k)
			k++
		case "2a":
			caller(tabmatkul, tabnilai, pilihan, j, k)
			outputAfterProcess(tabmatkul, tabnilai, pilihan, j, k)
		case "2b":
			caller(tabmatkul, tabnilai, pilihan, j, k)
			outputAfterProcess(tabmatkul, tabnilai, pilihan, j, k)
		case "3a":
			caller(tabmatkul, tabnilai, pilihan, j, k)
			outputAfterProcess(tabmatkul, tabnilai, pilihan, j, k)
			j--
		case "3b":
			caller(tabmatkul, tabnilai, pilihan, j, k)
			outputAfterProcess(tabmatkul, tabnilai, pilihan, j, k)
			k--
		case "3c":
			caller(tabmatkul, tabnilai, pilihan, j, k)
			outputAfterProcess(tabmatkul, tabnilai, pilihan, j, k)
		case "4a":
			caller(tabmatkul, tabnilai, pilihan, j, k)
		case "4b":
			caller(tabmatkul, tabnilai, pilihan, j, k)
		case "5":
			caller(tabmatkul, tabnilai, pilihan, j, k)
		case "6":
			caller(tabmatkul, tabnilai, pilihan, j, k)
		case "7":
			caller(tabmatkul, tabnilai, pilihan, j, k)
		case "8":
			caller(tabmatkul, tabnilai, pilihan, j, k)
		}

	}
	if pilihan == "9" {
		fmt.Println("===========================================================================================================================")
		fmt.Println("                                    Pengolahan Data Telah Selesai")
		fmt.Println("===========================================================================================================================")
	}
}

func outputAfterProcess(tabmatkul *matkulMaha, tabnilai *nilaiMaha, pilihan string, j, k int) {
	var i int
	totalperMatkul(tabnilai, j)
	hitungGrade(tabnilai, j)
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("          Data After Process          ")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	switch pilihan {
	case "1a":
		fmt.Println("      Data Mata Kuliah Mahasiswa                    ")
		for i = 0; i < j+1; i++ {
			fmt.Println("lokasi", i)
			fmt.Println(tabmatkul[i].nim, "\t", tabmatkul[i].nama, "\t", tabmatkul[i].m1, "\t", tabmatkul[i].m2, "\t", tabmatkul[i].m3, "\t", tabmatkul[i].m4, "\t", tabmatkul[i].m5, "\t", tabmatkul[i].m6, "\t", tabmatkul[i].m7, "\t", tabmatkul[i].m8, "\t", tabmatkul[i].sks)
		}
	case "1b":
		fmt.Println("Data Nilai Mahasiswa")
		OutputNilaiMaha(tabnilai, tabmatkul, k+1)
		k++

	case "2a":
		fmt.Println("Data Mata Kuliah Mahasiswa")
		for i = 0; i < j; i++ {
			fmt.Println(tabmatkul[i].nim, "\t", tabmatkul[i].nama, "\t", tabmatkul[i].m1, "\t", tabmatkul[i].m2, "\t", tabmatkul[i].m3, "\t", tabmatkul[i].m4, "\t", tabmatkul[i].m5, "\t", tabmatkul[i].m6, "\t", tabmatkul[i].m7, "\t", tabmatkul[i].m8, "\t", tabmatkul[i].sks)
		}
	case "2b":
		fmt.Println("Data Nilai Mahasiswa")
		OutputNilaiMaha(tabnilai, tabmatkul, k)

	case "3a":
		fmt.Println("Data Mata Kuliah Mahasiswa")
		for i = 0; i < j-1; i++ {
			fmt.Println("lokasi", i)
			fmt.Println(tabmatkul[i].nim, "\t", tabmatkul[i].nama, "\t", tabmatkul[i].m1, "\t", tabmatkul[i].m2, "\t", tabmatkul[i].m3, "\t", tabmatkul[i].m4, "\t", tabmatkul[i].m5, "\t", tabmatkul[i].m6, "\t", tabmatkul[i].m7, "\t", tabmatkul[i].m8, "\t", tabmatkul[i].sks)
		}
	case "3b":
		fmt.Println("Data Nilai Mahasiswa")
		OutputNilaiMaha(tabnilai, tabmatkul, k-1)
		k--

	case "3c":
		fmt.Println("Data Nilai Mahasiswa")
		OutputNilaiMaha(tabnilai, tabmatkul, k)
	}
}
func addDataMataMaha(tab *matkulMaha, j int) {
	var i int
	for i = j; i < j+1; i++ {
		fmt.Scan(&tab[i].nim, &tab[i].nama, &tab[i].m1, &tab[i].m2, &tab[i].m3, &tab[i].m4, &tab[i].m5, &tab[i].m6, &tab[i].m7, &tab[i].m8, &tab[i].sks)
	}
}

func addDataNilai(tab *nilaiMaha, k int) {
	var i int
	for i = k; i < k+1; i++ {
		fmt.Scan(&tab[i].namaku, &tab[i].kuis, &tab[i].alproku, &tab[i].matdisku, &tab[i].mavekku, &tab[i].bindoku, &tab[i].bingku, &tab[k].kalkulusku, &tab[i].sisdigku, &tab[i].pbdku)
		fmt.Scan(&tab[i].namats, &tab[i].uts, &tab[i].alprots, &tab[i].matdists, &tab[i].mavekts, &tab[i].bindots, &tab[i].bingts, &tab[k].kalkulusts, &tab[i].sisdigts, &tab[i].pbdts)
		fmt.Scan(&tab[i].namaus, &tab[i].uas, &tab[i].alprous, &tab[i].matdisus, &tab[i].mavekus, &tab[i].bindous, &tab[i].bingus, &tab[k].kalkulusus, &tab[i].sisdigus, &tab[i].pbdus)

	}

}

func editDataMatkul(tabmatkul *matkulMaha, lokasi int) {
	fmt.Scan(&tabmatkul[lokasi].nim, &tabmatkul[lokasi].nama, &tabmatkul[lokasi].m1, &tabmatkul[lokasi].m2, &tabmatkul[lokasi].m3, &tabmatkul[lokasi].m4, &tabmatkul[lokasi].m5, &tabmatkul[lokasi].m6, &tabmatkul[lokasi].m7, &tabmatkul[lokasi].m8, &tabmatkul[lokasi].sks)

}

func editDataNilai(tabnilai *nilaiMaha, lokasi int) {
	fmt.Scan(&tabnilai[lokasi].namaku, &tabnilai[lokasi].kuis, &tabnilai[lokasi].alproku, &tabnilai[lokasi].matdisku, &tabnilai[lokasi].mavekku, &tabnilai[lokasi].bindoku, &tabnilai[lokasi].bingku, &tabnilai[lokasi].kalkulusku, &tabnilai[lokasi].sisdigku, &tabnilai[lokasi].pbdku)
	fmt.Scan(&tabnilai[lokasi].namats, &tabnilai[lokasi].uts, &tabnilai[lokasi].alprots, &tabnilai[lokasi].matdists, &tabnilai[lokasi].mavekts, &tabnilai[lokasi].bindots, &tabnilai[lokasi].bingts, &tabnilai[lokasi].kalkulusts, &tabnilai[lokasi].sisdigts, &tabnilai[lokasi].pbdts)
	fmt.Scan(&tabnilai[lokasi].namaus, &tabnilai[lokasi].uas, &tabnilai[lokasi].alprous, &tabnilai[lokasi].matdisus, &tabnilai[lokasi].mavekus, &tabnilai[lokasi].bindous, &tabnilai[lokasi].bingus, &tabnilai[lokasi].kalkulusus, &tabnilai[lokasi].sisdigus, &tabnilai[lokasi].pbdus)

}

func deleteMatkul(tabmatkul *matkulMaha, idx, j int) {
	var i int
	for i = idx; i < j; i++ {
		tabmatkul[i] = tabmatkul[i+1]
	}
}

// Menghapus data nilai susatu mahasiswa berdasarkan lokasi as idx
func deleteDataNilai(tabnilai *nilaiMaha, idx, k int) {
	var i int
	for i = idx; i < k; i++ {
		tabnilai[i] = tabnilai[i+1]
	}
}
func deleteMatkulnilai(tabnilai *nilaiMaha, idx int, dihapus string) {
	switch dihapus {
	case "alpro":
		tabnilai[idx].alproku = 0
		tabnilai[idx].alprots = 0
		tabnilai[idx].alprous = 0
		tabnilai[idx].alproot = 0
		tabnilai[idx].gralpro = "E"
	case "matdis":
		tabnilai[idx].matdisku = 0
		tabnilai[idx].matdists = 0
		tabnilai[idx].matdisus = 0
		tabnilai[idx].matdisot = 0
		tabnilai[idx].grmatdis = "E"
	case "mavek":
		tabnilai[idx].mavekku = 0
		tabnilai[idx].mavekts = 0
		tabnilai[idx].mavekus = 0
		tabnilai[idx].mavekot = 0
		tabnilai[idx].grmavek = "E"
	case "bindo":
		tabnilai[idx].bindoku = 0
		tabnilai[idx].bindots = 0
		tabnilai[idx].bindous = 0
		tabnilai[idx].bindoot = 0
		tabnilai[idx].grbindo = "E"
	case "bing":
		tabnilai[idx].bingku = 0
		tabnilai[idx].bingts = 0
		tabnilai[idx].bingus = 0
		tabnilai[idx].bingot = 0
		tabnilai[idx].grbing = "E"
	case "kalkulus":
		tabnilai[idx].kalkulusku = 0
		tabnilai[idx].kalkulusts = 0
		tabnilai[idx].kalkulusus = 0
		tabnilai[idx].kalkulusot = 0
		tabnilai[idx].grkalkulus = "E"
	case "sisdig":
		tabnilai[idx].sisdigku = 0
		tabnilai[idx].sisdigts = 0
		tabnilai[idx].sisdigus = 0
		tabnilai[idx].sisdigot = 0
		tabnilai[idx].grsisdig = "E"

	case "pbd":
		tabnilai[idx].pbdku = 0
		tabnilai[idx].pbdts = 0
		tabnilai[idx].pbdus = 0
		tabnilai[idx].pbdot = 0
		tabnilai[idx].grpbd = "E"

	}

}
func listMahaOfMatkul(tabmatkul *matkulMaha, tabnilai *nilaiMaha, j int, matkulX string) {
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Daftar mahasiswa yang mengambil matkul tersebut")
	var i int
	switch matkulX {
	case "alpro":
		for i = 0; i < j; i++ {
			if tabmatkul[i].m1 != "----" && tabnilai[i].alproku != 0 {
				fmt.Println(tabmatkul[i].nama)
			}
		}
	case "matdis":
		for i = 0; i < j; i++ {
			if tabmatkul[i].m2 != "----" && tabnilai[i].matdisku != 0 {
				fmt.Println(tabmatkul[i].nama)
			}
		}
	case "mavek":
		for i = 0; i < j; i++ {
			if tabmatkul[i].m3 != "----" && tabnilai[i].mavekku != 0 {
				fmt.Println(tabmatkul[i].nama)
			}
		}
	case "bindo":
		for i = 0; i < j; i++ {
			if tabmatkul[i].m4 != "----" && tabnilai[i].bindoku != 0 {
				fmt.Println(tabmatkul[i].nama)
			}
		}
	case "bing":
		for i = 0; i < j; i++ {
			if tabmatkul[i].m5 != "----" && tabnilai[i].bingku != 0 {
				fmt.Println(tabmatkul[i].nama)
			}
		}
	case "kalkulus":
		for i = 0; i < j; i++ {
			if tabmatkul[i].m6 != "----" && tabnilai[i].kalkulusku != 0 {
				fmt.Println(tabmatkul[i].nama)
			}
		}
	case "sisdig":
		for i = 0; i < j; i++ {
			if tabmatkul[i].m7 != "----" && tabnilai[i].sisdigku != 0 {
				fmt.Println(tabmatkul[i].nama)
			}
		}
	case "pbd":
		for i = 0; i < j; i++ {
			if tabmatkul[i].m8 != "----" && tabnilai[i].pbdku != 0 {
				fmt.Println(tabmatkul[i].nama)
			}
		}
	}
}

func listMatkulOfMaha(tabmatkul *matkulMaha, tabnilai *nilaiMaha, idx int) {

	fmt.Println("Nama Mahasiswa", tabmatkul[idx].nama)
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	if tabmatkul[idx].m1 != "----" && tabnilai[idx].alproku != 0 {
		fmt.Println(tabmatkul[idx].m1)
	}
	if tabmatkul[idx].m2 != "----" && tabnilai[idx].matdisku != 0 {
		fmt.Println(tabmatkul[idx].m2)
	}
	if tabmatkul[idx].m3 != "----" && tabnilai[idx].mavekku != 0 {
		fmt.Println(tabmatkul[idx].m3)
	}
	if tabmatkul[idx].m4 != "----" && tabnilai[idx].bindoku != 0 {
		fmt.Println(tabmatkul[idx].m4)
	}
	if tabmatkul[idx].m5 != "----" && tabnilai[idx].bingku != 0 {
		fmt.Println(tabmatkul[idx].m5)
	}
	if tabmatkul[idx].m6 != "----" && tabnilai[idx].kalkulusku != 0 {
		fmt.Println(tabmatkul[idx].m6)
	}
	if tabmatkul[idx].m7 != "----" && tabnilai[idx].sisdigku != 0 {
		fmt.Println(tabmatkul[idx].m7)
	}
	if tabmatkul[idx].m8 != "----" && tabnilai[idx].pbdku != 0 {
		fmt.Println(tabmatkul[idx].m8)
	}
}

func outputTranskrip(tabnilai *nilaiMaha, tabmatkul *matkulMaha, k int) {
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("                Transkrip Nilai Mahasiswa           ")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("\t\t", "Alpro", "\t", "Madis", "\t", "Mavek", "\t", "Bindo", "\t", "Bing", "\t", "Kalku", "\t", "Sidig", "\t", "Pbd")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	OutputNilaiMaha(tabnilai, tabmatkul, k)
}

func caller(tabmatkul *matkulMaha, tabnilai *nilaiMaha, pilihan string, j, k int) {
	var lokasi, idx int
	switch pilihan {
	case "1a":
		fmt.Println("Masukan sesuai format data daftar matakuliah sebelumnya jika kosong isi dengan '----' untuk data daftar matakuliah")
		addDataMataMaha(tabmatkul, j)
	case "1b":
		fmt.Println("Masukan sesuai format data nilai mahasiswa sebelumnya,jika kosong isi dengan '--' ")
		fmt.Println("DATA YANG DIMASUKAN HARUS LENGKAP\nBaris pertama untuk kuis\nBaris kedua untuk uts\nBaris ketiga untuk uas\n")
		addDataNilai(tabnilai, k)
	case "2a":
		fmt.Println("Masukan nama mahasiswa data yang ingin diganti data matkulnya")
		var namaX string
		fmt.Scan(&namaX)
		lokasi = seqtabmatkul(tabmatkul, j, namaX)
		fmt.Println("Masukan data Mata kuliah yang baru, setiap kolom wajib diisi!")
		editDataMatkul(tabmatkul, lokasi)
	case "2b":
		fmt.Println("Masukan nama mahasiswa yang ingin diganti data nilainya")
		var namaX string
		fmt.Scan(&namaX)
		lokasi = seqtabnilai(tabnilai, k, namaX)
		fmt.Println("Masukan data Nilai yang baru,setiap kolom dan baris wajib diisi!")
		editDataNilai(tabnilai, lokasi)
	case "3a":
		fmt.Println("Masukan nama mahasiswa yang data mata kuliahnya ingin dihapus")
		var namaX string
		fmt.Scan(&namaX)
		idx = seqtabmatkul(tabmatkul, j, namaX)
		deleteMatkul(tabmatkul, idx, j)
	case "3b":
		fmt.Println("Masukan nama mahasiswa yang data nilainya yang ingin dihapus")
		var namaX string
		fmt.Scan(&namaX)
		idx = seqtabnilai(tabnilai, k, namaX)
		deleteDataNilai(tabnilai, idx, k)
	case "3c":
		fmt.Println("Masukan nama mahasiswa yang data matkulnya  ingin dihapus ")
		var namaX string
		fmt.Scan(&namaX)
		idx = seqtabnilai(tabnilai, k, namaX)
		fmt.Println("Masukan matkul yang ingin dihapus nilainya\n- alpro\n- matdis\n- mavek\n- bindo\n- bing\n- kalkulus\n- sisdig\n- pbd")
		var dihapus string
		fmt.Scan(&dihapus)
		deleteMatkulnilai(tabnilai, idx, dihapus)
	case "4a":
		fmt.Println("Masukan nama matkul yang ingin ditampilkan daftar mahasiswanya\n- alpro\n- matdis\n- mavek\n- bindo\n- bing\n- kalkulus\n- sisdig\n- pbd")
		var matkulX string
		fmt.Scan(&matkulX)
		listMahaOfMatkul(tabmatkul, tabnilai, j, matkulX)
	case "4b":
		fmt.Println("Masukan nama mahasiswa yang ingin ditampilkan daftar matkulnya")
		var namaX string
		fmt.Scan(&namaX)
		idx = seqtabmatkul(tabmatkul, j, namaX)
		listMatkulOfMaha(tabmatkul, tabnilai, idx)
	case "5":
		totalSemua(tabnilai, tabmatkul, k)
		insortMhs(tabnilai, tabmatkul, j)
		insortSKS(tabnilai, tabmatkul, k)
		printdata(tabnilai, tabmatkul, j)
	case "6":
		totalSemua(tabnilai, tabmatkul, k)
		selectAscnilai(tabnilai, tabmatkul, j)
		selectAscsks(tabnilai, tabmatkul, k)
		printdata(tabnilai, tabmatkul, j)
	case "7":
		fmt.Println("Masukan total nilai seorang mahasiswa")
		var nilaiX float64
		fmt.Scan(&nilaiX)
		fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("Nama Mahasiswa yang memiliki total nilai tersebut adalah:")
		fmt.Println(binNilaiNama(tabnilai, k, nilaiX))

	case "8":
		outputTranskrip(tabnilai, tabmatkul, k)
	}
}

func main() {
	var tabmatkul matkulMaha
	var tabnilai nilaiMaha
	var pilihan string
	fmt.Println("--Masukan Data Matkul Awal Mahasiswa dengan min 10 data:")
	fmt.Println("(data matkul awal dan untuk testing sudah saya taruh di file .txt pada folder)")
	inputDataMaha(&tabmatkul)
	fmt.Println(" ")
	fmt.Println("--Masukan Data Nilai Awal Mahasiswa dengan min 10 data:")
	fmt.Println("(data nilai awal dan untuk testing sudah saya taruh di file .txt pada folder)")
	inputNilaiMaha(&tabnilai)
	fmt.Println("==========================================================================================================================")
	fmt.Println("Daftar Mata kuliah yang Diambil Oleh Mahasiswa")
	fmt.Println("==========================================================================================================================")
	fmt.Println("Nim              Nama       M1      M2            M3       M4      M5       M6           M7             M8      totalSKs")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	OutputDataMaha(&tabmatkul)
	fmt.Println("===========================================================================================================================")
	fmt.Println("Data Nilai Mahasiswa")
	fmt.Println("===========================================================================================================================")
	fmt.Println("                 Alpro\tMatdis\tMavek\tBindo\tBing\tKalk\tSisdig\tPbd ")
	OutputdataAwalnilai(&tabnilai)
	OutputNilaiMaha(&tabnilai, &tabmatkul, 10)
	mainloop(&tabmatkul, &tabnilai, pilihan)

}

// Pengguna bisa menampilkan daftar mahasiswa terurut berdasarkan nilai dan jumlah sks yang diambil.
func insortMhs(tab *nilaiMaha, Tab *matkulMaha, J int) {
	var pass, i, j int
	var temp coursevalue
	var temp2 course
	//sort berdasarkan totalnilai
	pass = 1
	for pass <= J-1 {
		i = pass
		j = pass
		temp = tab[i]
		temp2 = Tab[j]
		for i > 0 && temp.totalAll > tab[i-1].totalAll {
			tab[i] = tab[i-1]
			Tab[j] = Tab[j-1]
			i--
			j--
		}
		tab[i] = temp
		Tab[j] = temp2
		pass++
	}
}
func insortSKS(tab *nilaiMaha, Tab *matkulMaha, K int) {
	var pass, i, j int
	var temp coursevalue
	var temp2 course
	//sort berdasarkan SKS
	pass = 1
	for pass <= K-1 {
		i = pass
		j = pass
		temp = tab[i]
		temp2 = Tab[j]
		for j > 0 && temp2.sks > Tab[j-1].sks {
			tab[i] = tab[i-1]
			Tab[j] = Tab[j-1]
			i--
			j--
		}
		tab[i] = temp
		Tab[j] = temp2
		pass++
	}
}

// selection sort ascending berdasarkan totalnilai
func selectAscnilai(tabnilai *nilaiMaha, tabmatkul *matkulMaha, J int) {
	var pass, i int
	var idx int
	var temp coursevalue
	var temp2 course
	pass = 1
	for pass <= J-1 {
		idx = pass - 1
		i = pass
		for i < J {
			if tabnilai[idx].totalAll > tabnilai[i].totalAll {
				idx = i

			}
			i++
		}
		temp = tabnilai[pass-1]
		temp2 = tabmatkul[pass-1]
		tabnilai[pass-1] = tabnilai[idx]
		tabmatkul[pass-1] = tabmatkul[idx]
		tabnilai[idx] = temp
		tabmatkul[idx] = temp2
		pass++

	}
}

// selection sort ascending berdasarkan sks
func selectAscsks(tabnilai *nilaiMaha, tabmatkul *matkulMaha, K int) {
	var pass, i int
	var idx int
	var temp coursevalue
	var temp2 course
	pass = 1
	for pass <= K-1 {
		idx = pass - 1
		i = pass
		for i < K {
			if tabmatkul[idx].sks > tabmatkul[i].sks {
				idx = i
			}
			i++
		}
		temp = tabnilai[pass-1]
		temp2 = tabmatkul[pass-1]
		tabnilai[pass-1] = tabnilai[idx]
		tabmatkul[pass-1] = tabmatkul[idx]
		tabnilai[idx] = temp
		tabmatkul[idx] = temp2
		pass++

	}
}
func printdata(tabnilai *nilaiMaha, tabmatkul *matkulMaha, M int) {
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Terurut berdasarkan nilai dan sks")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
	for i := 0; i < M; i++ {
		fmt.Print(tabnilai[i].namaku, "\t")
		fmt.Printf("%.2f", tabnilai[i].totalAll)
		fmt.Print("\t", tabmatkul[i].sks, "\n")
	}
}

// untuk mencari data matkul yang akan diedit dengan cara memasukan namaX yg kemudian akan direturn lokasi
func seqtabmatkul(tabmatkul *matkulMaha, j int, namaX string) int {
	var lokasi int
	lokasi = -1
	for i := 0; i < j; i++ {
		if tabmatkul[i].nama == namaX {
			lokasi = i
		}
	}
	return lokasi
}

// untuk mencari data matkul yang akan diedit dengan cara memasukan namaX yang kemudia akan return lokasi
func seqtabnilai(tabnilai *nilaiMaha, k int, namaX string) int {
	var lokasi int
	lokasi = -1
	for i := 0; i < k; i++ {
		if tabnilai[i].namaku == namaX {
			lokasi = i
		}
	}
	return lokasi
}

// untuk mencari nilai total semua seorang mahasiswa dengan binary berdasarkan sorting ascending
func binNilaiNama(tabnilai *nilaiMaha, k int, nilaiX float64) string {
	var l, r, found int
	var mid int
	found = -1
	l = 0
	r = k - 1
	for l <= r && found == -1 {
		mid = (l + r) / 2
		if nilaiX < tabnilai[mid].totalAll {
			r = mid - 1
		} else if nilaiX > tabnilai[mid].totalAll {
			l = mid + 1
		} else {
			found = mid
		}
	}
	return tabnilai[found].namaku
}

// Pengguna bisa menampilkan transkrip nilai mahasiswa.
func seqtranskrip(tab *nilaiMaha, J int) int {
	var i int
	var nama string
	i = -1
	fmt.Print("Masukkan Nama :")
	fmt.Scan(&nama)
	for i == -1 {
		i = 0
		for nama != tab[i].namaku && i <= J {
			i++
		}
		if i > J {
			i = -1
			fmt.Print("Nama tidak ditemukan, Coba Ulangi")
			fmt.Scan(&nama)
		}
	}
	return i
}

func printtranskrip(tab nilaiMaha, Tab matkulMaha, idx int) {

	fmt.Print("Nama : ", tab[idx].namaku, "\n")
	fmt.Print("NIM : ", Tab[idx].nim, "\n")
	fmt.Print("NO", "\t", "Mata Kuliah", "\t", "Nilai Akhir", "\t", "Grade", "\n")
	fmt.Printf("%s\t%s\t\t\t%.2f\t\t%s\n", "1", "Alpro", tab[idx].alproot, tab[idx].gralpro)
	fmt.Printf("%s\t%s\t\t\t%.2f\t\t%s\n", "2", "Matdis", tab[idx].matdisot, tab[idx].grmatdis)
	fmt.Printf("%s\t%s\t\t\t%.2f\t\t%s\n", "3", "Bindo", tab[idx].bindoot, tab[idx].grbindo)
	fmt.Printf("%s\t%s\t\t\t%.2f\t\t%s\n", "4", "Bing", tab[idx].bingot, tab[idx].grbing)
	fmt.Printf("%s\t%s\t\t%.2f\t\t%s\n", "5", "Kalkulus", tab[idx].kalkulusot, tab[idx].grkalkulus)
	fmt.Printf("%s\t%s\t\t\t%.2f\t\t%s\n", "6", "Sisdig", tab[idx].sisdigot, tab[idx].grsisdig)
	fmt.Printf("%s\t%s\t\t\t%.2f\t\t%s\n", "7", "PBD", tab[idx].pbdot, tab[idx].grpbd)
	if Tab[idx].sks == 23 {
		fmt.Printf("%s\t%s\t\t\t%.2f\t\t%s\n", "8", "Mavek", tab[idx].mavekot, tab[idx].grmavek)
	}
}

// menghitungg grade per matkul
func hitungGrade(tab *nilaiMaha, J int) {
	for i := 0; i < J; i++ {
		if tab[i].alproot >= 80.01 && tab[i].alproot <= 100 {
			tab[i].gralpro = "A"
		} else if tab[i].alproot >= 70.01 && tab[i].alproot <= 80 {
			tab[i].gralpro = "AB"
		} else if tab[i].alproot >= 65.01 && tab[i].alproot <= 70 {
			tab[i].gralpro = "B"
		} else if tab[i].alproot >= 60.01 && tab[i].alproot <= 65 {
			tab[i].gralpro = "BC"
		} else if tab[i].alproot >= 55.01 && tab[i].alproot <= 60 {
			tab[i].gralpro = "C"
		} else if tab[i].alproot >= 50.01 && tab[i].alproot <= 55 {
			tab[i].gralpro = "D"
		} else {
			tab[i].gralpro = "E"
		}
		if tab[i].matdisot >= 80.01 && tab[i].matdisot <= 100 {
			tab[i].grmatdis = "A"
		} else if tab[i].matdisot >= 70.01 && tab[i].matdisot <= 80 {
			tab[i].grmatdis = "AB"
		} else if tab[i].matdisot >= 65.01 && tab[i].matdisot <= 70 {
			tab[i].grmatdis = "B"
		} else if tab[i].matdisot >= 60.01 && tab[i].matdisot <= 65 {
			tab[i].grmatdis = "BC"
		} else if tab[i].matdisot >= 55.01 && tab[i].matdisot <= 60 {
			tab[i].grmatdis = "C"
		} else if tab[i].matdisot >= 50.01 && tab[i].matdisot <= 55 {
			tab[i].grmatdis = "D"
		} else {
			tab[i].grmatdis = "E"
		}
		if tab[i].mavekot >= 80.01 && tab[i].mavekot <= 100 {
			tab[i].grmavek = "A"
		} else if tab[i].mavekot >= 70.01 && tab[i].mavekot <= 80 {
			tab[i].grmavek = "AB"
		} else if tab[i].mavekot >= 65.01 && tab[i].mavekot <= 70 {
			tab[i].grmavek = "B"
		} else if tab[i].mavekot >= 60.01 && tab[i].mavekot <= 65 {
			tab[i].grmavek = "BC"
		} else if tab[i].mavekot >= 55.01 && tab[i].mavekot <= 60 {
			tab[i].grmavek = "C"
		} else if tab[i].mavekot >= 50.01 && tab[i].mavekot <= 55 {
			tab[i].grmavek = "D"
		} else {
			tab[i].grmavek = "E"
		}
		if tab[i].bindoot >= 80.01 && tab[i].bindoot <= 100 {
			tab[i].grbindo = "A"
		} else if tab[i].bindoot >= 70.01 && tab[i].bindoot <= 80 {
			tab[i].grbindo = "AB"
		} else if tab[i].bindoot >= 65.01 && tab[i].bindoot <= 70 {
			tab[i].grbindo = "B"
		} else if tab[i].bindoot >= 60.01 && tab[i].bindoot <= 65 {
			tab[i].grbindo = "BC"
		} else if tab[i].bindoot >= 55.01 && tab[i].bindoot <= 60 {
			tab[i].grbindo = "C"
		} else if tab[i].bindoot >= 50.01 && tab[i].bindoot <= 55 {
			tab[i].grbindo = "D"
		} else {
			tab[i].grbindo = "E"
		}
		if tab[i].bingot >= 80.01 && tab[i].bingot <= 100 {
			tab[i].grbing = "A"
		} else if tab[i].bingot >= 70.01 && tab[i].bingot <= 80 {
			tab[i].grbing = "AB"
		} else if tab[i].bingot >= 65.01 && tab[i].bingot <= 70 {
			tab[i].grbing = "B"
		} else if tab[i].bingot >= 60.01 && tab[i].bingot <= 65 {
			tab[i].grbing = "BC"
		} else if tab[i].bingot >= 55.01 && tab[i].bingot <= 60 {
			tab[i].grbing = "C"
		} else if tab[i].bingot >= 50.01 && tab[i].bingot <= 55 {
			tab[i].grbing = "D"
		} else {
			tab[i].grbing = "E"
		}
		if tab[i].kalkulusot >= 80.01 && tab[i].kalkulusot <= 100 {
			tab[i].grkalkulus = "A"
		} else if tab[i].kalkulusot >= 70.01 && tab[i].kalkulusot <= 80 {
			tab[i].grkalkulus = "AB"
		} else if tab[i].kalkulusot >= 65.01 && tab[i].kalkulusot <= 70 {
			tab[i].grkalkulus = "B"
		} else if tab[i].kalkulusot >= 60.01 && tab[i].kalkulusot <= 65 {
			tab[i].grkalkulus = "BC"
		} else if tab[i].kalkulusot >= 55.01 && tab[i].kalkulusot <= 60 {
			tab[i].grkalkulus = "C"
		} else if tab[i].kalkulusot >= 50.01 && tab[i].kalkulusot <= 55 {
			tab[i].grkalkulus = "D"
		} else {
			tab[i].grkalkulus = "E"
		}
		if tab[i].sisdigot >= 80.01 && tab[i].sisdigot <= 100 {
			tab[i].grsisdig = "A"
		} else if tab[i].sisdigot >= 70.01 && tab[i].sisdigot <= 80 {
			tab[i].grsisdig = "AB"
		} else if tab[i].sisdigot >= 65.01 && tab[i].sisdigot <= 70 {
			tab[i].grsisdig = "B"
		} else if tab[i].sisdigot >= 60.01 && tab[i].sisdigot <= 65 {
			tab[i].grsisdig = "BC"
		} else if tab[i].sisdigot >= 55.01 && tab[i].sisdigot <= 60 {
			tab[i].grsisdig = "C"
		} else if tab[i].sisdigot >= 50.01 && tab[i].sisdigot <= 55 {
			tab[i].grsisdig = "D"
		} else {
			tab[i].grsisdig = "E"
		}
		if tab[i].pbdot >= 80.01 && tab[i].pbdot <= 100 {
			tab[i].grpbd = "A"
		} else if tab[i].pbdot >= 70.01 && tab[i].pbdot <= 80 {
			tab[i].grpbd = "AB"
		} else if tab[i].pbdot >= 65.01 && tab[i].pbdot <= 70 {
			tab[i].grpbd = "B"
		} else if tab[i].pbdot >= 60.01 && tab[i].pbdot <= 65 {
			tab[i].grpbd = "BC"
		} else if tab[i].pbdot >= 55.01 && tab[i].pbdot <= 60 {
			tab[i].grpbd = "C"
		} else if tab[i].pbdot >= 50.01 && tab[i].pbdot <= 55 {
			tab[i].grpbd = "D"
		} else {
			tab[i].grpbd = "E"
		}
	}

}
func gradekeNilai(tab *nilaiMaha, K int) {
	for i := 0; i < K; i++ {
		if tab[i].gralpro == "A" {
			tab[i].alproot = 80.01
		}
		if tab[i].grmatdis == "A" {
			tab[i].matdisot = 80.01
		}
		if tab[i].grmavek == "A" {
			tab[i].mavekot = 80.01
		}
		if tab[i].grbindo == "A" {
			tab[i].bindoot = 80.01
		}
		if tab[i].grbing == "A" {
			tab[i].bingot = 80.01
		}
		if tab[i].grkalkulus == "A" {
			tab[i].kalkulusot = 80.01
		}
		if tab[i].grsisdig == "A" {
			tab[i].sisdigot = 80.01
		}
		if tab[i].grpbd == "A" {
			tab[i].pbdot = 80.01
		}
	}
}
func gradeAll(tab *nilaiMaha, J int) {
	for i := 0; i < J; i++ {
		if tab[i].totalAll == 80.01 {
			tab[i].grAll = "A"
		} else if tab[i].totalAll >= 70.01 && tab[i].totalAll <= 80 {
			tab[i].grAll = "AB"
		} else if tab[i].totalAll >= 65.01 && tab[i].totalAll <= 70 {
			tab[i].grAll = "B"
		} else if tab[i].totalAll >= 60.01 && tab[i].totalAll <= 65 {
			tab[i].grAll = "BC"
		} else if tab[i].totalAll >= 55.01 && tab[i].totalAll <= 60 {
			tab[i].grAll = "C"
		} else if tab[i].totalAll >= 50.01 && tab[i].totalAll <= 55 {
			tab[i].grAll = "D"
		} else {
			tab[i].grAll = "E"
		}
	}
}
func hitungIPK(tab *nilaiMaha, K int) {
	for i := 0; i < K; i++ {
		if tab[i].grAll == "A" {
			tab[i].ipk = 4
		} else if tab[i].grAll == "AB" && tab[i].totalAll >= 80.01 {
			tab[i].ipk = 3.8
		} else if tab[i].grAll == "AB" {
			tab[i].ipk = 3.75
		} else if tab[i].grAll == "B" {
			tab[i].ipk = 3.5
		} else if tab[i].grAll == "BC" {
			tab[i].ipk = 3.3
		} else if tab[i].grAll == "C" {
			tab[i].ipk = 3.25
		} else if tab[i].grAll == "D" {
			tab[i].ipk = 3
		} else if tab[i].grAll == "E" {
			tab[i].ipk = 2.75
		}
	}
}

// menghitung total nilai per matkul
func totalperMatkul(Tab *nilaiMaha, J int) {
	//tipe quiz, uts, uas yg diinputkan masih string
	var i int
	for i = 0; i < J; i++ {
		Tab[i].alproot = (((Tab[i].alproku * 30) / 100) + ((Tab[i].alprots * 35) / 100) + ((Tab[i].alprous * 35) / 100))
		Tab[i].matdisot = (((Tab[i].matdisku * 30) / 100) + ((Tab[i].matdists * 35) / 100) + ((Tab[i].matdisus * 35) / 100))
		Tab[i].mavekot = (((Tab[i].mavekku * 30) / 100) + ((Tab[i].mavekts * 35) / 100) + ((Tab[i].mavekus * 35) / 100))
		Tab[i].bindoot = (((Tab[i].bindoku * 30) / 100) + ((Tab[i].bindots * 35) / 100) + ((Tab[i].bindous * 35) / 100))
		Tab[i].bingot = (((Tab[i].bingku * 30) / 100) + ((Tab[i].bingts * 35) / 100) + ((Tab[i].bingus * 35) / 100))
		Tab[i].kalkulusot = (((Tab[i].kalkulusku * 30) / 100) + ((Tab[i].kalkulusts * 35) / 100) + ((Tab[i].kalkulusus * 35) / 100))
		Tab[i].sisdigot = (((Tab[i].sisdigku * 30) / 100) + ((Tab[i].sisdigts * 35) / 100) + ((Tab[i].sisdigus * 35) / 100))
		Tab[i].pbdot = (((Tab[i].pbdku * 30) / 100) + ((Tab[i].pbdts * 35) / 100) + ((Tab[i].pbdus * 35) / 100))
	}
}
