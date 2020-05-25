package constants

const (
	// Binary Nodes Operators
	OpEqual          = "=="
	OpDifferent      = "!="
	OpOrSymbol       = "||"
	OpOrVerbose      = "or"
	OpAndSymbol      = "&&"
	OpAndVerbose     = "and"
	OpLess           = "<"
	OpGreater        = ">"
	OpLessOrEqual    = "<="
	OpGreaterOrEqual = ">="
	OpAdd            = "+"
	OpSubtract       = "-"
	OpMultiply       = "*"
	OpDivide         = "/"
	OpModulo         = "%"
	OpIn             = "in"
	OpNotIn          = "not in"
	OpExponent       = "**"
	OpContains       = "contains"
	OpStartsWith     = "startsWith"
	OpEndsWith       = "endsWith"
	OpMatches        = "matches"
	OpRange          = ".."
	OpBitwiseAnd     = "&"
	OpBitwiseOr      = "|"
	OpBitwiseXor     = "^"
	OpBitwiseLShift  = "<<"
	OpBitwiseRShift  = ">>"

	// Unary Nodes Operators
	OpPositive   = "+"
	OpNegative   = "-"
	OpNotSymbol  = "!"
	OpNotVerbose = "not"
	OpBitwiseNot = "~"

	SingleOperators       = "+-/%#,?:~^"
	DoubleFirstOperators  = "&|=*<>!"
	DoubleSecondOperators = "&|=*<>"
)
