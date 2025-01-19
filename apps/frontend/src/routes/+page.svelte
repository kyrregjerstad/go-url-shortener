<script lang="ts">
	import { superForm } from 'sveltekit-superforms';

	let { data, form } = $props();
	const {
		form: formData,
		errors,
		enhance,
		message
	} = superForm(data.form, {
		taintedMessage: null,
		onUpdate({ form, result }) {
			if (form.valid && result?.type === 'success') {
				// Access the shortUrl from the result data
				console.log('Short URL:', result.data?.shortUrl);
			}
		}
	});
</script>

<main class="container mx-auto max-w-2xl px-4 py-8">
	<h1 class="mb-8 text-center text-4xl font-bold">URL Shortener</h1>

	<form method="POST" use:enhance class="space-y-4">
		<div>
			<label for="url" class="block text-sm font-medium text-gray-700">Enter URL to shorten</label>
			<input
				type="url"
				id="url"
				name="url"
				bind:value={$formData.url}
				required
				class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
				placeholder="https://example.com"
			/>
			{#if $errors.url}
				<p class="mt-1 text-sm text-red-600">{$errors.url}</p>
			{/if}
		</div>

		<button
			type="submit"
			class="w-full rounded-md bg-blue-600 px-4 py-2 text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50"
		>
			Shorten URL
		</button>
	</form>

	{#if $message}
		<div class="mt-4 rounded-md bg-green-100 p-4 text-green-700">
			{$message}
		</div>
	{/if}

	{#if form?.shortUrl}
		<div class="mt-8 rounded-md bg-green-50 p-4">
			<h2 class="mb-2 text-lg font-semibold">Your shortened URL:</h2>
			<div class="flex items-center gap-2">
				<a
					href={form.shortUrl}
					target="_blank"
					rel="noopener noreferrer"
					class="break-all text-blue-600 hover:underline"
				>
					{form.shortUrl}
				</a>
				<button
					onclick={() => navigator.clipboard.writeText(form.shortUrl)}
					class="p-2 text-gray-500 hover:text-gray-700"
					title="Copy to clipboard"
				>
					📋
				</button>
			</div>
		</div>
	{/if}
</main>
