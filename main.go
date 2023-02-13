package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/Saffrontea/har-analyze/types"
)

var (
	har           types.Har
	requests      []*types.Request
	responses     []*types.Response
	statusCode    map[string]int
	statusEntry   map[int]int
	zeroByteEntry []int
	tooLongEntry  []int
	mimeCount     map[string]int
)

func main() {
	harFile, err := os.Open("har/ja.wikipedia.org.har")
	if err != nil {
		log.Fatal("Can't open file.")
		return
	}
	defer func(harFile *os.File) {
		err := harFile.Close()
		if err != nil {
			log.Panic("Oops! Can't close file.")
		}
	}(harFile)
	f, err := io.ReadAll(harFile)
	if err != nil {
		log.Fatal("Can't read file.")
		return
	}

	err = json.Unmarshal(f, &har)
	if err != nil {
		log.Fatal("Can't parse HAR file...")
		return
	}
	//fmt.Print(har)
	fmt.Println("Inspecting HAR file...")
	requests = make([]*types.Request, len(har.Log.Entries))
	responses = make([]*types.Response, len(har.Log.Entries))
	for i := 0; i < len(har.Log.Entries); i++ {
		requests[i] = &har.Log.Entries[i].Request
		responses[i] = &har.Log.Entries[i].Response
	}
	inspectHARFile()
	fmt.Println("Creator: " + har.Log.Creator.Name + ", Version: " + har.Log.Creator.Version)
	for i := 0; i < len(har.Log.Entries); i++ {
		fmt.Print("Entry:" + strconv.Itoa(i) + " ")
		fmt.Print("Method: " + har.Log.Entries[i].Request.Method + " ")
		fmt.Print("URL: " + har.Log.Entries[i].Request.URL + "\n")
		fmt.Print("MIME: " + har.Log.Entries[i].Response.Content.MimeType + " ")
		fmt.Print("Status: " + strconv.Itoa(har.Log.Entries[i].Response.Status) + " ")
		fmt.Print("Byte-Length: " + strconv.Itoa(har.Log.Entries[i].Response.TransferSize) + " ")
		fmt.Print("Time: " + strconv.FormatFloat(har.Log.Entries[i].Time, 'f', -1, 32) + "ms\n")
		fmt.Println()
	}
	fmt.Println("------------Stat-----------")
	fmt.Println("1. Status Code")
	for code, cnt := range statusCode {
		fmt.Printf("Code: %s    %d times\n", code, cnt)
	}
	fmt.Println()
	fmt.Println("2. Zero Byte Connection (Without Local Cache)")
	for _, ent := range zeroByteEntry {
		fmt.Printf("%s : returns zero byte\n", har.Log.Entries[ent].Request.URL)
	}
	if len(zeroByteEntry) == 0 {
		fmt.Println("No Zero Byte Connection Found.")
	}
	fmt.Println()
	fmt.Println("3. Too Long Connection")
	for _, ent := range tooLongEntry {
		fmt.Printf("%s connection time is too long!\n", har.Log.Entries[ent].Request.URL)
	}
	if len(tooLongEntry) == 0 {
		fmt.Println("No Too Long Connection Found.")
	}
}

func inspectHARFile() {
	statusCode = make(map[string]int)
	statusEntry = make(map[int]int)
	mimeCount = make(map[string]int)
	for i := 0; i < len(har.Log.Entries); i++ {
		//Check Status Code

		statusCode[strconv.Itoa(responses[i].Status)]++
		statusEntry[i] = responses[i].Status
		//        Check Zero Byte files
		if responses[i].TransferSize == 0 && responses[i].Content.Size == 0 {
			zeroByteEntry = append(zeroByteEntry, i)
		}
		//Check long Time connection
		if har.Log.Entries[i].Time > 1000.0 {
			tooLongEntry = append(tooLongEntry, i)
		}
		mimeCount[responses[i].Content.MimeType]++
	}
}
