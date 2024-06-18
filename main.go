package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Rastgele sayı üreteci için seed ayarla
	rand.Seed(time.Now().UnixNano())

	// 1 ile 20 arasında rastgele bir sayı tut
	secretNumber := rand.Intn(20) + 1

	var guess int
	var attempts int

	fmt.Println("1 ile 20 arasında bir sayı tuttum, tahmin etmeye çalışın!")

	// Kullanıcı doğru tahmin yapana kadar döngüye devam et
	for {
		fmt.Print("Tahmininiz: ")
		fmt.Scanln(&guess)
		attempts++

		if guess < secretNumber {
			fmt.Println("Daha büyük bir sayı girin.")
		} else if guess > secretNumber {
			fmt.Println("Daha küçük bir sayı girin.")
		} else {
			fmt.Printf("Tebrikler! %d tahminde doğru sayıyı buldunuz: %d\n", attempts, secretNumber)
			break
		}
	}
}
