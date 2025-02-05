package main
import (
 "fmt"
 "net"
)
func main() {
	var num_of_ports int
    fmt.Println("Enter the number of ports to scan (min 1, max 65535): ")
    fmt.Scanln(&num_of_ports)

    var ip_address string
    fmt.Println("Enter the IP address to scan: ")
    fmt.Scanln(&ip_address)
	for i := 1; i <= num_of_ports; i++ {
		address := fmt.Sprintf("%s:%d", ip_address, i)
		conn, err := net.Dial("tcp", address)
		fmt.Println(i)
		if err != nil {
			continue
		}
 		conn.Close()
 		fmt.Printf("%d open\n", i)
 		}
}