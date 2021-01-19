package confformat

import "fmt"

const (
exampleJSON = `{"name":"Example2","age":98}`
exampleYAML = `name: Example3
age: 97
`
)
// UnmarshallAll takes data in various formats and converts them into structs 
func UnmarshalAll() error {
	j := JSONData{}
	y := YAMLData{}
	
	if err := j.Decode([]byte(exampleJSON)); err != nil {
		return err
	}
	fmt.Println("JSON Unmarshal = ", j)

	if err := y.Decode([]byte(exampleYAML)); err != nil {
		return err
	}
	fmt.Println("Yaml Unmarshal = ", y)
	return nil	
}
