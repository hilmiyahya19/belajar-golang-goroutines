package belajar_golang_goroutines

import (
	"fmt"     // Digunakan untuk menampilkan output ke console
	"runtime" // Digunakan untuk mengakses informasi runtime Go (CPU, thread, goroutine)
	"sync"    // Digunakan untuk sinkronisasi goroutine menggunakan WaitGroup
	"testing" // Digunakan untuk membuat unit test di Go
	"time"    // Digunakan untuk simulasi delay menggunakan Sleep
)

func TestGetGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{} // Membuat WaitGroup untuk menunggu semua goroutine selesai

	for i := 0; i < 100; i++ {
		group.Add(1) // Menambahkan 1 tugas ke WaitGroup
		go func() {
			time.Sleep(3 * time.Second) // Mensimulasikan pekerjaan goroutine selama 3 detik
			group.Done()               // Menandai bahwa goroutine telah selesai
		}()
	}

	totalCPU := runtime.NumCPU() // Mengambil jumlah core CPU yang tersedia
	fmt.Println("Total CPU", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1) // Mengambil jumlah maksimum OS thread yang digunakan Go runtime
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine() // Mengambil jumlah goroutine yang sedang aktif
	fmt.Println("Total Goroutine", totalGoroutine)

	group.Wait() // Menunggu sampai semua goroutine selesai dieksekusi
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{} // Membuat WaitGroup untuk sinkronisasi goroutine

	for i := 0; i < 100; i++ {
		group.Add(1) // Menambahkan jumlah goroutine yang akan dijalankan
		go func() {
			time.Sleep(3 * time.Second) // Mensimulasikan pekerjaan goroutine
			group.Done()               // Mengurangi counter WaitGroup
		}()
	}

	totalCPU := runtime.NumCPU() // Mengambil total core CPU
	fmt.Println("Total CPU", totalCPU)

	runtime.GOMAXPROCS(20) // Mengatur jumlah maksimum OS thread yang boleh digunakan Go runtime menjadi 20
	totalThread := runtime.GOMAXPROCS(-1) // Mengambil kembali nilai GOMAXPROCS yang sedang aktif
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine() // Mengambil jumlah goroutine yang sedang berjalan
	fmt.Println("Total Goroutine", totalGoroutine)
	
	group.Wait() // Menunggu semua goroutine selesai sebelum test berakhir
}

// Kesimpulan:
// Kode ini menunjukkan cara Go mengelola concurrency menggunakan goroutine dan runtime scheduler.
// runtime.NumCPU digunakan untuk mengetahui jumlah core CPU, sedangkan runtime.GOMAXPROCS mengatur
// jumlah maksimum OS thread yang dapat digunakan untuk mengeksekusi goroutine secara paralel.
// Meskipun goroutine dapat dibuat dalam jumlah besar, eksekusi paralelnya tetap dibatasi oleh
// GOMAXPROCS, sehingga Go mampu menjalankan concurrency secara efisien dan terkontrol.