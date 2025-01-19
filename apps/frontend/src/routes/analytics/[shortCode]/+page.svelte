<script lang="ts">
	let { data } = $props();

	type Visit = {
		timestamp: string;
		userAgent: string;
		ipAddress: string;
		referer?: string;
	};
</script>

<main class="container mx-auto max-w-4xl px-4 py-8">
	<h1 class="mb-8 text-center text-4xl font-bold">URL Analytics</h1>

	<div class="space-y-6">
		<div class="flex items-center justify-between">
			<h2 class="text-2xl font-semibold">Visit History</h2>
			<p class="text-gray-600">Total Visits: {data.analytics.visits.length}</p>
		</div>

		<div class="overflow-x-auto">
			<table class="min-w-full divide-y divide-gray-200">
				<thead class="bg-gray-50">
					<tr>
						<th
							class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
						>
							Time
						</th>
						<th
							class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
						>
							User Agent
						</th>
						<th
							class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
						>
							IP Address
						</th>
						<th
							class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-500"
						>
							Referer
						</th>
					</tr>
				</thead>
				<tbody class="divide-y divide-gray-200 bg-white">
					{#each data.analytics.visits as visit}
						<tr>
							<td class="whitespace-nowrap px-6 py-4">
								{new Date(visit.timestamp).toLocaleString()}
							</td>
							<td class="max-w-xs truncate px-6 py-4">
								<span title={visit.userAgent}>{visit.userAgent}</span>
							</td>
							<td class="whitespace-nowrap px-6 py-4">{visit.ipAddress}</td>
							<td class="max-w-xs truncate px-6 py-4">
								{#if visit.referer}
									<a
										href={visit.referer}
										class="text-blue-600 hover:underline"
										target="_blank"
										rel="noopener noreferrer"
									>
										{visit.referer}
									</a>
								{:else}
									<span class="text-gray-400">Direct</span>
								{/if}
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	</div>
</main>
