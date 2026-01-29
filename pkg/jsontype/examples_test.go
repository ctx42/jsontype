package jsontype_test

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/ctx42/convert/pkg/convert"

	"github.com/ctx42/jsontype/pkg/jsontype"
)

func ExampleValue_MarshalJSON() {
	jType := jsontype.New(uint(42))
	data, _ := json.Marshal(jType)

	fmt.Println(string(data))
	// Output:
	// {"type":"uint","value":42}
}

func ExampleValue_UnmarshalJSON() {
	data := []byte(`{"type": "uint", "value": 42}`)
	gType := &jsontype.Value{}
	_ = json.Unmarshal(data, gType)

	fmt.Printf("%[1]v (%[1]T)\n", gType.GoValue())
	// Output:
	// 42 (uint)
}

func ExampleValue_MarshalJSON_time() {
	// Create a `jsontype.Value` and marshall it.
	tim := time.Date(2000, 1, 2, 3, 4, 5, 600000000, time.UTC)
	jType := jsontype.New(tim)
	data, _ := json.Marshal(jType)

	// Unmarshall the value and show its Go type.
	gType := &jsontype.Value{}
	_ = json.Unmarshal(data, gType)

	fmt.Printf("  marshalled: %s\n", string(data))
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
	// jsontype: value out of range: from float64 to uint8
}

func ExampleRegister_custom() {
	// Custom converter for a type named "seconds" representing duration in seconds.
	cnv := func(value float64) (time.Duration, error) {
		return time.Duration(value) * time.Second, nil
	}

	// Register converter.
	jsontype.Register("seconds", convert.ToAnyAny(cnv))

	// Custom type named "seconds" representing duration in seconds.
	data := []byte(`{"type": "seconds", "value": 42}`)

	gType := &jsontype.Value{}
	err := json.Unmarshal(data, gType)

	_ = err // Check error.

	fmt.Printf("unmarshalled: %[1]v (%[1]T)\n", gType.GoValue())
	// Output:
	// unmarshalled: 42s (time.Duration)
}

func ExampleFromMap() {
	m := map[string]any{
		"type":  "uint",
		"value": uint(42),
	}

	val, err := jsontype.FromMap(m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v (%T)\n", val.GoValue(), val.GoValue())
	// Output: 42 (uint)
}
