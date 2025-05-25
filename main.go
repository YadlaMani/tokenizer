package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)
func formatNumber(num float64) string{
	if num==float64(int64(num)){
		return fmt.Sprintf("%.1f",num);
	}
	return fmt.Sprintf("%g",num);
} 
func formatNumberEvaluate(num float64) string {
	if num == float64(int64(num)) {
		return fmt.Sprintf("%.f", num)
	}
	return fmt.Sprintf("%g", num)
}
func trimWhitespace(s string) string {
	result := ""
	for _, ch := range s {
		if !unicode.IsSpace(ch) {
			result += string(ch)
		}
	}
	return result
}
func evalBoolean(src string, unaries []rune) string {
	
	val := false
	if len(unaries)==0 {
		return src
	}
	if src == "true" {
		val = true
	} else if src == "false" || src == "nil" {
		val = false
	}

	
	for i := len(unaries) - 1; i >= 0; i-- {
		if unaries[i] == '!' {
			val = !val
		}
	}

	
	if val {
		return "true"
	}
	return "false"
}


func main() {
	tokens := map[rune]string{
		'(': "LEFT_PAREN",
		')': "RIGHT_PAREN",
		'{': "LEFT_BRACE",
		'}': "RIGHT_BRACE",
		'*': "STAR",
		'+': "PLUS",
		'.': "DOT",
		'-': "MINUS",
		'/': "SLASH",
		';': "SEMICOLON",
		',': "COMMA",
		'=': "EQUAL",
		'!': "BANG",
		'<': "LESS",
		'>':"GREATER",
		
	}
	reserved:=map[string]string{
		"and": "AND",
		"or":  "OR",
		"if":  "IF",
		"else": "ELSE",
		"while": "WHILE",
		"for": "FOR",
		"return": "RETURN",
		"function": "FUNCTION",
		"class": "CLASS",
		"var": "VAR",
		"let": "LET",
		"const": "CONST",
		"true": "TRUE",
		"false": "FALSE",
		"null": "NULL",
		"this": "THIS",
		"super": "SUPER",
		"break": "BREAK",
		"continue": "CONTINUE",
		"switch": "SWITCH",
		"case": "CASE",
		"default": "DEFAULT",
		"import": "IMPORT",
		"print":"PRINT",
		"nil": "NIL",
		"fun":"FUN",
		
	}
	bracketPairs := map[byte]byte{
	'(': ')',
	'{': '}',
	'[': ']',
}

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]
	if command != "tokenize" && command!="evaluate"{
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	if command == "evaluate" {
	src:=string(fileContents)
	src=strings.TrimSpace(src)
	hasNeg:=false
	extractUnaries := func(s string) ([]rune, string) {
		unaries := []rune{}
		for len(s) > 0 && (s[0] == '-' || s[0] == '!') {
			if s[0] == '-'{
				hasNeg=!hasNeg
				
			}else{
				unaries = append(unaries, rune(s[0]))
			}
			
			s = s[1:]
			s = strings.TrimSpace(s)
		}
		return unaries, s
	}
	unaries, src := extractUnaries(src)
	for {
		if len(src) < 2 {
			break
		}
		open := src[0]
		close := src[len(src)-1]
		expectedClose, ok := bracketPairs[open]
		if !ok || close != expectedClose {
			break
		}
		src = src[1 : len(src)-1]
		src = strings.TrimSpace(src)
	}
	if strings.HasPrefix(src, "\"") && strings.HasSuffix(src, "\"") && len(src) >= 2 {
		str := src[1 : len(src)-1]
		fmt.Println(str)
		os.Exit(0)
	}
	moreUnaries, src := extractUnaries(src)
	unaries = append(unaries, moreUnaries...)
	switch src{
	case "true","false","nil":
		fmt.Printf(evalBoolean(src,unaries))
		os.Exit(0)
	}
	parts := strings.Fields(src)
	if len(parts) == 1 {
		value := parts[0]
		num,err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid number format: %s\n", value)
			os.Exit(1)
		}
		if len(unaries)>0 {
			if len(unaries)%2==0{
				fmt.Println("true");
			}else{
				fmt.Println("false");
			}

		}else{
			if hasNeg{
				fmt.Printf("-%s\n", formatNumberEvaluate(num))
			}else{
				fmt.Println(formatNumberEvaluate(num))
			}
		}
		os.Exit(0)
	}else  {
		left, err1 := strconv.ParseFloat(parts[0], 64)
		op := parts[1]
		right, err2 := strconv.ParseFloat(parts[2], 64)

		if err1 != nil || err2 != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid number format.\n")
			os.Exit(1)
		}

		var result float64
		switch op {
		case "+":
			result = left + right
		case "-":
			result = left - right
		case "*":
			result = left * right
		case "/":
			if right == 0 {
				fmt.Fprintf(os.Stderr, "Error: Division by zero.\n")
				os.Exit(1)
			}
			result = left / right
		default:
			fmt.Fprintf(os.Stderr, "Error: Unknown operator: %s\n", op)
			os.Exit(1)
		}

		fmt.Println(formatNumberEvaluate(result))
		os.Exit(0)
	}

	
	fmt.Fprintf(os.Stderr, "Error: Unable to evaluate expression: %s\n", src)
	os.Exit(1)



} else {
	haderror := false
	line := 1
	for i := 0; i < len(fileContents); i++ {
		ch := rune(fileContents[i])
		if ch == '"' {
			i++
			str := ""
			for i < len(fileContents) && fileContents[i] != '"' {
				str += string(fileContents[i])
				i++
			}
			if i < len(fileContents) && fileContents[i] == '"' {
				fmt.Printf("STRING \"%s\" %s\n", str,str)
				
			}else{
				fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", line)
				haderror=true
			}
			continue

		}
		if unicode.IsLetter(ch) || ch=='_'{
			str:=""
			for i < len(fileContents) && (unicode.IsDigit(rune(fileContents[i])) || unicode.IsLetter(rune(fileContents[i])) || fileContents[i] == '_') {
				str+=string(fileContents[i])
				i++
			}
			if reservedToken,ok:=reserved[str];ok{
				fmt.Printf("%s %s null\n", reservedToken, str)
				
			}else{
				fmt.Printf("IDENTIFIER %s null\n", str)
			}
			
			i--
			continue
			

		}
		if unicode.IsDigit(ch){
			start:=i
			for i < len(fileContents) && unicode.IsDigit(rune(fileContents[i])) {
				i++
			}
			if i<len(fileContents) && fileContents[i] == '.' {
				i++
			  if i< len(fileContents) && unicode.IsDigit(rune(fileContents[i])) {
				for i< len(fileContents) && unicode.IsDigit(rune(fileContents[i])) {
					i++
				}


			}else{
				fmt.Fprintf(os.Stderr, "[line %d] Error: Invalid number format (dot not followed by digits).\n", line)
		haderror = true
		continue
			}
		}
		numStr := string(fileContents[start:i])
val, err := strconv.ParseFloat(numStr, 64)
if err != nil {
	fmt.Fprintf(os.Stderr, "[line %d] Error: Invalid number: %s\n", line, numStr)
	haderror = true
} else {
	fmt.Printf("NUMBER %s %s\n", numStr, formatNumber((val)))
}

i--
continue
		}
		if ch=='/'{
			if i+1 < len(fileContents) && fileContents[i+1] == '/' {
				i += 2
				for i <len(fileContents) && fileContents[i]!='\n'{
					i++
				}
				line++
				continue
			}else{
				fmt.Printf("%s %c null\n", tokens[ch], ch)
				continue
			}
		}
		//
		if ch=='=' || ch=='<' || ch=='>' || ch=='!' {
			if i+1 < len(fileContents) && fileContents[i+1] == '=' {
				fmt.Printf("%s %c= null\n",tokens[ch]+"_EQUAL",ch)
				i++ 
			} else {
				fmt.Printf("%s %c null\n",tokens[ch],ch)
			}
			continue
		}
		
		if tokenType, ok := tokens[ch]; ok {
			fmt.Printf("%s %c null\n", tokenType, ch)
		} else if ch == ' ' || ch == '\n' || ch == '\t' {
			if(ch=='\n'){
				line++;
			}
			continue
		}else{
			fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n",line, ch)
			haderror=true
				
		}
	}


	fmt.Println("EOF  null")
	if(haderror){
		os.Exit(65)
	}
	os.Exit(0)
}
	
}
