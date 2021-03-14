package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

func main() {
	welcome()

	for {
		showMenu()
		menu := readMenu()

		switch menu {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing logs...")
			showLogs()
		case 3:
			fmt.Println("Exit")
			os.Exit(0)
		default:
			fmt.Println("Don't recognize this command")
			os.Exit(-1)
		}
	}
}

func welcome() {
	name := "Kirk"
	version := 1.1

	fmt.Println("Hello Mr/Ms.", name)
	fmt.Println("Program Version:", version)
}

func showMenu() {
	fmt.Println("Menu:")
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("3 - Quit Program")
}

func readMenu() int {
	var readMenu int
	fmt.Scan(&readMenu)
	fmt.Println("Chosen 'menu':", readMenu)
	fmt.Println("")

	return readMenu
}

func startMonitoring() {
	fmt.Println("Start monitoring...")

	sites := readSitesFromFile()

	for i := 0; i < monitoring; i++ {
		for _, site := range sites {
			fmt.Println("Monitoring:", site)
			siteTest(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func siteTest(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was loaded successfully")
		writeLogs(site, true)
	} else {
		fmt.Println("Site:", site, "not found.")
		writeLogs(site, false)
	}
}

func readSitesFromFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func writeLogs(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("An error has occurred:", err)
	}
	fmt.Println(string(file))
}
