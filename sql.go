package main

// Arguments to format are: [1]: type name
const valueMethod = `func (i %[1]s) Value() (driver.Value, error) {
	return i.String(), nil
}
`

const scanMethod = `func (i *%[1]s) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of %[1]s: %%[1]T(%%[1]v)", value)
	}

	val, err := %[1]sString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}

`
const nullableImplementation = `
type Null%[1]s struct {
	%[1]s      %[1]s
	Valid      bool
}

func NewNull%[1]s(val interface{}) (x Null%[1]s) {
	x.Scan(val) // yes, we ignore this error, it will just be an invalid value.
	return
}

// Scan implements the Scanner interface.
func (x *Null%[1]s) Scan(value interface{}) (err error) {
	if value == nil {
		x.%[1]s, x.Valid = %[1]s(0), false
		return
	}

	err = x.%[1]s.Scan(value)
	x.Valid = (err == nil)
	return
}

// Value implements the driver Valuer interface.
func (x Null%[1]s) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	return x.%[1]s.String(), nil
}
`

func (g *Generator) addValueAndScanMethod(typeName string) {
	g.Printf("\n")
	g.Printf(valueMethod, typeName)
	g.Printf("\n\n")
	g.Printf(scanMethod, typeName)
	g.Printf("\n\n")
	g.Printf(nullableImplementation, typeName)
}
