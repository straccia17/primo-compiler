import { createEffect } from "core/signal";

export function template(ctx) {
	const text = document.createTextNode("");
	createEffect(() => text.data = ctx["name"]());
	return {
		mount(container) {
			container.appendChild(text);
		}
	}
}