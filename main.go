package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/xuri/excelize/v2"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func YeniTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

type Trie struct {
	root *TrieNode
}

func YeniTrie() *Trie {
	return &Trie{root: YeniTrieNode()}
}

func (t *Trie) Insert(word string) {
	current := t.root
	word = strings.ToLower(word)
	for _, ch := range word {
		if _, ok := current.children[ch]; !ok {
			current.children[ch] = YeniTrieNode()
		}
		current = current.children[ch]
	}
	current.isEnd = true
}

func suggestionsRecursive(node *TrieNode, prefix string, results *[]string) {
	if node.isEnd {
		*results = append(*results, prefix)
	}
	for ch, child := range node.children {
		suggestionsRecursive(child, prefix+string(ch), results)
	}
}

func (t *Trie) AutoComplete(prefix string) []string {
	current := t.root
	prefix = strings.ToLower(prefix)
	for _, ch := range prefix {
		if nextNode, ok := current.children[ch]; ok {
			current = nextNode
		} else {
			return []string{}
		}
	}
	results := []string{}
	suggestionsRecursive(current, prefix, &results)
	return results
}

var trie = YeniTrie()

func autocompleteHandler(c *fiber.Ctx) error {
	prefix := c.Query("prefix")
	if prefix == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "prefix parametresi gerekli",
		})
	}

	suggestions := trie.AutoComplete(prefix)
	return c.JSON(fiber.Map{
		"prefix":      prefix,
		"suggestions": suggestions,
	})
}

func main() {
	f, err := excelize.OpenFile("idioms_and_meanings.xlsx")
	if err != nil {
		log.Fatalf("Excel dosyası açılamadı: %v", err)
	}
	defer f.Close()

	uniqueSozler := make(map[string]bool)
	sheets := f.GetSheetList()

	for _, sheet := range sheets {
		rows, err := f.GetRows(sheet)
		if err != nil {
			continue
		}

		for i, row := range rows {
			if i == 0 || len(row) == 0 {
				continue
			}
			soz := strings.TrimSpace(row[0])
			if soz != "" && !uniqueSozler[soz] {
				uniqueSozler[soz] = true
				trie.Insert(soz)
			}
		}
	}

	app := fiber.New(fiber.Config{
		AppName: "Atasözü Arama Servisi",
	})

	app.Use(cors.New())
	app.Get("/autocomplete", autocompleteHandler)
	app.Static("/", "./")

	port := 8080
	fmt.Printf("Sunucu http://localhost:%d adresinde çalışıyor...\n", port)
	app.Listen(fmt.Sprintf(":%d", port))
}
