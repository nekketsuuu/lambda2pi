%{

package parser

import "github.com/nekketsuuu/lambda2pi"

%}


%union {
       ident lambda2pi.LambdaIdent
       term lambda2pi.Lambda
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
		$$ = lambda2pi.LApp{ First: $1, Second: $2 }
	}

abstr:
	atomic
|	LAMBDA IDENT DOT expr
	{
		$$ = lambda2pi.LAbs{ Var: $2, Body: $4 }
	}

atomic:
	LPAR expr RPAR
	{
		$$ = $2
	}
|	IDENT
	{
		$$ = lambda2pi.LVar{ Name: $1 }
	}


%%
