package model

import (
	"time"

	"github.com/gorilla/websocket"
)

type UserConnectionInfo struct {
	Conn     *websocket.Conn `json:"-"`
	ID       string          `json:"id"`
	EmpID    string          `json:"empId"`
	RoleID   string          `json:"roleId"`
	OrgID    string          `json:"orgID"`
	StnID    string          `json:"stnID"`
	DeptID   string          `json:"deptId"`
	CommID   string          `json:"commId"`
	Username string          `json:"username"`
}
type RegistrationMessage struct {
	OrgID    string `json:"orgId"`    // องค์กร
	Username string `json:"username"` // ชื่อผู้ใช้
	EmpID    string `json:"empId"`    // รหัสพนักงาน
	RoleID   string `json:"roleId"`   // ตำแหน่ง
	DeptID   string `json:"deptId"`   // แผนก
	CommID   string `json:"commId"`   // หน่วยงานย่อย/สายงาน
	StnID    string `json:"stnId"`    // สถานี/สาขา
}

// Notification คือข้อมูลการแจ้งเตือนหลัก
type Notification struct {
	ID          int         `json:"id"`
	OrgID       string      `json:"orgId"`
	SenderType  string      `json:"senderType"`  // "SYSTEM" or "USER"
	SenderPhoto string      `json:"senderPhoto"` // เพิ่มใหม่
	Sender      string      `json:"sender"`
	Message     string      `json:"message"`
	EventType   string      `json:"eventType"`
	RedirectUrl string      `json:"redirectUrl"`
	CreatedAt   time.Time   `json:"createdAt"`
	CreatedBy   string      `json:"createdBy"` // เพิ่มใหม่
	ExpiredAt   time.Time   `json:"expiredAt"` // เพิ่มใหม่
	Data        []Data      `json:"data"`
	Recipients  []Recipient `json:"recipients"` // ใช้ตอนสร้างเท่านั้น
}

// Recipient คือเป้าหมายผู้รับ
type Recipient struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
