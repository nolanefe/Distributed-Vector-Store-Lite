# Distributed Vector Store Lite #

A distributed vector database where concurrent Go routines act as partitioned data shards to execute hyper-fast cosine similarity searches for scalable AI retrieval. This project serves as a lightweight, concurrent alternative to standard SQL databases, specifically engineered to keep high-dimensional embeddings in memory without heavy I/O bottlenecks.

## Core Architecture ##

* **Thread-Safe Sharding:** Utilizes `sync.RWMutex` to allow highly concurrent read operations (searches) while safely locking the partition during data ingestion.
* **Math Execution:** Implements linear scan cosine similarity logic to compute the directional distance between vectors with `O(N * D)` time complexity.
* **Concurrency Model:** Built entirely on Go's lightweight routines to manage memory safely and efficiently without relying on external database bloat.

## Tech Stack ##

* **Language:** Go (Golang)
* **Architecture:** In-Memory Sharding, Concurrent Programming
* **Domain:** Vector Databases, AI Data Infrastructure

## Build and Execution ##

This project requires the Go compiler to be installed on your system.

# 1. Clone the repository
git clone https://github.com/nolanefe/Distributed-Vector-Store-Lite.git
cd Distributed-Vector-Store-Lite

# 2. Run the node and execute the search pipeline
go run main.go
