package handler

import (
	"sort"
	"strings"
)

func PostBook(book *Book) (map[string]int, error) {
	resp := make(map[string]int)
	
	inputs := strings.Split(book.Content, " ")
	resp = sortWords(processWords(inputs))
	return resp, nil
}

func processWords(words []string) map[string]int {
	resp := make(map[string]int)
	for _, w:= range words {
		if c, exists := resp[w]; exists{
			resp[w] = c + 1
		}else{
			resp[w] = 1
		}
	}

	return resp
}

func sortWords(input map[string]int) map[string]int {
	resp := make(map[string]int)

	keys := make([]string, 0, len(input))
    for k := range input {
        keys = append(keys, k)
    }

	sort.Slice(keys, func(i, j int) bool {
        return input[keys[i]] > input[keys[j]]
    })

	for i, k:= range keys{
		if i > 9{
			break
		}
		resp[k] = input[k]
	}

	return resp
}