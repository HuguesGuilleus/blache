const REGION_SIZE: number = 16 * 32;

enum TileType {
	bloc = "bloc",
	biome = "biome",
	height = "height",
}

// Zoom change values
enum Zoom {
	in,
	out,
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
		this.moveMouse();
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
		for (let x = -this.size; x < this.canvas.width + this.size; x += this.size) {
			for (let z = -this.size; z < this.canvas.height + this.size; z += this.size) {
				this.drawImage(x, z);
			}
		}
	}
	// Draw one Image
	async drawImage(canvasX: number, canvasZ: number) {
		const it: number = this.iteration;
		const l: number = this.size;
		let dx: number = canvasX - Math.abs(this.posX % l);
		let dz: number = canvasZ - Math.abs(this.posZ % l);
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
	// return the coord of a region.
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
	// Set the zoom of the viewer.
	zoomSet(z: Zoom) {
		switch (z) {
			case Zoom.in:
				this.size = Math.min(this.size * 2, REGION_SIZE * 10);
				break;
			case Zoom.out:
				this.size = Math.max(this.size / 2, REGION_SIZE / 16);
				break;
		}
	}
	// Move from key event
	moveKey(key) {
		const f = {
			'ArrowLeft': () => this.posX -= REGION_SIZE / 2,
			'ArrowRight': () => this.posX += REGION_SIZE / 2,
			'ArrowUp': () => this.posZ -= REGION_SIZE / 2,
			'ArrowDown': () => this.posZ += REGION_SIZE / 2,
			'-': () => this.zoomSet(Zoom.out),
			'+': () => this.zoomSet(Zoom.in),
			'0': (() => {
				this.posX = 0;
				this.posZ = 0;
				this.size = REGION_SIZE;
			}),
			's': () => this.canvas.toBlob(b => userDownload(b,
				`${document.location.hostname}_${
				this.stdCoord(this.posX, 0)}.${
				this.stdCoord(this.posZ, 0)}.png`)),
		}[key];
		if (!f) return;
		f();
		this.drawAll();
	}
	// Set the hanlder on the canvas to set the drag move.
	moveMouse() {
		let lastWheel: Date = new Date();
		this.canvas.addEventListener('wheel', event => {
			if (event.deltaY === 0) return;
			let now: Date = new Date();
			if (now.valueOf() - lastWheel.valueOf() > 100) {
				this.zoomSet(event.deltaY > 0 ? Zoom.out : Zoom.in);
				this.drawAll();
			}
		});
		this.canvas.addEventListener("mousemove", e => this.printCoords(e.x, e.y));
	}
	// Print the current coords into #coords
	printCoords(mx: number, mz: number) {
		let rx = this.stdCoord(this.posX, mx);
		let rz = this.stdCoord(this.posZ, mz);
		let bx = (this.posX + mx) * REGION_SIZE / view.size;
		let bz = (this.posZ + mz) * REGION_SIZE / view.size;
		document.getElementById('coordsRegion').textContent = `(${rx},${rz})`;
		document.getElementById('coordsBloc').textContent = `(${bx},${bz})`;
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
document.getElementById('tileTypeHeight').addEventListener("click", () => view.changeTileType(TileType.height));

window.addEventListener("keydown", event => view.moveKey(event.key));
