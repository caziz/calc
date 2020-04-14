package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
import . "github.com/caziz/calc/lex"
import . "github.com/caziz/calc/parse"

func main() {
	repl()
}

func repl() {
	env := make(map[string]int)
	for {
		newEnv, result, err := readEval(env)
		env = newEnv
		if err != nil {
			fmt.Println(err)
		}
		if result != "" {
			fmt.Println(result)
		}
	}
}

func readEval(env map[string]int) (map[string]int, string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("[calc] ")
	lexeme, err := reader.ReadString('\n')
	if err != nil {
		return env, "", err
	}
	tokens, err := Lex(lexeme)
	if err != nil {
		return env, "", err
	}
	id := ""
	if len(tokens) > 2 && tokens[0].Type == Id && tokens[1].Type == Equals {
		id = tokens[0].Id
		tokens = tokens[2:]
	}
	e, err := Parse(tokens)
	if err != nil {
		return env, "", err
	}
	val := e.Eval(env)
	str := strconv.Itoa(val)
	if id == "" {
		return env, str, nil
	}
	env[id] = val
	return env, "", nil
}