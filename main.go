package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func readJson() []User {
	jsonFile, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer jsonFile.Close()

	data, _ := io.ReadAll(jsonFile)

	var users []User

	json.Unmarshal(data, &users)

	return users
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

}

func getPrimeList(users []User) []int {
	var primeIDs []int

	for _, user := range users {
		if isPrime(user.UserId) {
			primeIDs = append(primeIDs, user.UserId)
		}
	}

	return primeIDs
}

func main() {

	avl := NewAVL()
	rbt := NewRedBlackTree()

	users := readJson()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(users), func(i, j int) { users[i], users[j] = users[j], users[i] })

	for _, u := range users {
		avl.root = avl.Insert(avl.root, &u)
		rbt.Insert(&u)
	}

	fmt.Println("Árvore AVL após a inserção:")
	avl.PrintTree()
	fmt.Println("Árvore preta e vermelha após a inserção:")
	rbt.PrintTree()

	primes := getPrimeList(users)
	for _, prime := range primes {
		avl.Remove(avl.root, prime)
		rbt.Remove(prime)
	}

	fmt.Println("Árvore AVL após a remoção de nodos dos quais o userId é um número primo:")
	avl.PrintTree()
	fmt.Println("Árvore preta e vermelha após a remoção de nodos dos quais o userId é um número primo:")
	rbt.PrintTree()
}
