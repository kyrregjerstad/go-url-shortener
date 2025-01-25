<script lang="ts">
	import InputCard from './InputCard.svelte';

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

<main class="container mx-auto h-[calc(100dvh-11rem)] max-w-2xl px-4 py-8">
	<h1 class="mb-8 text-center text-4xl font-bold">Link Shortener</h1>

	<form method="POST" use:enhance>
		<InputCard value={$formData.url} error={$errors.url} />
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
