var img;
var theta=0;
function setup(){
	img=loadImage("arrow.png");
	createCanvas(500,500);
	frameRate(5);
}
function draw(){
	background(0);
	push();
		translate(width/2,height/2);
		rotate(radians(-theta));
		image(img,0,0);
	pop();
	theta++;
}