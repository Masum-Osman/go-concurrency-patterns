# Go Concurrency Patterns Learning Repository

This repository contains practical examples of Go concurrency patterns, demonstrating various approaches to handling concurrent operations in Go. Each pattern is implemented with clear examples and detailed explanations.

## Credits and Inspiration

This repository is an expanded and detailed version of the patterns found in [This Repo](https://github.com/lotusirous/go-concurrency-patterns). While the original repository provides excellent base examples, this version adds:
- Detailed documentation for each pattern
- Extended real-world use cases
- Comprehensive explanations
- Additional code comments and examples

## Table of Contents

| Pattern | Description |
|---------|-------------|
| [1. Basic Goroutine](1-boring/main.go) | Introduction to goroutines with a simple example showing concurrent execution using `go` keyword. Demonstrates basic concurrency without channels. |
| [2. Channel Communication](2-chan/main.go) | Shows how to use channels for communication between goroutines. Implements basic send and receive operations with channels. |
| [3. Generator Pattern](3-generator/main.go) | Demonstrates the generator pattern where a function returns a channel that generates values. Shows channel closure and range operations. |
| [4. Fan-In Pattern](4-fanin/readme.md) | Combines multiple input channels into a single output channel. Useful for merging streams of data from multiple sources. |
| [5. Restore Sequence Pattern](5-restore-sequence/readme.md) | Implements synchronized fan-in pattern where senders wait for acknowledgment before sending the next message. Ensures ordered processing of messages. |
| [6. Mixed Sequence Pattern](6-mixed-sequence/readme.md) | Combines both acknowledged and unacknowledged message processing in a single system. Useful for handling different priority levels of data. |
| [7. Select with Timeout](7-select-timeout/readme.md) | Shows how to implement timeouts in Go using the select statement. Essential for preventing indefinite blocking in concurrent operations. |
| [8. Quit Signal Pattern](8-quit-signal/readme.md) | Demonstrates graceful shutdown of goroutines using quit channels. Important for clean application shutdown and resource cleanup. |
| [9. Channel Chain Pattern](9-channel-chain/readme.md) | Illustrates how to chain multiple goroutines together using channels. Useful for creating processing pipelines. |

## Purpose

This repository serves as a learning resource for understanding common concurrency patterns in Go. Each pattern is accompanied by:
- Practical examples
- Detailed explanations
- Real-world use cases
- Best practices and considerations

## How to Use

1. Clone the repository
2. Navigate to any pattern directory
3. Read the associated readme (if available)
4. Run the example using `go run main.go`
5. Study the code and comments to understand the pattern

## Learning Path

The patterns are arranged in increasing order of complexity. It's recommended to:
1. Start with basic goroutines (1-boring)
2. Move on to channel basics (2-chan)
3. Progress through more complex patterns
4. Experiment with combining patterns for real-world scenarios

## Prerequisites

- Basic knowledge of Go programming
- Go installed on your system
- Understanding of basic concurrency concepts

## Contributing

Feel free to:
- Add more examples
- Improve documentation
- Suggest new patterns
- Report issues
- Submit improvements