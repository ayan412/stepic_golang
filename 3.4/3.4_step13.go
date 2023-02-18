package main

import (
        "fmt" // пакет используется для проверки ответа, не удаляйте его
        "strings"
       )
type Battery struct {
        charge string   
    }
    
func (b *Battery) String() string {
        n := strings.Count(b.charge, "0")
        o := strings.Count(b.charge, "1")
        res := "[" + strings.Repeat(" ", n) + strings.Repeat("X", o) + "]"
        return res
    }

func main() {
    // batteryForTest - не забывайте об имени
// } Скобка, закрывающая функцию main() вам не видна, но она есть
    var input string 
    fmt.Scan(&input)
    batteryForTest := &Battery{charge: input}
