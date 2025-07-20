# Introduction
You are playing a CLI game. You are to respond with one of the listed commands. Do not add any additional text to your response. Only the listed commands are possible.

# Game Concepts

## Board
There is a 9x9 game board of cells represented. Here is an example of the board with a filled cell at position [6][0]:
```
[
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [1,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
]
```

## Pieces
You will be given 3 pieces each round to place on the board. Here is an example of a piece that forms an "L" shape:
```
[
  [1,0]
  [1,0]
  [1,1]
]
```

Here is an example of the piece placed on an otherwise empty board at position [0][0]:
```
[
  [1,0,0,0,0,0,0,0,0]
  [1,0,0,0,0,0,0,0,0]
  [1,1,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
]
```

## Completions
When you place a piece, the board will be evaluated for the following completions:
1. Row
2. Column
3. Box

All the cells that are part of a completion will be freed up. All the cells that are not part of a completion will remain on the board.

Here is an example that contains 1 of each completion:
```
[
  [1,1,1,0,0,0,0,0,0]
  [1,1,1,0,0,0,0,0,0]
  [1,1,1,1,1,1,1,1,1]
  [1,0,0,0,0,0,0,0,0]
  [1,0,0,0,0,0,0,0,0]
  [1,0,0,0,0,0,0,0,0]
  [1,0,0,0,0,0,0,0,0]
  [1,0,0,0,0,0,0,0,0]
  [1,0,0,0,0,0,0,0,0]
]
```

### Commands
1. placePiece(int piece, int row, int column)

Example pieces:
```
[
  [1],
  [1,1,1],
  [
    [0,1,0]
    [0,1,0]
    [1,1,1]
  ]
]
```

Example board:
```
[
  [0,1,0,0,0,0,0,0,0]
  [0,1,0,0,0,0,0,0,0]
  [0,1,0,0,0,0,0,0,0]
  [0,1,0,0,0,0,0,0,0]
  [0,1,0,0,0,0,0,0,0]
  [0,1,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
]
```

So `placePiece(2,6,0)` would result in:
```
[
  [0,1,0,0,0,0,0,0,0]
  [0,1,0,0,0,0,0,0,0]
  [0,1,0,0,0,0,0,0,0]
  [0,1,0,0,0,0,0,0,0]
  [0,1,0,0,0,0,0,0,0]
  [0,1,0,0,0,0,0,0,0]
  [0,1,0,0,0,0,0,0,0]
  [0,1,0,0,0,0,0,0,0]
  [1,1,1,0,0,0,0,0,0] 
]
```

After evaluating for completions:
```
[
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [0,0,0,0,0,0,0,0,0]
  [1,0,1,0,0,0,0,0,0] 
]
```

# Objective
Stay alive for as many rounds as possible. The game ends when none of the available pieces can be placed.
