package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
	out         io.Writer
	game        Game
}

// PlayerPrompt is the text asking the user for the number of players
const PlayerPrompt = "Please enter the number of players: "

// BadPlayerInputErrMsg is the text telling the user they did bad things
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"

// BadWinnerInputMsg is the text telling the user they declated the winner wrong
const BadWinnerInputMsg = "invalid winner input, expect format of 'PlayerName wins'"

// ErrorPlayerNumberPrompt tells the user the entered in the value wrong
const ErrorPlayerNumberPrompt = "ERROR: Please enter the number of players as a number:"

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayers, err := strconv.Atoi(cli.readLine())

	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers, cli.out)

	winnerInput := cli.readLine()
	winner, err := extractWinner(winnerInput)

	if err != nil {
		fmt.Fprint(cli.out, BadWinnerInputMsg)
		return
	}

	cli.game.Finish(winner)
}

func extractWinner(userInput string) (string, error) {
	if !strings.Contains(userInput, " wins") {
		return "", errors.New(BadWinnerInputMsg)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
