const REGION_SIZE: number = 16 * 32;

class Viewer {
	canvas;
	ctx: CanvasRenderingContext2D;
	posX: number = 0;
	posZ: number = 0;
	size: number = REGION_SIZE; // the size of one region.
	tileType: string = "biome";
	iteration: number = 0;// the number of draw
	constructor(id: string, h: string) {
		this.canvas = document.getElementById('canvas2d');
		if (this.canvas === null) throw "id no match";

		this.ctx = this.canvas.getContext('2d');
		if (this.ctx === null) throw "no canvas contex";

		this.hashSet(h);
		this.resize();
		window.addEventListener("load", () => this.resize());
		window.addEventListener("resize", () => this.resize());
	}
	resize() {
		this.canvas.width = window.innerWidth;
		this.canvas.height = window.innerHeight;
		this.drawAll();
	}
	// Draw all the image
	drawAll() {
		document.location.hash = this.hashGet();
		this.ctx.fillStyle = "black";
		this.ctx.fillRect(0, 0, this.canvas.width, this.canvas.height);
		for (let x = 0; x < this.canvas.width; x += this.size) {
			for (let z = 0; z < this.canvas.height; z += this.size) {
				this.drawImage(x, z);
			}
		}
	}
	// Draw one Image
	async drawImage(canvasX: number, canvasZ: number) {
		const it: number = this.iteration;
		const l: number = this.size;
		try {
			let x: number = this.stdCoord(this.posX, canvasX);
			let z: number = this.stdCoord(this.posZ, canvasZ);
			let dx: number = canvasX - this.posX % l;
			let dz: number = canvasZ - this.posZ % l;
			// Dessine l'image
			let img = await download(x, z, this.tileType);
			if (this.iteration !== it) return;
			this.ctx.imageSmoothingEnabled = false;
			this.ctx.drawImage(img, dx, dz, l, l);
			// Grille
			this.ctx.strokeStyle = "red";
			this.ctx.lineWidth = 3.0;
			this.ctx.stroke(new Path2D(`M${dx} ${dz} v${l} h${l} v${-l} z`));
		} catch (error) {
			if (this.iteration !== it) return;
			this.ctx.fillStyle = "black";
			this.ctx.fillRect(canvasX, canvasZ, this.size, this.size);
		}
	}
	stdCoord(pos: number, canvas: number): number {
		return Math.floor((pos + canvas) / this.size);
	}
	hashGet(): string {
		return btoa(JSON.stringify({
			x: this.posX,
			z: this.posZ,
			s: this.size,
		}));
	}
	hashSet(h: string): void {
		if (h === "") return;
		let o = JSON.parse(atob(h.slice(1)));
		if ('x' in o) this.posX = o.x;
		if ('z' in o) this.posZ = o.z;
		if ('s' in o) this.size = o.s;
		this.drawAll();
	}
}

// Download one image.
function download(x: number, z: number, v: string): Promise<HTMLImageElement> {
	return new Promise<HTMLImageElement>((resolve, reject) => {
		let i = new Image();
		i.src = `${v}/${x}.${z}.png`;
		i.onload = () => {
			resolve(i);
		};
		i.onerror = () => reject(null);
	});
}

const view: Viewer = new Viewer('canvas2d', document.location.hash);

window.addEventListener("keydown", event => {
	let f = {
		'ArrowLeft': () => view.posX -= view.size / 2,
		'ArrowRight': () => view.posX += view.size / 2,
		'ArrowUp': () => view.posZ -= view.size / 2,
		'ArrowDown': () => view.posZ += view.size / 2,
		'-': () => view.size /= 2,
		'+': () => view.size *= 2,
		'0': (() => {
			view.posX = 0;
			view.posZ = 0;
			view.size = REGION_SIZE;
		}),
	}[event.key];
	if (!f) return;
	f();
	view.drawAll();
});
