package main

import (
	"fmt"
	"log"
	"os"
)

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF
	WHITESPACE
	LEFT_BRACE
	RIGHT_BRACE
	LEFT_BRACKET
	RIGHT_BRACKET
	COLON
	COMMA
	STRING
	NUMBER
	BOOLEAN
)

var tokenNames = map[TokenType]string{
	ILLEGAL:       "ILLEGAL",
	EOF:           "EOF",
	WHITESPACE:    "WHITESPACE",
	LEFT_BRACE:    "{",
	RIGHT_BRACE:   "}",
	LEFT_BRACKET:  "[",
	RIGHT_BRACKET: "]",
	COLON:         ":",
	COMMA:         ",",
	STRING:        "STRING",
	NUMBER:        "NUMBER",
	BOOLEAN:       "BOOLEAN",
}

type Token struct {
	Type  TokenType
	Value string
}

func newToken(tokenType TokenType, value string) *Token {
	return &Token{
		Type:  tokenType,
		Value: value,
	}
}

func GetTokenList(fileName string) []*Token {
	tokens := []*Token{}
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	for i, b := range fileData {
		switch b {
		case '{':
			tokens = append(tokens, newToken(LEFT_BRACE, string(b)))
		case '}':
			tokens = append(tokens, newToken(RIGHT_BRACE, string(b)))
		case '[':
			tokens = append(tokens, newToken(LEFT_BRACKET, string(b)))
		case ']':
			tokens = append(tokens, newToken(RIGHT_BRACKET, string(b)))
		case ',':
			tokens = append(tokens, newToken(COMMA, string(b)))
		case ':':
			tokens = append(tokens, newToken(COLON, string(b)))
		case '"':
			readStringValue, newi, err := readString(i, fileData)
			if err != nil {
				panic(err)
			}
			i += newi
			tokens = append(tokens, newToken(STRING, readStringValue))

		}

	}
	tokens = append(tokens, newToken(EOF, ""))
	return tokens
}

func readString(i int, fileData []byte) (token string, count int, err error) {
	count = i
	for {
		if count >= len(fileData) {
			return "", 0, fmt.Errorf("unterminated string")
		}
		count++
		if fileData[count] == '"' {
			break
		}
	}
	return string(fileData[i:count]), count, nil
}
