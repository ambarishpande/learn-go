package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Creating a new 'deck' Type
// This is like saying extend all functionality of []string
type deck []string

func newDeck() deck {
	cardSuites := []string{"Hearts", "Diamonds", "Spades", "Cloves"}
	cardTypes := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cards := deck{}
	for _, suit := range cardSuites {
		for _, value := range cardTypes {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Attaching a Receiver function to the deck type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Method creates a hand of handSize cards and also returns remaing cards.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Converts a deck to a comma seperated string.
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) (deck, error) {
	data, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
	return deck(strings.Split(string(data), ",")), error
}

func (d deck) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := range d {
		random := r.Intn(len(d) - 1)
		d[i], d[random] = d[random], d[i]
	}
}
