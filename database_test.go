package tdb

import(
    "fmt"
)

func ExamplePrepare() {
    BindStrValue(":name:", "test' 1");
    w := Prepare("name: :name:");
    fmt.Println(w)
    // Output: name: "test\' 1"

}