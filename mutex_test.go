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