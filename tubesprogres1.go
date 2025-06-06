/*
Nomor Topik dan Judul Topik: Nomor 4 dan Aplikasi Pengelolaan Data Sampah dan Daur Ulang

Deskripsi Problem yang diberikan: Aplikasi ini digunakan untuk mencatat data pengelolaan sampah dan
daur ulang di suatu wilayah. Data utama yang digunakan adalah daftar
jenis sampah, jumlah sampah yang dikumpulkan, dan data proses daur
ulang. Pengguna aplikasi adalah petugas pengelolaan sampah atau
masyarakat yang ingin mencatat kontribusi daur ulang mereka.

Daftar Anggota Tim: KEVIN ALDRIK JONATHAN TADJONGGA, MUHAMMAD SHEVA WARDHANA

Penjelasan :
A.

B.

C.


*/

package main

import "fmt"

const nMax int = 1234

type detailSampah struct {
	namaSampah   string
	jumlahSampah int
}

type jenisSampah struct {
	namaJenis string
	nSampah   int
	info      [nMax]detailSampah
}

func main() {
	var sampah [nMax]jenisSampah
	var menu int
	var nJenis int

	for {
		home()
		fmt.Scan(&menu)
		switch menu {
		case 1:
			readData(&sampah, &nJenis)
		case 2:
			editData(&sampah, nJenis)
		case 3:
			deleteData(&sampah, &nJenis)
		case 4:
			searchMenu(sampah, nJenis)
		case 5:
			writeData(sampah, nJenis)
		case 6:
			addRecycleData()
		case 7:
			showStatistics(sampah, nJenis)
		case 8:
			sortMenu(&sampah, nJenis)
		case 9:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

func home() {
	fmt.Println()
	fmt.Println("=== Aplikasi Pengelolaan Data Sampah dan Daur Ulang ===")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("1. Tambah Data Sampah")
	fmt.Println("2. Ubah Data Sampah")
	fmt.Println("3. Hapus Data Sampah")
	fmt.Println("4. Cari Data Sampah")
	fmt.Println("5. Tampilkan Data Sampah")
	fmt.Println("6. Tambah Data Daur Ulang")
	fmt.Println("7. Lihat Statistik Pengelolaan")
	fmt.Println("8. Urutkan Data Sampah")
	fmt.Println("9. Keluar")
	fmt.Println("--------------------------------------------------------")
	fmt.Print("Pilih menu (1-9): ")
}

func readData(s *[nMax]jenisSampah, nJenis *int) {
	fmt.Print("Banyak jenis sampah yang ingin ditambahkan: ")
	fmt.Scan(nJenis)
	for i := 0; i < *nJenis; i++ {
		fmt.Println("Pilih jenis sampah:")
		fmt.Println("1. Organik")
		fmt.Println("2. Non_Organik")
		fmt.Println("3. B3")
		fmt.Print("Masukkan pilihan (1-3): ")
		var pilihan int
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			s[i].namaJenis = "Organik"
		case 2:
			s[i].namaJenis = "Non_Organik"
		case 3:
			s[i].namaJenis = "B3"
		default:
			fmt.Println("Pilihan tidak valid, diatur ke 'Lainnya'")
			s[i].namaJenis = "Lainnya"
		}

		fmt.Print("Banyak data sampah pada jenis ini: ")
		fmt.Scan(&s[i].nSampah)
		for j := 0; j < s[i].nSampah; j++ {
			fmt.Print("  Masukkan nama sampah: ")
			fmt.Scan(&s[i].info[j].namaSampah)
			fmt.Print("  Masukkan jumlah sampah: ")
			fmt.Scan(&s[i].info[j].jumlahSampah)
		}
		fmt.Println()
	}
}

func writeData(s [nMax]jenisSampah, nJenis int) {
	for i := 0; i < nJenis; i++ {
		fmt.Printf("\n%s\n-----------------------------------------------\n", s[i].namaJenis)
		for j := 0; j < s[i].nSampah; j++ {
			fmt.Printf("   %-10s : %d \n", s[i].info[j].namaSampah, s[i].info[j].jumlahSampah)
		}
	}
}

// Fungsi tambahan akan diisi di bawah ini:
func editData(s *[nMax]jenisSampah, nJenis int) {
	fmt.Println("[Fungsi editData belum diimplementasikan]")
}

func deleteData(s *[nMax]jenisSampah, nJenis *int) {
	fmt.Println("[Fungsi deleteData belum diimplementasikan]")
}

func searchMenu(s [nMax]jenisSampah, nJenis int) {
	var keyword string
	fmt.Print("Masukkan jenis sampah yang ingin dicari: ")
	fmt.Scan(&keyword)

	// Sequential Search
	found := false
	for i := 0; i < nJenis; i++ {
		if s[i].namaJenis == keyword {
			fmt.Printf("\nData ditemukan untuk jenis: %s\n", s[i].namaJenis)
			for j := 0; j < s[i].nSampah; j++ {
				fmt.Printf("   %-10s : %d\n", s[i].info[j].namaSampah, s[i].info[j].jumlahSampah)
			}
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}

func addRecycleData() {
	fmt.Println("[Fungsi addRecycleData belum diimplementasikan]")
}

func showStatistics(s [nMax]jenisSampah, nJenis int) {
	fmt.Println("[Fungsi showStatistics belum diimplementasikan]")
}

func sortMenu(s *[nMax]jenisSampah, nJenis int) {
	var pilihan int
	fmt.Println("Urutkan berdasarkan:")
	fmt.Println("1. Nama Jenis Sampah (Selection Sort)")
	fmt.Println("2. Jumlah Sampah (Insertion Sort)")
	fmt.Print("Masukkan pilihan (1-2): ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		// Selection Sort berdasarkan namaJenis
		for i := 0; i < nJenis-1; i++ {
			minIdx := i
			for j := i + 1; j < nJenis; j++ {
				if s[j].namaJenis < s[minIdx].namaJenis {
					minIdx = j
				}
			}
			s[i], s[minIdx] = s[minIdx], s[i]
		}
		fmt.Println("Data berhasil diurutkan berdasarkan nama jenis.")
	case 2:
		// Insertion Sort berdasarkan jumlah total sampah per jenis
		for i := 1; i < nJenis; i++ {
			key := s[i]
			j := i - 1
			sumKey := totalSampah(key)
			for j >= 0 && totalSampah(s[j]) > sumKey {
				s[j+1] = s[j]
				j--
			}
			s[j+1] = key
		}
		fmt.Println("Data berhasil diurutkan berdasarkan jumlah sampah.")
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func totalSampah(j jenisSampah) int {
	total := 0
	for i := 0; i < j.nSampah; i++ {
		total += j.info[i].jumlahSampah
	}
	return total
}
