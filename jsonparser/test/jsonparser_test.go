package jsonparser_test

import (
	"encoding/json"
	"testing"

	"github.com/tarekbadrshalaan/goStuff/jsonparser"
)

//!+test
//go test -v
func TestJSONParser(t *testing.T) {
	var tests = []struct {
		data        []byte
		keys        []string
		expectedkey string
		expectedval string
	}{
		{[]byte(`{"json": {"errors": [], "data": {"url": "www.example.com", "count": 0, "id": "123", "name": "foo"}}}`), []string{"json", "data"}, "name", "foo"},
		{[]byte(`{"widget": {
					"debug": "on",
					"text": {
						"data": "Click Here",
						"size": 36,
						"style": "bold",
						"name": "text1",
						"hOffset": 250,
						"vOffset": 100,
						"alignment": "center",
						"onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"}}}`),
			[]string{"widget", "text"}, "style", "bold"},
		{[]byte(`{
				"id": "0001", 
				"name": "Cake", 
				"batters":
					{"batter": 
						{ "id": "1001", "type": "Regular"}
					}
				}`),
			[]string{"batters", "batter"}, "id", "1001"},
	}

	for i, test := range tests {
		res, err := jsonparser.JSONParser(test.data, test.keys...)
		if err != nil {
			t.Errorf("TestJSONParser error found %v", err)
		}
		expected := test.expectedval
		actual := res[test.expectedkey]
		if actual != expected {
			t.Errorf("index : %d , the result not as expected\nexpected %v\nactual %v", i, expected, actual)
		}
	}
}

func TestGetkey(t *testing.T) {
	var tests = []struct {
		data        []byte
		keys        []string
		expectedkey string
		expectedval string
	}{
		{[]byte(`{"json": {"errors": [], "data": {"url": "www.example.com", "count": 0, "id": "123", "name": "foo"}}}`), []string{"json", "data"}, "name", "foo"},
		{[]byte(`{"widget": {
					"debug": "on",
					"text": {
						"data": "Click Here",
						"size": 36,
						"style": "bold",
						"name": "text1",
						"hOffset": 250,
						"vOffset": 100,
						"alignment": "center",
						"onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"}}}`),
			[]string{"widget", "text"}, "style", "bold"},
		{[]byte(`{
				"id": "0001", 
				"name": "Cake", 
				"batters":
					{"batter": 
						{ "id": "1001", "type": "Regular"}
					}
				}`),
			[]string{"batters", "batter"}, "id", "1001"},
	}

	for i, test := range tests {
		var objmap map[string]interface{}
		if err := json.Unmarshal(test.data, &objmap); err != nil {
			t.Errorf("TestGetkey, json.Unmarshal error found %v", err)
		}
		res, err := jsonparser.Getkey(objmap, test.keys...)
		if err != nil {
			t.Errorf("index : %d ,jsonparser.Getkey error found %v", i, err)
		}
		expected := test.expectedval
		actual := res[test.expectedkey]
		if actual != expected {
			t.Errorf("index : %d , the result not as expected\nexpected %v\nactual %v", i, expected, actual)
		}
	}
}

func TestParserstring(t *testing.T) {
	var tests = []struct {
		data        []byte
		keys        []string
		expectedval string
	}{
		{[]byte(`{"json": {"errors": [], "data": {"url": "www.example.com", "count": 0, "id": "123", "name": "foo"}}}`), []string{"json", "data", "name"}, "foo"},
		{[]byte(`{"widget": {
					"debug": "on",
					"text": {
						"data": "Click Here",
						"size": 36,
						"style": "bold",
						"name": "text1",
						"hOffset": 250,
						"vOffset": 100,
						"alignment": "center",
						"onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"}}}`),
			[]string{"widget", "text", "style"}, "bold"},
		{[]byte(`{"id": "0001", 
					"name": "Cake", 
					"batters":
						{"batter": 
							{ "id": "1001", "type": "Regular"}
						}
					}`),
			[]string{"batters", "batter", "id"}, "1001"},
	}

	for i, test := range tests {
		res, err := jsonparser.JSONParserstring(test.data, test.keys...)
		if err != nil {
			t.Errorf("TestJSONParser error found %v", err)
		}
		if res != test.expectedval {
			t.Errorf("index : %d , the result not as expected\nexpected %v\nactual %v", i, test.expectedval, res)
		}
	}
}

func TestGetkeystring(t *testing.T) {
	var tests = []struct {
		data        []byte
		keys        []string
		expectedval string
	}{
		{[]byte(`{"json": {"errors": [], "data": {"url": "www.example.com", "count": 0, "id": "123", "name": "foo"}}}`), []string{"json", "data", "name"}, "foo"},
		{[]byte(`{"widget": {
					"debug": "on",
					"text": {
						"data": "Click Here",
						"size": 36,
						"style": "bold",
						"name": "text1",
						"hOffset": 250,
						"vOffset": 100,
						"alignment": "center",
						"onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"}}}`),
			[]string{"widget", "text", "style"}, "bold"},
		{[]byte(`{"id": "0001", 
					"name": "Cake", 
					"batters":
						{"batter": 
							{ "id": "1001", "type": "Regular"}
						}
					}`),
			[]string{"batters", "batter", "id"}, "1001"},
	}

	for i, test := range tests {
		var objmap map[string]interface{}
		if err := json.Unmarshal(test.data, &objmap); err != nil {
			t.Errorf("TestGetkeystring, json.Unmarshal error found %v", err)
		}
		res, err := jsonparser.Getkeystring(objmap, test.keys...)
		if err != nil {
			t.Errorf("index : %d ,jsonparser.Getkey error found %v", i, err)
		}
		if res != test.expectedval {
			t.Errorf("index : %d , the result not as expected\nexpected %v\nactual %v", i, test.expectedval, res)
		}
	}
}

//!-tests

//!+bench
//go test -v  -bench=.
func BenchmarkJSONParser(b *testing.B) {
	data := []byte(`{"json": {
						"errors": [], 
						"data": {
								"url": "www.example.com", "count": 0, "id": "123", "name": "foo",
								"widget": {
									"debug": "on",
									"text": {
										"data": "Click Here",
										"size": 36,
										"style": "bold",
										"batters":{ 
											"batter": {
													"id": "1001",
													"type": "Regular"
											}
										}
									}
								}
							}
						}
					}`)

	for index := 0; index < b.N; index++ {
		res, err := jsonparser.JSONParser(data, "json", "data", "widget", "text", "batters", "batter")
		if err != nil {
			b.Errorf("TestJSONParser error found %v", err)
		}
		expected := "Regular"
		actual := res["type"]
		if actual != expected {
			b.Errorf("BenchmarkJSONParser, the result not as expected\nexpected %v\nactual %v", expected, actual)
		}

	}
}

func BenchmarkGetkey(b *testing.B) {
	data := []byte(`{"json": {
						"errors": [], 
						"data": {
								"url": "www.example.com", "count": 0, "id": "123", "name": "foo",
								"widget": {
									"debug": "on",
									"text": {
										"data": "Click Here",
										"size": 36,
										"style": "bold",
										"batters":{ 
											"batter": {
													"id": "1001",
													"type": "Regular"
											}
										}
									}
								}
							}
						}
					}`)

	var objmap map[string]interface{}
	if err := json.Unmarshal(data, &objmap); err != nil {
		b.Errorf("BenchmarkGetkey, json.Unmarshal error found %v", err)
	}
	for index := 0; index < b.N; index++ {
		res, err := jsonparser.Getkey(objmap, "json", "data", "widget", "text", "batters", "batter")
		if err != nil {
			b.Errorf("BenchmarkGetkey error found %v", err)
		}
		expected := "Regular"
		actual := res["type"]
		if actual != expected {
			b.Errorf("BenchmarkGetkey, the result not as expected\nexpected %v\nactual %v", expected, actual)
		}

	}
}

func BenchmarkParserstring(b *testing.B) {
	data := []byte(`{"json": {
						"errors": [], 
						"data": {
								"url": "www.example.com", "count": 0, "id": "123", "name": "foo",
								"widget": {
									"debug": "on",
									"text": {
										"data": "Click Here",
										"size": 36,
										"style": "bold",
										"batters":{ 
											"batter": {
													"id": "1001",
													"type": "Regular"
											}
										}
									}
								}
							}
						}
					}`)

	for index := 0; index < b.N; index++ {
		res, err := jsonparser.JSONParserstring(data, "json", "data", "widget", "text", "batters", "batter", "type")
		if err != nil {
			b.Errorf("BenchmarkParserstring error found %v", err)
		}
		if res != "Regular" {
			b.Errorf("BenchmarkParserstring, the result not as expected\nexpected %v\nactual %v", "Regular", res)
		}

	}
}

func BenchmarkGetkeystring(b *testing.B) {
	data := []byte(`{"json": {
						"errors": [], 
						"data": {
								"url": "www.example.com", "count": 0, "id": "123", "name": "foo",
								"widget": {
									"debug": "on",
									"text": {
										"data": "Click Here",
										"size": 36,
										"style": "bold",
										"batters":{ 
											"batter": {
													"id": "1001",
													"type": "Regular"
											}
										}
									}
								}
							}
						}
					}`)
	var objmap map[string]interface{}
	if err := json.Unmarshal(data, &objmap); err != nil {
		b.Errorf("BenchmarkGetkeystring, json.Unmarshal error found %v", err)
	}

	for index := 0; index < b.N; index++ {
		res, err := jsonparser.Getkeystring(objmap, "json", "data", "widget", "text", "batters", "batter", "type")
		if err != nil {
			b.Errorf("BenchmarkGetkeystring error found %v", err)
		}
		if res != "Regular" {
			b.Errorf("BenchmarkGetkeystring, the result not as expected\nexpected %v\nactual %v", "Regular", res)
		}

	}
}

//!-bench
