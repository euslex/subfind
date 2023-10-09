package main

import (
    "bufio"
    "fmt"
    "net/http"
    "os"
    "sync"
)

func main() {
    // Open the subdomain file for reading
    file, err := os.Open("subdomain.txt")
    if err != nil {
        fmt.Println("Error opening subdomain file:", err)
        return
    }
    defer file.Close()

    // Create a scanner to read lines from the file
    scanner := bufio.NewScanner(file)

    // Create a wait group to wait for all goroutines to finish
    var wg sync.WaitGroup

    // Iterate over each subdomain and check the status code
    for scanner.Scan() {
        subdomain := scanner.Text()
        wg.Add(1)
        go checkSubdomainStatus(subdomain, &wg)
    }

    // Wait for all goroutines to finish
    wg.Wait()
}

func checkSubdomainStatus(subdomain string, wg *sync.WaitGroup) {
    defer wg.Done()

    // Construct the full URL with the subdomain
    url := "http://" + subdomain

    // Send an HTTP GET request to the URL
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("%s: Error making HTTP request: %v\n", subdomain, err)
        return
    }
    defer resp.Body.Close()

    // Print the status code for the subdomain
    fmt.Printf("%s: Status Code %d\n", subdomain, resp.StatusCode)
}
