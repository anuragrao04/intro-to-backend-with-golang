###### Introduction To Backend With Golang

### Welcome To Class 3

---

### TODO

1. APIs - An Introduction
2. HTTP & HTTP Methods
3. GIN - A Library In Golang To Build APIs
4. How Do You Install Libraries In Go?
5. The Unix Time Format

---

### APIs

- Stand For **A**pplication **P**rogramming **I**nterface
- Is an interface/language between two applications
- In our case, the frontend & backend
- Usually have a fixed format: `json`, `xml` etc.

---

### HTTP

- **H**yper **T**ext **T**ransfer **P**rotocol
- The protocol to talk on the web
- Other protocols exist, this is the one used for APIs and web pages

---

### HTTP Methods

- The HTTP protocol consists of a _verb_ and a *URL*
- When you access a website, say https://mail.google.com/mail/u/0/#inbox, you are sending a *GET* request (verb) to the */mail/u/0/#inbox* URL
- URLs are also called *routes* in this context. We will be using this word.

---

- These verbs have meaning - but they are suggestive. They do not enforce anything but they are useful in terms of naming different functions of your backend API.
- For example, when you *GET /user/id1234*, you are probably getting the details of the user with user ID id1234

---

- Other verbs that are most commonly used are *POST* *DELETE*, *UPDATE*, etc.
- We will mostly work with *POST* and *GET* requests

---

### GIN

https://gin-gonic.com/docs/

- GIN is a library in Go that helps you write APIs
- It's open source - that means you can go look at the source code and change it

---

### Let's Get Our Hands Dirty

1. Create a new *empty* folder. **name it appropriately, you know who you are**
2. Do a `go mod init day3` in a terminal in that folder. (Remember opening a folder on vscode -> then opening terminal)

---

3. **This Is How You Install A Library:** Run `go get github.com/gin-gonic/gin` in your terminal.
4. Here, `github.com/gin-gonic/gin` is the name of the *module* for gin. We are using this module as a library in our module named *class3* (remember modules and packages from last class)

---

5. Create a file `main.go` with the following contents:

```go
package main
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", pingHandler)
	r.Run()
}

func pingHandler(c *gin.Context) {
	// send back 'pong' in plain text
	c.String(200, "pong")
}
```

---

6. Run this code using:

```shell
go run main.go
```

---

7. Go to your browser and go to the url `http://localhost:8080/ping`. You should see a pong

---

### Purpose

- This is the 'hello world' of backend APIs
- This is also coded into the backends to verify if the backend program is running properly. For example, we have a ping route in my superlit backend. You can visit: `http://10.2.80.90:6969/ping` to receive back a pong!

---

### Moving On...

---

### How Do You Keep Track Of Time If You Are A Computer?

---

### The Unix Time Format

- Just start counting milli seconds and store it in an integer
- From when? Random - January 1st 1970

---

### Limitation?

- The 32 bit integer
- Binary Counting?

---

### Back To Golang

##### In Class Assignment

- Make a route to return the current time in the unix time format
- The route must be *GET /time* which will return the current unix time in plain text format
- The unix time looks like this: `1729068727`
- Google it & RTFM

---


