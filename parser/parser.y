%{

package parser

import "github.com/nekketsuuu/lambda2pi"

%}


%union {
       ident lambda2pi.LambdaIdent
       term lambda2pi.Lambda
}

%type<term>	expr app atomic

%token	LPAR RPAR
%token	LAMBDA DOT
%token<ident>	IDENT


%%


expr:
	app
	{
		$$ = $1
		if l, isYyLex := yylex.(*yyLex); isYyLex {
			l.term = $$
		}
	}
|	LAMBDA IDENT DOT expr
	{
		$$ = lambda2pi.LAbs{ Var: $2, Body: $4 }
	}

app:
	atomic
	{
		$$ = $1
	}
|	app atomic
	{
		$$ = lambda2pi.LApp{ First: $1, Second: $2 }
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
