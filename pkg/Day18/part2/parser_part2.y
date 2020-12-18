%{

package Day18p2

import (
	"fmt"
	"unicode"
)


%}

%union{
	val int
}


%type <val> expr number
%token <val> DIGIT

%left '*'
%left '+'
%%

list	: /* empty */
	| list stat '\n'
	;

stat	:    expr
		{
      return $1
		}
	;

expr	:    '(' expr ')'
{ $$  =  $2 }
|    expr '+' expr
{ $$  =  $1 + $3 }
|    expr '*' expr
{ $$  =  $1 * $3 }
	|    number
	;

number	:    DIGIT
	|    number DIGIT
		{ $$ = 10 * $1 + $2 }
	;

%%

type CalcLex struct {
	s string
	pos int
}


func (l *CalcLex) Lex(lval *CalcSymType) int {
	var c rune = ' '
	for c == ' ' {
		if l.pos == len(l.s) {
			return 0
		}
		c = rune(l.s[l.pos])
		l.pos += 1
	}

	if unicode.IsDigit(c) {
		lval.val = int(c) - '0'
		return DIGIT
	}
	return int(c)
}

func (l *CalcLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}


func SolvePart2(s []string, c chan int)  {
    part2 := 0

    for _, ss := range s {
      part2 += CalcParse(&CalcLex{s: ss+"\n"})
      }


	c <- part2
    }
