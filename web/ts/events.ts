// Use to get some element into the dom.
function $(id: string): HTMLElement {
	const e = document.getElementById(id);
	if (!e) throw `Element '${id}' no found in DOM`;
	return e;
}

// Create the canvas and add user event handler.
function main() {
	const canvas = new Canvas(
		<HTMLCanvasElement>$("canvas2d"),
		coordURL.parse(location.hash),
	),
		coordsBloc = $("coordsBloc"),
		coordsRegion = $("coordsRegion"),
		coordsStruct = $("coordsStruct"),
		url = <HTMLAnchorElement>$("url");

	$("tileTypeBloc").addEventListener(
		"click",
		() => canvas.type = UserTileType.bloc,
	);
	$("tileTypeBiome").addEventListener(
		"click",
		() => canvas.type = UserTileType.biome,
	);
	$("tileTypeHeight").addEventListener(
		"click",
		() => canvas.type = UserTileType.height,
	);
	$("savePNG").addEventListener("click", () => canvas.userDownload());

	url.addEventListener("click", (event) => {
		event.preventDefault();
		navigator.clipboard.writeText(url.href);
	});
	canvas.onDrawAll.push(() => {
		url.href = coordURL.url(canvas);
		url.innerText = coordURL.hash(canvas);
	});

	type enbaleProperties =
		| "enabledFrontier"
		| "enabledStructure"
		| "enabledWater";
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

	window.addEventListener("keydown", (e) => {
		switch (e.key.toLowerCase()) {
			case "b":
				canvas.type = UserTileType.bloc;
				break;
			case "h":
				canvas.type = UserTileType.height;
				break;
			case "n":
				canvas.type = UserTileType.biome;
				break;
			case "f":
				return switchFrontier();
			case "w":
				return switchWater();
			case "x":
				return switchStructure();
			case "s":
				return canvas.userDownload();
			case "-":
				return canvas.zoomOut(
					canvas.canvasElement.width / 2,
					canvas.canvasElement.height / 2,
				);
			case "+":
				return canvas.zoomIn(
					canvas.canvasElement.width / 2,
					canvas.canvasElement.height / 2,
				);
			case "arrowleft":
				canvas.positionX -= canvas.size / 4;
				break;
			case "arrowright":
				canvas.positionX += canvas.size / 4;
				break;
			case "arrowup":
				canvas.positionZ -= canvas.size / 4;
				break;
			case "arrowdown":
				canvas.positionZ += canvas.size / 4;
				break;
			case "0":
				canvas.positionX = canvas.positionZ = 0;
				canvas.size = REGION_SIZE;
				break;
			default:
				return;
		}

		coordsBloc.innerText = "";
		coordsRegion.innerText = "";
		coordsStruct.innerText = "";
		canvas.drawAll();
	});

	canvas.canvasElement.addEventListener("mousemove", (event) => {
		if (event.buttons) {
			canvas.positionX -= event.movementX;
			canvas.positionZ -= event.movementY;
			canvas.drawAll();
		}

		const x = Math.trunc(
			(canvas.positionX + event.x) * REGION_SIZE / canvas.size,
		),
			z = Math.trunc((canvas.positionZ + event.y) * REGION_SIZE / canvas.size),
			regionCoordX = Math.floor(x / REGION_SIZE),
			regionCoordZ = Math.floor(z / REGION_SIZE),
			insideRegionX = x - regionCoordX * REGION_SIZE, // we do not use modulo because negatives values.
			insideRegionZ = z - regionCoordZ * REGION_SIZE;

		coordsBloc.innerText = `${x}, ${z}`;
		coordsRegion.innerText = `${regionCoordX}, ${regionCoordZ}`;
		coordsStruct.innerText = "";
		for (
			const s of (canvas.structures.get(
				Coordinate.stringer(regionCoordX, regionCoordZ),
			) ?? [])
		) {
			if (
				Math.abs(insideRegionX - s.x * 16) < 10 &&
				Math.abs(insideRegionZ - s.z * 16) < 10
			) {
				coordsStruct.innerText = s.name;
				break;
			}
		}
	});

	canvas.canvasElement.addEventListener("wheel", (event) => {
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
		once: true,
	})
	: main();
