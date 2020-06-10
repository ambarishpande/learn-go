package main

import "fmt"

func main() {
	s := "ambarish"
	enc := encrypt(s)
	fmt.Println(enc)
	dec := decrypt(enc)
	fmt.Println(dec)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%d %#U, %#x\n", s[i], s[i], s[i])
	}

}

func encrypt(s string) string {
	bs := []byte(s)
	for i := 0; i < len(bs); i++ {
		bs[i] = bs[i] + 2
	}
	return string(bs)
}

func decrypt(s string) string {
	bs := []byte(s)
	for i := 0; i < len(bs); i++ {
		bs[i] = bs[i] - 2
	}
	return string(bs)
}
