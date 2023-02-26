package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"syscall"
    "golang.org/x/term"
	"strings"
)

func main() {
    // prompt the user to enter a password
    fmt.Print("Enter Password: ")
    bytePassword, err := term.ReadPassword(int(syscall.Stdin))
    if err != nil {
        log.Println(err)
        return
    }
    password := string(bytePassword)
    fmt.Println("")

    // check the strength of the password
    score := 0
    length := len(password)
    if length >= 12 {
        score += 2
    } else if length >= 8 {
        score += 1
    }
    if containsNumber(password) {
        score += 1
    }
    if containsUppercaseLetter(password) {
        score += 1
    }
    if containsSymbol(password) {
        score += 1
    }

    // provide feedback to the user
    fmt.Printf("Password strength score: %d/5\n", score)
    switch {
    case score == 5:
        fmt.Println("Excellent! Your password is very strong.")
    case score >= 3:
        fmt.Println("Good, but there's still room for improvement.")
    case score >= 1:
        fmt.Println("Weak! Consider making your password longer and/or more complex.")
    default:
        fmt.Println("Password is too short! Please enter a password with at least 8 characters.")
    }

    // read the contents of the specified file in chunks
    filename := "/usr/share/wordlists/rockyou.txt"
    file, err := os.Open(filename)
    if err != nil {
        log.Println(err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Buffer(nil, 1024*1024)

    found := false
    for scanner.Scan() {
        line := scanner.Text()
        if containsSubstring(line, password) {
            fmt.Println("Your password appears in a public wordlist")
            found = true
            break
        }
    }

    if !found {
        fmt.Println("Couldn't find your password in my database.")
    }
}

func containsNumber(password string) bool {
    for _, c := range password {
        if c >= '0' && c <= '9' {
            return true
        }
    }
    return false
}

func containsUppercaseLetter(password string) bool {
    for _, c := range password {
        if c >= 'A' && c <= 'Z' {
            return true
        }
    }
    return false
}

func containsSymbol(password string) bool {
    symbols := "!@#$%^&*()_+-={}[]\\|;:'\"<>,.?/"
    for _, c := range password {
        if strings.ContainsRune(symbols, c) {
            return true
        }
    }
    return false
}

func containsSubstring(s, substr string) bool {
    return strings.Contains(s, substr)
}
