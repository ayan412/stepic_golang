package main
import "fmt"
var workArray [10]uint8
var n uint8
func main() {
   //for ind := range workArray {
   //workArray[ind] = fmt.Scan(&n)
   //fmt.Println(workArray[ind])}
   for i := 0; i <= len(workArray); i++ {
      workArray[i], _ = fmt.Scan(&n)
      fmt.Println(workArray)
   }
   fmt.Println(workArray)
}

