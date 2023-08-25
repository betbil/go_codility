package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("5/3", 5/3)              //1
	fmt.Println("5%3", 5%3)              //2
	fmt.Println(Shuffle("hello, world")) // Output: a shuffled version of "hello, world"

}

func Shuffle(s string) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	runes := []rune(s)
	r.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})

	return string(runes)
}

func RandomChar() rune {
	//rand.Seed(time.Now().UnixNano())
	//NewRand(NewSource(seed))
	rr := rand.New(rand.NewSource(time.Now().Unix()))
	chooseUpper := rr.Intn(2) == 0
	if chooseUpper {
		return rune(rr.Intn(26) + 'A')
	} else {
		return rune(rr.Intn(26) + 'a')
	}
}
