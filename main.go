package main

import (
    "encoding/json"
    "fmt"
    ffjson "xstackeye/ffjson-demo/idl/ffjson"
)

type Response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func stdDemo() {
    fmt.Println("Marshal&Unmarshal using encoding/json:")
    res2D := &Response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println("Marshal result:", string(res2B))

    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := Response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println("Unmarshal result:", res)
    fmt.Println("res.Fruits[0]:", res.Fruits[0])
}

func ffjsonDemo() {
    fmt.Println("Marshal&Unmarshal using ffjson:")
    res2D := &ffjson.Response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := res2D.MarshalJSON()
    fmt.Println("Marshal result:", string(res2B))

    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := ffjson.Response2{}
    res.UnmarshalJSON([]byte(str))
    fmt.Println("Unmarshal result:", res)
    fmt.Println("res.Fruits[0]:", res.Fruits[0])
}

func main() {
    stdDemo()
    fmt.Println("--------------------------")
    ffjsonDemo()
}