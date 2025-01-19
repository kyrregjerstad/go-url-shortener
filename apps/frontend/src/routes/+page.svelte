<script lang="ts">
	import { shortenUrl } from '$lib/api';

	let url = '';
	let shortUrl = '';
	let error = '';
	let loading = false;

	async function handleSubmit() {
		if (!url) return;

		loading = true;
		error = '';

		try {
			shortUrl = await shortenUrl(url);
		} catch (e) {
			error = e instanceof Error ? e.message : 'Something went wrong';
		} finally {
			loading = false;
		}
	}
</script>

<main class="container mx-auto max-w-2xl px-4 py-8">
	<h1 class="mb-8 text-center text-4xl font-bold">URL Shortener</h1>

	<form on:submit|preventDefault={handleSubmit} class="space-y-4">
		<div>
			<label for="url" class="block text-sm font-medium text-gray-700">Enter URL to shorten</label>
			<input
				type="url"
				id="url"
				bind:value={url}
				required
				class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
				placeholder="https://example.com"
			/>
		</div>

		<button
			type="submit"
			disabled={loading}
			class="w-full rounded-md bg-blue-600 px-4 py-2 text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50"
		>
			{loading ? 'Shortening...' : 'Shorten URL'}
		</button>
	</form>

	{#if error}
		<div class="mt-4 rounded-md bg-red-100 p-4 text-red-700">
			{error}
		</div>
	{/if}

	{#if shortUrl}
		<div class="mt-8 rounded-md bg-green-50 p-4">
			<h2 class="mb-2 text-lg font-semibold">Your shortened URL:</h2>
			<div class="flex items-center gap-2">
				<a
					href={shortUrl}
					target="_blank"
					rel="noopener noreferrer"
					class="break-all text-blue-600 hover:underline"
				>
					{shortUrl}
				</a>
				<button
					on:click={() => navigator.clipboard.writeText(shortUrl)}
					class="p-2 text-gray-500 hover:text-gray-700"
					title="Copy to clipboard"
				>
					📋
				</button>
			</div>
		</div>
	{/if}
</main>
