package jsontype_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ctx42/jsontype/pkg/jsontype"
)

func ExampleNew_uint() {
	// Create a `jsontype.Value` and marshall it.
	jType := jsontype.New(uint(42))
	data, _ := json.Marshal(jType)
	fmt.Printf("  marshalled: %s\n", string(data))

	// Unmarshall the value and show its Go type.
	gType := &jsontype.Value{}
	_ = json.Unmarshal(data, gType)
	fmt.Printf("unmarshalled: %[1]v (%[1]T)\n", gType.GoValue())

	// Output:
	//   marshalled: {"type":"uint","value":42}
	// unmarshalled: 42 (uint)
}

func ExampleNew_time() {
	// Create a `jsontype.Value` and marshall it.
	jType := jsontype.New(time.Date(2000, 1, 2, 3, 4, 5, 600000000, time.UTC))
	data, _ := json.Marshal(jType)
	fmt.Printf("  marshalled: %s\n", string(data))

	// Unmarshall the value and show its Go type.
	gType := &jsontype.Value{}
	_ = json.Unmarshal(data, gType)
	fmt.Printf("unmarshalled: %[1]v (%[1]T)\n", gType.GoValue())

	// Output:
	// marshalled: {"type":"time.Time","value":"2000-01-02T03:04:05.6Z"}
	// unmarshalled: 2000-01-02 03:04:05.6 +0000 UTC (time.Time)
}

func ExampleValue_UnmarshalJSON_safe() {
	data := []byte(`{"type":"byte","value":99999}`)

	gType := &jsontype.Value{}
	err := json.Unmarshal(data, gType)

	fmt.Println(err)
	// Output:
	// decodeByte: requires float64 value in range of uint8: invalid range
}
