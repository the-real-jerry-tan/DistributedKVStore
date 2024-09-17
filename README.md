
# Distributed Key-Value Store (Go)

## Overview

This project is a **distributed key-value store** built in Go. It demonstrates distributed system concepts, particularly how to implement a simple key-value store with support for multiple nodes. The key-value store allows nodes to join the cluster, store key-value pairs, and elect a leader through consensus mechanisms.

The project utilizes Go's powerful concurrency model, including goroutines and channels, to handle multiple nodes interacting with the store. It also uses thread-safe operations with mutexes to ensure data integrity when multiple processes are accessing the store concurrently.

## Key Features

- **Key-Value Storage**: Implements `Set` and `Get` operations to allow users to store and retrieve key-value pairs efficiently.
- **Node Management**: Nodes can join or leave the cluster dynamically. Each node maintains its own state and interacts with the store via the network.
- **Leader Election**: A simple consensus algorithm is used to elect a leader among the nodes. This is a basic version of leader-based consensus.
- **Concurrency**: The system takes full advantage of Go's goroutines and channels for handling concurrent requests to the key-value store.
- **Thread-Safe Operations**: The store ensures that concurrent reads and writes are properly managed using mutex locks to avoid data corruption.

## Project Structure

```
.
├── main.go           # Entry point of the application
├── store.go          # Contains the key-value store logic
├── cluster.go        # Handles node and cluster operations
├── node.go           # Defines node behaviors, joining and leaving clusters
├── raft.go           # Basic implementation of leader election
├── go.mod            # Go module definition for dependencies
└── go.sum            # Checksums for dependencies
```

## How to Run the Project

### Step 1: Clone the Repository

Clone the repository from GitHub:

```bash
git clone https://github.com/the-real-jerry-tan/DistributedKVStore.git
cd DistributedKVStore/src
```

### Step 2: Initialize the Go Module

If you haven't already initialized the Go module, run:

```bash
go mod init DistributedKVStore
```

### Step 3: Build and Run the Project

Run the project by building and running the `main.go` file:

```bash
go build
./main
```

This will start the server on `localhost:8080`.

### Step 4: Test Key-Value Operations

Use `curl` to test basic operations like storing and retrieving key-value pairs:

- **Set a key-value pair**:
   ```bash
   curl "http://localhost:8080/set?key=myKey&value=myValue"
   ```

- **Get a key-value pair**:
   ```bash
   curl "http://localhost:8080/get?key=myKey"
   ```

### Step 5: Test Node Management

You can simulate nodes joining and leaving the cluster:

- **Node joins cluster**:
   This can be simulated by adding more nodes to the cluster in the code. Each node will log messages about its status in the cluster.

- **Leader Election**:
   Once nodes are added to the cluster, the system will elect a leader and log the results of the election.

## Future Enhancements

- **Raft Consensus Algorithm**: The current leader election is basic and could be improved by implementing a full Raft consensus algorithm. This would make the system more robust and fault-tolerant across distributed nodes.
- **Persistent Storage**: Currently, the key-value store is in-memory. Adding persistent storage (e.g., using a database or file-based system) would make the store more resilient to crashes.
- **Enhanced Networking**: Implement more sophisticated networking with RPC (Remote Procedure Calls) or gRPC for more complex cluster communication between nodes.

## License

This project is © 2024 Jerry Tan. All Rights Reserved.
