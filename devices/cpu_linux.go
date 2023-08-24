//go:build linux

package devices

import (
	"net"
)

func macAddress() (PhysicalID string, err error) {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		return defaultPhysicalID, err
	}
	PhysicalID = defaultPhysicalID
	for _, inter := range interfaces {
		if inter.Name == "en0" || inter.Name == "eth0" {
			//fmt.Println(inter.Name)
			mac := inter.HardwareAddr //获取本机MAC地址
			//fmt.Println("MAC = ", mac)
			PhysicalID = mac.String()
			break
		}
	}
	return
}

// CpuPhysicalID linux用第一块网卡的mac地址代替
func CpuPhysicalID() (PhysicalID string, err error) {
	return macAddress()
}
