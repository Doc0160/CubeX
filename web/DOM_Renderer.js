
var DOM_Renderer =
    (function(){
        function DOM_Renderer(){
            $("#board").html("");
        }

        DOM_Renderer.prototype.Render = function(cubex){
            $("#board").html("");
            for(var x = 0; x < cubex.size ; x++){
                tmp = '<div class="line">';
                for(var y = 0; y < cubex.size ; y++) {
                    var classes = "cell";
                    var text = cubex.map[x][y];
                    switch(cubex.map[x][y]) {
                        case 0:
                            classes += ' floor';
                            text = "F";
                            break;
                        case 1:
                            classes += ' red_spawn';
                            text = 'S';
                            break;
                        case 2:
                            classes += ' blue_spawn';
                            text = 'S';
                            break;
                        case 3:
                            classes += ' green_spawn';
                            text = 'S';
                            break;
                        case 4:
                            classes += ' yellow_spawn';
                            text = 'S';
                            break;
                        case 254:
                            classes += ' wall';
                            text = 'W';
                            break;
                        case 255:
                            classes += ' hole';
                            text = 'H';
                            break;
                    }
                    switch(cubex.game[x][y]) {
                        case 1:
                            classes += " red_cube"
                            text = '■';
                            break;
                        case 2:
                            classes += " blue_cube"
                            text = '■';
                            break;
                        case 3:
                            classes += " green_cube"
                            text = '■';
                            break;
                        case 4:
                            classes += " yellow_cube"
                            text = '■';
                            break;
                    }
                    tmp += '<div class="'+ classes +'">' + text + '</div>';
                }
                $("#board").append(tmp + '</div>');
            }
            console.log(cubex);
        }
        return DOM_Renderer;
    })();


var WebGL_Renderer =
    (function(){
        function WebGL_Renderer(){
            this.group = new THREE.Group();
            
            this.planeRed = new THREE.MeshBasicMaterial( { color: 0xeFF0000, overdraw: 0.5 } );
            this.planeBlue = new THREE.MeshBasicMaterial( { color: 0xe0000FF, overdraw: 0.5 } );
            this.planeGreen = new THREE.MeshBasicMaterial( { color: 0xe00FF00, overdraw: 0.5 } );
            this.planeYellow = new THREE.MeshBasicMaterial( { color: 0xeFFFF00, overdraw: 0.5 } );

            this.cubeMaterial = new THREE.MeshBasicMaterial( { vertexColors: THREE.FaceColors, overdraw: 0.5 } );
            
            this.red_box = new THREE.BoxGeometry(w*cr, w*cr, w*cr);
            this.blue_box = new THREE.BoxGeometry(w*cr, w*cr, w*cr);
            this.green_box = new THREE.BoxGeometry(w*cr, w*cr, w*cr);
            this.yellow_box = new THREE.BoxGeometry(w*cr, w*cr, w*cr);
            
            this.wall = new THREE.BoxGeometry(w*wr, w*wr, w*wr);
            
            this.floor1 = new THREE.BoxGeometry(w, w, w);
            this.floor2 = new THREE.BoxGeometry(w, w, w);
            
            for ( var i = 0; i < this.red_box.faces.length; i++ ) {
	            this.red_box.faces[i].color.setHex(0xFF0000);
	            this.green_box.faces[i].color.setHex(0x00FF00);
	            this.blue_box.faces[i].color.setHex(0x0000FF);
	            this.yellow_box.faces[i].color.setHex(0xFFFF00);
                
	            this.wall.faces[i].color.setHex(0xFFFFFF);
                
	            this.floor1.faces[i].color.setHex(0x1ABC9C);
	            this.floor2.faces[i].color.setHex(0x3498DB);
            }

            this.camera = new THREE.PerspectiveCamera(70, 1, 1, 1000);
	        this.camera.rotation.z = Math.PI/2;
	        this.camera.position.z = 100;
            
	        this.scene = new THREE.Scene();
            
	        this.renderer = new THREE.WebGLRenderer();
	        this.renderer.setClearColor(0xADD8E6);
	        this.renderer.setPixelRatio(window.devicePixelRatio);
	        this.renderer.setSize(438, 438);
	        this.renderer.sortObjects = false;
            $("#board").html("").el().append(this.renderer.domElement);
        }

        WebGL_Renderer.prototype.Floor = function(x, y) {
            var cube;
            if ((x+y) % 2 == 0) {
                var cube = new THREE.Mesh(this.floor1, this.cubeMaterial);
            } else {
                var cube = new THREE.Mesh(this.floor2, this.cubeMaterial);
            }
            cube.position.x = w*x - (11*w/2 - w/2);
            cube.position.y = w*y - (11*w/2 - w/2);
            cube.position.z = -w;
            return cube
        }
        WebGL_Renderer.prototype.Cube = function(color, x, y) {
            var cube = new THREE.Mesh(color, this.cubeMaterial);
            cube.position.x = w*x - (11*w/2 - w/2);
            cube.position.y = w*y - (11*w/2 - w/2);
            return cube
        }
        WebGL_Renderer.prototype.Plane = function(color, x, y) {
            var geometry = new THREE.PlaneBufferGeometry(w*cr, w*cr);
            var plane = new THREE.Mesh(geometry, color);
            plane.position.x = w*x - (11*w/2 - w/2);
            plane.position.y = w*y - (11*w/2 - w/2);
            plane.position.z = 1 - w/2;
            return plane;
        }
        WebGL_Renderer.prototype.Render = function(cubex){
            this.scene.remove(this.group);
            this.group = new THREE.Group();
            for (var x = 0; x < cubex.map.length; x++) {
                for (var y = 0; y < cubex.map[x].length; y++) {
                    switch(cubex.map[x][y]){
                        case 255:
                            break;
                        case 254:
                            this.group.add(this.Cube(this.wall, x, y));
                            this.group.add(this.Floor(x, y));
                            break;
                        case 1:
                            this.group.add(this.Plane(this.planeRed, x, y));
                            this.group.add(this.Floor(x, y));
                            break;
                        case 2:
                            this.group.add(this.Plane(this.planeBlue, x, y));
                            this.group.add(this.Floor(x, y));
                            break;
                        case 3:
                            this.group.add(this.Plane(this.planeGreen, x, y));
                            this.group.add(this.Floor(x, y));
                            break;
                        case 4:
                            this.group.add(this.Plane(this.planeYellow, x, y));
                            this.group.add(this.Floor(x, y));
                            break;
                        default:
                            this.group.add(this.Floor(x, y));
                            break;
                    }
                    switch(cubex.game[x][y]){
                        case 1:
                            this.group.add(this.Cube(this.red_box, x, y));
                            break;
                        case 2:
                            this.group.add(this.Cube(this.blue_box, x, y));
                            break;
                        case 3:
                            this.group.add(this.Cube(this.green_box, x, y));
                            break;
                        case 4:
                            this.group.add(this.Cube(this.yellow_box, x, y));
                            break;
                        default:
                            break;
                    }
                }
            }
            this.scene.add(this.group);
            this.camera.rotation.x = cubex.mouseX/9600;
            this.camera.rotation.y = -cubex.mouseY/9600;
            this.renderer.render(this.scene, this.camera);
        }
        return WebGL_Renderer;
    })();


