<script lang="ts">
	import type { Product } from '$lib/types.js';
	import type { PageData } from './$types.js';

	export let data: PageData;

	const products = data.data as Product[];

	console.log(data);
</script>

<div class="px-40 py-10">
	<h1 class="text-4xl">Hele wolt sin katalog >:)</h1>
	<div class="flex flex-wrap justify-center gap-40">
		{#each products as product}
			<div class="h-96 w-72 bg-blue-200 p-2 rounded-md flex flex-col gap-2">
				<img class="rounded-md" src={product.Imagelink} alt={`Bilde av ${product.Title}`} />
				<div class="flex flex-col justify-between h-full">
					<span>{product.Title}</span>
					<div class="text-sm text-black/50">{product.Category} - {product.Subcategory}</div>
				</div>
			</div>
		{/each}
	</div>
</div>
