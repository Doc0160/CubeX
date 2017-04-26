/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Github: Doc0160 $
   $Notice: (C) Copyright 2017 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

import (
    "github.com/gorilla/websocket"
    "fmt"
)

var _ = fmt.Println

type WebPlayer struct {
    WebGame *WebGame
    Color GamePiece
    Websocket *websocket.Conn
}

func NewWebPlayer(wg *WebGame, ws *websocket.Conn, color GamePiece) *WebPlayer {
    p := WebPlayer{}
    p.WebGame = wg
    p.Websocket = ws
    p.Color = color
    return &p
}

func (p *WebPlayer) Write(i interface{}) {
    p.Websocket.WriteJSON(i)
}

func (p *WebPlayer) ReadLoop() {
    var m CommonCommand
    for {
        
        err := p.Websocket.ReadJSON(&m)
        if err != nil {
            p.WebGame.RemovePlayer(p.Websocket)
            return
        }
        
        fmt.Println(m)
        switch m.Type {
        case "connection":
            SendFullGame(p.Websocket, p.WebGame.Game)
            
        case "direction":
            if p.WebGame.Game.Who == p.Color {
                p.WebGame.Game.MovePieces(m.Direction)
                p.WebGame.Game.NextTurn()
                if p.WebGame.Game.Turns % 5 == 0 {
                    p.WebGame.Game.SpawnPiece(p.Color)
                }
                for p.WebGame.Game.Who != p.Color &&
                    p.WebGame.Players[p.WebGame.Game.Who-1] == nil {
                    p.WebGame.Game.MovePieces(RandDirection())
                    p.WebGame.Game.NextTurn()
                    if p.WebGame.Game.Turns % 5 == 0 {
                        p.WebGame.Game.SpawnPiece(p.WebGame.Game.Who)
                    }
                }
                SendFullGame(p.Websocket, p.WebGame.Game)
            }
            
        default:
            fmt.Println(m)
        }
    }
}

type WebGame struct {
    Game *Game
    Players [4]*WebPlayer
    
    Register chan *WebPlayer
    Unregister chan *WebPlayer

    Broadcast chan FullGame
}

func NewWebGame() *WebGame {
    g := WebGame{}
    g.Game = NewGame(11)
    g.Game.SpawnPiece(RedPiece)
    g.Game.SpawnPiece(BluePiece)
    g.Game.SpawnPiece(GreenPiece)
    g.Game.SpawnPiece(YellowPiece)
    return &g
}

func (g *WebGame) AddPlayer(ws *websocket.Conn) GamePiece {
    for k, v := range g.Players {
        if v == nil {
            g.Players[k] = NewWebPlayer(g, ws, GamePiece(k+1))
            g.Players[k].ReadLoop()
            return GamePiece(k+1)
        }
    }
    ws.Close()
    return Nothing
}

func (g *WebGame) RemovePlayer(p *websocket.Conn) {
    for k, v := range g.Players {
        if v != nil && v.Websocket == p {
            g.Players[k] = nil
        }
    }
}

func (g *WebGame) SendBroadcast(s FullGame){
    for k, v := range g.Players {
        if v != nil {
            //g.Players[k].Websocket.WriteJSON(&s)
            g.Players[k].Write(&s)
        }
    }
}

/*func (g *WebGame) Loop() {
    var m CommonCommand
    for {
        for k, v := range g.Players {
            if v != nil {
                err := g.Players[k].Websocket.ReadJSON(&m)
                if err != nil {
                    fmt.Println(k, err)
                    g.RemovePlayer(v.Websocket)
                }
                fmt.Println(k, m)
            }
        }
    }
}*/
