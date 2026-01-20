package belajar_golang_goroutines // Package untuk mempelajari konsep goroutine dan concurrency di Go

import (
	"fmt"          // Digunakan untuk menampilkan output ke console
	"sync"         // Menyediakan primitive concurrency seperti WaitGroup
	"sync/atomic"  // Menyediakan operasi atomic untuk mencegah race condition
	"testing"      // Digunakan agar fungsi dapat dijalankan sebagai unit test
)

func TestAtomic(t *testing.T) { // Fungsi unit test untuk menguji penggunaan atomic pada goroutine
	var x int64 = 0            // Variabel counter bertipe int64 (wajib untuk operasi atomic)
	group := sync.WaitGroup{} // WaitGroup untuk menunggu seluruh goroutine selesai

	for i := 1; i <= 1000; i++ { // Loop untuk membuat 1000 goroutine
		group.Add(1) // Menambah jumlah goroutine yang harus ditunggu (WAJIB sebelum go func)
		go func() {  // Menjalankan proses secara concurrent menggunakan goroutine
			defer group.Done() // Menandai goroutine selesai setelah fungsi ini berakhir
			for j := 1; j <= 100; j++ { // Loop increment counter sebanyak 100 kali
				atomic.AddInt64(&x, 1) // Menambah nilai x secara aman (thread-safe)
			}
		}()
	}

	group.Wait() // Menunggu hingga semua goroutine memanggil Done()
	fmt.Println("Final Counter:", x) // Menampilkan hasil akhir counter setelah semua goroutine selesai

	// KESIMPULAN:
	// Kode ini mendemonstrasikan penggunaan goroutine, sync.WaitGroup, dan atomic operation
	// untuk melakukan increment counter secara concurrent tanpa race condition.
	// WaitGroup memastikan seluruh goroutine selesai sebelum program berakhir,
	// sementara atomic.AddInt64 menjamin operasi penambahan nilai berjalan aman
	// meskipun diakses oleh banyak goroutine secara bersamaan.
}
