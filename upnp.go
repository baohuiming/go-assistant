package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/huin/goupnp"
	"github.com/huin/goupnp/dcps/av1"
)

func Play() {
	go serve()
	pushToDlna()
	// 1分钟后关闭程序
	time.Sleep(1 * time.Minute)
	os.Exit(0)
}

func pushToDlna() {
	configP := ReadConfig()
	deviceName := configP.DLNA.DeviceName

	// Discover UPnP devices on the local network
	devices, err := goupnp.DiscoverDevices(av1.URN_AVTransport_1)
	if err != nil {
		log.Fatal("Failed to discover devices:", err)
	}

	var clientP *av1.AVTransport1

	// Connect to the device which is named DLNA-Feixun-R1
	for _, device := range devices {
		if device.Root.Device.FriendlyName == deviceName {
			log.Println("Found device:", device.Root.Device.FriendlyName)
			// Connect
			d, err := av1.NewAVTransport1ClientsByURL(device.Location)
			if err != nil {
				log.Fatal("Failed to create client:", err)
			}
			clientP = d[0]
			break
		}
	}
	if clientP == nil {
		log.Fatal("Failed to find device")
	}
	log.Println("Connected to device:", deviceName)

	url := "http://" + getLocalIp() + ":8080/output.mp3"
	err = clientP.SetAVTransportURI(0, url, "")
	if err != nil {
		log.Fatal("Failed to set URI:", err)
	}
	err = clientP.Play(0, "1")
	if err != nil {
		log.Fatal("Failed to play:", err)
	}
}

// 启动一个简单的HTTP服务器提供MP3文件
func serve() {
	http.HandleFunc("/output.mp3", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "output.mp3")
	})

	log.Println("Server started on http://" + getLocalIp() + ":8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func getLocalIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		log.Panicln("Error getting addresses:", err)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			ip := ipnet.IP.To4()
			if ip != nil && ipnet.IP.IsGlobalUnicast() {
				return ip.String()
			}
		}
	}
	return ""
}
