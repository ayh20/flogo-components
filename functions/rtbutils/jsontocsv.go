package rtbutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"sort"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnjsontocsv{})
}

type fnjsontocsv struct {
}

func (fnjsontocsv) Name() string {
	return "jsontocsv"
}

func (fnjsontocsv) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeBool}, false
}

func (fnjsontocsv) Eval(params ...interface{}) (interface{}, error) {

	r := bytes.NewReader([]byte(params[0].(string)))
	d := json.NewDecoder(r)
	d.UseNumber()
	var obj interface{}
	if err := d.Decode(&obj); err != nil {
		return nil, err
	}

	results, _ := JSON2CSV(obj)

	pts, err := allPointers(results)
	if err != nil {
		return "", err
	}
	sort.Sort(pts)
	keys := pts.Strings()
	header := pts.DotNotations(false)

	var output string

	// write header ... don't do the ATM
	var writeheader bool = params[1].(bool)
	if writeheader {
		for col, val := range header {
			if col > 0 {
				output = output + "," + val
			} else {
				if output == "" {
					output = output + val
				} else {
					output = output + "\n" + val
				}
			}
		}
	}

	// process the flattened data... it could be one instance or multiple
	for _, result := range results {
		record := toRecord(result, keys)
		for col, val := range record {
			if col > 0 {
				output = output + "," + val
			} else {
				if output == "" {
					output = output + val
				} else {
					output = output + "\n" + val
				}
			}
		}
	}

	return output, nil
}

// JSON2CSV converts JSON to CSV.
func JSON2CSV(data interface{}) ([]KeyValue, error) {
	results := []KeyValue{}
	v := valueOf(data)
	switch v.Kind() {
	case reflect.Map:
		if v.Len() > 0 {
			result, err := flatten(v)
			if err != nil {
				return nil, err
			}
			results = append(results, result)
		}
	case reflect.Slice:
		if isObjectArray(v) {
			for i := 0; i < v.Len(); i++ {
				result, err := flatten(v.Index(i))
				if err != nil {
					return nil, err
				}
				results = append(results, result)
			}
		} else if v.Len() > 0 {
			result, err := flatten(v)
			if err != nil {
				return nil, err
			}
			if result != nil {
				results = append(results, result)
			}
		}
	default:
		return nil, errors.New("unsupported JSON structure")
	}

	return results, nil
}

func isObjectArray(obj interface{}) bool {
	value := valueOf(obj)
	if value.Kind() != reflect.Slice {
		return false
	}

	len := value.Len()
	if len == 0 {
		return false
	}
	for i := 0; i < len; i++ {
		if valueOf(value.Index(i)).Kind() != reflect.Map {
			return false
		}
	}

	return true
}
func allPointers(results []KeyValue) (pointers pointers, err error) {
	set := make(map[string]bool, 0)
	for _, result := range results {
		for _, key := range result.Keys() {
			if !set[key] {
				set[key] = true
				pointer, err := New(key)
				if err != nil {
					return nil, err
				}
				pointers = append(pointers, pointer)
			}
		}
	}
	return
}
func toRecord(kv KeyValue, keys []string) []string {
	record := make([]string, 0, len(keys))
	for _, key := range keys {
		if value, ok := kv[key]; ok {
			record = append(record, toString(value))
		} else {
			record = append(record, "")
		}
	}
	return record
}
func toString(obj interface{}) string {
	return fmt.Sprintf("%v", obj)
}
