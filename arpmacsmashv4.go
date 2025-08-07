package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "regexp"
    "strings"
)

// Helper function to normalize MAC address formats
func normalizeMacAddress(mac string) string {
    mac = strings.ToLower(mac)
    mac = strings.ReplaceAll(mac, ":", "")
    mac = strings.ReplaceAll(mac, "-", "")
    mac = strings.ReplaceAll(mac, ".", "")
    return mac
}

func main() {
    // Define command-line flags
    arpFile := flag.String("arpfile", "", "Path to the ARP table file")
    macFile := flag.String("macfile", "", "Path to the MAC address file")
    outputFile := flag.String("output", "", "Path to the output file (optional)")

    flag.Parse()

    if *arpFile == "" || *macFile == "" {
        fmt.Println("Both -arpfile and -macfile must be specified")
        flag.Usage()
        return
    }

    // Open the ARP file
    arpF, err := os.Open(*arpFile)
    if err != nil {
        fmt.Printf("Error opening ARP file: %v\n", err)
        return
    }
    defer arpF.Close()

    // Open the MAC file
    macF, err := os.Open(*macFile)
    if err != nil {
        fmt.Printf("Error opening MAC file: %v\n", err)
        return
    }
    defer macF.Close()

    // Open the output file if specified
    var output *os.File
    if *outputFile != "" {
        output, err = os.Create(*outputFile)
        if err != nil {
            fmt.Printf("Error creating output file: %v\n", err)
            return
        }
        defer output.Close()
    } else {
        output = os.Stdout
    }

    // Process ARP file and build a lookup table for MAC addresses
    arpScanner := bufio.NewScanner(arpF)
    macRegex := regexp.MustCompile(`(?:[0-9a-fA-F]{2}[:.-]?){5}[0-9a-fA-F]{2}`)
    arpLookup := make(map[string]string)

    for arpScanner.Scan() {
        line := arpScanner.Text()
        macMatches := macRegex.FindAllString(line, -1)
        for _, mac := range macMatches {
            normalizedMac := normalizeMacAddress(mac)
            arpLookup[normalizedMac] = line
        }
    }

    if err := arpScanner.Err(); err != nil {
        fmt.Printf("Error reading ARP file: %v\n", err)
        return
    }

    // Debug: Print ARP Lookup size
    fmt.Printf("Loaded %d MAC addresses from ARP file into lookup table\n", len(arpLookup))

    // Process MAC file and match entries with the ARP table
    macScanner := bufio.NewScanner(macF)
    for macScanner.Scan() {
        line := macScanner.Text()
        macMatches := macRegex.FindAllString(line, -1)

        for _, mac := range macMatches {
            normalizedMac := normalizeMacAddress(mac)
            if arpEntry, found := arpLookup[normalizedMac]; found {
                fmt.Fprintf(output, "%s %s\n", arpEntry, line)
            }
        }
    }

    if err := macScanner.Err(); err != nil {
        fmt.Printf("Error reading MAC file: %v\n", err)
    }
}
