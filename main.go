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
    "math/rand"
    "net/http"
    "github.com/gorilla/handlers"
    "github.com/gorilla/websocket"
    "fmt"
)
var _ = fmt.Print

type Direction uint8
const (
    NOWHERE Direction = iota
    UP
    DOWN
    LEFT
    RIGHT
)

func AssetServe(w http.ResponseWriter, r *http.Request) {
    if data, err := Asset(r.URL.Path[1:]);
    err == nil {
        switch {
        case r.URL.Path[len(r.URL.Path)-2:] == "js":
            w.Header().Set("Content-Type", "application/json")
        case r.URL.Path[len(r.URL.Path)-3:] == "ttf":
            w.Header().Set("Content-Type", "application/x-font-truetype")
        case r.URL.Path[len(r.URL.Path)-3:] == "css":
            w.Header().Set("Content-Type", "text/css")
        default:
            println(r.URL.Path[len(r.URL.Path)-3:])
        }
        w.Write(data)
    } else {
        w.Write([]byte(r.URL.Path))
    }
}

func IndexServe(w http.ResponseWriter, r *http.Request){
    if data, err := Asset("index.html");
    err == nil {
        w.Write(data)
    } else {
        panic(err)
    }
}

type CommonCommand struct {
    Type string `json:"type"`

    Direction Direction
}

type FullGame struct {
    Type string `json:"type"`
    Map [][]MapPiece `json:"map"`
    Game [][]GamePiece `json:"game"`
    Size int `json:"size"`
    Turns int `json:"turns"`
    Who string `json:"who"`
}

func SendFullGame(ws *websocket.Conn, g *Game) {
    err := ws.WriteJSON(FullGame{
        Type: "FullGame",
        Map: g.Map.board,
        Game: g.Board.board,
        Size: g.Size,
        Turns: g.Turns,
        Who: g.Who.String(),
    })
    if err != nil {
        fmt.Println("write:", err)
    }
}

func main() {

  g := NewWebGame()

    r := http.NewServeMux()
    var upgrader = websocket.Upgrader{}
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
            
        case "/":
            IndexServe(w, r)

        case "/vue.js":
            if data, err := Asset("vue.js");
            err == nil {
                w.Write(data)
            } else {
                panic(err)
            }

        case "/ws":
                c, err := upgrader.Upgrade(w, r, nil)
                if err != nil {
                    fmt.Println("upgrade:", err)
                    return
                }

                g.AddPlayer(c)

        default:
            AssetServe(w, r)
        }
    })
    
    server := &http.Server{
        Addr: ":8080",
        Handler: handlers.CompressHandler(r),
    }
    
	server.ListenAndServe()
}

func RandDirection() Direction {
    switch rand.Uint32() % 4 {
        case 0:
            return DOWN
        case 1:
            return UP
        case 2:
            return LEFT
        case 3:
            return RIGHT
    }
    return UP
}
