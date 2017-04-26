/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Github: Doc0160 $
   $Notice: (C) Copyright 2017 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

type MapPiece rune
const (
    Floor MapPiece = 0
    RedSpawn = 1
    BlueSpawn = 2
    GreenSpawn = 3
    YellowSpawn = 4

    Wall = 254
    Hole = 255
)

type Map struct {
    size int
    board [][]MapPiece
}

func NewMap(_s int) * Map {
    var b Map
    b.size = _s
    for _s := 0; _s < b.size; _s++ {
        b.board = append(b.board, []MapPiece{})
        for __s := 0; __s < b.size; __s++ {
            b.board[_s] = append(b.board[_s], Floor)
        }
    }
    var d int = int(b.size)/2
    var s int = int(b.size)-1
    for v := 0; v < d-1; v++ {
        b.board[v]  [0]   = Hole
        b.board[s-v][0]   = Hole
        b.board[v]  [s]   = Hole
        b.board[s-v][s]   = Hole
        b.board[0]  [v]   = Hole
        b.board[0]  [s-v] = Hole
        b.board[s]  [v]   = Hole
        b.board[s]  [s-v] = Hole
    }
    
    b.board[d][d] = Hole
    b.board[d-1][d] = Wall
    b.board[d+1][d] = Wall
    b.board[d][d-1] = Wall
    b.board[d][d+1] = Wall

    for v := 0; v < 3; v ++ {
        b.board[v+d-1][0] = Wall
        b.board[v+d-1][s] = Wall
        b.board[0][v+d-1] = Wall
        b.board[s][v+d-1] = Wall
    }
    
    b.board[1][1] = RedSpawn
    b.board[1][s-1] = BlueSpawn
    b.board[s-1][s-1] = GreenSpawn
    b.board[s-1][1] = YellowSpawn
    return &b
}

func (b *Map) SafePlace(x, y uint8, p MapPiece) {
    if b.board[x][y] == Floor {
        b.board[x][y] = p
        return
    }
    panic("already placed")
}
