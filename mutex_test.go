package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
	"sync"
)

// mutex digunakan untuk menangani 1 variable yg disharing, yg diakses oleh beberapa goroutine secara bersamaan, biar aman lakukan locking dan unlocking sebelum mengakses variable tsb
func TestMutex(t *testing.T) {
	counter := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				counter = counter + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Counter:", counter)
}

// RWMutex digunakan untuk menangani 1 variable yg disharing, yg diakses oleh beberapa goroutine secara bersamaan, biar aman lakukan locking dan unlocking sebelum mengakses variable tsb
// bedanya RWMutex ini ada 2 mode, read dan write
// read digunakan ketika hanya membaca data saja, sehingga beberapa goroutine bisa membaca data secara bersamaan
// write digunakan ketika menulis data, sehingga hanya 1 goroutine saja yang bisa menulis data, dan goroutine lain harus menunggu sampai proses penulisan selesai
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance  int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println("Balance:", account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance:", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex        // Mutex untuk mengunci akses ke data user agar thread-safe
	Name    string    // Nama user
	Balance int       // Saldo user
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock() // Mengunci mutex milik user ini
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock() // Membuka kembali mutex user
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount // Mengubah saldo (tidak aman tanpa lock)
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock() // Goroutine mengunci user1 lebih dulu
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second) // Memberi waktu goroutine lain berjalan

	user2.Lock() // Lalu mengunci user2
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

// Membuat dua user dengan mutex masing-masing
func TestDeadlock(t *testing.T) {
	user1 := UserBalance{Name: "Eko", Balance: 1000000}
	user2 := UserBalance{Name: "Budi", Balance: 1000000}

	// pemicu deadlock, karena kedua goroutine mengunci user secara berlawanan
	// Goroutine 1: lock Eko → lock Budi
	// Goroutine 2: lock Budi → lock Eko
	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	// Memberi waktu agar goroutine berjalan (karena Test tidak WaitGroup)
	time.Sleep(10 * time.Second)

	// Print hasil akhir (bisa tidak konsisten jika deadlock terjadi)
	fmt.Println("User 1:", user1.Name, "Balance:", user1.Balance)
	fmt.Println("User 2:", user2.Name, "Balance:", user2.Balance)
}

// Deadlock dapat terjadi pada kode di atas karena dua goroutine melakukan penguncian (lock) pada dua resource yang sama (mutex milik user1 dan user2) dengan urutan yang berbeda, sehingga tercipta kondisi saling menunggu (circular wait). Ketika goroutine pertama berhasil mengunci user1 lalu menunggu user2, sementara goroutine kedua sudah mengunci user2 dan menunggu user1, tidak ada goroutine yang bisa melanjutkan eksekusi karena masing-masing mutex masih terkunci. Kondisi ini diperparah oleh adanya time.Sleep yang memberi kesempatan goroutine lain untuk berjalan, sehingga potensi deadlock semakin besar. Meskipun pada eksekusi tertentu deadlock tidak selalu muncul, secara logika dan desain kode ini tetap tidak aman dan berisiko mengalami deadlock kapan saja.