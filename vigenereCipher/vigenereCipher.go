package vigenereCipher

import "fmt"

const characters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Cipher struct {
	letterMap map[string]int
}

func (c *Cipher) PrintLetterMap() {
	fmt.Println(c.letterMap, len(c.letterMap))
}

func (c *Cipher) createLetterMap() {
	c.letterMap = make(map[string]int)
	for k, v := range characters {
		letter := string(v)
		c.letterMap[letter] = k
	}
}

func NewCipher() *Cipher {
	c := &Cipher{}
	c.createLetterMap()

	return c
}
