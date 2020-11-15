package main

func main() {
	lg := getLogger()

	for {
		for _, i := range strSplit(getOutput("netsh wlan show interface"), "\n") {
			if strIn("State", i) && strIn("disconnected", i) {
				lg.info("disconnected, connection...")
				system("netsh wlan connect ssid=TP-Link_5G name=TP-Link_5G")
			}
		}

		sleep(3)
	}
}
