/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-25 23:04:48
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

type IPPort struct {
	IP   string
	Port int
}

type UserIPPorts struct {
	sync.RWMutex
	mapping map[int][]IPPort
}

var userIPPorts = &UserIPPorts{
	mapping: make(map[int][]IPPort),
}

func main() {
	// 示例操作，需在各方法内部调用对应的锁定和解锁函数
	go addUserIPPort(1, "192.168.0.1", 8080)
	go addUserIPPort(1, "192.168.0.2", 9000)
	go printUserIPPorts(1)

	// 阻塞主线程，等待所有协程完成操作
	time.Sleep(time.Hour)
}

// 添加用户的 IP 端口
func addUserIPPort(userID int, ip string, port int) {
	userIPPorts.Lock()
	defer userIPPorts.Unlock()

	ipPort := IPPort{ip, port}
	userIPPorts.mapping[userID] = append(userIPPorts.mapping[userID], ipPort)
}

// 获取用户的全部 IP 端口
func getUserIPPorts(userID int) []IPPort {
	userIPPorts.RLock()
	defer userIPPorts.RUnlock()

	return userIPPorts.mapping[userID]
}

// 删除用户的某个 IP 端口
func removeUserIPPort(userID int, ip string, port int) {
	userIPPorts.Lock()
	defer userIPPorts.Unlock()

	found := false
	var updatedIPPorts []IPPort
	for _, p := range userIPPorts.mapping[userID] {
		if p.IP == ip && p.Port == port {
			found = true
			continue
		}
		updatedIPPorts = append(updatedIPPorts, p)
	}
	if found {
		userIPPorts.mapping[userID] = updatedIPPorts
	}
}

// 更新用户的某个 IP 端口（这里假设只更新第一个匹配的）
func updateUserIPPort(userID int, oldIP string, oldPort int, newIP string, newPort int) {
	userIPPorts.Lock()
	defer userIPPorts.Unlock()

	for i, p := range userIPPorts.mapping[userID] {
		if p.IP == oldIP && p.Port == oldPort {
			userIPPorts.mapping[userID][i] = IPPort{newIP, newPort}
			return
		}
	}
}

// 打印用户的全部 IP 端口
func printUserIPPorts(userID int) {
	ipPorts := getUserIPPorts(userID)
	fmt.Printf("User %d has IP ports: %+v\n", userID, ipPorts)
}
