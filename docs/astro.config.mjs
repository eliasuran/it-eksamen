import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

// https://astro.build/config
export default defineConfig({
	integrations: [
		starlight({
			title: 'Wolt API',
			social: {
				github: 'https://github.com/eliasuran/it-eksamen',
			},
			sidebar: [
				{
					label: 'Introduksjon',
					items: [
						// Each item here is one entry in the navigation menu.
						{ label: 'Introduksjon', link: '/introduksjon/01-introduksjon/' },
						{ label: 'Rask start', link: '/introduksjon/02-raskstart/' },
						{ label: 'Oppsett beskrivelse', link: '/introduksjon/03-oppsett-beskrivelse/' },
					],
				},
				{
					label: 'API referanser',
					autogenerate: { directory: 'reference' },
				},
			],
		}),
	],
});
