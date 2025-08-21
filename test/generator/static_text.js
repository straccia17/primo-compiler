export function template(ctx) {
	const text = document.createTextNode("Hello, World!");
	return {
		mount(container) {
			container.appendChild(text);
		}
	}
}