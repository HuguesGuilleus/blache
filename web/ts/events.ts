// Use to get some element into the dom.
function $(id: string): HTMLElement {
	const e = document.getElementById(id);
	if (!e) throw `Element '${id}' no found in DOM`;
	return e;
}

// Create the canvas and add user event handler.
function main() {
	const canvas = new Canvas(<HTMLCanvasElement>$("canvas2d")),
		coordsBloc = $("coordsBloc"),
		coordsRegion = $("coordsRegion"),
		coordsStruct = $("coordsStruct");

	$("tileTypeBloc").addEventListener("click", () =>
		canvas.changeType(UserTileType.bloc)
	);
	$("tileTypeBiome").addEventListener("click", () =>
		canvas.changeType(UserTileType.biome)
	);
	$("tileTypeHeight").addEventListener("click", () =>
		canvas.changeType(UserTileType.height)
	);
	$("savePNG").addEventListener("click", () => canvas.userDownload());

	type enbaleProperties = "enabledFrontier" | "enabledStructure" | "enabledWater";
	function button(prop: enbaleProperties): () => void {
		const input = <HTMLInputElement>$(prop);
		canvas[prop] = input.checked;
		input.addEventListener("change", () => {
			canvas[prop] = input.checked;
			canvas.drawAll();
		});
		return () => {
			input.checked = canvas[prop] = !canvas[prop];
			canvas.drawAll();
		};
	}
	const switchFrontier = button("enabledFrontier"),
		switchStructure = button("enabledStructure"),
		switchWater = button("enabledWater");

	window.addEventListener("keydown", e => {
		switch (e.key) {
			case "f": return switchFrontier();
			case "w": return switchWater();
			case "x": return switchStructure();
			case "s": return canvas.userDownload();

			case "-": return canvas.zoomOut(
				canvas.canvasElement.width / 2,
				canvas.canvasElement.height / 2
			);
			case "+": return canvas.zoomIn(
				canvas.canvasElement.width / 2,
				canvas.canvasElement.height / 2
			);
			case "ArrowLeft":
				canvas.positionX -= canvas.size / 4;
				break;
			case "ArrowRight":
				canvas.positionX += canvas.size / 4;
				break;
			case "ArrowUp":
				canvas.positionZ -= canvas.size / 4;
				break;
			case "ArrowDown":
				canvas.positionZ += canvas.size / 4;
				break;
			case "0":
				canvas.positionX = canvas.positionZ = 0;
				canvas.size = REGION_SIZE;
				break;

			default: return;
		}

		coordsBloc.innerText = "";
		coordsRegion.innerText = "";
		coordsStruct.innerText = "";
		canvas.drawAll();
	});

	canvas.canvasElement.addEventListener("mousemove", event => {
		const x = Math.trunc((canvas.positionX + event.x) * REGION_SIZE / canvas.size),
			z = Math.trunc((canvas.positionZ + event.y) * REGION_SIZE / canvas.size),
			regionCoordX = Math.floor(x / REGION_SIZE),
			regionCoordZ = Math.floor(z / REGION_SIZE),
			insideRegionX = x - regionCoordX * REGION_SIZE, // we do not use modulo because negatives values.
			insideRegionZ = z - regionCoordZ * REGION_SIZE;

		coordsBloc.innerText = `${x}, ${z}`;
		coordsRegion.innerText = `${regionCoordX}, ${regionCoordZ}`;
		coordsStruct.innerText = '';
		for (const s of (canvas.structures.get(Coordinate.stringer(regionCoordX, regionCoordZ)) ?? [])) {
			if (Math.abs(insideRegionX - s.x * 16) < 10 && Math.abs(insideRegionZ - s.z * 16) < 10) {
				coordsStruct.innerText = s.name;
				break;
			}
		}

		if (!event.buttons) return;
		canvas.positionX -= event.movementX;
		canvas.positionZ -= event.movementY;
		canvas.drawAll();
	});

	canvas.canvasElement.addEventListener("wheel", event => {
		const d = event.deltaY;
		if (d > 0) {
			canvas.zoomOut(event.x, event.y);
		} else if (d < 0) {
			canvas.zoomIn(event.x, event.y);
		}
	});

	window.addEventListener("resize", () => canvas.resize());
	window.addEventListener("load", () => canvas.resize());
}

document.readyState == "loading"
	? document.addEventListener("DOMContentLoaded", main, {
		once: true
	})
	: main();
