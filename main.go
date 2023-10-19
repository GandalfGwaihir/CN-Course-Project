package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func sendGetRequest(clientSocket net.Conn) {
    request := "GET /get HTTP/1.1\r\nHost: localhost\r\n\r\n"
    clientSocket.Write([]byte(request))

    response := make([]byte, 8192)
    clientSocket.Read(response)
    fmt.Printf("Response from the server:\n%s\n", string(response))
}

func sendGetRequestByName(clientSocket net.Conn) {
    name := ""
    fmt.Print("Enter name: ")
    reader := bufio.NewReader(os.Stdin)
    name, _ = reader.ReadString('\n')
    name = strings.TrimSpace(name)

    // Construct the GET request with the entered name as a query parameter
    request := fmt.Sprintf("GET /get2?name=%s HTTP/1.1\r\nHost: localhost\r\n\r\n", name)
    clientSocket.Write([]byte(request))

    response := make([]byte, 8192) // Adjust buffer size as needed
    clientSocket.Read(response)
    fmt.Printf("Response from the server:\n%s\n", string(response))
}

func sendPostRequest(clientSocket net.Conn) {
    prn := ""
    name := ""
    cgpa := ""
    activity := ""
    date := ""
    aadhar := ""
    phone := ""
    email := ""

    // Consume the newline character left in the input buffer
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter PRN number: ")
    prn, _ = reader.ReadString('\n')
    prn = strings.TrimSpace(prn)

    fmt.Print("Enter name: ")
    name, _ = reader.ReadString('\n')
    name = strings.TrimSpace(name)

    fmt.Print("Enter CGPA: ")
    cgpa, _ = reader.ReadString('\n')
    cgpa = strings.TrimSpace(cgpa)

    fmt.Print("Enter activity: ")
    activity, _ = reader.ReadString('\n')
    activity = strings.TrimSpace(activity)

    fmt.Print("Enter date (YYYY-MM-DD): ")
    date, _ = reader.ReadString('\n')
    date = strings.TrimSpace(date)

    fmt.Print("Enter Aadhar number: ")
    aadhar, _ = reader.ReadString('\n')
    aadhar = strings.TrimSpace(aadhar)

    fmt.Print("Enter phone number: ")
    phone, _ = reader.ReadString('\n')
    phone = strings.TrimSpace(phone)

    fmt.Print("Enter email: ")
    email, _ = reader.ReadString('\n')
    email = strings.TrimSpace(email)

    // Construct the POST data
    data := fmt.Sprintf("prn=%s&name=%s&cgpa=%s&activity=%s&date=%s&aadhar=%s&phone=%s&email=%s", prn, name, cgpa, activity, date, aadhar, phone, email)

    request := fmt.Sprintf("POST /post HTTP/1.1\r\nHost: localhost\r\nContent-Length: %d\r\n\r\n%s", len(data), data)
    clientSocket.Write([]byte(request))

    response := make([]byte, 8192)
    clientSocket.Read(response)
    fmt.Printf("Response from the server:\n%s\n", string(response))
}

func sendDeleteRequest(clientSocket net.Conn) {
    name := ""
    fmt.Print("Enter name to delete: ")
    reader := bufio.NewReader(os.Stdin)
    name, _ = reader.ReadString('\n')
    name = strings.TrimSpace(name)

    // Construct the DELETE request with the name as a query parameter
    request := fmt.Sprintf("DELETE /delete?name=%s HTTP/1.1\r\nHost: localhost\r\n\r\n", name)
    clientSocket.Write([]byte(request))

    response := make([]byte, 8192)
    clientSocket.Read(response)
    fmt.Printf("Response from the server:\n%s\n", string(response))
}

func sendPutRequest(clientSocket net.Conn) {
    name := ""
    newActivity := ""

    // Consume the newline character
    reader := bufio.NewReader(os.Stdin)
    reader.ReadString('\n')

    fmt.Print("Enter name to update: ")
    name, _ = reader.ReadString('\n')
    name = strings.TrimSpace(name)

    fmt.Print("Enter new activity: ")
    newActivity, _ = reader.ReadString('\n')
    newActivity = strings.TrimSpace(newActivity)

    // Construct the PUT request with the name and new activity as query parameters
    request := fmt.Sprintf("PUT /put?name=%s&activity=%s HTTP/1.1\r\nHost: localhost\r\n\r\n", name, newActivity)
    clientSocket.Write([]byte(request))

    response := make([]byte, 8192)
    clientSocket.Read(response)
    fmt.Printf("Response from the server:\n%s\n", string(response))
}

func main() {
    clientSocket, err := net.Dial("tcp", "127.0.0.1:8080") // Replace with the server's IP address and port
    if err != nil {
        fmt.Printf("Error in connection: %v\n", err)
        os.Exit(1)
    }
    defer clientSocket.Close()

    fmt.Println("Choose an option:")
    fmt.Println("1. Send a GET request")
    fmt.Println("2. Send a GET request by name")
    fmt.Println("3. Send a POST request")
    fmt.Println("4. Send a DELETE request")
    fmt.Println("5. Send a PUT request")

    var choice int
    fmt.Scan(&choice)

    switch choice {
    case 1:
        sendGetRequest(clientSocket)
    case 2:
        sendGetRequestByName(clientSocket)
    case 3:
        sendPostRequest(clientSocket)
    case 4:
        sendDeleteRequest(clientSocket)
    case 5:
        sendPutRequest(clientSocket)
    default:
        fmt.Println("Invalid option")
    }
}
