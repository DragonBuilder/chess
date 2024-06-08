package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/DragonBuilder/chess/services"
	"github.com/urfave/cli/v2"
)

func main() {

	const name = "chess"

	app := &cli.App{
		Name:  name,
		Usage: "Technologise chess app.",
		Commands: []*cli.Command{
			{
				Name:        "put",
				Usage:       "Put a piece on the board and find out the moves a Pawn, King or Queen chess piece can do when on the board.",
				UsageText:   fmt.Sprintf("%s put <piece> <position>", name),
				Description: "\npiece - The chess piece. Can be Pawn, King or Queen\nposition - The position on the board. Can range from A1-A8 till H1-H8.",
				Args:        true,
				Action: func(ctx *cli.Context) error {
					if ctx.Args().Len() != 2 {
						return fmt.Errorf("expected [<piece> <position>]. got %v", ctx.Args().Slice())
					}

					piece := ctx.Args().Get(0)
					position := ctx.Args().Get(1)

					possibleMoves, err := services.Put(piece, position)
					if err != nil {
						return err
					}

					fmt.Println(strings.Join(possibleMoves, ", "))

					return nil

				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
