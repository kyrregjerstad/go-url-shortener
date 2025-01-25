<script lang="ts">
	import { toast } from 'svelte-sonner';
	import InputCard from './InputCard.svelte';

	import { superForm } from 'sveltekit-superforms';

	let { data, form } = $props();
	let {
		form: formData,
		errors,
		enhance,
		submitting
	} = superForm(data.form, {
		taintedMessage: null,
		onUpdate({ form, result }) {
			console.log('result', result);
			if (form.valid && result?.type === 'success') {
				// Access the shortUrl from the result data

				if (result.data?.shortUrl) {
					copyToClipboard(result.data?.shortUrl);
					toast.success('🚀 Link shortened and copied to clipboard!');
				} else {
					toast.error('🚨 Error: ' + result.data?.error);
				}
			}

			if (form.valid && result?.type === 'failure') {
				toast.error('🚨 Error: ' + result.data?.error);
			}
		}
	});

	function copyToClipboard(url: string) {
		navigator.clipboard.writeText(url);
	}
</script>

<main class="container mx-auto h-[calc(100dvh-11rem)] max-w-2xl px-4 py-8">
	<h1 class="mb-8 text-center text-4xl font-bold">Link Shortener</h1>

	<form method="POST" use:enhance>
		<InputCard value={$formData.url} error={$errors.url} isSubmitting={$submitting} />
	</form>

	{#if form?.shortUrl}
		<div class="border-border bg-card text-card-foreground mt-4 rounded-lg border p-4 shadow-sm">
			<h2 class="mb-2 text-lg font-medium">Your shortened URL:</h2>
			<div class="flex items-center gap-2">
				<a
					href={form.shortUrl}
					target="_blank"
					rel="noopener noreferrer"
					class="text-primary hover:text-primary/90 break-all"
				>
					{form.shortUrl}
				</a>
				<button
					onclick={() => copyToClipboard(form.shortUrl)}
					class="text-muted-foreground hover:bg-accent hover:text-accent-foreground rounded-md p-2"
					title="Copy to clipboard"
				>
					📋
				</button>
			</div>
		</div>
	{/if}
</main>
