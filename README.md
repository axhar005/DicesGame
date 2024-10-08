
# DicesGame

Welcome to **DicesGame**, a fun and strategic dice game where players compete to accumulate points by rolling dice and making smart decisions.

## Table of Contents

- [How to Play](#how-to-play)
- [Scoring Rules](#scoring-rules)
- [Controls](#controls)
- [Installation](#installation)
- [Getting Started](#getting-started)
- [Contributing](#contributing)
- [License](#license)

## How to Play

The objective of DicesGame is to roll the dice and select them to maximize your score. You can reroll non-selected dice to try and improve your score, or end your turn when you are satisfied with the points you've earned.

Only the numbers 1 and 5 have direct point values. For other numbers, you need to roll at least three of a kind to earn points. If all dice in a roll score points, you can reroll all the dice again for a chance to earn even more points. The game ends when the player decides to quit or achieves a winning score.

### Actions:

- **Select Dice**: Use the spacebar to select or deselect a die.
- **Add Dice + Reroll**: Press **E** to add the selected dice to your score and reroll the remaining dice.
- **Add Dice + End Turn**: Press **F** to add the selected dice to your score and end your turn.
- **Show Help**: Press **H** to display the help menu. Press **H** again to close it.
- **Quit the Game**: Press **Q** to quit the game.

## Scoring Rules

Points are earned based on the number rolled and specific combinations of dice:

- **Single Die Rolls**:
  - Rolling a **1** gives **100 points**.
  - Rolling a **5** gives **50 points**.

- **Three of a Kind**:
  - If you roll three of the same number, you earn the dice number × 100 points.
  - Example:
	```
	+---+ +---+ +---+
	| 4 | | 4 | | 4 |
	+---+ +---+ +---+
	Score: 4 × 100 = 400 points
	```

- **Three ones (special case)**:
  - Rolling three ones gives **1000 points**.
	```
	+---+ +---+ +---+
	| 1 | | 1 | | 1 |
	+---+ +---+ +---+
	Score: 1000 points
	```

- **Four or More of a Kind**:
  - For each additional matching die (after three of a kind), multiply the score by 2.
  - Example with four **4s**:
	```
	+---+ +---+ +---+ +---+
	| 4 | | 4 | | 4 | | 4 |
	+---+ +---+ +---+ +---+
	Score: (4 × 100) × 2 = 800 points
	```

## Controls

Here is a quick summary of the controls for the game:

| Key   | Action                       |
|-------|------------------------------|
| Space | Select or deselect a die      |
| E     | Add selected dice + reroll    |
| F     | Add selected dice + end turn  |
| H     | Show or hide the help menu    |
| Q     | Quit the game                 |

## Installation

To play DicesGame, follow these simple steps:

1. Install **Golang** and its dependencies. Make sure Go is installed by running:
   ```bash
   go version
   ```

2. Install the necessary Go dependencies for the game:
   ```bash
   go get -u ./...
   ```

3. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/yourusername/DicesGame.git
   ```

4. Navigate to the game directory:
   ```bash
   cd DicesGame
   ```

5. Compile or run the game:
   ```bash
   go run .
   ```

## Getting Started

Once you have installed the game, start it and follow the on-screen instructions to begin playing. Try to accumulate as many points as possible by rolling and selecting dice according to the scoring rules.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
