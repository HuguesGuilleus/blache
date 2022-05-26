/// Prase and create string with "#tileType.size.x.z" format to share in hash
/// in URL.
const coordURL = {
	/// Parse the URL form hash.
	parse(hash: string): [UserTileType, number, number, number] {
		const [typeString = "", zoom = "", x = "", z = ""] = hash.substr(1).split(":");

		const type = typeString === "biome" ? UserTileType.biome
			: (typeString === "height" ? UserTileType.height : UserTileType.bloc);

		return [type, parseInt(zoom) || REGION_SIZE, parseInt(x) || 0, parseInt(z) || 0]
	},

	/// Return the url from current location and param.
	url(canvas: Canvas): string {
		return location.origin + location.pathname + this.hash(canvas);
	},

	/// Return only hash location with canvas coord.
	hash(canvas: Canvas): string {
		return "#" + [
			canvas.type,
			canvas.size,
			Math.floor(canvas.positionX),
			Math.floor(canvas.positionZ),
		].join(":");
	},
}
