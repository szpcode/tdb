README
======

### What is Tools Databse for golang (tdb)?
-----------------
Tools that prevent SQL injection attacks for dynamic numbers of parameters.

### Usage

Example
```go
package main


import "github.com/szpcode/tdb"
import "fmt"

func main() {

    tdb.BindIntValue(":id:", 12)
    sql1 := tdb.Prepare("SELECT * FROM users WHERE id = :id:")

    fmt.Println("sql1: "+ sql1)

    var ids = make([]int,0)
    ids = append(ids, 10)
    ids = append(ids, 12)

    tdb.BindIntValues(":ids:", ids)
    sql2 := tdb.Prepare("SELECT * FROM users WHERE id IN (:ids:)")

    fmt.Println("sql2: "+ sql2)

    tdb.BindStrValue(":surname:", "%Now ' 1 = 1 %")
    sql3 := tdb.Prepare("SELECT * FROM users WHERE surname LIKE :surname:")

    fmt.Println("sql3: "+ sql3)

    var names = make([]string,0)
    names = append(names, "Nowak")
    names = append(names, "Kowa\" 1 = 1 \"lski")
    names = append(names, "Robinson")
    names = append(names, "Richardson")

    tdb.BindStrValues(":names:", names)
    sql4 := tdb.Prepare("SELECT * FROM users WHERE surname IN (:names:)")

    fmt.Println("sql4: "+ sql4)


    tdb.BindStrValue(":surname:", "Kowalski")
    tdb.BindStrValue(":name:", "Jan")
    tdb.BindIntValue(":age:", 23)

    var personality = make([]string,0)
    personality = append(personality, "care'free")
    personality = append(personality, "charming")
    personality = append(personality, "clever")

    tdb.BindStrValues(":personality:", personality)


    sql5 := tdb.Prepare("SELECT * FROM users WHERE surname = :surname: AND name = :name: AND age = :age: AND personality IN (:personality:)")
    fmt.Println("sql5: "+ sql5)
}
```
