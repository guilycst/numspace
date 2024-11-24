package main

import (
	"fmt"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
)

type Vocabulary struct {
	mu        sync.Mutex
	m         map[string]int
	docTokens [][]string
}

func (s *Vocabulary) setNE(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.m[key]; !exists {
		s.m[key] = len(s.m)
	}
}

func (s *Vocabulary) SortedKeys() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	keys := make([]string, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func Tokenize(doc string) []string {
	doc = strings.ToLower(doc)
	doc = strings.ReplaceAll(doc, ".", "")
	doc = strings.ReplaceAll(doc, ",", "")
	return strings.Fields(doc)
}

func chunkTokens(tokens []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(tokens); i += chunkSize {
		end := i + chunkSize
		if end > len(tokens) {
			end = len(tokens)
		}
		chunks = append(chunks, tokens[i:end])
	}
	return chunks
}

func NewVocabulary(documents ...string) *Vocabulary {
	vocab := Vocabulary{m: make(map[string]int), docTokens: make([][]string, len(documents))}

	// Max workers allowed
	maxWorkers := runtime.NumCPU() / 2
	chunkSize := 1000

	// Channels for work and worker control
	chunkCh := make(chan []string, maxWorkers)
	var workerCount int32     // Tracks active workers
	var idleWorkerCount int32 // Tracks idle workers
	var wg sync.WaitGroup     // Tracks worker lifetimes

	// Dynamic worker function
	worker := func() {
		defer atomic.AddInt32(&workerCount, -1) // Decrement worker count on exit
		defer wg.Done()

		for workChunk := range chunkCh {
			// Process chunk
			for _, word := range workChunk {
				vocab.setNE(word)
			}
			// Mark worker as idle
			atomic.AddInt32(&idleWorkerCount, 1)
		}
	}

	// Function to spawn a new worker if below the cap
	spawnWorkerIfNeeded := func() {
		// Check if we need to spawn a worker
		if atomic.CompareAndSwapInt32(&idleWorkerCount, 0, 1) { // Mark one worker as needed
			if atomic.LoadInt32(&workerCount) < int32(maxWorkers) {
				atomic.AddInt32(&workerCount, 1)
				wg.Add(1)
				go worker()
			}
		}
	}
	spawnWorkerIfNeeded() // Start first worker

	for i, doc := range documents {
		vocab.docTokens[i] = Tokenize(doc)
		for _, chunk := range chunkTokens(vocab.docTokens[i], chunkSize) {
			spawnWorkerIfNeeded()
			chunkCh <- chunk
		}
	}
	close(chunkCh)

	// Wait for all workers to finish
	wg.Wait()
	return &vocab
}

func BuildVectors(vocab *Vocabulary) [][]float64 {
	vectors := make([][]float64, len(vocab.docTokens))
	for i, tokens := range vocab.docTokens {
		vectors[i] = make([]float64, len(vocab.m))
		for _, token := range tokens {
			if index, exists := vocab.m[token]; exists {
				vectors[i][index]++
			}
		}
	}
	return vectors
}

func main() {
	docs := []string{
		"This is a test document.",
		"This document is another test.",
		"A completely different document.",
	}

	vocab := NewVocabulary(docs...)
	for word, index := range vocab.m {
		fmt.Printf("Word: %s, Index: %d\n", word, index)
	}

	// Sorted keys
	sortedKeys := vocab.SortedKeys()
	fmt.Println("Sorted Vocabulary:", sortedKeys)
}
