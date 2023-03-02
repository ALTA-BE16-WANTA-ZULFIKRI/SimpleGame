package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	random := rand.NewSource(time.Now().UnixMilli())
	rnd := rand.New(random)
	randomint := rnd.Intn(100)
	fmt.Println(randomint)

	var tanya int
	var input int
	var isRunning = true
	for isRunning {

		fmt.Println("selamat datang dipermainan tebak angka 1 sampai 100 ?")
		fmt.Println("Tekan 1 jika anda ingin bermain")
		fmt.Println("Tekan 999 jika ingin exit")
		fmt.Scanln(&tanya)

		if tanya == 1 {

			for isRunning {
				fmt.Println("Masukan angka tebakan anda")
				fmt.Scanln(&input)
				if input == randomint {
					fmt.Println("selamat tebakan anda benar")
					isRunning = false
				} else if input == 999 {
					isRunning = false
					fmt.Println("Terima kasih")
				} else if input != randomint {
					fmt.Println("silahkan coba lagi")
				}
			}

		} else if tanya == 2 {
			fmt.Println("Terima kasih ")

		} else {

			fmt.Println("silahkan coba lagi")
		}

	}
}
