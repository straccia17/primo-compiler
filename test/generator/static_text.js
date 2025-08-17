export function template() {
	const text = document.createTextNode("Hello, World!");
	return {
		mount(container) {
			container.appendChild(text);
		}
	}
}