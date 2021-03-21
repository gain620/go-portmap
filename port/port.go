package port

import (
	"net"
	"strconv"
	"sync"
	"time"
)

type ScanResult struct {
	Port  int
	State string
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: port}
	addr := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, addr, time.Millisecond*100)
	if err != nil {
		result.State = "CLOSED"
		return result
	}

	defer conn.Close()
	result.State = "OPEN"
	return result
}

func AsyncScans(protocol, hostname string, i int, results *[]ScanResult, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	*results = append(*results, ScanPort(protocol, hostname, i))
	mutex.Unlock()
}

func WellKnownScan(protocol, hostname string) []ScanResult {
	var results []ScanResult
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 1; i < 100; i++ {
		wg.Add(1)
		AsyncScans(protocol, hostname, i, &results, &wg, &mutex)
	}

	wg.Wait()
	return results
}

func PerfectScan() {
	// var results []ScanResult
}
