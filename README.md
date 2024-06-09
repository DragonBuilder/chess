## Features
1. Be able to create a chessboard. A board consists of an 8x8 grid of squares.
2. Be able to put a chess piece on the board.
3. Be able to get all possible moves for a piece put in a specific position on the board.

## Installation

### Linux

1. Download the binary `chess` from [here](https://github.com/DragonBuilder/chess/releases/tag/v0.1) 
2. Run `./chess --help` to know how run the executable.

## Building from source
1. Clone the repo.
2. run `go build -o chess ./...
3. This will create the executable named `chess`.
4. Run `./chess --help` to know how run the executable.

## Running the app
The CLL supports a `put` command that takes prints out the possible moves of a chess piece when put on a board. It takes 2 arguments :- 
   1. Name of the chess piece
   2. The position on the board

### Example usages := 
 
```bash
./chess put Pawn A1
A2
```
 
```bash
./chess put King C5
C6, D6, D5, D4, C4, B4, B5, B6
```