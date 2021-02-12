const REGION_SIZE: number = 16 * 32;
const blackImage = function(): ImageData {
	const raw = new Uint8ClampedArray(REGION_SIZE * REGION_SIZE * 4);
	for (let i: number = 0; i < raw.length; i += 4) {
		raw[i + 3] = 255;
	}
	return new ImageData(raw, REGION_SIZE, REGION_SIZE);
}();

enum TileType {
	bloc = 'bloc',
	biome = 'biome',
	height = 'height',
}

interface Struct {
	x: number,
	z: number,
	name: string,
	color: string,
}

type Img = ImageData | HTMLImageElement;

class Viewer {
	canvas_el: HTMLCanvasElement;
	canvas_ctx: CanvasRenderingContext2D;
	coordsRegion: HTMLElement;
	coordsBloc: HTMLElement;
	coordsStruct: HTMLElement;
	regions = {
		[TileType.bloc]: new Map<string, Img | null>(),
		[TileType.biome]: new Map<string, Img | null>(),
		[TileType.height]: new Map<string, Img | null>(),
	};
	struct = new Map<string, Struct[] | undefined>();
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
		this.coordsStruct = $('coordsStruct');
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
					break;
				case '+':
					this.zoomIn(this.canvas_el.width / 2, this.canvas_el.height / 2);
					break;
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
			this.coordsBloc.innerText = '?';
			this.coordsRegion.innerText = '?';
			this.drawAll();
		});
		this.canvas_el.addEventListener('mousemove', event => {
			const x = Math.trunc((this.posX + event.x) * REGION_SIZE / this.size),
				z = Math.trunc((this.posZ + event.y) * REGION_SIZE / this.size),
				rx = Math.floor(x / REGION_SIZE),
				rz = Math.floor(z / REGION_SIZE),
				px = x - rx * REGION_SIZE,
				pz = z - rz * REGION_SIZE;

			this.coordsBloc.innerText = `${x}, ${z}`;
			this.coordsRegion.innerText = `${rx}, ${rz}`;
			this.coordsStruct.innerText = '';
			for (let s of (this.struct.get(Coordinate.toString(rx, rz)) ?? [])) {
				if (Math.abs(px - s.x * 16) < 10 && Math.abs(pz - s.z * 16) < 10) {
					this.coordsStruct.innerText = s.name;
					break;
				}
			}

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
			this.drawAll();
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
	}
	private zoomOut(w: number, h: number) {
		if (this.size < REGION_SIZE / 16) return;
		this.posX = this.posX / 2 - w / 2;
		this.posZ = this.posZ / 2 - h / 2;
		this.size /= 2;
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
		for (let k of this.regions[this.type].keys()) {
			const [x, z] = Coordinate.parse(k);
			if (
				(x + 1) * S > X
				&& x * S < X + W
				&& (z + 1) * S > Z
				&& z * S < Z + H
			) {
				this.getTile(k, x, z, T)
			}
		}
	}

	// Get the region tile and the region's structures, then draw all this info.
	private getTile(k: string, x: number, z: number, t: TileType) {
		this.drawTile(x, z, t,
			this.getRegion(k, x, z, t),
			this.getStructures(k, x, z),
		)
	}

	// Return an image for the region. If the region is'nt yet fetch, fetch it
	// and call View.getTile() when the image is loaded.
	private getRegion(k: string, x: number, z: number, t: TileType): Img {
		const img = this.regions[t].get(k);
		if (img != null) return img;

		this.regions[t].set(k, blackImage);
		const i = new Image(REGION_SIZE, REGION_SIZE);
		i.onload = () => {
			this.regions[t].set(k, i);
			this.getTile(k, x, z, t);
		};
		i.src = `${t}/${x}.${z}.png`;

		return blackImage;
	}

	// Return the list of structure, if the list is not yet fetch, fetch it
	// then call View.getTile().
	private getStructures(k: string, x: number, z: number): Struct[] {
		const l = this.struct.get(k);
		if (Array.isArray(l)) return l;

		this.struct.set(k, []);
		fetch(`structs/${x}.${z}.json`)
			.then(rep => rep.json())
			.then((l: Struct[]) => {
				this.struct.set(k, l.map(s => {
					s.color = `hsl(
						${s.name.split('').reduce((s, c) => (s + c.charCodeAt(0)) % 256, 0)
						}, 100%, 50%)`;
					return s;
				}));
				this.getTile(k, x, z, this.type);
			});

		return [];
	}

	// Draw the region image into the screen.
	private drawTile(xa: number, za: number, t: TileType, surface: Img, struct: Struct[]) {
		if (t != this.type) return;
		const S = this.size,
			xr = xa * S - this.posX,
			zr = za * S - this.posZ;

		if (surface instanceof HTMLImageElement) {
			this.canvas_ctx.imageSmoothingEnabled = false;
			this.canvas_ctx.drawImage(surface, xr, zr, S, S);
		} else {
			this.canvas_ctx.putImageData(surface, xr, zr, 0, 0, S, S);
		}

		this.canvas_ctx.strokeStyle = 'orange';
		this.canvas_ctx.lineWidth = 2.5;
		this.canvas_ctx.stroke(new Path2D(`M${xr} ${zr} v${S} h${S} v${-S} z`));

		for (let s of struct) {
			this.canvas_ctx.fillStyle = s.color;
			this.canvas_ctx.fillRect(
				xr + s.x * 16 * S / REGION_SIZE,
				zr + s.z * 16 * S / REGION_SIZE,
				8, 8);
		}
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
