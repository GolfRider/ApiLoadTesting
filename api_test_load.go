package main

import (
	"fmt"
	"net/http"
	"os"
	"os/user"
	"time"

	"github.com/tsenart/vegeta/lib"
)

const (
	RATE          = 100               // req per second
	TEST_DURATION = 120 * time.Second // test duration


	API_HOST = "https://<API_SERVER_URL>"
	API_KEY  = "API_KEY_123456789"
)


// rate : req/sec
// duration : in sec
func loadTest(rate int, duration int) {

	// write to file
	report_file_name := fmt.Sprintf("api_test_report_%d%s", rate, ".txt")
	usr, err := user.Current()
	checkError(err)
	userHome := usr.HomeDir
	reportFile, err := os.Create(userHome + "/" + report_file_name)
	checkError(err)
	defer reportFile.Close()

	// set common headers
	header := make(http.Header)
	header.Set("Content-Type", "application/json")
	header.Set("Authorization", "APIKEY "+API_KEY)

	// API Targets
	getBooks := vegeta.Target{
		Method: "GET",
		URL:    API_HOST + "/books",
		Header: header,
	}

	getBookById := vegeta.Target{
		Method: "GET",
		URL:    API_HOST + "/books/isbn_12345",
		Header: header,
	}

	targeter := vegeta.NewStaticTargeter(getBookById, getBooks)  // list of api targets

	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics

	for res := range attacker.Attack(targeter, uint64(rate), time.Duration(duration)*time.Second) {
		metrics.Add(res)
	}
	metrics.Close()
	reporter := vegeta.NewTextReporter(&metrics)

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
	reporter.Report(reportFile)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	duration := 120         // in seconds
	loadTest(500, duration) // 10k rps
}
