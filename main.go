package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)
func formatNumber(num float64) string{
	if num==float64(int64(num)){
		return fmt.Sprintf("%.1f",num);
	}
	return fmt.Sprintf("%g",num);
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

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]
	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	haderror:=false;
	line:=1
	for i := 0; i < len(fileContents); i++{
		ch:=rune(fileContents[i])
		if ch=='"'{
			i++
			str:=""
			for i < len(fileContents) && fileContents[i] != '"' {
				str+=string(fileContents[i])
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
