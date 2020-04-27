// Download d for the user in the file name.
function userDownload(b: Blob, name: string) {
	let u = URL.createObjectURL(b);
	let a = document.createElement('a');
	document.body.appendChild(a);
	a.href = u;
	a.download = name;
	a.click();
	a.remove();
	URL.revokeObjectURL(u);
}
