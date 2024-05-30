import { defineConfig } from 'vite';
import { sveltekit } from '@sveltejs/kit/vite';

export default defineConfig({
	plugins: [sveltekit()],
	preview: {
		port: 3000,
		strictPort: true
	},
	server: {
		port: 3000,
		strictPort: true,
		host: true,
		origin: 'http://0.0.0.0:3000'
	}
});
