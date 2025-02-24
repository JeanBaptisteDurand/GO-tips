package main

import "fmt"

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type MongoDBNumberStore struct {
	numberList []int
}

func (m *MongoDBNumberStore) GetAll() ([]int, error) {
	return m.numberList, nil
}

func (m *MongoDBNumberStore) Put(n int) error {
	m.numberList = append(m.numberList, n)
	return nil
}

type RedisNumberStore struct {
	numberList []int
}

func (r *RedisNumberStore) GetAll() ([]int, error) {
	return r.numberList, nil
}

func (r *RedisNumberStore) Put(n int) error {
	r.numberList = append(r.numberList, n/2)
	return nil
}

type ApiServer struct {
	numberStore NumberStorer
}

func main() {
	apiServer := ApiServer {
		// numberStore: &MongoDBNumberStore{},
		numberStore: &RedisNumberStore{},
	}

	err := apiServer.numberStore.Put(8)
	if err != nil {
		panic(err)
	}

	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(numbers)
}