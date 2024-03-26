/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-25 23:20:52
 */
package main

import (
	"fmt"
	"sync"
)

type Endpoint struct {
	Ip   string
	Port int
}

type Session interface {
	Add(userId string, endpoint Endpoint) bool
	Remove(userId string, endpoint Endpoint) bool
	Get(userId string) []Endpoint
	Print(userId string)
	PrintAll()
}

type MemorySession struct {
	store map[string][]Endpoint
	mutex sync.RWMutex
}

func NewMemorySession() *MemorySession {
	return &MemorySession{
		store: make(map[string][]Endpoint),
		mutex: sync.RWMutex{},
	}
}

func (ms *MemorySession) Add(userId string, endpoint Endpoint) bool {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	userEndpoints, exists := ms.store[userId]
	if exists {
		for _, e := range userEndpoints {
			if e.Ip == endpoint.Ip && e.Port == endpoint.Port {
				return false
			}
		}
	}
	ms.store[userId] = append(ms.store[userId], endpoint)
	return true
}

func (ms *MemorySession) Remove(userId string, endpoint Endpoint) bool {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	userEndpoints, exists := ms.store[userId]
	if !exists {
		return false
	}
	for i, e := range userEndpoints {
		if e.Ip == endpoint.Ip && e.Port == endpoint.Port {
			ms.store[userId] = append(userEndpoints[:i], userEndpoints[i+1:]...)
			return true
		}
	}
	return false
}

func (ms *MemorySession) Get(userId string) []Endpoint {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()

	endpoints, exists := ms.store[userId]
	if exists {
		return endpoints
	}
	return []Endpoint{}
}

func (ms *MemorySession) Print(userId string) {
	endpoints := ms.Get(userId)
	fmt.Printf("UserId: %s\nEndpoints:\n", userId)
	for _, ep := range endpoints {
		fmt.Printf("\tIp: %s\tPort: %d\n", ep.Ip, ep.Port)
	}
}

func (ms *MemorySession) PrintAll() {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()

	for userId, endpoints := range ms.store {
		fmt.Printf("UserId: %s\nEndpoints:\n", userId)
		for _, ep := range endpoints {
			fmt.Printf("\tIp: %s\tPort: %d\n", ep.Ip, ep.Port)
		}
		fmt.Println()
	}
}

func main() {
	session := NewMemorySession()
	ep1 := Endpoint{"192.168.0.1", 8080}
	ep2 := Endpoint{"192.168.0.2", 9000}

	// 添加操作
	success := session.Add("user1", ep1)
	if success {
		fmt.Println("Endpoint added successfully.")
	} else {
		fmt.Println("Endpoint already exists.")
	}

	// 打印操作
	session.Print("user1")

	// 移除操作
	success = session.Remove("user1", ep1)
	if success {
		fmt.Println("Endpoint removed successfully.")
	} else {
		fmt.Println("Endpoint doesn't exist.")
	}

	success = session.Add("user2", ep2)

	// 再次打印操作
	session.Print("user1")

	session.PrintAll()
}
