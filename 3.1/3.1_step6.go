/*
В ходе анализа результатов переписи населения информация была сохранена в объекте типа map:

groupCity := map[int][]string{
	10:   []string{...}, // города с населением 10-99 тыс. человек
	100:  []string{...}, // города с населением 100-999 тыс. человек
	1000: []string{...}, // города с населением 1000 тыс. человек и более
}
При подготовке важного отчета о городах с населением 100-999 тыс. человек был подготовлен другой объект типа map:

cityPopulation := map[string]int{
	город: население города в тысячах человек,
}
Однако база данных с информацией о точной численности населения содержала ошибки, поэтому в cityPopulation в т.ч. была сохранена информация о городах, которые входят в другие группы из groupCity.

Ваша программа имеет доступ к обоим указанным отображениям, требуется исправить cityPopulation, чтобы в ней была сохранена информация только о городах из группы groupCity[100].

Функция main() уже объявлена, доступ к отображениям осуществляется по указанным именам. По результатам выполнения ничего больше делать не требуется, проверка будет осуществлена автоматически.


 */

// В САМОМ РЕШЕНИИ НА СТЕПИКЕ ДОСТАТОЧНО УКАЗАТЬ ЦИКЛ БЕЗ ВВОДА В ТЕЛО main СВОИХ НАБОРОВ КАРТ
package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   []string{"город1", "город2", "город3"}, // города с населением 10-99 тыс. человек
		100:  []string{"город4", "город5", "город6"}, // города с населением 100-999 тыс. человек
		1000: []string{"город7", "город8", "город9"}, // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"город4": 100,
		"город5": 107,
		"город6": 100,
		"город1": 10,
		"город9": 1000,
	}

	for k, v := range groupCity {
		if k != 100 {
			for _, city := range v {
				delete(cityPopulation, city)
			}
		}
	}
	fmt.Println(cityPopulation)

}
