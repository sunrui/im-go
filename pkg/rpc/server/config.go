/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-27 11:57:09
 */

package server

// Config 配置
type Config struct {
	CertFile string // 证书文件
	KeyFile  string // 证书私钥文件
	Port     int    // 端口
}
