%{

// This parsing rules has 3 shift/reduce conflicts,
// however it doesn't matter because shift rules are preceded.
//
// Auto-generated comments are not mixed up with godoc comments
// by the empty comment below.

//
package parser

import "github.com/nekketsuuu/lambda2pi/lib/syntax"

%}


%union {
       ident syntax.LambdaIdent
       term syntax.Lambda
}

%type<term>	top expr abstr atomic

%token	LPAR RPAR
%token	LAMBDA DOT
%token<ident>	IDENT


%%

top:
	expr
	{
		$$ = $1
		if l, isYyLex := yylex.(*yyLex); isYyLex {
			l.term = $$
		}
	}

expr:
	abstr
	{
		$$ = $1
	}
|	expr abstr
	{
		$$ = syntax.LApp{ First: $1, Second: $2 }
	}

abstr:
	atomic
|	LAMBDA IDENT DOT expr
	{
		$$ = syntax.LAbs{ Var: $2, Body: $4 }
	}

atomic:
	LPAR expr RPAR
	{
		$$ = $2
	}
|	IDENT
	{
		$$ = syntax.LVar{ Name: $1 }
	}


%%
