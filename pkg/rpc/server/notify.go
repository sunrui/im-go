/*
 * Copyright © 2024 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2024-03-27 11:51:45
 */

package server

// Notifier 通知
type Notifier interface {
	OnError(err error) // 错误
	onClose()          // 关闭
}
