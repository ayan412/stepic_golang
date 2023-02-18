func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
                                            // все полученные значения имеют тип пустого интерфейс    
    val1, ok := value1.(float64)
    if !ok {
        fmt.Printf("value=%v: %T\n", value1, value1)
        return
    }
    val2, ok := value2.(float64) 
    if !ok {
        fmt.Printf("value=%v: %T\n", value2, value2)
        return
    }
    if val, ok := operation.(string); ok {
		switch val {
		case "+":
            result := val1 + val2
			fmt.Printf("%.4f\n", result)
		case "-":
			result := val1 - val2
			fmt.Printf("%.4f\n", result)
		case "*":
			result := val1 * val2
			fmt.Printf("%.4f\n", result)
		case "/":
			result := val1 / val2
			fmt.Printf("%.4f\n", result)
		default:
			fmt.Println("неизвестная операция")
			return
		}
	} else {
		fmt.Println("неизвестная операция")
		return
	}
    
}
