package tutorials

import "fmt"

func SwithWithFallThrough() {
	day, month := 1, 1

	switch {
	case day == 1:
		fmt.Println("Day is Monday!")
		fallthrough // fallthrough allows multiple switch-case statements to be checked even when matching a particular case; so this allows the next case statement to be checked
	case month == 1:
		fmt.Println("Month is January!")
		fallthrough // this fallthrough statement allows the default case to be executed as well
	default:
		fmt.Println("Utter woke nonsense!")
		// fallthrough // however fallthrough statement here gives compilation error; since there is no more case statement after this
	}
}
