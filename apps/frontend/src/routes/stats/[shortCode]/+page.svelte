<script lang="ts">
	let { data } = $props();
</script>

<main class="container mx-auto max-w-2xl px-4 py-8">
	<h1 class="mb-8 text-center text-4xl font-bold">URL Statistics</h1>

	<div class="space-y-6 rounded-lg border p-6 shadow-sm">
		<div>
			<h2 class="text-lg font-semibold text-gray-900">Original URL</h2>
			<a
				href={data.stats.longUrl}
				class="text-blue-600 hover:underline"
				target="_blank"
				rel="noopener noreferrer"
			>
				{data.stats.longUrl}
			</a>
		</div>

		<div>
			<h2 class="text-lg font-semibold text-gray-900">Created At</h2>
			<p class="text-gray-600">{new Date(data.stats.createdAt).toLocaleString()}</p>
		</div>

		<div>
			<h2 class="text-lg font-semibold text-gray-900">Total Visits</h2>
			<p class="text-gray-600">{data.stats.visits}</p>
		</div>

		{#if data.stats.lastVisit}
			<div>
				<h2 class="text-lg font-semibold text-gray-900">Last Visit</h2>
				<p class="text-gray-600">{new Date(data.stats.lastVisit).toLocaleString()}</p>
			</div>
		{/if}
	</div>
</main>
