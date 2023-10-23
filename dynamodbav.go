package main

const dynamodbAvMethods = `
// MarshalDynamoDBAttributeValue implements the attributevalue.Marshaler interface for %[1]s
func (i %[1]s) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: i.String()}, nil
}

// UnmarshalDynamoDBAttributeValue implements the attributevalue.Unmarshaler interface for %[1]s
func (i *%[1]s) UnmarshalDynamoDBAttributeValue(value types.AttributeValue) error {
	avS, ok := value.(*types.AttributeValueMemberS)
	if !ok {
		return fmt.Errorf("%[1]s should be a AttributeValueMemberS, got %%T", value)
	}

	var err error
	*i, err = %[1]sString(avS.Value)
	return err
}
`

func (g *Generator) buildDynamodbAvMethods(runs [][]Value, typeName string) {
	g.Printf(dynamodbAvMethods, typeName)
}
