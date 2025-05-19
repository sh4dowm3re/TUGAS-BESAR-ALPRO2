package main

import "fmt"
const nMax int = 1234

type detailSampah struct{
	namaSampah string
	jumlahSampah int
}

type jenisSampah struct{
	namaJenis string
	nSampah int
	info [nMax]detailSampah
}

func main(){
	var sampah [nMax]jenisSampah
	var menu int
	var nJenis int

	for {
		home()
		fmt.Scan(&menu)
		switch menu {
		case 1 :
			readData(&sampah, &nJenis)
//		case 2 :
//			rereadData(&sampah, nSampah, nSampah)
		case 5 :
			writeData(sampah, nJenis)
		case 6 :
			return
		}
	}
}

func home(){
	fmt.Println("")
	fmt.Println("Alikasi Pengelolaan Data Sampah dan Daur Ulang")
	fmt.Println("-----------------------------------------------")
	fmt.Println("1. Tambah Data Sampah")
	fmt.Println("2. Ubah Data Sampah")
	fmt.Println("3. Hapus Data Sampah")
	fmt.Println("4. Cari Data Sampah")
	fmt.Println("5. Tampilkan Data Sampah")
	fmt.Println("6. Exit")
	fmt.Println("-----------------------------------------------")
	fmt.Println("")
}

func readData(s *[nMax]jenisSampah, nJenis *int){

	if

	fmt.Print("Banyak jenis sampah yang di tambahkan : ")
	fmt.Scan(nJenis)

	for i:= 0; i < *nJenis; i++{
		fmt.Print("Masukan nama jenis sampah : ")
		fmt.Scan(&s[i].namaJenis)
		fmt.Print("Banyak data sampah yang di tambahkan : ")
		fmt.Scan(&s[i].nSampah)
		for j:= 0; j < s[i].nSampah; j++{
			fmt.Print("Masukan nama sampah : ")
			fmt.Scan(&s[i].info[j].namaSampah)
			fmt.Print("Masukan jumlah sampah : ")
			fmt.Scan(&s[i].info[j].jumlahSampah)
		}
		fmt.Println("")
	}
}

func writeData(s [nMax]jenisSampah, nJenis int){
	for i:=0;i<nJenis;i++{
		fmt.Printf("\n%s\n-----------------------------------------------\n", s[i].namaJenis)
		for j:=0; j<s[i].nSampah;j++{
			fmt.Printf("   %-10s : %d \n", s[i].info[j].namaSampah, s[i].info[j].jumlahSampah)
		}
	}
}

/*func rereadData(s *[nMax]sampah){
	var a int
	for {
		fmt.Print("Data yang ingin di ubah : ")
		fmt.Scan(&a)
		if (s[a].nama == "" && s[a].jumlah == 0){
			fmt.Println("Data kosong")
		} else {
			fmt.Scan(&s[a].nama)
			fmt.Scan(&s[a].jumlah)
			return
		}
	}	
} */