package iputil

import "net"

func GetAllIPsInSubnet(subnet, mask string) []net.IP {
	ip := net.ParseIP(subnet)
	ipNet := net.IPNet{
		IP:   ip,
		Mask: net.IPMask(net.ParseIP(mask).To4()),
	}
	ipRange := make([]net.IP, 0)
	for ip = ipNet.IP.Mask(ipNet.Mask); ipNet.Contains(ip); incIP(ip) {
		ipRange = append(ipRange, net.IP{ip[0], ip[1], ip[2], ip[3]})
	}
	return ipRange
}

// 递增IP地址
func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
