package main

// Arguments to format are: [1]: type name
const allValuesMethod = `// %[1]sValues returns all values of the enum
func %[1]sValues() []%[1]s {
	return _%[1]sValues
}
`

// Arguments to format are: [1]: type name
const validMethodLoop = `// IsValid returns "true" if the value is listed in the enum definition. "false" otherwise
func (i %[1]s) IsValid() bool {
	return IsAValid%[1]s(string(i))
}
`

// Arguments to format are: [1]: type name
const validFuncLoop = `// IsAValid%[1]s returns "true" if the value is listed in the enum definition. "false" otherwise
func IsAValid%[1]s(s string) bool {
	for _, v := range _%[1]sValues {
		if s == string(v) {
			return true
		}
	}
	return false
}
`

// Arguments to format are: [1]: type name
const altStringValuesMethod = `func (%[1]s) Values() []string {
	return %[1]sStrings()
}
`

func (g *Generator) buildBasicExtras(values []Value, typeName string) {
	// At this moment, either "g.declareIndexAndNameVars()" or "g.declareNameVars()" has been called

	// Print the slice of values
	g.Printf("\nvar _%sValues = []%s{", typeName, typeName)
	for _, value := range values {
		g.Printf("\t%s, ", value.name)
	}
	g.Printf("}\n\n")

	// Print the map between name and value
	g.printValueMap(values, typeName)

	// Print the basic extra methods
	g.Printf(allValuesMethod, typeName)
	g.Printf(validMethodLoop, typeName)
	g.Printf(validFuncLoop, typeName)
}

func (g *Generator) printValueMap(values []Value, typeName string) {
	g.Printf("\nvar _%sNameToValueMap = map[string]%s{\n", typeName, typeName)

	var n int

	for _, value := range values {
		g.Printf("\t_%sConcat[%d:%d]: %s,\n", typeName, n, n+len(value.name), value.name)
		n += len(value.name)
	}

	g.Printf("}\n\n")
}
