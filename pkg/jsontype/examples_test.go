package jsontype_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ctx42/convert/pkg/xcast"

	"github.com/ctx42/jsontype/pkg/jsontype"
)

func ExampleNew_uint_marshal() {
	jType := jsontype.New(uint(42))
	data, _ := json.Marshal(jType)

	fmt.Println(string(data))
	// Output:
	// {"type":"uint","value":42}
}

func ExampleNew_uint_unmarshal() {
	data := []byte(`{"type": "uint", "value": 42}`)

	gType := &jsontype.Value{}
	_ = json.Unmarshal(data, gType)

	fmt.Printf("%[1]v (%[1]T)\n", gType.GoValue())
	// Output:
	// 42 (uint)
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
	data := []byte(`{"type": "uint8", "value":99999}`)

	gType := &jsontype.Value{}
	err := json.Unmarshal(data, gType)

	fmt.Println(err)
	// Output:
	// jsontype: float64 value out of range for uint8
}

func ExampleFromConv() {
	reg := jsontype.NewRegistry()
	reg.Register(jsontype.Int, jsontype.FromConv(xcast.Float32ToInt))

	data := []byte(`{"type": "int", "value":42}`)

	gType := &jsontype.Value{}
	err := json.Unmarshal(data, gType)

	_ = err // Check error.

	fmt.Printf("unmarshalled: %[1]v (%[1]T)\n", gType.GoValue())
	// Output:
	// unmarshalled: 42 (int)
}

func ExampleRegister() {
	// Custom type named "seconds" representing duration in seconds.
	data := []byte(`{"type": "seconds", "value": 42}`)

	// Custom decoder for the "seconds" type.
	dec := func(value float64) (time.Duration, error) {
		return time.Duration(value) * time.Second, nil
	}

	// Register decoder.
	jsontype.Register("seconds", jsontype.FromConv(dec))

	gType := &jsontype.Value{}
	err := json.Unmarshal(data, gType)

	_ = err // Check error.

	fmt.Printf("unmarshalled: %[1]v (%[1]T)\n", gType.GoValue())
	// Output:
	// unmarshalled: 42s (time.Duration)
}
