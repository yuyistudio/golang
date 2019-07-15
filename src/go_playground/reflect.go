package main

import (
	"fmt"
	"encoding/json"
	"reflect"
)

type CoordInfo struct {
	City     int64   `json:"city"`
	District int64   `json:"district"`
	Business []int64 `json:"business"`
	score    float32   `json:"score"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	LastTime float64 `json:"last_time"`
	lastTime2 float64 `json:"last_time"`
}

type ResidentInfo struct {
	CoordsInfo []*CoordInfo `json:"usp"`
}

func testJson() {
	profileInfoStr := `{"modify_time": 1492001002, "usp": [{"city": 23948102, "district": 23995653, "business": [23995653], "score": 2.0221428571428572, "lat": 29.338312, "lng": 105.292373, "type": 1, "last_time": 1491998976.0}, {"city": 23948102, "district": 23995653, "business": [23995653], "score": 0.3057142857142857, "lat": 29.336307, "lng": 105.275973, "type": 2, "last_time": 1491314784.0}, {"city": 23948102, "district": 23995653, "business": [23995653], "score": 0.22928571428571426, "lat": 29.340795, "lng": 105.2933, "type": 0, "last_time": 1489505281.0}, {"city": 23948102, "district": 23995653, "business": [23995653], "score": 0.22928571428571426, "lat": 29.334936, "lng": 105.274753, "type": 0, "last_time": 1491619297.0}]}`
	profileInfo := new(ResidentInfo)
	if err := json.Unmarshal([]byte(profileInfoStr), profileInfo); err != nil {
		fmt.Printf("failed to unmarshal json `%s`, reason `%v`\n", profileInfoStr, err.Error())
		return
	} else {
		fmt.Printf("OK %v\n", profileInfo.CoordsInfo[0].score)
	}
}

func formatSliceValue(vt reflect.Value) string {
	rt := vt.Type()
	valueType := rt.Elem()
	if valueType.Kind() != reflect.Ptr && valueType.Kind() != reflect.Struct {
		return fmt.Sprintf("%#v", vt.Interface())
	}
	res := fmt.Sprintf("[]%s{", valueType.String())
	for i := 0; i < vt.Len(); i++ {
		f := vt.Index(i)
		res += FormatValue(f) + ", "
	}
	if res[len(res)-2:] == ", " {
		res = res[:len(res)-2]
	}
	res += "}"
	return res
}

func FormatValue(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Slice:
		if v.IsNil() {
			return "nil"
		} else {
			return formatSliceValue(v)
		}
	case reflect.Ptr:
		if v.IsNil() {
			return "nil"
		} else {
			return "&" + FormatValue(reflect.Indirect(v))
		}
	case reflect.Struct:
		return formatStructValue(v)
	default:
		return fmt.Sprintf("%#v", v.Interface())
	}
}

func formatStructValue(vt reflect.Value) string {
	rt := vt.Type()
	res := rt.Name() + "("
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		if f.PkgPath != "" {
			continue
		}
		fv := vt.Field(i)
		fName := f.Name
		res = res + fmt.Sprintf("%s:%s, ", fName, FormatValue(fv))
	}
	if res[len(res)-2:] == ", " {
		res = res[:len(res)-2]
	}
	res += ")"
	return res
}

func ToDebugString(v interface{}) string {
	return FormatValue(reflect.ValueOf(v))
}

func main() {
	var a CoordInfo
	a.City = 33
	a.score = 100
	a.lastTime2 = 99
	fmt.Println(ToDebugString(a))
}