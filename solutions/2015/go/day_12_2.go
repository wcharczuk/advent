package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/blendlabs/go-exception"
)

func test(label string, actual, expected interface{}) {
	if actual != expected {
		fmt.Printf("test %s failed! actual: %v expected: %v\n", label, actual, expected)
		os.Exit(1)
	}
}

func main() {
	// test("array", processFile("../testdata/day12_test_1"), 6)
	// test("sub object w/ red", processFile("../testdata/day12_test_2"), 4)
	// test("top level red", processFile("../testdata/day12_test_3"), 0)
	// test("array red", processFile("../testdata/day12_test_4"), 6)

	codeFile := "../testdata/day12"
	finalTotal := processFile(codeFile)
	fmt.Println("Final Total", finalTotal)
}

func processFile(codeFile string) float64 {
	if f, err := os.Open(codeFile); err == nil {
		defer f.Close()

		contents, readErr := ioutil.ReadAll(f)
		if readErr != nil {
			fmt.Printf("Error reading file: %v\n", readErr)
			return -1
		}

		object := map[string]interface{}{}
		jsonErr := json.Unmarshal(contents, &object)
		if jsonErr == nil {
			return processJsonObject(object)
		}

		array := []interface{}{}
		jsonErr = json.Unmarshal(contents, &array)
		if jsonErr == nil {
			return processJsonArray(array)
		}

		objectArray := []map[string]interface{}{}
		jsonErr = json.Unmarshal(contents, &objectArray)
		if jsonErr == nil {
			subTotal := 0.0
			for _, subObj := range objectArray {
				subTotal += processJsonObject(subObj)
			}
			return subTotal
		}
	} else {
		fmt.Printf("Error opening file: %v\n", err)
		return -1
	}
	return -1
}

func processJsonObject(object interface{}) float64 {
	primitiveResult, wasPrimitive := processPrimitive(object)
	if wasPrimitive {
		return primitiveResult
	}

	if objectTyped, isArray := object.([]interface{}); isArray {
		return processJsonArray(objectTyped)
	}

	if objectMap, isMap := object.(map[string]interface{}); isMap {
		return processJsonMap(objectMap)
	}

	fmt.Printf("Failed to process object %T\n", object)
	fmt.Printf("Field Value: %#v\n", object)
	fmt.Printf("Location: \n%s\n", exception.GetStackTrace())
	return 0
}

func processJsonArray(array []interface{}) float64 {
	total := 0.0

	for _, arrayValue := range array {
		if arrayValueAsInteger, isPrimitive := processPrimitive(arrayValue); isPrimitive {
			total = total + arrayValueAsInteger
		} else if arrayValueAsMap, isMap := arrayValue.(map[string]interface{}); isMap {
			total = total + processJsonMap(arrayValueAsMap)
		} else if arrayValueAsArray, isArray := arrayValue.([]interface{}); isArray {
			total = total + processJsonArray(arrayValueAsArray)
		}
	}

	return total
}

func processJsonMap(object map[string]interface{}) float64 {
	subTotal := 0.0
	for _, field := range object {

		fieldPrimitiveResult, fieldWasPrimitive := processPrimitive(field)
		if fieldWasPrimitive {
			subTotal = subTotal + fieldPrimitiveResult
			continue
		}

		if typed, isTyped := field.(string); isTyped {
			if typed == "red" {
				return 0
			}
			continue
		}

		if typed, isTyped := field.(map[string]interface{}); isTyped {
			subTotal = subTotal + processJsonObject(typed)
			continue
		}

		if typed, isTyped := field.([]interface{}); isTyped {
			subTotal = subTotal + processJsonArray(typed)
			continue
		}

		fmt.Printf("Failed to process field %T\n", field)
		fmt.Printf("Field Value: %#v\n", field)
		fmt.Printf("Location: \n%s\n", exception.GetStackTrace())
	}
	return subTotal
}

func processPrimitive(field interface{}) (float64, bool) {
	if typed, isTyped := field.(uint8); isTyped {
		return float64(typed), true
	}
	if typed, isTyped := field.(uint16); isTyped {
		return float64(typed), true
	}
	if typed, isTyped := field.(uint); isTyped {
		return float64(typed), true
	}
	if typed, isTyped := field.(uint64); isTyped {
		return float64(typed), true
	}
	if typed, isTyped := field.(int8); isTyped {
		return float64(typed), true
	}
	if typed, isTyped := field.(int16); isTyped {
		return float64(typed), true
	}
	if typed, isTyped := field.(int); isTyped {
		return float64(typed), true
	}
	if typed, isTyped := field.(int64); isTyped {
		return float64(typed), true
	}
	if typed, isTyped := field.(float32); isTyped {
		return float64(typed), true
	}
	if typed, isTyped := field.(float64); isTyped {
		return float64(typed), true
	}

	return 0, false
}
