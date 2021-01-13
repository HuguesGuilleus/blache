// Download d for the user in the file name.
function userDownload(b: Blob, name: string) {
	const u: string = URL.createObjectURL(b);
	const a: HTMLAnchorElement = document.createElement('a');
	document.body.appendChild(a);
	a.href = u;
	a.download = name;
	a.click();
	a.remove();
	URL.revokeObjectURL(u);
}
