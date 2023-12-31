package main

import (
	"fmt"
)

func parse(tokens []*Token) (interface{}, error) {

	var parsedJSON interface{}
	var err error
	for i, token := range tokens {
		switch token.Type {
		case LEFT_BRACE:
			count := 0
			parsedJSON, count, err = readObject(i, tokens)
			if err != nil {
				return nil, err
			}
			i += count
		case EOF:
			break
		}
	}

	return parsedJSON, nil

}

func readObject(i int, tokens []*Token) (interface{}, int, error) {
	var obj map[string]interface{}
	i++
	// if we reach } immediately
	if tokens[i].Type == RIGHT_BRACE {
		return obj, i, nil
	}
	//looping through the strings
	for {
		token := tokens[i]
		// key should be a string

		if token.Type != STRING {
			return obj, 0, fmt.Errorf("Expected a key String but got %s at %d", token.Value, i)
		}
		key := token.Value

		i++
		// Now checking if there is a Colon between Key and Value
		if token.Type != COLON {
			return obj, 0, fmt.Errorf("Expected a key String but got %s at %d", token.Value, i)
		}
		i++
		count := 0
		value, count, err := parseToken()
		i += count
	}
	return obj, i
}
