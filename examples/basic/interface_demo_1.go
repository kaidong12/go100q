package basic

import "fmt"

func InterfaceDemo1_type_assertion() {
	var t interface{}
	//t = functionOfSomeType()
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
}

func InterfaceDemo1_type_assertion_with_ok(x interface{}) string {
	if x == nil {
		return "NULL"
	} else if _, ok := x.(int); ok {
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(uint); ok {
		return fmt.Sprintf("%d", x)
	} else if b, ok := x.(bool); ok {
		if b {
			return "TRUE"
		}
		return "FALSE"
	} else if s, ok := x.(string); ok {
		fmt.Printf("s is: ", s)
		return s
		//return sqlQuoteString(s) // (not shown)
	} else {
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}
