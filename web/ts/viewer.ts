const REGION_SIZE: number = 16 * 32;
const blackImage = function(): ImageData {
	const raw = new Uint8ClampedArray(REGION_SIZE * REGION_SIZE * 4);
	for (let i: number = 0; i < raw.length; i += 4) {
		raw[i] = raw[i + 0] = raw[i + 1] = raw[i + 2] = 0;
		raw[i + 3] = 255;
	}
	return new ImageData(raw, REGION_SIZE, REGION_SIZE);
}();

enum TileType {
	bloc = 'bloc',
	biome = 'biome',
	height = 'height',
}

type Img = ImageData | HTMLImageElement;

class Viewer {
	canvas_el: HTMLCanvasElement;
	canvas_ctx: CanvasRenderingContext2D;
	coordsRegion: HTMLElement;
	coordsBloc: HTMLElement;
	regions = {
		[TileType.bloc]: new Map(),
		[TileType.biome]: new Map(),
		[TileType.height]: new Map(),
	};
	type: TileType = TileType.bloc;
	posX: number = 0;
	posZ: number = 0;
	size: number = REGION_SIZE;

	constructor() {
		// Use to get some element into the dom.
		function $(id: string): HTMLElement {
			const e = document.getElementById(id);
			if (!e) throw `Element '${id}' no found in DOM`;
			return e;
		}

		// Elements
		this.coordsRegion = $('coordsRegion');
		this.coordsBloc = $('coordsBloc');
		this.canvas_el = <HTMLCanvasElement>$('canvas2d');
		const c = this.canvas_el.getContext('2d');
		if (!c) throw 'Can not create canvas';
		this.canvas_ctx = c;
		$('tileTypeBloc')
			.addEventListener('click', () => this.typeChange(TileType.bloc));
		$('tileTypeBiome')
			.addEventListener('click', () => this.typeChange(TileType.biome));
		$('tileTypeHeight')
			.addEventListener('click', () => this.typeChange(TileType.height));

		// Download regions list
		fetch('regions.json')
			.then(rep => rep.json())
			.then((rs: string[]) => {
				for (let c of rs) {
					this.regions[TileType.bloc].set(c, null);
					this.regions[TileType.biome].set(c, null);
					this.regions[TileType.height].set(c, null);
				}
				this.drawAll();
			});

		window.addEventListener('load', () => this.resize());
		window.addEventListener('resize', () => this.resize());
		window.addEventListener('keydown', e => {
			switch (e.key) {
				case '-':
					this.zoomOut(this.canvas_el.width / 2, this.canvas_el.height / 2);
					return;
				case '+':
					this.zoomIn(this.canvas_el.width / 2, this.canvas_el.height / 2);
					return;
				case 'ArrowLeft':
					this.posX -= this.size / 4;
					break;
				case 'ArrowRight':
					this.posX += this.size / 4;
					break;
				case 'ArrowUp':
					this.posZ -= this.size / 4;
					break;
				case 'ArrowDown':
					this.posZ += this.size / 4;
					break;
				case '0':
					this.posX = this.posZ = 0;
					this.size = REGION_SIZE;
					break;
				case 's':
					this.canvas_el.toBlob(b => userDownload(b ?? new Blob(),
						`${document.location.hostname}_${
						Math.trunc(REGION_SIZE / this.size)}_${
						Math.floor(this.posX / this.size)}.${
						Math.floor(this.posZ / this.size)}.png`));
					return;
				default:
					return;
			}
			this.drawAll();
		});
		this.canvas_el.addEventListener('mousemove', event => {
			const x = Math.trunc((this.posX + event.x) * REGION_SIZE / this.size);
			const z = Math.trunc((this.posZ + event.y) * REGION_SIZE / this.size);
			this.coordsBloc.innerText = `${x}, ${z}`;
			this.coordsRegion.innerText = `${Math.floor(x / REGION_SIZE)}, ${Math.floor(z / REGION_SIZE)}`;
			if (!event.buttons) return;
			this.posX -= event.movementX;
			this.posZ -= event.movementY;
			this.drawAll();
		});
		this.canvas_el.addEventListener('wheel', event => {
			const d = event.deltaY;
			if (d > 0) {
				this.zoomOut(event.x, event.y);
			} else if (d < 0) {
				this.zoomIn(event.x, event.y);
			}
		});
	}

	// Edit the type of tile
	private typeChange(t: TileType) {
		if (t != this.type) {
			this.type = t;
			this.drawAll();
		}
	}
	private zoomIn(w: number, h: number) {
		if (this.size > REGION_SIZE * 16) return;
		this.posX = this.posX * 2 + w;
		this.posZ = this.posZ * 2 + h;
		this.size *= 2;
		this.drawAll();
	}
	private zoomOut(w: number, h: number) {
		if (this.size < REGION_SIZE / 16) return;
		this.posX = this.posX / 2 - w / 2;
		this.posZ = this.posZ / 2 - h / 2;
		this.size /= 2;
		this.drawAll();
	}

	// Change the size of the canvas.
	private resize() {
		this.canvas_el.width = window.innerWidth;
		this.canvas_el.height = window.innerHeight;
		this.drawAll();
	}

	// Draw all regions.
	private drawAll() {
		this.canvas_ctx.fillStyle = 'black';
		this.canvas_ctx.fillRect(0, 0, this.canvas_el.width, this.canvas_el.height);

		const X = this.posX;
		const Z = this.posZ;
		const S = this.size;
		const W = this.canvas_el.width;
		const H = this.canvas_el.height;
		const T = this.type;

		// Draw regions
		for (let cs of this.regions[this.type].keys()) {
			const [x, z] = Coordinate.parse(cs);
			if (
				(x + 1) * S > X
				&& x * S < X + W
				&& (z + 1) * S > Z
				&& z * S < Z + H
			) {
				this.drawImage(x, z, T, this.getOrDownload(x, z, T));
			}
		}
	}

	// Return an image for the region. If the region is'nt yet fetch, fetch it
	// and call View.drawImage in the future.
	private getOrDownload(x: number, z: number, t: TileType): Img {
		const k = Coordinate.toString(x, z);
		const img = this.regions[t].get(k);
		if (img != null) return img;

		this.regions[t].set(k, blackImage);
		const i = new Image(REGION_SIZE, REGION_SIZE);
		i.onload = () => {
			this.regions[t].set(k, i);
			this.drawImage(x, z, t, i);
		};
		i.src = `${t}/${x}.${z}.png`;

		return blackImage;
	}

	// Draw the region image into the screen.
	private drawImage(xa: number, za: number, t: TileType, img: Img) {
		if (t != this.type) return;
		const S = this.size;
		const xr = xa * S - this.posX;
		const zr = za * S - this.posZ;
		if (img instanceof HTMLImageElement) {
			this.canvas_ctx.imageSmoothingEnabled = false;
			this.canvas_ctx.drawImage(img, xr, zr, S, S);
		} else {
			this.canvas_ctx.putImageData(img, xr, zr, 0, 0, S, S);
		}
		this.canvas_ctx.strokeStyle = 'orange';
		this.canvas_ctx.lineWidth = 2.5;
		this.canvas_ctx.stroke(new Path2D(`M${xr} ${zr} v${S} h${S} v${-S} z`));
	}
}

(function() {
	function main() {
		new Viewer();
	}
	document.readyState == 'loading' ? document.addEventListener('DOMContentLoaded', main, {
		once: true
	}) : main();
}());
