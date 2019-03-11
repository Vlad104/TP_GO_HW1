package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type Stack struct {
	buffer []int
	poiner int
}

func (st *Stack) Init(size int) {
	st.buffer = make([]int, size)
	st.poiner = 0
}

func (st *Stack) Push(value int) {
	st.buffer[st.poiner] = value
	st.poiner++
}

func (st *Stack) Pop() (int) {
	st.poiner--
	value := st.buffer[st.poiner]
	return value
}

// чтение входной строки
func getInput() (string) {
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	return myscanner.Text()
}

// проверка символа на число
func isDigit(symbol byte) bool {
	return symbol <= '9' && symbol >= '0'
}

// достаем из строки цифры
func makeNumber(str string, strCounter int) (int, int, error) {
	var buffer string
	// записываем цифры в буфер, пока они идут друг за другом (те одно число)
	for strCounter < len(str) {
		symbol := str[strCounter]
		if isDigit(symbol) {
			buffer += string(symbol)
			strCounter++;
		} else {
			break
		}
	}
	result, err := strconv.Atoi(buffer)
	return result, strCounter, err
}

// сам калькулятор
func RPN(input string) (int, error) {
	var stack Stack
	maxStackSize := len(input)
	stack.Init(maxStackSize)
	result := 0
	isEnd := false
	stringCounter := 0
	for stringCounter < len(input) && !isEnd {
		symbol := input[stringCounter]
		switch symbol {
		case ' ':
			break
		case '+':
			val1 := stack.Pop()
			val2 := stack.Pop()
			stack.Push(val2 + val1)
		case '-':
			val1 := stack.Pop()
			val2 := stack.Pop()
			stack.Push(val2 - val1)
		case '*':
			val1 := stack.Pop()
			val2 := stack.Pop()
			stack.Push(val2 * val1)
		case '/':
			val1 := stack.Pop()
			val2 := stack.Pop()
			stack.Push(val2 / val1)
		case '=':
			result = stack.Pop()
			isEnd = true
		default:
			// выделяем число из строки
			num, newStringCounter, err := makeNumber(input, stringCounter)
			if err != nil {
				return 0, err
			}
			stringCounter = newStringCounter;
			stack.Push(num)
		}
		stringCounter++
	}

	return result, nil
}

func main() {
	input := getInput()
	res, err := RPN(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}