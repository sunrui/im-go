/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-26 15:01:55
 */

package client

// Config 配置
type Config struct {
	CertFile    string // 证书文件
	ServerName  string // 服务器名称
	Ip          string // ip
	Port        int    // 端口
	AccessToken string // 访问令牌
}
