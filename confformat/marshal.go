package confformat

import "fmt"

// MarshalAll takes data stored in structus and converts them to the various data formats 
func MarshalAll() error {
j:= JSONData {
	Name: "NAME2",
	Age: 50,
}

y := YAMLData{
	Name: "NAME3",
	Age: 35,
}

jsonRes, err := j.ToJSON()
if err != nil {
	return err
}
fmt.Println("JSON Marshal=", jsonRes.String())

yamlRes, err := y.ToYAML()
if err != nil {
	return err
}
fmt.Println("YAML  Marshal=", yamlRes.String())
return nil 
}
