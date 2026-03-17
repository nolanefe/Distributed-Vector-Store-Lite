package store

import (
	"math"
	"sync"
)

type Vector []float32
type Document struct {
	ID        string
	Embedding Vector
}

// VectorShard holds a collection of documents.
// It uses a read/write lock so multiple searches can happen safely at the same time.
type VectorShard struct {
	mu   sync.RWMutex
	docs []Document
}

// NewVectorShard creates a new, empty shard.
func NewVectorShard() *VectorShard {
	return &VectorShard{
		docs: make([]Document, 0),
	}
}

// AddDocument safely adds a new document to the shard.
// It uses a write lock to prevent memory errors if a search is happening simultaneously.
func (vs *VectorShard) AddDocument(doc Document) {
	vs.mu.Lock()
	defer vs.mu.Unlock()
	vs.docs = append(vs.docs, doc)
}

// cosineSimilarity calculates the distance between two vectors.
// A score closer to 1 means the vectors are very similar.
func cosineSimilarity(a, b Vector) float32 {
	var dotProduct, normA, normB float32
	for i := range a {
		dotProduct += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}
	if normA == 0 || normB == 0 {
		return 0
	}
	return dotProduct / (float32(math.Sqrt(float64(normA))) * float32(math.Sqrt(float64(normB))))
}

type SearchResult struct {
	DocID string
	Score float32
}

// Search scans the shard to find the similarity score for a given query.
// It uses a read lock, so multiple queries can be processed at the same time.
func (vs *VectorShard) Search(query Vector, topK int) []SearchResult {
	vs.mu.RLock()
	defer vs.mu.RUnlock()

	var results []SearchResult
	
	// Calculate the similarity score for every document in this shard
	for _, doc := range vs.docs {
		score := cosineSimilarity(query, doc.Embedding)
		results = append(results, SearchResult{DocID: doc.ID, Score: score})
	}
	
	return results
}