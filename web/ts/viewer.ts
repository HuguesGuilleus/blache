const REGION_SIZE: number = 16 * 32;

enum TileType {
	bloc = "bloc",
	biome = "biome",
}

class Viewer {
	canvas;
	ctx: CanvasRenderingContext2D;
	posX: number = 0;
	posZ: number = 0;
	size: number = REGION_SIZE; // the size of one region.
	tileType: TileType;
	iteration: number = 0; // the number of draw
	listRegions: string[] = [];
	constructor(id: string, h: string) {
		this.canvas = document.getElementById(id);
		if (this.canvas === null) throw "id no match";

		this.ctx = this.canvas.getContext('2d');
		if (this.ctx === null) throw "no canvas contex";

		this.initListRegions();

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
		this.iteration++;
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
		let dx: number = canvasX - this.posX % l;
		let dz: number = canvasZ - this.posZ % l;
		try {
			let x: number = this.stdCoord(this.posX, canvasX);
			let z: number = this.stdCoord(this.posZ, canvasZ);
			this.regionsExist(x, z);
			// Draw image
			let img = await download(x, z, this.tileType);
			if (this.iteration !== it) return;
			this.ctx.imageSmoothingEnabled = false;
			this.ctx.drawImage(img, dx, dz, l, l);
			// Grid
			this.ctx.strokeStyle = "red";
			this.ctx.lineWidth = 3.0;
		} catch (error) {
			// Grid
			if (this.iteration !== it) return;
			this.ctx.strokeStyle = "orangered";
			this.ctx.lineWidth = 0.2;
		}
		this.ctx.stroke(new Path2D(`M${dx} ${dz} v${l} h${l} v${-l} z`));
	}
	stdCoord(pos: number, canvas: number): number {
		return Math.floor((pos + canvas) / this.size);
	}
	// Download the regions list.
	async initListRegions() {
		this.listRegions = await (await fetch("regions.json")).json();
		this.drawAll();
	}
	// Test if a region exist with its coord.
	regionsExist(x: number, z: number) {
		if (!this.listRegions.includes(`(${x},${z})`)) {
			throw "region does not exist";
		}
	}
	hashGet(): string {
		return btoa(JSON.stringify({
			x: this.posX,
			z: this.posZ,
			s: this.size,
			t: this.tileType,
		}));
	}
	hashSet(h: string): void {
		let c;
		try {
			c = JSON.parse(atob(h.replace(/^#/, ""))) || {};
		} catch (_) {
			c = {};
		}
		this.tileType = (c.t in TileType) ? c.t : TileType.bloc;
		this.posX = Number(c.x) || 0;
		this.posZ = Number(c.z) || 0;
		this.size = Number(c.s) || REGION_SIZE;
		this.drawAll();
	}
	// Change the tile type
	changeTileType(t: TileType) {
		if (t == this.tileType) return;
		this.tileType = t;
		this.drawAll();
	}
}

// Download one image.
function download(x: number, z: number, t: TileType): Promise<HTMLImageElement> {
	return new Promise<HTMLImageElement>((resolve, reject) => {
		let i = new Image();
		i.src = `${t}/${x}.${z}.png`;
		i.onload = () => resolve(i);
		i.onerror = () => reject(null);
	});
}

const view: Viewer = new Viewer('canvas2d', document.location.hash);

document.getElementById('tileTypeBloc').addEventListener("click", () => view.changeTileType(TileType.bloc));
document.getElementById('tileTypeBiome').addEventListener("click", () => view.changeTileType(TileType.biome));

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
