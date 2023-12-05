package main

// import (
// 	"bufio"
// 	"errors"
// 	"fmt"
// 	"io"
// 	"os"
// )

// func day1() (int, int, error) {
// 	answerP1, err := d1p1()
// 	if err != nil {
// 		return 0, 0, err
// 	}

// 	answerP2, err := d1p2()
// 	if err != nil {
// 		return 0, 0, err
// 	}

// 	return answerP1, answerP2, err
// }

// func d1p1() (int, error) {
// 	file, err := os.Open("inputs/d1p1.txt")
// 	if err != nil {
// 		return 0, errors.New("could not read file input")
// 	}

// 	defer file.Close()
// 	reader := bufio.NewReader(file)

// 	for {
// 		line, isPrefix, err := reader.ReadLine()
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}

// 			return 0, errors.New("there was an issue reading the file")
// 		}

// 		if isPrefix {
// 			return 0, errors.New("line too long to parse")
// 		}

// 		fmt.Println(line)
// 	}

// 	return 0, nil
// }

// func d1p2() (int, error) {
// 	file, err := os.Open("inputs/d1p2.txt")
// 	if err != nil {
// 		return 0, errors.New("could not read file input")
// 	}

// 	defer file.Close()
// 	reader := bufio.NewReader(file)

// 	for {
// 		line, isPrefix, err := reader.ReadLine()
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}

// 			return 0, errors.New("there was an issue reading the file")
// 		}

// 		if isPrefix {
// 			return 0, errors.New("line too long to parse")
// 		}

// 		fmt.Println(line)
// 	}

// 	return 0, nil
// }
