package main

import "fmt"
const nMax int = 1234

type detailSampah struct{
	namaSampah string
	jumlahSampah int
}
type tabDetailSampah [nMax]detailSampah

type jenisSampah struct{
	namaJenis string
	nSampah int
	info tabDetailSampah
	nSampahTetap int
}
type tabJenisSampah [nMax]jenisSampah

func main(){
	var s tabJenisSampah
	s[0].namaJenis = "Organik"
	s[1].namaJenis = "Non-Organik"
	s[2].namaJenis = "B3"

	s[0].nSampahTetap = 0
	s[1].nSampahTetap = 0
	s[2].nSampahTetap = 0
	var a int
	var indexTetap int = 0

	for{
		fmt.Scan(&a)
		switch a {
		case 1 :
			bacaData(&s, a-1, &indexTetap)
			writeData(s)
		case 2 :
			bacaData(&s, a-1, &indexTetap)
			writeData(s)
		case 3 :
			bacaData(&s, a-1, &indexTetap)
			writeData(s)
		case 4 :
			return
		}
	}
}

func bacaData(s *tabJenisSampah, a int, indexTetap *int){
	var temp detailSampah
	var key string

	fmt.Print("Banyak data sampah yang di tambahkan : ")
	fmt.Scan(&s[a].nSampah) //banyak data sampah

	for i:= 0; i < s[a].nSampah; i++{
		//Scan nama dan jumlah sampah
		fmt.Print("Masukan nama sampah : ")
		fmt.Scan(&temp.namaSampah)
		fmt.Print("Masukan jumlah sampah : ")
		fmt.Scan(&temp.jumlahSampah)

		key = temp.namaSampah
		fmt.Println(cek(*s,a,i,key))
		if cek(*s,a,i,key) != -1 {
			s[a].info[cek(*s,a,i,key)].namaSampah = temp.namaSampah
			s[a].info[cek(*s,a,i,key)].jumlahSampah += temp.jumlahSampah
			s[a].nSampah--
			i--
		} else {
			s[a].info[*indexTetap].namaSampah = temp.namaSampah
			s[a].info[*indexTetap].jumlahSampah = temp.jumlahSampah
			*indexTetap++
		}
	}
	s[a].nSampahTetap += s[a].nSampah
}

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

func writeData(s tabJenisSampah){
	for i:=0;i<3;i++{
		fmt.Printf("\n[ %s ]\n-----------------------------------------------\n", s[i].namaJenis)
		for j:=0; j<s[i].nSampahTetap;j++{
			fmt.Printf(" ->  %-10s : %d \n", s[i].info[j].namaSampah, s[i].info[j].jumlahSampah)
		}
	}
}

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
