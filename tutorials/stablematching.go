package tutorials

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var NamesOfMen = make(map[int]string)
var NamesOfWomen = make(map[int]string)

// what number of pairs to match
func readInt(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		n, err := strconv.Atoi(text)
		if err == nil && n > 0 {
			return n
		}
		fmt.Println("Please enter a valid positive number")
	}
}

func readPreferences(n int, group string, names []string) [][]int {
	reader := bufio.NewReader(os.Stdin)
	prefs := make([][]int, n) // initialize an array of length n, each element of which is an 2D integer map

	for index, val := range names {
		fmt.Println(index, val)
	}

	for i := 0; i < n; i++ {
		for {
			if group == "man" {
				fmt.Printf("Enter %s's preferences as space-separated indices from (0 to %d):\n", NamesOfMen[i], n-1)
				fmt.Print("Women candidates: \t", NamesOfWomen)
			} else {
				fmt.Printf("Enter %s's preferences as space-separated indices from (0 to %d):\n", NamesOfWomen[i], n-1)
				fmt.Println("Men candidates: \t", NamesOfMen)
			}

			fmt.Println()
			line, _ := reader.ReadString('\n')
			parts := strings.Fields(line) // splitting words based on space

			if len(parts) != n {
				fmt.Printf("Requires %d strict preferences. Try again!", n)
				continue
			}

			valid := true
			seen := make(map[int]bool) // initializes a map object with keys of type int and values of type bool
			row := make([]int, n)

			for j, p := range parts { // j stores the index, p stores the iterated element
				idx, err := strconv.Atoi(p)
				if err != nil || idx < 0 || idx >= n || seen[idx] {
					valid = false
					break
				}
				seen[idx] = true
				row[j] = idx
			}
			if valid {
				prefs[i] = row
				break
			}
			fmt.Println("Invalid preferences. Try again!")
		}

	}
	return prefs

}

func galeShapley(n int, menPrefs, womenPrefs [][]int) [][2]int {
	fmt.Println(menPrefs)
	fmt.Println(womenPrefs)

	// Invert women's preferences for faster comparison
	womenRank := make([][]int, n)
	for w := 0; w < n; w++ {
		womenRank[w] = make([]int, n)
		for rank, m := range womenPrefs[w] {
			womenRank[w][m] = rank
		}
	}

	freeMen := make([]int, n)
	for i := range freeMen {
		freeMen[i] = i
	}

	nextProposal := make([]int, n) // next woman to propose to for each man
	engagedTo := make([]int, n)    // stores pairs of woman to man
	for i := range engagedTo {
		engagedTo[i] = -1
	}

	for len(freeMen) > 0 {
		m := freeMen[0]
		w := menPrefs[m][nextProposal[m]]
		nextProposal[m]++
		if engagedTo[w] == -1 {
			engagedTo[w] = m
			freeMen = freeMen[1:]
		} else {
			mPrime := engagedTo[w]
			if womenRank[w][m] < womenRank[w][mPrime] {
				engagedTo[w] = m
				freeMen[0] = mPrime
			}
		}
	}

	result := make([][2]int, n)
	for w, m := range engagedTo {
		result[m] = [2]int{m, w}
	}
	return result
}

func assignNames(num int, group string) []string {
	for {
		fmt.Printf("Enter names of %s separated by space\n", group)

		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')

		names := strings.Fields(line)

		if len(names) != num {
			fmt.Printf("Error! Requires %d names! Try again\n", num)
			continue
		}

		if group == "men" {
			for idx, mName := range names {
				NamesOfMen[idx] = mName
			}
		} else {
			for idx, wName := range names {
				NamesOfWomen[idx] = wName
			}
		}
		return names
	}
}

func ComputeStableMatching() {
	n := readInt("Enter number of pairs: ")

	// assign names for men and women
	menNames := assignNames(n, "men")
	womenNames := assignNames(n, "women")

	fmt.Println("\n-- Enter men's preferences --")
	menPrefs := readPreferences(n, "man", womenNames)
	fmt.Println("\n-- Enter women's preferences --")
	womenPrefs := readPreferences(n, "woman", menNames)

	pairs := galeShapley(n, menPrefs, womenPrefs)
	fmt.Println("\nStable matching pairs (man, woman):")
	for _, pair := range pairs {
		fmt.Printf("%s <-> %s\n", menNames[pair[0]], womenNames[pair[1]])
	}
}
