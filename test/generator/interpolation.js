export function template() {
	const text = document.createTextNode("");
	return {
		mount(container) {
			container.appendChild(text);
		},
		update(changes) {
			text.data = changes["name"] ?? "";
		}
	}
}