
function webgl_support() { 
    try{
        var canvas = document.createElement( 'canvas' ); 
        return !! window.WebGLRenderingContext && ( 
            canvas.getContext( 'webgl' ) || canvas.getContext( 'experimental-webgl' ) );
    }catch( e ) { return false; } 
};

function websocket_support() {
    return window["WebSocket"];
}

var CubeX =
    (function(){
        function CubeX(){
            this.size = -1;
            this.who = "Rainbow";
            this.turns = -1;
            this.game = [[]];
            this.map = [[]];
            this.mouseX = 0;
            this.mouseY = 0;
            $("#board").on("click", Bind(function(self, event){
                event.preventDefault();
                var rect = $("#board").el().getBoundingClientRect();            
                self.mouseX = event.clientX - rect.left - 440/2;
                self.mouseY = event.clientY - rect.top - 440/2;
                if(Math.abs(self.mouseX)>Math.abs(self.mouseY)){
                    if(self.mouseX>0) {
                        self.SendDirection(Directions.RIGHT);
                    } else {
                        self.SendDirection(Directions.LEFT);
                    }
                } else if(Math.abs(self.mouseX)<Math.abs(self.mouseY)) {
                    if(self.mouseY>0) {
                        self.SendDirection(Directions.DOWN);
                    } else {
                        self.SendDirection(Directions.UP);
                    }
                } else {
                    
                }
            }, this));
            if(webgl_support()) {
                this.renderer = new WebGL_Renderer();
            } else {
                this.renderer = new DOM_Renderer();
            }
            $("#board").on("mousemove", Bind(function(self){
                var rect = $("#board").el().getBoundingClientRect();            
                self.mouseX = event.clientX - rect.left - 440/2;
                self.mouseY = event.clientY - rect.top - 440/2;
                self.renderer.Render(self);
            }, this));
            
            this.websocket = new WebSocketJSON();
            this.websocket.AddDisconnectHandler(Bind(function(self){
                $("#message").html("Disconnected").noClass("danger");
            }, this));
            this.websocket.AddConnectHandler(Bind(function(self){
                $("#message").html("Connected").noClass("success");
                setTimeout(function() {
                    self.websocket.Send({
                        type: "connection",
                    });
                }, 0);
            }, this));
            this.websocket.AddMessageHandler("FullGame", Bind(function(self, socket, data_view){
                console.log(data_view);
                self.size = data_view.size;
                self.map = data_view.map;
                self.game = data_view.game;
                self.who = data_view.who;
                self.turns = data_view.turns;
                $("#game").html("" + self.who + ", " + self.turns);
                self.Render();
            }, this));
            this.websocket.Connect("ws://" + document.location.host + "/ws");
            this.Render();
        }

        CubeX.prototype.Render = function(){
            this.renderer.Render(this);
        };

        CubeX.prototype.SendDirection = function(d) {
            if(this.websocket.Connected()) {
                this.websocket.Send({
                    type: "direction",
                    direction: d,
                });
            } else {
                console.log('no websocks');
            }
        };

        return CubeX;
    })();

var t = 0;
var w = 10;
var cr = 0.5;
var wr = 0.8;

var Directions = {
    UP: 1,
    DOWN: 2,
    LEFT: 3,
    RIGHT: 4,
};

window.onload = function () {
    var cubex = new CubeX();
};

