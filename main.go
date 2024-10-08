package main

import "fmt"
import "log"
import "math/rand"
import "strings"
import "github.com/eiannone/keyboard"


const GREEN = "\033[32m"
const RED = "\033[31m"
const CYAN	= "\033[1;36m"
const RESET = "\033[0m"

const DICE_NUMBER = 6
const MAXSCORE = 4000

const rule = `                                         HELP
--------------------------------------------------------------------------------------
        space = Select
        e     = add selected dice + reroll
        f     = add selected dice + end turn
        h     = help or quit help
        q     = quit
        +---+
        | 1 | = 100 points
        +---+
        +---+
        | 5 | = 50 points
        +---+
        +---+ +---+ +---+
        | 1 | | 1 | | 1 | = 1000 points exception
        +---+ +---+ +---+
        +---+ +---+ +---+
        | 4 | | 4 | | 4 | = dice num * 100 points
        +---+ +---+ +---+
        +---+ +---+ +---+ +---+
        | 4 | | 4 | | 4 | | 4 | = (dice num * 100) * 2 points etc
        +---+ +---+ +---+ +---+
`
var (
	d1 = []string{
		"+-------+",
		"|       |",
		"|   O   |",
		"|       |",
		"+-------+",
	}

	d2 = []string{
		"+-------+",
		"| O     |",
		"|       |",
		"|     O |",
		"+-------+",
	}

	d3 = []string{
		"+-------+",
		"| O     |",
		"|   O   |",
		"|     O |",
		"+-------+",
	}

	d4 = []string{
		"+-------+",
		"| O   O |",
		"|       |",
		"| O   O |",
		"+-------+",
	}

	d5 = []string{
		"+-------+",
		"| O   O |",
		"|   O   |",
		"| O   O |",
		"+-------+",
	}

	d6 = []string{
		"+-------+",
		"| O   O |",
		"| O   O |",
		"| O   O |",
		"+-------+",
	}
)

const (
	HELP = iota
	WINNER
	GAME
)

func roll(diceNumber int) []int {
	var newSet []int
	for j := 0; j < diceNumber; j++ {
		newSet = append(newSet, rand.Intn(6)+1)
	}
	return newSet
}

func asciiDice(dice int) []string {
	switch dice {
	case 1:
		return d1
	case 2:
		return d2
	case 3:
		return d3
	case 4:
		return d4
	case 5:
		return d5
	case 6:
		return d6
	default:
		return d1
	}
}

func printDices(dices [][]int, selectedDices []int, selected int) {
	for i, row := range dices {
		if row == nil{
			continue
		}
		for line := 0; line < 6; line++ {
			fmt.Print("           ")
			for j, val := range row {
				if i == 0{
					if line == 5 {
						if j == selected {
							fmt.Print("    ^    ")
						} else {
							fmt.Print("           ")
						}
					} else {
						if contains(selectedDices, j) {
							fmt.Print(GREEN, asciiDice(val)[line], "  ", RESET)
						} else {
							fmt.Print(RED, asciiDice(val)[line], "  ", RESET)
						}
					}
				}else{
					if line < 5{
						fmt.Print(CYAN, asciiDice(val)[line], "  ", RESET)
					}
				}
			}
			fmt.Println()
		}
	}
}

func leaderboard(playerScore [2]int, playerName [2]string, currentPlayer int, tmpScore int) {
	player := ">>> " + playerName[currentPlayer] + " <<<"
	fmt.Print("\033[H\033[2J")
	fmt.Println(strings.Repeat("-", 86))
	fmt.Printf("%s\n", centerText("Player Turn", 86))
	fmt.Printf("%s\n", centerText(player, 86))
	fmt.Println(strings.Repeat("-", 86))
	fmt.Printf("                        %-18s: %5d / %-5d\n", "Score " + playerName[0], playerScore[0], MAXSCORE)
	fmt.Printf("                        %-18s: %5d / %-5d\n", "Score " + playerName[1], playerScore[1], MAXSCORE)
	fmt.Println(strings.Repeat("-", 86))
	fmt.Printf("                        %-18s: %5d\n", "Potential score ", tmpScore)
	fmt.Println(strings.Repeat("-", 86))
}

func helper(){
	fmt.Println(strings.Repeat("-", 86))
	fmt.Printf("%s", rule)
	fmt.Println(strings.Repeat("-", 86))
}

func winner(playerName string){
	fmt.Println(strings.Repeat("-", 86))
	fmt.Printf("%s\n", centerText("Winner : " + playerName, 86))
	fmt.Printf("%s\n", centerText("Press any key to replay.", 86))
	fmt.Println(strings.Repeat("-", 86))
}

func scoreCount(slice []int) int{
	var count int
	countDice := make(map[int]int)
	
	for _, value := range slice {
		countDice[value]++
	}

	for key, value := range countDice {
		multiplier := 100 * key
		switch value {
		case 1, 2:
			if (key == 1){
				count += 100 * value
				continue
			}else if (key == 5){
				count += 50 * value
				continue
			}
			return 0
		case 3:
			if (key == 1){count += 1000; continue}
			count += multiplier
		case 4:
			if (key == 1){count += 1100; continue}
			count += multiplier * 2
		case 5:
			if (key == 1){count += 1200; continue}
			count += multiplier * 4
		case 6:
			if (key == 1){count += 2000; continue}
			count += multiplier * 8
		}
	}

	return count
}

func takeDicesAndReroll(dices *[][]int, selectedDices *[]int) []int{
	emptyRowIndex := -1
	selectLen := len(*selectedDices)
	diceLen := len((*dices)[0])

	// security
	if *selectedDices == nil || dices == nil{
		return nil
	}

	// find the empty row
	for i, row := range *dices {
		if row == nil{
			emptyRowIndex = i;
			break
		}
	}
	if emptyRowIndex == -1{
		return nil
	}

	// moving dices to the empty row
	for i, dice := range (*dices)[0] {
		if contains((*selectedDices), i){
			(*dices)[emptyRowIndex] = append((*dices)[emptyRowIndex], dice)
		}
	}

	// reroll dices
	(*dices)[0] = roll(diceLen - selectLen)
	*selectedDices = (*selectedDices)[:0]
	return (*dices)[emptyRowIndex]
}

func changePlayer(currentPlayer *int){
	if *currentPlayer == 0 {
		*currentPlayer = 1 
	}else if *currentPlayer == 1 {
		 *currentPlayer = 0 
	}
}

func reset(dices *[][]int, selectedDices *[]int, tmpScore *int, currentPlayer *int){
	if (dices != nil) { *dices = make([][]int, 8); (*dices)[0] = roll(6) }
	if (selectedDices != nil) { *selectedDices = make([]int, 0) }
	if (tmpScore != nil) { *tmpScore = 0 }
	if (tmpScore != nil) { changePlayer(currentPlayer) }
}

func resetGame(dices *[][]int, selectedDices *[]int, tmpScore *int, currentPlayer *int, state *int, selected *int, playerScore *[2]int){
	if (dices != nil) { *dices = make([][]int, 8); (*dices)[0] = roll(6) }
	if (selectedDices != nil) { *selectedDices = make([]int, 0) }
	if (tmpScore != nil) { *tmpScore = 0 }
	if (currentPlayer != nil) { *currentPlayer = 0; }
	if (state != nil) { *state = HELP }
	if (selected != nil) { *selected = 0 }
	if (playerScore != nil) { *playerScore = [2]int{} }
}

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	var tmpScore int
	var playerScore [2]int
	var selectedDices []int
	state := HELP
	selected := 0
	currentPlayer := 0;
	dices := make([][]int, 8)
	dices[0] = roll(DICE_NUMBER)
	playerName := [2]string{"Bob", "Robert"}

	for {
		fmt.Print("\033[H\033[2J")
		switch state {
		case HELP:
			helper()
		case WINNER:
			leaderboard(playerScore, playerName, currentPlayer, tmpScore)
			if currentPlayer - 1 < 0 {
				winner(playerName[currentPlayer]) 
			}else{
				winner(playerName[currentPlayer-1])
			}
		case GAME:
			leaderboard(playerScore, playerName, currentPlayer, tmpScore)
			printDices(dices, selectedDices, selected)
		}

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
		switch state {
		case HELP:
			if char != 'q' && (key > 0 || char > 0){
				state = GAME
				continue
			}
		case GAME:
			if (key == keyboard.KeyArrowRight || char == 'd'){
				if selected < len(dices[0])-1 {
					selected++
				}
			} else if (key == keyboard.KeyArrowLeft || char == 'a'){
				if selected > 0 {
					selected--
				}
			} else if key == keyboard.KeySpace{
				if(dices[0] != nil){
					if contains(selectedDices, selected) {
						selectedDices = remove(selectedDices, selected)
					} else {
						selectedDices = append(selectedDices, selected)
					}
				}
			} else if char == 'e'{
				if (len(selectedDices) <= 0) { continue }
				selected = 0
				row := takeDicesAndReroll(&dices, &selectedDices)
				if (row == nil){
					reset(&dices, &selectedDices, &tmpScore, &currentPlayer)
					continue
				}
				score := scoreCount(row)
					if score > 0 {
						tmpScore += score
					}else{
						reset(&dices, &selectedDices, &tmpScore, &currentPlayer)
					}
				if len(dices[0]) == 0{
					reset(&dices, &selectedDices, nil, nil)
				}
			} else if char == 'f'{
				selected = 0
				row := takeDicesAndReroll(&dices, &selectedDices)
				if (row != nil){
					tmpScore += scoreCount(row)
					if tmpScore > 0 { playerScore[currentPlayer] += tmpScore }
				}
				reset(&dices, &selectedDices, &tmpScore, &currentPlayer)
			} else if char == 'h'{
				state = HELP
			}
			if (playerScore[0] >= MAXSCORE || playerScore[1] >= MAXSCORE){ state = WINNER }
		case WINNER:
			if key > 0 || char > 0{
				resetGame(&dices, &selectedDices, &tmpScore, &currentPlayer, &state, &selected, &playerScore)
			}
		}

		if char == 'q' || key == keyboard.KeyEsc {
			break
		}
	}
}
