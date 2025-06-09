package main

import (
	"fmt"
)

//maksimal array
const NMAX int = 5

//tipe data Sampah yang berisi nama, banyak dan daur ulang sampah
type Sampah struct{
	namaSampah	    string	
	banyakSampah	int
	daurUlang 	 	string
	index			int
}
type tabSampah [NMAX]Sampah //array sampah

//tipe data Sistem yang berisi banyak data sampah yang di input, nama jenis sampah, dan juga data sampah
type Sistem struct{
	nSampah 		int 
	data 			tabSampah
	namaJenis		string
	no				int
}
type tabSistem [3]Sistem // array sampah yang dimana hanya memiliku 3 jenis


func main(){
	//deklarasi variabel
	var sampah 		tabSistem
	var chooseMenu	int

	//nama jenis
	sampah[0].namaJenis = "Organik"
	sampah[1].namaJenis = "Non-Organik"
	sampah[2].namaJenis = "B3"
	sampah[0].no = 1
	sampah[1].no = 2
	sampah[2].no = 3

	//tampilan menu apk
	for {
		menuDisplay()
		fmt.Scan(&chooseMenu)
		fmt.Println("")
		switch chooseMenu{
		case 1 :
			fiturAddData(&sampah)			
		case 2 :
			fiturEditData(&sampah)
		case 3 :
			fiturRemoveData(&sampah)
		case 4 :
			fiturPrintData(&sampah)
		case 5 :
			fiturSortingData(&sampah)
		case 6 :
			fiturSearchingData(&sampah)
		case 7 :
			fiturStatisticData(&sampah)
		case 8 :
			return
		}
	}
}


//SECTION - Add Data
func fiturAddData(s *tabSistem){
	var a 			int
	for {
		Header("adddata1")
		printData(s)
		Footer("jenis")
		fmt.Scan(&a)

		//menu
		switch a {
		case 1 :
			addData(s, a-1)
		case 2 :
			addData(s, a-1)
		case 3 :
			addData(s, a-1)
		case 4 :
			return
		}
	}
}

func addData(s *tabSistem, a int){
	//deklarasi variabel
	var nama 		string
	var banyak 		int
	var daurUlang	int
	var metode		string


	//input nama dan banyak sampah	
	Header("adddata2")
	Footer("nama")
	fmt.Scan(&nama)
	Header("adddata2")
	Footer("banyak")
	fmt.Scan(&banyak)

	Header("adddata2")
	Footer("recycle")
	fmt.Scan(&daurUlang)

	for daurUlang != 1 && daurUlang != 2 && daurUlang != 3{
		Header("adddata2")
		Footer("recycle")
		fmt.Scan(&daurUlang)
	}

	switch daurUlang{
	case 1 :
		metode = "Reduce"
	case 2 :
		metode = "Reuse"
	case 3 :
		metode = "Recycle"
	}

	//jika  n masih kurang dari NMAX, maka data masih bisa di tambahkan
	// CATATAN : masih harus membuat jika namanya sama, maka jumlahnya di tambahkan
	if s[a].nSampah < NMAX {
		s[a].data[s[a].nSampah] = Sampah{namaSampah: nama, banyakSampah: banyak, index: s[a].nSampah, daurUlang: metode}
		check(s, a)
		s[a].nSampah++
		fmt.Println("Data berhasil ditambahkan.")
	} else {
		Header("adddata2")
		Footer("full")
	}
}

func check(s *tabSistem, a int){
	var key = s[a].data[s[a].nSampah]

	for i := 0; i < 3; i++{
		for j := 0; j < s[i].nSampah; j++ {
			if s[i].data[j].namaSampah == key.namaSampah{
				s[i].data[j].banyakSampah += key.banyakSampah
				s[a].nSampah--
			}
		}
	}
}


//!SECTION


//SECTION - Edit Data

func fiturEditData(s *tabSistem){
	var index 		int
	var nama 		string
	var banyak 		int
	var confirm 	int
	var jenis 		int

	Header("edit2")
	printData(s)
	Footer("jenis")
	fmt.Scan(&jenis)
	Header("edit2")
	Footer("index")
	fmt.Scan(&index)

	if index >= 0 && index < s[jenis-1].nSampah {

		Header("edit1")
		Footer("nama")
		fmt.Scan(&nama)
		Header("edit1")
		Footer("banyak")
		fmt.Scan(&banyak)

		for{
			fmt.Printf("\n\n╔╡%-10s╞════════════════════════════════════════════════════════════════════════╗\n","   Edit Data   ")
			fmt.Printf("║═══════╤══════════════════╤══════════════════════╤═════════════════╤═════════════════════║\n")
			fmt.Printf("║ %-3s | %-16s | %-20s | %-14s | %-12s ║\n"," No. ","   Jenis Sampah", "     Nama Sampah", " Banyak Sampah ", " Metode Daur Ulang ")
			fmt.Printf("║═══════╪══════════════════╪══════════════════════╪═════════════════╪═════════════════════║\n")
			fmt.Printf("║ %-3s | %-16s | %-20s | %-14d  | %-12s        ║\n","     ",s[jenis-1].namaJenis ,s[jenis-1].data[index].namaSampah, s[jenis-1].data[index].banyakSampah, s[jenis-1].data[index].daurUlang)
			fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
			fmt.Printf("║%45s%-44s║\n"," V","")
			fmt.Printf("║═══════╤══════════════════╤══════════════════════╤═════════════════╤═════════════════════║\n")
			fmt.Printf("║ %-3s | %-16s | %-20s | %-14d  | %-12s ║\n","     ",s[jenis-1].namaJenis , nama, banyak, "                   ")
			fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
			fmt.Printf("║%-89s║\n"," [Confirm] / [Not] : (1/0)")
			fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════╝\n")
			fmt.Printf("╚══> ")
			fmt.Scan(&confirm)

			switch confirm{
			case 1 :
				if index >= 0 && index < s[0].nSampah {
					s[jenis-1].data[index].namaSampah = nama 
					s[jenis-1].data[index].banyakSampah = banyak
					Header("edit1")
					Footer("edit1")
				} else {
					Header("edit")
					Footer("not-found")
				}
				return
			case 0 :
				Header("edit2")
				Footer("edit2")
				return
			}
		}
	}else {
		Header("edit")
		Footer("not-found")
	}
}

//!SECTION

//SECTION - Remove Data
func fiturRemoveData(s *tabSistem){
	var index 		int
	var a 			int
	var confirm 	int

	Header("remove")
	Footer("jenis")
	fmt.Scan(&a)
	Header("remove")
	Footer("index")
	fmt.Scan(&index)

	if index >= 0 && index < s[a-1].nSampah {

		for {
			fmt.Printf("\n\n╔╡%-10s╞════════════════════════════════════════════════════════════════════════╗\n","   Remove Data   ")
			fmt.Printf("║═══════╤══════════════════╤══════════════════════╤═════════════════╤═════════════════════║\n")
			fmt.Printf("║ %-3s | %-16s | %-20s | %-14s | %-12s ║\n"," No. ","   Jenis Sampah", "     Nama Sampah", " Banyak Sampah ", " Metode Daur Ulang ")
			fmt.Printf("║═══════╪══════════════════╪══════════════════════╪═════════════════╪═════════════════════║\n")
			fmt.Printf("║  %d.%d  | %-16s | %-20s | %-14d  | %-12s        ║\n",s[a-1].no,s[a-1].data[index].index,s[a-1].namaJenis ,s[a-1].data[index].namaSampah, s[a-1].data[index].banyakSampah, s[a-1].data[index].daurUlang)
			fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
			fmt.Printf("║%-89s║\n"," [Confirm] / [Not] : (1/0)")
			fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════╝\n")
			fmt.Printf("╚══> ")
			fmt.Scan(&confirm)

			switch confirm{
			case 1 :
				for i := index; i < s[a-1].nSampah-1; i++ {
					s[a-1].data[i] = s[a-1].data[i+1]
				}
				s[a-1].nSampah--
				Header("remove")
				Footer("remove1")
				return
			case 0 :
				Header("remove")
				Footer("remove2")
				return
			}
		}
	} else {
		Header("remove")
		Footer("not-found")
	}
}

//!SECTION

//SECTION - Add Recycle Data

func fiturAddRecycleData(s *tabSistem){
	//deklarasi variabel
	var index		int
	var metode 		string
	var jenis 		int

	//input
	fmt.Print("Jenis data : ")
	fmt.Scan(&jenis)
	fmt.Print("Index data sampah: ")
	fmt.Scan(&index)
	fmt.Print("Metode daur ulang: ")
	fmt.Scan(&metode)
	
	//
	if index >= 0 && index < s[jenis-1].nSampah {
		s[jenis-1].data[index].daurUlang = metode
		fmt.Println("Metode daur ulang dicatat.")
	} else {
		fmt.Println("Index tidak valid.")
	}
}

//!SECTION

//SECTION - Statistic Data
//NOTE - DISPLAY 

func fiturStatisticData(s *tabSistem){
	total := 0
	didaurUlang := 0
	for i := 0; i < 3; i++{
		for j := 0; j < s[i].nSampah; j++ {
			total += s[i].data[j].banyakSampah
			if s[i].data[j].daurUlang != "" {
				didaurUlang += s[i].data[j].banyakSampah
			}
		}
	}
	Header("statistic")
	fmt.Println("Total Sampah:", total)
	fmt.Println("Total Didaur Ulang:", didaurUlang)
}

//!SECTION

//SECTION - Searching 

func fiturSearchingData(s *tabSistem){
	var key string

	Header("search1")
	Footer("search1")
	fmt.Scan(&key)
	searchingDisplay(*s, key)
}

func searchingDisplay(s tabSistem, key string){
	var isThere bool = false

	for i := 0; i < 3; i++{
		for j := 0; j < s[i].nSampah; j++ {
			if s[i].data[j].namaSampah == key{
				Header("search2")
				fmt.Printf("║  %d.%d  | %-16s | %-20s | %-15d | %-19s ║\n",s[i].no,s[i].data[j].index,s[i].namaJenis,s[i].data[j].namaSampah,s[i].data[j].banyakSampah,s[i].data[j].daurUlang)
				Footer("search2")
				isThere = true
			}
		}
	}

	if isThere == false {
		Header("search1")
		Footer("not-found")
	}

}

//!SECTION

//SECTION - Sorting Data
//FIXME - perlu di cek

func fiturSortingData(s *tabSistem) {

    for i := 0; i < len(s); i++ {
        for j := 1; j < s[i].nSampah; j++ {
            temp := s[i].data[j]
            k := j - 1
            for k >= 0 && s[i].data[k].namaSampah > temp.namaSampah {
                s[i].data[k+1] = s[i].data[k]
                k--
            }
            s[i].data[k+1] = temp
        }
    }
	fiturPrintData(s)
}


//!SECTION

//SECTION - Print Data

func fiturPrintData(s *tabSistem){
	Header("display")
	printData(s)
	Footer("")
}

func printData(s *tabSistem){
	for i := 0; i < len(s); i++ {
		for j := 0; j < s[i].nSampah; j++ {
			fmt.Printf("║  %d.%d  | %-16s | %-20s | %-14d  |  %-12s       ║\n",s[i].no,j,s[i].namaJenis ,s[i].data[j].namaSampah, s[i].data[j].banyakSampah, s[i].data[j].daurUlang)
		}
	}
}

//!SECTION


//SECTION - Display Menu

func menuDisplay(){
	fmt.Printf("\n\n╔╡%-10s╞══════════════════════════════════╗\n","   MENU")
	fmt.Printf("║══════════════════════════════════════════════║\n")
	fmt.Printf("║%-46s║\n"," 1. Add Data")
	fmt.Printf("║%-46s║\n"," 2. Edit Data")
	fmt.Printf("║%-46s║\n"," 3. Remove Data")
	fmt.Printf("║%-46s║\n"," 4. Display Data")
	fmt.Printf("║%-46s║\n"," 5. Sort Data")
	fmt.Printf("║%-46s║\n"," 6. Search Data")
	fmt.Printf("║%-46s║\n"," 7. Display Statistic Data")
	fmt.Printf("║%-46s║\n"," 8. Exit")
	fmt.Printf("║──────────────────────────────────────────────║\n")
	fmt.Printf("║%-46s║\n","Choose (1/2/3/4/5/6/7/8) :")
	fmt.Printf("║══════════════════════════════════════════════╝\n")
	fmt.Printf("╚══> ")
}
//!SECTION

//SECTION - Header
func Header(title string){
	switch title {
	case "adddata1" : //NOTE - H add data
		fmt.Printf("\n\n╔╡%-10s╞═════════════════════════════════════════════════════════════════════════╗\n","   Add Data   ")
		fmt.Printf("║═══════╤══════════════════╤══════════════════════╤═════════════════╤═════════════════════║\n")
		fmt.Printf("║ %-3s | %-16s | %-20s | %-14s | %-12s ║\n"," No. ","   Jenis Sampah", "     Nama Sampah", " Banyak Sampah ", " Metode Daur Ulang ")
		fmt.Printf("║═══════╪══════════════════╪══════════════════════╪═════════════════╪═════════════════════║\n")

	case "adddata2" :
		fmt.Printf("\n\n╔╡%-10s╞═════════════════════════════════════════════════════════════════════════╗\n","   Add Data   ")

	case "display" :
		fmt.Printf("\n\n╔╡%-10s╞═════════════════════════════════════════════════════════════════════╗\n","   Display Data   ")
		fmt.Printf("║═══════╤══════════════════╤══════════════════════╤═════════════════╤═════════════════════║\n")
		fmt.Printf("║ %-3s | %-16s | %-20s | %-14s | %-12s ║\n"," No. ","   Jenis Sampah", "     Nama Sampah", " Banyak Sampah ", " Metode Daur Ulang ")
		fmt.Printf("║═══════╪══════════════════╪══════════════════════╪═════════════════╪═════════════════════║\n")

	case "edit1" : //NOTE - H edit
		fmt.Printf("\n\n╔╡%-10s╞════════════════════════════════════════════════════════════════════════╗\n","   Edit Data   ")

	case "edit2" :
		fmt.Printf("\n\n╔╡%-10s╞════════════════════════════════════════════════════════════════════════╗\n","   Edit Data   ")
		fmt.Printf("║ %-3s | %-16s | %-20s | %-14s | %-12s ║\n"," No. ","   Jenis Sampah", "     Nama Sampah", " Banyak Sampah ", " Metode Daur Ulang ")
		fmt.Printf("║═══════╪══════════════════╪══════════════════════╪═════════════════╪═════════════════════║\n")
	
	case "search1" : //NOTE - H search
		fmt.Printf("\n\n╔╡%-10s╞══════════════════════════════════════════════════════════════════════╗\n","   Search Data   ")

	case "search2" :
		fmt.Printf("\n\n╔╡%-10s╞══════════════════════════════════════════════════════════════════════╗\n","   Search Data   ")
		fmt.Printf("║ %-3s | %-16s | %-20s | %-14s | %-12s ║\n"," No. ","   Jenis Sampah", "     Nama Sampah", " Banyak Sampah ", " Metode Daur Ulang ")
		fmt.Printf("║═══════╪══════════════════╪══════════════════════╪═════════════════╪═════════════════════║\n")

	case "remove" :
		fmt.Printf("\n\n╔╡%-10s╞══════════════════════════════════════════════════════════════════════╗\n","   Remove Data   ")

	case "statistic" :
		fmt.Printf("\n\n╔╡%-10s╞═══════════════════════════════════════════════════════════════════╗\n","   Statistic Data   ")
	
	}
}

//!SECTION

//SECTION - Footer
func Footer(title string){
	switch title {

	case "" ://NOTE - footer
		fmt.Printf("╚═════════════════════════════════════════════════════════════════════════════════════════╝\n")

	case "jenis" :
		fmt.Printf("║═══════╧══════════════════╧══════════════════════╧═════════════════╧═════════════════════║\n")
		fmt.Printf("║%-89s║\n"," Pilih jenis sampah : (1/2/3/4)")
		fmt.Printf("║%-89s║\n"," 1. Organik")
		fmt.Printf("║%-89s║\n"," 2. Non-Organik")
		fmt.Printf("║%-89s║\n"," 3. B3")
		fmt.Printf("║%-89s║\n"," 4. Exit")
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════╝\n")
		fmt.Printf("╚══> ")
	
	case "nama" :
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
		fmt.Printf("║%-89s║\n"," Masukan nama sampah :")
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════╝\n")
		fmt.Printf("╚══> ")

	case "banyak" :
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
		fmt.Printf("║%-89s║\n"," Masukan banyak sampah :")
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════╝\n")
		fmt.Printf("╚══> ")
	
	case "index" :
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
		fmt.Printf("║%-89s║\n"," Masukan index sampah :")
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════╝\n")
		fmt.Printf("╚══> ")

	case "search1" :
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
		fmt.Printf("║%-89s║\n"," Masukan nama sampah yang ingin di cari :")
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════╝\n")
		fmt.Printf("╚══> ")

	case "search2" :
		fmt.Printf("╚═══════╧══════════════════╧══════════════════════╧═════════════════╧═════════════════════╝\n")


	case "not-found" : //NOTE - not found
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
		fmt.Printf("║%-89s║\n"," Data Not Found")
		fmt.Printf("╚═════════════════════════════════════════════════════════════════════════════════════════╝\n")

	case "remove1" :
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
		fmt.Printf("║%-89s║\n"," Data deleted successfuly")
		fmt.Printf("╚═════════════════════════════════════════════════════════════════════════════════════════╝\n")

	case "remove2" :
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
		fmt.Printf("║%-89s║\n"," Data not deleted")
		fmt.Printf("╚═════════════════════════════════════════════════════════════════════════════════════════╝\n")

	case "edit1" :
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
		fmt.Printf("║%-89s║\n"," Data successfuly edited")
		fmt.Printf("╚═════════════════════════════════════════════════════════════════════════════════════════╝\n")

	case "edit2" :
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
		fmt.Printf("║%-89s║\n"," Data not changed")
		fmt.Printf("╚═════════════════════════════════════════════════════════════════════════════════════════╝\n")
	
	case "full" :
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
		fmt.Printf("║%-89s║\n"," Data full!")
		fmt.Printf("╚═════════════════════════════════════════════════════════════════════════════════════════╝\n")

	case "recycle" :
		fmt.Printf("║═══════╧══════════════════╧══════════════════════╧═════════════════╧═════════════════════║\n")
		fmt.Printf("║%-89s║\n"," Pilih metode daur ulang sampah : (1/2/3/4)")
		fmt.Printf("║%-89s║\n"," 1. Reduce")
		fmt.Printf("║%-89s║\n"," 2. Reuse")
		fmt.Printf("║%-89s║\n"," 3. Recycle")
		fmt.Printf("║%-89s║\n"," 4. Exit")
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════╝\n")
		fmt.Printf("╚══> ")
	}
}
//!SECTION

//!SECTION
