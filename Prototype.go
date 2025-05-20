package main

import "fmt"
const nMax int = 1234 //nMax

//Array detail sampah yang berisi nama dan jumlah sampah
type detailSampah struct{
	namaSampah string
	jumlahSampah int
}
type tabDetailSampah [nMax]detailSampah

//array jenis sampah yang berisi nama jenis sampah dan array detail sampah
type jenisSampah struct{
	namaJenis string
	nSampah int
	info tabDetailSampah
	nSampahTetap int
}
type tabJenisSampah [nMax]jenisSampah

func main(){
	//deklarasi variabel
	var s tabJenisSampah 
	var a int
	var indexTetap int = 0

	//mengisi data nama jenis
	s[0].namaJenis = "Organik"
	s[1].namaJenis = "Non-Organik"
	s[2].namaJenis = "B3"

	//mengisi nsampah tetap (masih bug)
	s[0].nSampahTetap = 0
	s[1].nSampahTetap = 0
	s[2].nSampahTetap = 0

	//pengulangan switch case
	for{
		//memilih jenis sampah yang ingin di input
		fmt.Println("1. Organik")
		fmt.Println("2. Non-Organik")
		fmt.Println("3. B3")
		fmt.Println("4. Exit")
		fmt.Print("Masukan Jenis Sampah (1/2/3/4) : ")
		fmt.Scan(&a)
		
		switch a {
		case 1 : //organik
			bacaData(&s, a-1, &indexTetap)
			writeData(s)
		case 2 : //non
			bacaData(&s, a-1, &indexTetap)
			writeData(s)
		case 3 : //B3
			bacaData(&s, a-1, &indexTetap)
			writeData(s)
		case 4 : //keluar
			return
		}
	}
}

//function membaca data (masih prototype dan banyak bug) || pada function ini, berhasil membaca data, namun bila membaca kembali, data akan bertambah dan tidak masuk di cetakan
func bacaData(s *tabJenisSampah, a int, indexTetap *int){
	var temp detailSampah //var sementara yang menampung nilai
	var key string // key untuk search

	//memasukan banyak data sampah yang ingin di tambahkan
	fmt.Print("Banyak data sampah yang di tambahkan : ")
	fmt.Scan(&s[a].nSampah) //banyak data sampah

	for i:= 0; i < s[a].nSampah; i++{
		//Scan nama dan jumlah sampah
		fmt.Print("Masukan nama sampah : ")
		fmt.Scan(&temp.namaSampah)
		fmt.Print("Masukan jumlah sampah : ")
		fmt.Scan(&temp.jumlahSampah)

		key = temp.namaSampah //inisialisasi key
		fmt.Println(cek(*s,a,i,key)) //hanya mengecek index (nanti di hilangakan)

		//melakukan pengecekan apakah data nama yang di masukan sekarang sudah ada di dalam array utama atau tidak
		if cek(*s,a,i,key) != -1 { // jika ada, maka data nilai jumlah akan bertambah
			s[a].info[cek(*s,a,i,key)].namaSampah = temp.namaSampah
			s[a].info[cek(*s,a,i,key)].jumlahSampah += temp.jumlahSampah
			s[a].nSampah--
			i--
		} else { //jika tidak maka akan menjadi data baru
			s[a].info[*indexTetap].namaSampah = temp.namaSampah
			s[a].info[*indexTetap].jumlahSampah = temp.jumlahSampah
			*indexTetap++
		}
	}
	s[a].nSampahTetap += s[a].nSampah //sepertinya ada masalah di sini
}

//belum selesai
/*
func readData(s *tabJenisSampah, a int){
	fmt.Print("Banyak data sampah yang di tambahkan : ")
	fmt.Scan(&s[a].nSampah)

for j:= 0; j < s[a].nSampah; j++{

		fmt.Print("Masukan nama sampah : ")
		fmt.Scan(&s[a].info[j].namaSampah)
		fmt.Print("Masukan jumlah sampah : ")
		fmt.Scan(&s[a].info[j].jumlahSampah)
		key := s[a].info[j].namaSampah
		fmt.Println(cek(*s,a,j,key))
		if cek(*s,a,j,key) != -1 {
			s[a].info[cek(*s,a,j,key)].jumlahSampah += s[a].info[j].jumlahSampah
			j--
			s[a].nSampah--
		}
	}
}*/

//mencetak data
func writeData(s tabJenisSampah){
	for i:=0;i<3;i++{
		fmt.Printf("\n[ %s ]\n-----------------------------------------------\n", s[i].namaJenis)
		for j:=0; j<s[i].nSampahTetap;j++{
			fmt.Printf(" ->  %-10s : %d \n", s[i].info[j].namaSampah, s[i].info[j].jumlahSampah)
		}
	}
}

//searching
func cek(s tabJenisSampah, a int, i int, key string)int{
	//Pencarian Id menggunakan binary search
	var left, mid, right int
	left = 0
	right = i
	for left <= right {
		mid = (left + right) / 2
		if s[a].info[mid].namaSampah == key {
			return mid // Ditemukan
		} else if  s[a].info[mid].namaSampah < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1 // Tidak ditemukan
}

/*func search(){
	for i := 0; i < n; i++ {
		minIdx := i
		for k := i + 1; k < n; j++ {
			if s[a].info[j].namaSampah < s[a].info[minIdx].namaSampah {
				minIdx = k
			}
		}
		// Tukar elemen
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
}
}*/
