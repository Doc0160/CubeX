
var $ =
    (function(){
        "use strict";
        var
        version = "0.01",
        domElement = function(selector) {
            if(!selector){
                return this;
            }
            this.selector = selector || null;
            this.element = null;
            this.init();
        };
        domElement.prototype.init = function() {
            if(typeof this.selector === "string") {
                switch (this.selector[0]) {
                    case '<':
                        var matches = this.selector.match(/<([\w-]*)>/);
                        if (matches === null || matches === undefined) {
                            throw 'Invalid Selector / Node';
                            return false;
                        }
                        var nodeName = matches[0].replace('<', '').replace('>', '');
                        this.element = document.createElement(nodeName);
                        break;
                    default:
                        this.element = document.querySelector(this.selector);
                }
            } else if(this.selector.nodeType) {
                this.element = this.selector
            } else if(typeof(this.selector) === "function") {
                window.onload = this.selector;
            } else {
                throw 'unsuported type';
            }
        };

        // class
        domElement.prototype.addClass = function(c) {
            if(!this.element.classList.contains(c))
                this.element.classList.add(c);
            return this;
        };
        domElement.prototype.removeClass = function(c) {
            if(this.element.classList.contains(c))
                this.element.classList.remove(c);
            return this;
        };
        domElement.prototype.toggleClass = function(c) {
            this.element.classList.toggle(c);
            return this;
        };
        domElement.prototype.noClass = function(c) {
            if(c == undefined) c = "";
            this.element.className = c;
            return this;
        };
        
        domElement.prototype.el = function() {
            return this.element;
        };
        domElement.prototype.on = function(event, callback) {
            var evt = this.eventHandler.bindEvent(event, callback, this.element);
            return this;
        };
        domElement.prototype.off = function(event) {
            var evt = this.eventHandler.unbindEvent(event, this.element);
            return this;
        };
        domElement.prototype.val = function(newVal) {
            return (newVal !== undefined ? this.element.value = newVal : this.element.value);
        };
        domElement.prototype.append = function(html) {
            if(html instanceof domElement) {
                html = html.el();
            }
            if(html.nodeType) {
                this.element.appendChild(html);
            } else {
                this.element.innerHTML = this.element.innerHTML + html;
            }
            return this;
        };
        domElement.prototype.prepend = function(html) {
            if(html instanceof domElement) {
                html = html.el();
            }
            if(html.nodeType) {
                section.insertBefore(html, this.element.firstChild);
            } else {
                this.element.innerHTML = html + this.element.innerHTML;
            }
            return this;
        };
        domElement.prototype.html = function(html) {
            if (html === undefined) {
                return this.element.innerHTML;
            }
            this.element.innerHTML = html;
            return this;
        };
        domElement.prototype.eventHandler = {
            events: [],
            bindEvent: function(event, callback, targetElement) {
                this.unbindEvent(event, targetElement);
                targetElement.addEventListener(event, callback, false);
                this.events.push({
                    type: event,
                    event: callback,
                    target: targetElement
                });
            },
            findEvent: function(event) {
                return this.events.filter(function(evt) {
                    return (evt.type === event);
                }, event)[0];
            },
            unbindEvent: function(event, targetElement) {
                var foundEvent = this.findEvent(event);
                if (foundEvent !== undefined) {
                    targetElement.removeEventListener(event, foundEvent.event, false);
                }
                this.events = this.events.filter(function(evt) {
                    return (evt.type !== event);
                }, event);
            }
        };

        // return
        return function(selector) {
            var el = new domElement(selector);
            return el;
        };
    })();


// websocket json
var WebSocketJSON =
    (function(){
        function WebSocketJSON() {
            this.MessageHandlers = { };
		    this.Socket = null;
        }
        // Will return true if the socket is also in the process of connecting
        WebSocketJSON.prototype.Connected = function(){
		    return this.Socket != null;
	    }
        
	    WebSocketJSON.prototype.AddConnectHandler = function(handler){
		    this.AddMessageHandler("__OnConnect__", handler);
	    }

	    WebSocketJSON.prototype.AddDisconnectHandler = function(handler){
		    this.AddMessageHandler("__OnDisconnect__", handler);
	    }

	    WebSocketJSON.prototype.AddMessageHandler = function(message_name, handler){
		    // Create the message handler array on-demand
		    if (!(message_name in this.MessageHandlers))
			    this.MessageHandlers[message_name] = [ ];
		    this.MessageHandlers[message_name].push(handler);
	    }

        WebSocketJSON.prototype.Connect = function(address){
		    // Disconnect if already connected
		    if (this.Connected())
			    this.Disconnect();

		    console.log(this, "Connecting to " + address);

		    this.Socket = new WebSocket(address);
		    this.Socket.onopen = Bind(OnOpen, this);
		    this.Socket.onmessage = Bind(OnMessage, this);
		    this.Socket.onclose = Bind(OnClose, this);
		    this.Socket.onerror = Bind(OnError, this);
	    }
        
	    WebSocketJSON.prototype.Disconnect = function(){
		    console.log(this, "Disconnecting");
		    if (this.Connected())
			    this.Socket.close();
	    }

	    WebSocketJSON.prototype.Send = function(msg){
            if(msg.type === undefined) {
                msg.type = "undefined";
            }
		    if (this.Connected())
			    this.Socket.send(JSON.stringify(msg));
	    }
        
	    function CallMessageHandlers(self, message_name, data_view){
		    if (message_name in self.MessageHandlers){
			    var handlers = self.MessageHandlers[message_name];
			    for (var i in handlers)
			        handlers[i](self, data_view);
		    }
	    }

	    function OnOpen(self, event){
		    console.log(self, "Connected");
		    CallMessageHandlers(self, "__OnConnect__");
	    }

	    function OnClose(self, event){
		    // Clear all references
		    self.Socket.onopen = null;
		    self.Socket.onmessage = null;
		    self.Socket.onclose = null;
		    self.Socket.onerror = null;
		    self.Socket = null;

		    console.log(self, "Disconnected");
		    CallMessageHandlers(self, "__OnDisconnect__");
	    }

	    function OnError(self, event){
		    console.log(self, "Connection Error ");
	    }

	    function OnMessage(self, event){
            var data_view = JSON.parse(event.data);
            var id = data_view.type;
            CallMessageHandlers(self, id, data_view);
	    }
        
        return WebSocketJSON;
    })();
