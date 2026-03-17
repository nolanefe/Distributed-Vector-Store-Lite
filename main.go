package main

import (
	"fmt"
	"github.com/nolanefe/Distributed-Vector-Store-Lite/store"
)

func main() {
	fmt.Println("Initializing Vector Store...")

	// Create a new shard
	shard := store.NewVectorShard()

	// Add some sample document embeddings
	shard.AddDocument(store.Document{ID: "financial_doc_1", Embedding: store.Vector{0.1, 0.2, 0.3}})
	shard.AddDocument(store.Document{ID: "market_data_A", Embedding: store.Vector{0.9, 0.8, 0.7}})
	shard.AddDocument(store.Document{ID: "trading_algo_v2", Embedding: store.Vector{0.12, 0.22, 0.35}})

	fmt.Println("Data successfully stored.")

	// Perform a search
	query := store.Vector{0.1, 0.25, 0.3}
	fmt.Println("Running Cosine Similarity Search...")
	
	// Search for the query
	results := shard.Search(query, 2)

	// Print the results
	fmt.Printf("Search Results: %+v\n", results)
}