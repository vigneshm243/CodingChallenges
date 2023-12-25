package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	// "os"
	"strconv"
	"strings"
	"unicode"
)

type Options struct {
	countBytes bool
	countLines bool
	countWords bool
	countChars bool
}

type stats struct {
	bytes uint64
	lines uint64
	words uint64
	chars uint64
}

func CalculateWordCount(reader *bufio.Reader) stats {
	var prevChar rune
	var bytesCount uint64
	var linesCount uint64
	var wordsCount uint64
	var charsCount uint64

	for {
		charRead, bytesRead, err := reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		bytesCount += uint64(bytesRead)
		charsCount++
		//If new line is encountered line count is increased
		if charRead == '\n' {
			linesCount++
		}
		// Incrementing the word count as and when we come across a space and prev character is not a space
		if !unicode.IsSpace(prevChar) && unicode.IsSpace(charRead) {
			wordsCount++
		}

		prevChar = charRead
	}

	return stats{bytes: bytesCount, words: wordsCount, lines: linesCount, chars: charsCount}
}

func PrintWordCount(wordCount stats, commandLineOptions Options, fileNames []string) {
	var formattedOp []string

	if commandLineOptions.countLines {
		formattedOp = append(formattedOp, strconv.FormatUint(wordCount.lines, 10))
	}
	if commandLineOptions.countWords {
		formattedOp = append(formattedOp, strconv.FormatUint(wordCount.words, 10))
	}
	if commandLineOptions.countBytes {
		formattedOp = append(formattedOp, strconv.FormatUint(wordCount.bytes, 10))
	}
	if commandLineOptions.countChars {
		formattedOp = append(formattedOp, strconv.FormatUint(wordCount.chars, 10))
	}
	if len(fileNames) == 1 {
		formattedOp = append(formattedOp, fileNames[0])
	}
	
	fmt.Println(strings.Join(formattedOp, "\t"))
}