const Coordinate = {
	parse(s: string): [number, number] {
		const t: string[] = s.split('.');
		if (t.length != 2) throw `Invalid syntax for coordinates string`;
		return [parseInt(t[0]), parseInt(t[1])];
	},
	toString(x: number, z: number): string {
		return x + '.' + z;
	},
}
