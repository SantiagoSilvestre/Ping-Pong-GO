# Go Concurrency Example: Ping Pong with Channels

## 📌 Overview

This project demonstrates basic concurrency in Go using:

- Goroutines
- Channels
- Synchronization through blocking operations

The program prints:

```
ping
pong
ping
pong
...
```

Continuously, in a controlled alternating sequence.

---

## 🧠 How It Works

### 1.Ping & Pong

Two goroutines continuously send messages to their respective channels:

```go
func ping(c chan string) {
	for {
		c <- "ping"
	}
}

func pong(c chan string) {
	for {
		c <- "pong"
	}
}
```

- `ping` sends `"ping"` to `c1`
- `pong` sends `"pong"` to `c2`
- Both run infinitely

---

### 2. Imprimir

```go
func imprimir(c1, c2 chan string) {
	for {
		ping := <-c1
		fmt.Println(ping)
		time.Sleep(time.Second * 1)

		pong := <-c2
		fmt.Println(pong)
		time.Sleep(time.Second * 1)
	}
}
```

This function:

1. Receives from `c1` → prints `"ping"`
2. Waits 1 second
3. Receives from `c2` → prints `"pong"`
4. Waits 1 second

👉 This enforces **strict alternation**

---

### 3. Main Function

```go
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go ping(c1)
	go pong(c2)
	go imprimir(c1, c2)

	var input string
	fmt.Scanln(&input)
}
```

- Starts all goroutines
- Keeps program running until user input

---

## ⚠️ Important Notes

### Blocking Channels

Channels are **unbuffered**, meaning:

- `c <- value` blocks until someone reads
- `<-c` blocks until someone writes

This ensures synchronization between goroutines.

---

### CPU Usage Consideration

The producers run infinite loops without delay:

```go
for {
	c <- "ping"
}
```

This is safe because the send operation blocks, but in real-world systems, it's better to use:

- `time.Sleep`
- or `time.Ticker`

---

## 🚀 Improvements (Next Steps)

You can evolve this example by:

- Using `time.Ticker` for controlled intervals
- Adding `context.Context` for graceful shutdown
- Using `select` for multiplexing channels

---

## 🎯 Key Takeaways

- Goroutines run concurrently
- Channels synchronize communication
- Order must be controlled explicitly if needed
- Go concurrency is simple but powerful

---

## 📦 Output Example

```
ping
pong
ping
pong
...
```

---

Happy coding 🚀
