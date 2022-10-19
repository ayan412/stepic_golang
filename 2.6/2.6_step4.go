func main() {
// package main уже объявлен, нужные импорты уже есть
var a,b int
//fmt.Scan(&a,&b)
fmt.Scan(&a, &b)
i,err := divide(a,b)
if err == nil {
fmt.Println(i)
} else {
fmt.Println("ошибка")
}
}