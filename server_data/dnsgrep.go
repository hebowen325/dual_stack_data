// a lightweight utility to scan a sorted file for a substring at the start of each line

package main

import (
        . "dnsgrep/DNSBinarySearch"
        "fmt"
        "io"
        "io/ioutil"
        "os"
        "sort"
        "strings"
        "sync"

        "github.com/golang/example/stringutil"
        "github.com/jessevdk/go-flags"
)

type Options struct {
        IPv4File string `short:"i" long:"Ipfile" description:"A large file containing IPv4 addresses" required:"true"`
}

// command line parsing
var options Options
var parser = flags.NewParser(&options, flags.Default)
var IPaddr []string
var limits = Limits{
        MaxScan:        1000,    // 100MB
        MaxOutputLines: 1000000, // 1,000,000 lines
}
var filesize int
var wg sync.WaitGroup
var x string

func GetStringBuffer(f *os.File, offset int) (string, error) {
        _, err := f.Seek(int64(offset), 0)
        if err != nil {
                return "", err
        }
        var returnBuf []byte
        if offset+500 >= filesize {
                returnBuf = make([]byte, filesize-offset)
                _, err = io.ReadAtLeast(f, returnBuf, filesize-offset)
        } else {
                returnBuf = make([]byte, 500)
                _, err = io.ReadAtLeast(f, returnBuf, 500)
        }
        if err != nil {
                return "", err
        }
        return string(returnBuf), nil
}
func getNextLine(str string) string {
        // get the start of the next line
        lines := strings.Split(str, "\n")
        if len(lines) < 2 {
                // we expect the input file to be sufficiently large that we do not need to handle the EOF/start edge cases
                // we also expect that every line is less than 500 chars, that could also trigger this case
                return ""
        }

        // take out what we are going to compare
        // (the first line, after the next newline char, up to the length of the line we are trying to find)
        return lines[1]
}
func getLineDetails(f *os.File, offset int, searchStr string) (compareLine string, err error) {
        // get the string buffer at this offset
        stringBuffer, err := GetStringBuffer(f, offset)
        if err != nil {
                return "", err
        }

        // get the next line
        fullLine := getNextLine(stringBuffer)
        compareLine = fullLine
        if fullLine == "" {
                return "", fmt.Errorf("Failed to get next line from string buffer: %s\n", stringBuffer)
        }

        // filter out up to the length of the search string
        if len(compareLine) > len(searchStr) {
                compareLine = compareLine[0:len(searchStr)]
        }

        return
}
func DNSSearch(start int, end int, interval int) {
        var domain_name string
        var IPv4 string
        var IPv6 string
        f, err := os.Open(x + "/fdns_aaaa/Record.txt")
        if err != nil {
                return
        }
        fi, err := f.Stat()
        if err != nil {
                return
        }
        for i := start; i <= end; i += interval {
                if i < end-1 && strings.Compare(strings.Split(IPaddr[i], ",")[0], strings.Split(IPaddr[i+1], ",")[0]) == 0 {
                        continue
                }
                domain_name = stringutil.Reverse(strings.Split(IPaddr[i], ",")[0])
                IPv4 = stringutil.Reverse(strings.Split(IPaddr[i], ",")[1])
                for t := i - 1; t >= 0; t-- {
                        if strings.Compare(strings.Split(IPaddr[t], ",")[0], strings.Split(IPaddr[i], ",")[0]) == 0 {
                                IPv4 = IPv4 + "," + stringutil.Reverse(strings.Split(IPaddr[t], ",")[1])
                        } else {
                                break
                        }
                }
                var searchStr string = stringutil.Reverse(domain_name)
                // open the file & get it's size
                filesize = int(fi.Size())
                foundFileLocation := sort.Search(int(fi.Size()), func(j int) bool {
                        // use the intermediary function to get the line details at the offset we are currently considering
                        searchLineCompare, err := getLineDetails(f, j, searchStr)
                        if err != nil {
                                // this should trigger an error in the next phase causing us to fail out quickly
                                return false
                        }
                        // substring compare
                        if strings.Compare(searchStr, searchLineCompare) > 0 {
                                return false
                        } else {
                                return true
                        }
                }) // end sort.Search
                stringBuffer, err := GetStringBuffer(f, foundFileLocation)
                fullLine := getNextLine(stringBuffer)
                if err != nil {
                        fmt.Printf("%s;%s; \n", domain_name, IPv4)
                        continue
                }
                if fullLine == "" {
                        fmt.Printf("%s;%s; \n", domain_name, IPv4)
                        continue
                }
                output, err := DNSBinarySearch(strings.Split(fullLine, ";")[1], domain_name, limits)
                if err != nil {
                        fmt.Printf("%s;%s; \n", domain_name, IPv4)
                        continue
                } else {
                        flag := false
                        IPv6 = ""
                        for _, result := range output {
                                if strings.Split(result, ",")[1] == domain_name {
                                        if flag == true {
                                                IPv6 = IPv6 + ","
                                        }
                                        flag = true
                                        IPv6 = IPv6 + strings.Split(result, ",")[0]
                                }
                        }
                        if flag == false {
                                fmt.Printf("%s;%s; \n", domain_name, IPv4)
                        }
                }
        }
        f.Close()
        wg.Done()
}
func main() {
        // command line parsing
        _, err := parser.Parse()
        if err != nil {
                panic(err)
        }
        // increase our limits x10 as we're running this locally
        x = "/home/sgl/bwhe/DNSGrep"
        files, err := ioutil.ReadDir(x + "/fdns_a/")
        if err != nil {
                panic(err)
        }
        // 获取文件，并输出它们的名字
        for _, file := range files {
                if x+"/fdns_a/"+file.Name() == options.IPv4File {
                        data, err := ioutil.ReadFile(x + "/fdns_a/" + file.Name())
                        if err != nil {
                                fmt.Println("File reading error", err)
                                return
                        }
                        IPaddr = strings.Split(string(data), "\n")
                        for i := 0; i < 64; i++ {
                                wg.Add(1)
                                go DNSSearch(i, len(IPaddr)-2, 64)
                        }
                        wg.Wait()
                }
        } // main.go is really just a wrapper around this function
}
