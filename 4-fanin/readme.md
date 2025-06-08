# Fan-In Concurrency Pattern in Go

## Overview

The **fan-in pattern** merges multiple input channels into a single output channel. This allows multiple independent goroutines to send messages concurrently, with all results being consumed through a unified stream. This pattern is particularly useful in systems that involve aggregating or multiplexing data from distributed or parallel sources.

---

## Real-World Use Cases

### 1. **Log Aggregation**

In distributed systems or microservice architectures, different services generate logs independently. The fan-in pattern can be used to collect logs from multiple channels (or sources) and aggregate them into a single logging service for indexing or analysis.

**Example:**

* Each microservice sends logs to its own channel.
* A fan-in goroutine merges them into a central log consumer.

---

### 2. **Real-Time Event Streaming**

In applications like real-time dashboards or monitoring systems, events from various parts of the system must be streamed together to a frontend or processing engine.

**Example:**

* Stream user interactions, server events, and error alerts.
* Use fan-in to aggregate and send to a WebSocket server or event processor.

---

### 3. **Concurrent API Aggregation**

When calling multiple external APIs concurrently, each API’s result can be sent to a separate channel. A fan-in pattern collects these into one unified stream for further processing or response construction.

**Example:**

* Parallel calls to pricing, inventory, and shipping APIs.
* Merge results into one output channel for combining into the final user response.

---

### 4. **IoT Sensor Data Collection**

In an IoT system with many sensors (temperature, humidity, pressure, etc.), each sensor might send data on its own channel. Fan-in collects all sensor readings for centralized processing.

**Example:**

* Separate goroutines reading sensor data.
* A fan-in routine pushes all data to a central processing loop.

---

### 5. **Worker Pool Output Coordination**

In a worker pool pattern where multiple workers process jobs concurrently, each worker can write its output to a separate channel. Fan-in helps in unifying the outputs for downstream consumption (e.g., writing results to a database).

**Example:**

* Each worker returns processed data.
* Fan-in merges results into a single sink channel.

---

### 6. **Chatroom Message Relay**

In real-time chat applications, messages may come from various clients. A fan-in pattern is ideal for merging messages from all clients into a single stream for broadcasting or logging.

**Example:**

* Each client connection has a receiving goroutine and channel.
* Fan-in merges messages from all clients for broadcasting.

---

### 7. **Telemetry and Metrics Gathering**

In monitoring systems, different components may emit metrics (CPU, memory, latency, etc.) concurrently. Fan-in can unify these streams for batch processing, visualization, or alerting.

**Example:**

* Each component reports metrics on a dedicated channel.
* Fan-in collects into one stream for Prometheus or similar tools.


### 8. **Database Change Stream Merging**

In applications using multiple databases or partitions (e.g., sharded systems), each database instance may emit change events (like inserts/updates). Fan-in can merge those change streams into one for downstream consumers (e.g., cache sync, analytics pipelines).

**Example:**

* Each DB shard emits changes via a channel.
* Fan-in merges all change logs into a single channel for processing.

---

### 9. **File Processing Pipeline**

When processing multiple files in parallel (e.g., batch image resizing, log parsing), each file processor can write its output to a channel. Fan-in merges the processed data for writing to disk or uploading to a service.

**Example:**

* Goroutines parse different files concurrently.
* Fan-in aggregates parsed results to a single writer.

---

### 10. **Financial Data Aggregation**

In financial applications, data feeds from different stock exchanges or instruments may arrive concurrently. Fan-in can aggregate these feeds for live dashboards or algorithmic trading systems.

**Example:**

* Goroutines handle NYSE, NASDAQ, and crypto feeds.
* Fan-in merges all into one stream for real-time price updates.

---

### 11. **Web Scraping Aggregation**

In web crawlers or scrapers, each worker scraping a URL can emit results (HTML, metadata, errors) to its channel. Fan-in consolidates those results for storage or further analysis.

**Example:**

* Scraper goroutines emit parsed content.
* Fan-in delivers the results to a single output processor.

---

### 12. **Audio/Video Stream Multiplexing**

In media servers or surveillance systems, you might receive media packets from different camera/audio sources. Fan-in can combine these into a synchronized pipeline for encoding or archiving.

**Example:**

* Multiple RTSP feeds send packets concurrently.
* Fan-in merges them into a transcoding queue.

---

### 13. **Sensor Fusion Systems**

In robotics or autonomous systems, multiple sensors (GPS, LiDAR, IMU, cameras) provide asynchronous data. Fan-in helps consolidate this for fusion algorithms that rely on multi-sensor input.

**Example:**

* Each sensor has a goroutine and channel.
* Fan-in feeds combined data into a decision engine.

---

### 14. **CI/CD Pipeline Logs Aggregation**

In a distributed CI/CD system, each step or job (e.g., lint, test, deploy) can emit logs independently. Fan-in merges these logs into a single view for monitoring or archiving.

**Example:**

* Job containers output logs to different channels.
* Fan-in delivers a unified stream to the CI dashboard.

---

### 15. **Real-Time Multiplayer Game Event Handling**

In multiplayer games, each player or server shard might emit events (movement, chat, state changes). Fan-in aggregates all events to update the game state centrally.

**Example:**

* Player goroutines generate input events.
* Fan-in merges into a single game state updater channel.

