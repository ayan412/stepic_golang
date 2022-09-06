/*
Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно. У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.

Если значение On == false, то оба метода вернут false.

Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true), если его нет, то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.

Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на экземпляр этой структуры с именем testStruct в функции main, в дальнейшем программа проверит результат.

 */

package main

import "fmt"

type Gun struct {
	On bool
	Ammo, Power int
}
func (s *Gun) Shoot() (bool) {
	if s.On == false || s.Ammo == 0 {
		return false
	} else {
		s.Ammo = s.Ammo - 1
		return true
	}
}

func (r *Gun) RideBike() (bool) {
	if r.On == false || r.Power == 0 {
		return false
	} else {
		r.Power = r.Power - 1
		return true
	}
}

func main() {
	testStruct := new(Gun)
	//var t Gun
	//testStruct := &t
	fmt.Println(*testStruct)
}




