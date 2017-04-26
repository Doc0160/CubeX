/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Github: Doc0160 $
   $Notice: (C) Copyright 2017 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

type GamePiece rune
const (
    Nothing GamePiece = 0
    RedPiece = 1
    BluePiece = 2
    GreenPiece = 3
    YellowPiece = 4
    DeadPiece = 255
)

func (p *GamePiece) String() string {
    switch *p {
    case RedPiece:
        return "Red"
    case GreenPiece:
        return "Green"
    case BluePiece:
        return "Blue"
    case YellowPiece:
        return "Yellow"
    default:
        return "kaka"
    }
}

type Board struct {
    size int
    board [][]GamePiece
}

func NewBoard(s int) *Board {
    var gb Board
    gb.size = s
    for s := 0; s < gb.size; s++ {
        gb.board = append(gb.board, []GamePiece{})
        for _s := 0; _s < gb.size; _s++ {
            gb.board[s] = append(gb.board[s], Nothing)
        }
    }
    return &gb
}

func (b* Board) Clone() *Board {
    var ngb *Board = NewBoard(b.size)
    for x := 0; x < b.size; x ++ {
        for y := 0; y < b.size; y ++ {
            ngb.board[x][y] = b.board[x][y]
        }
    }
    return ngb
}
