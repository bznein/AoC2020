%{

package Day18p1

import (
	"fmt"
	"unicode"
  "time"

	"github.com/bznein/AoC2020/pkg/timing"
)


%}

%union{
	val int
}


%type <val> expr number
%token <val> DIGIT

%left '*' '+'
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


func SolvePart1(s []string, c chan int) {
	defer timing.TimeTrack(time.Now())
    part1 := 0

    for _, ss := range s {
      part1 += CalcParse(&CalcLex{s: ss+"\n"})
      }


	c <- part1
    }
