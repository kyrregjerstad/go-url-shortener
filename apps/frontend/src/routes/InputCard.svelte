<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Card from '$lib/components/ui/card';
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';

	let { value = $bindable(), error } = $props();
</script>

<Card.Root
	class="animate-border w-full border border-transparent [background:linear-gradient(45deg,#172033,theme(colors.slate.800)_50%,#172033)_padding-box,conic-gradient(from_var(--border-angle),theme(colors.slate.900/.38)_20%,_theme(colors.indigo.900/.50)_20%,_theme(colors.indigo.900/.50)_30%,_theme(colors.indigo.500/.20)_94%,_theme(colors.slate.500/.48))_border-box]"
>
	<div class="bg-background/90 rounded-lg">
		<Card.Content class="space-y-1">
			<Label for="url">Link</Label>
			<Input type="url" id="url" name="url" bind:value required placeholder="https://example.com" />
		</Card.Content>
		<Card.Footer>
			{#if error}
				<p class="mt-1 text-sm text-red-600">{error}</p>
			{/if}
			<Button type="submit" class="w-full" disabled={value.length === 0}>Shorten Link</Button>
		</Card.Footer>
	</div>
</Card.Root>
