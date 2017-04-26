/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Github: Doc0160 $
   $Notice: (C) Copyright 2017 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

import "sync"

type Game struct {
    sync.Mutex
    Size int
    Map *Map
    Board *Board
    TurnsToSpawn int
    Turns int
    Who GamePiece
}

func NewGame(size int) *Game {
    var g Game
    g.Size = size
    g.Map = NewMap(size)
    g.Board = NewBoard(size)
    g.TurnsToSpawn = 1
    g.Who = RedPiece
    return &g
}

func (g *Game) NextTurn() {
    g.Lock()
    g.Who++
    for g.howMany(g.Who) == 0 {
        if g.Who > YellowPiece {
            g.Who = RedPiece
            g.Turns++
        } else {
            g.Who ++
        }
    }
    g.Unlock()
}

func (g* Game) WhoHasWon() GamePiece {
    var result GamePiece = Nothing
    g.Lock()
    for x := 0; x < g.Size; x++ {
        for y := 0; y < g.Size; y++ {
            if g.Board.board[x][y] != Nothing {
                if result == Nothing {
                    result = g.Board.board[x][y]
                }
                if result != g.Board.board[x][y] {
                    return Nothing
                }
            }
        }
    }
    g.Unlock()
    return result
}

func (g *Game) howMany(p GamePiece) int {
    var result = 0
    for x := 0; x < g.Size; x++ {
        for y := 0; y < g.Size; y++ {
            if g.Board.board[x][y] == p {
                result++
            }
        }
    }
    return result
}


func (g *Game) SpawnPiece(p GamePiece) {
    g.Lock()
    var s MapPiece
    switch p {
    case RedPiece:
        s = RedSpawn
    case BluePiece:
        s = BlueSpawn
    case YellowPiece:
        s = YellowSpawn
    case GreenPiece:
        s = GreenSpawn
    }
    for x := 0; x < g.Size; x ++ {
        for y := 0; y < g.Size; y ++ {
            if g.Map.board[x][y] == s {
                g.Board.board[x][y] = p
            }
        }
    }
    g.Unlock()
}

func (g *Game) MovePieces(d Direction) {
    g.Lock()
    var ngb *Board = g.Board.Clone()
    //dx, dy := 0, 0
    
    switch {

    case d == UP:
        //dx, dy = -1, 0
        for y := 0; y < g.Size; y++ {
            for x := g.Size-1; x > 0; x-- {
                if g.Board.board[x][y] == g.Who {
                    println(x, y, "/", g.Who)
                    var p = g.Board.board[x][y]
                    g.Board.board[x][y] = Nothing
                    x--
                    for ; x >= 0; x-- {
                        switch {
                            case g.Board.board[x][y] == RedPiece ||
                                g.Board.board[x][y] == BluePiece ||
                                g.Board.board[x][y] == GreenPiece ||
                                g.Board.board[x][y] == YellowPiece:
                            p, g.Board.board[x][y] = g.Board.board[x][y], p
                            
                            case g.Map.board[x][y] == Floor ||
                                g.Map.board[x][y] == RedSpawn ||
                                g.Map.board[x][y] == BlueSpawn ||
                                g.Map.board[x][y] == GreenSpawn ||
                                g.Map.board[x][y] == YellowSpawn:
                            g.Board.board[x][y] = p
                            p = Nothing
                            ngb = g.Board.Clone()

                        case g.Map.board[x][y] == Wall:
                            p = Nothing
                            g.Board = ngb.Clone()
                            
                        case g.Map.board[x][y] == Hole:
                            p = Nothing
                        }
                    }
                }
            }
            ngb = g.Board.Clone()
        }

    case d == DOWN:
        //dx, dy = 1, 0
        for y := 0; y < g.Size; y++ {
            for x := 0; x < g.Size; x++ {
                if g.Board.board[x][y] == g.Who {
                    println(x, y, "/", g.Who)
                    var p = g.Board.board[x][y]
                    g.Board.board[x][y] = Nothing
                    x++
                    for ; x < g.Size; x++ {
                        switch {
                            case g.Board.board[x][y] == RedPiece ||
                                g.Board.board[x][y] == BluePiece ||
                                g.Board.board[x][y] == GreenPiece ||
                                g.Board.board[x][y] == YellowPiece:
                            p, g.Board.board[x][y] = g.Board.board[x][y], p
                            
                            case g.Map.board[x][y] == Floor ||
                                g.Map.board[x][y] == RedSpawn ||
                                g.Map.board[x][y] == BlueSpawn ||
                                g.Map.board[x][y] == GreenSpawn ||
                                g.Map.board[x][y] == YellowSpawn:
                            g.Board.board[x][y] = p
                            p = Nothing
                            ngb = g.Board.Clone()
                            
                        case g.Map.board[x][y] == Wall:
                            g.Board = ngb.Clone()
                            
                        case g.Map.board[x][y] == Hole:
                            p = Nothing
                        }
                    }
                }
            }
            ngb = g.Board.Clone()
        }

    case d == RIGHT:
        //dx, dy = 0, 1
        for x := 0; x < g.Size; x++ {
            for y := 0; y < g.Size; y++ {
                if g.Board.board[x][y] == g.Who {
                    println(x, y, "/", g.Who)
                    var p = g.Board.board[x][y]
                    g.Board.board[x][y] = Nothing
                    y++
                    for ; y < g.Size; y++ {
                        switch {
                            case g.Board.board[x][y] == RedPiece ||
                                g.Board.board[x][y] == BluePiece ||
                                g.Board.board[x][y] == GreenPiece ||
                                g.Board.board[x][y] == YellowPiece:
                            p, g.Board.board[x][y] = g.Board.board[x][y], p
                            
                            case g.Map.board[x][y] == Floor ||
                                g.Map.board[x][y] == RedSpawn ||
                                g.Map.board[x][y] == BlueSpawn ||
                                g.Map.board[x][y] == GreenSpawn ||
                                g.Map.board[x][y] == YellowSpawn:
                            g.Board.board[x][y] = p
                            p = Nothing
                            ngb = g.Board.Clone()
                            
                        case g.Map.board[x][y] == Wall:
                            g.Board = ngb.Clone()
                            
                        case g.Map.board[x][y] == Hole:
                            p = Nothing
                        }
                    }
                }
            }
            ngb = g.Board.Clone()
        }
        
    case d == LEFT:
        //dx, dy = 0, -1
        for x := 0; x < g.Size; x++ {
            for y := g.Size-1; y > 0; y-- {
                if g.Board.board[x][y] == g.Who {
                    println(x, y, "/", g.Who)
                    var p = g.Board.board[x][y]
                    g.Board.board[x][y] = Nothing
                    y--
                    for ; y >= 0; y-- {
                        switch {
                            case g.Board.board[x][y] == RedPiece ||
                                g.Board.board[x][y] == BluePiece ||
                                g.Board.board[x][y] == GreenPiece ||
                                g.Board.board[x][y] == YellowPiece:
                            p, g.Board.board[x][y] = g.Board.board[x][y], p
                            
                            case g.Map.board[x][y] == Floor ||
                                g.Map.board[x][y] == RedSpawn ||
                                g.Map.board[x][y] == BlueSpawn ||
                                g.Map.board[x][y] == GreenSpawn ||
                                g.Map.board[x][y] == YellowSpawn:
                            g.Board.board[x][y] = p
                            p = Nothing
                            ngb = g.Board.Clone()

                        case g.Map.board[x][y] == Wall:
                            g.Board = ngb.Clone()
                            
                        case g.Map.board[x][y] == Hole:
                            p = Nothing
                        }
                    }
                }
            }
            ngb = g.Board.Clone()
        }
        
    }
    //var _ = dx
    //var _ = dy
    var _ = ngb
    /*
    for x := 0; x < int(g.Size); x++ {
        for y := 0; y < int(g.Size); y++ {
            if ngb.board[x][y] == g.Who {
                g.safeMove(x, y, x+dx, y+dy)
            }
        }
    }
    //*/
    g.Unlock()
}

func (g*Game) safeMove(ox,oy, x,y int) (b bool) {
    if x >= 0 && x < g.Size {
        if y >= 0 && y < g.Size {
            switch {
                
                case g.Map.board[x][y] == Floor ||
                    g.Map.board[x][y] == RedSpawn ||
                    g.Map.board[x][y] == BlueSpawn ||
                    g.Map.board[x][y] == GreenSpawn ||
                    g.Map.board[x][y] == YellowSpawn:
                if g.Board.board[x][y] != Nothing {
                    if !g.safeMove(x,y,
                        int(x)-(int(ox)-int(x)),
                        int(y)-(int(oy)-int(y)) ) {
                            return false
                        }
                }
                var p = g.Board.board[ox][oy]
                g.Board.board[ox][oy] = Nothing
                g.Board.board[x][y] = p
                return true
                
            case g.Map.board[x][y] == Hole:
                g.Board.board[ox][oy] = Nothing
                return true
                
            }
        }
    }
    return false
}
