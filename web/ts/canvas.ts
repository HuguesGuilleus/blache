/// The with or height of one in block unit.
const REGION_SIZE: number = 16 * 32;

/// The type of a tile.
type TileType = UserTileType | WaterTileType;

/// The user type of a tile, used in background.
enum UserTileType {
	bloc = "bloc",
	biome = "biome",
	height = "height",
}

enum WaterTileType {
	water = "water",
}

/// A minecraft structure.
interface Structure {
	/// X coordonate (west-/est+ axis), in block unit, inside the region.
	x: number,
	/// Z coordonare (sud+/north- axis), in block unit, inside the region.
	z: number,
	// The type of the structure.
	name: string,
	// The color to display a square on the map.
	color: string,
}

/// A region coordonate.
class Coordinate {
	/// X coordonate (west-/est+ axis), in region unit.
	public readonly x: number;
	/// Z coordonare (sud+/north- axis), in region unit.
	public readonly z: number;
	/// The string of the coordonate, used to download the tile or structure
	/// list, or check if a region have tile or not. Format: `x.z`
	public readonly s: string;
	/// Parse the string in format `x.z` to create new Coordonate object.
	constructor(s: string) {
		const t: string[] = s.split('.');
		if (t.length != 2) throw `Invalid syntax for coordinates string`;
		this.s = s;
		this.x = parseInt(t[0]);
		this.z = parseInt(t[1]);
	}

	static stringer(x: number, z: number): string {
		return x + "." + z;;
	}
}

/// The canvas where to display the minecraft map. Do not not manage user
/// events.
class Canvas {
	// The canvas where to display tile.
	public readonly canvasElement: HTMLCanvasElement;
	private readonly canvasContext: CanvasRenderingContext2D;

	/// The position of left top window corner in block unit. After change,
	/// call drawAll method.
	public positionX: number = 0;
	public positionZ: number = 0;

	/// The size in pixel of one tile. Use zoomIn or zoomOut method to
	/// change value.
	public size: number = REGION_SIZE;


	private type: TileType = UserTileType.bloc;

	/// To display or not frontiers, structures and water. After change,
	/// call drawAll method.
	public enabledFrontier: boolean = true;
	public enabledStructure: boolean = true;
	public enabledWater: boolean = true;

	/// The list of existing regions, fetched from regions.json.
	private existingRegions = new Set<Coordinate>();

	/// Regions tiles.
	private readonly regions = {
		[UserTileType.bloc]: new Map<string, HTMLImageElement>(),
		[UserTileType.biome]: new Map<string, HTMLImageElement>(),
		[UserTileType.height]: new Map<string, HTMLImageElement>(),
		[WaterTileType.water]: new Map<string, HTMLImageElement>(),
	};

	/// List of minecraft structure.
	public readonly structures = new Map<string, Structure[]>();

	constructor(canvas: HTMLCanvasElement) {
		const ctx = canvas.getContext('2d');
		if (!ctx) throw "Fail to create 2D draw context";

		this.canvasElement = canvas;
		this.canvasContext = ctx;

		this.blackAll();
		this.fetchRegion();
	}

	// Change the size of the canvas.
	public resize(): void {
		this.canvasElement.width = window.innerWidth;
		this.canvasElement.height = window.innerHeight;
		this.drawAll();
	}

	/// Change the tile (if diffrent), and call darwAll.
	public changeType(type: TileType): void {
		if (type != this.type) {
			this.type = type;
			this.drawAll();
		}
	}
	public zoomIn(w: number, h: number): void {
		if (this.size > REGION_SIZE * 16) return;
		this.positionX = this.positionX * 2 + w;
		this.positionZ = this.positionZ * 2 + h;
		this.size *= 2;
		this.drawAll();
	}
	public zoomOut(w: number, h: number): void {
		if (this.size < REGION_SIZE / 16) return;
		this.positionX = this.positionX / 2 - w / 2;
		this.positionZ = this.positionZ / 2 - h / 2;
		this.size /= 2;
		this.drawAll();
	}

	/// Download regions list then call this.drawAll();
	private fetchRegion(): void {
		fetch('regions.json')
			.then(rep => rep.json())
			.then((regionList: string[]) => {
				this.existingRegions.clear();
				for (const region of regionList) {
					this.existingRegions.add(new Coordinate(region));
				}
				this.drawAll();
			});
	}

	/// Draw all regions, use it after chnage a properties.
	public drawAll(): void {
		this.blackAll();

		const { width, height } = this.canvasElement,
			{ size, positionX, positionZ } = this;

		for (const coord of this.existingRegions.keys()) {
			const { x, z } = coord;
			if (
				(x + 1) * size > positionX
				&& x * size < positionX + width
				&& (z + 1) * size > positionZ
				&& z * size < positionZ + height
			) {
				this.drawRegion(coord);
			}
		}
	}

	/// Reset in black all background.
	private blackAll(): void {
		this.canvasContext.fillStyle = "black";
		this.canvasContext.fillRect(0, 0, this.canvasElement.width, this.canvasElement.height);
		this.canvasContext.imageSmoothingEnabled = false;
	}

	/// Draw one region of the coord. Draw the tile (block, biome or height),
	/// oif need water frontier and structure.
	private drawRegion(coord: Coordinate): void {
		const relativeX = coord.x * this.size - this.positionX,
			relativeZ = coord.z * this.size - this.positionZ;

		this.drawImage(this.getTile(coord, this.type), relativeX, relativeZ);
		if (this.enabledWater) {
			this.drawImage(this.getTile(coord, WaterTileType.water), relativeX, relativeZ);
		}

		if (this.enabledFrontier && this.size > 32) {
			if (this.size > 512) {
				this.canvasContext.strokeStyle = "dodgerblue";
				this.canvasContext.lineWidth = 1;
				this.canvasContext.beginPath();
				for (let i = 1; i < 32; i++) {
					this.canvasContext.moveTo(relativeX + this.size / 32 * i, relativeZ);
					this.canvasContext.lineTo(relativeX + this.size / 32 * i, relativeZ + this.size);
					this.canvasContext.moveTo(relativeX, relativeZ + this.size / 32 * i);
					this.canvasContext.lineTo(relativeX + this.size, relativeZ + this.size / 32 * i);
				}
				this.canvasContext.closePath();
				this.canvasContext.stroke();
			}
			this.canvasContext.strokeStyle = "orange";
			this.canvasContext.lineWidth = 2.5;
			this.canvasContext.beginPath();
			this.canvasContext.moveTo(relativeX, relativeZ);
			this.canvasContext.lineTo(relativeX, relativeZ + this.size);
			this.canvasContext.lineTo(relativeX + this.size, relativeZ + this.size);
			this.canvasContext.lineTo(relativeX + this.size, relativeZ);
			this.canvasContext.closePath();
			this.canvasContext.stroke();
		}

		if (this.enabledStructure) {
			for (let s of this.getStructures(coord)) {
				this.canvasContext.fillStyle = s.color;
				this.canvasContext.fillRect(
					relativeX + s.x * 16 * this.size / REGION_SIZE,
					relativeZ + s.z * 16 * this.size / REGION_SIZE,
					8, 8);
			}
		}
	}

	/// Draw image if complete.
	private drawImage(image: HTMLImageElement, relativeX: number, relativeZ: number): void {
		if (image.complete) {
			this.canvasContext.drawImage(image, relativeX, relativeZ, this.size, this.size);
		}
	}

	/// Get the tile from this.regions[type] or dowload it. Return the image
	/// maybe not complete.
	private getTile(coord: Coordinate, type: TileType): HTMLImageElement {
		let img: HTMLImageElement | undefined = this.regions[type].get(coord.s);

		if (!img) {
			img = new Image();
			this.regions[type].set(coord.s, img);
			img.onload = () => this.drawRegion(coord);
			img.src = type + "/" + coord.s + ".png";
		}

		return img;
	}

	/// Get the structures for the coord, if not already in this.structures,
	/// fetch then call this.drawTile().
	private getStructures(coord: Coordinate): Structure[] {
		const key = coord.s,
			list = this.structures.get(key);
		if (list) {
			return list;
		} else {
			this.structures.set(key, []);
			fetch("structs/" + key + ".json")
				.then(rep => rep.json())
				.then((list: Structure[]) => {
					for (const struct of list) {
						struct.color = "hsl(" +
							struct.name.split('').reduce((s, c) => (s + c.charCodeAt(0)) % 256, 0)
							+ ", 100%, 50%)";
					}
					this.structures.set(key, list);
					this.drawRegion(coord);
				});
			return [];
		}
	}

	/// Download in PNG the canvas.
	public userDownload(): void {
		this.canvasElement.toBlob(b => {
			if (!b) return;

			const url: string = URL.createObjectURL(b),
				a: HTMLAnchorElement = document.createElement('a');

			a.href = url;
			a.download = document.location.hostname + "_" +
				Math.trunc(REGION_SIZE / this.size) + "scale_" +
				Math.floor(this.positionX / this.size) + "_" +
				Math.floor(this.positionZ / this.size)
				+ ".png";

			document.body.appendChild(a);
			a.click();
			a.remove();
			URL.revokeObjectURL(url);
		});
	}

}
