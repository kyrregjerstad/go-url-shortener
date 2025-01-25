<script lang="ts">
	import {
		VisSingleContainer,
		VisXYContainer,
		VisLine,
		VisAxis,
		VisTooltip,
		VisDonut
	} from '@unovis/svelte';
	import type { Tooltip } from '@unovis/ts';

	let { data } = $props();

	type Visit = {
		timestamp: string;
		userAgent: string;
		ipAddress: string;
		referer?: string;
	};

	type ChartDataPoint = {
		x: number;
		totalVisits: number;
		uniqueVisits: number;
	};

	type PieDataPoint = {
		name: string;
		value: number;
	};

	// Format data for the line chart
	const hourlyVisits: ChartDataPoint[] = data.analytics.hourlyData
		.map((d) => ({
			x: new Date(d.timestamp).getTime(),
			totalVisits: d.totalVisits,
			uniqueVisits: d.uniqueVisits
		}))
		.sort((a, b) => a.x - b.x);

	// Format data for the browser pie chart
	const browserData: PieDataPoint[] = Object.entries(data.analytics.browserStats).map(
		([name, value]) => ({
			name,
			value
		})
	);

	// Format data for the referrer pie chart
	const referrerData: PieDataPoint[] = Object.entries(data.analytics.referrerStats).map(
		([name, value]) => ({
			name,
			value
		})
	);

	// Tooltip configuration
	const tooltipConfig = {
		triggers: {
			trigger: 'hover'
		}
	};

	// Color scale for the charts
	const colors = ['#3b82f6', '#ef4444', '#10b981', '#f59e0b', '#6366f1'];

	// Format date for axis
	function formatDate(timestamp: number) {
		return new Date(timestamp).toLocaleString(undefined, {
			month: 'short',
			day: 'numeric',
			hour: 'numeric'
		});
	}

	// Line chart configuration
	const x = (d: ChartDataPoint) => d.x;
	const y = (d: ChartDataPoint) => d.totalVisits;
	const lineColor = '#3b82f6';

	console.log('hourlyVisits', hourlyVisits);
</script>

<main class="container mx-auto max-w-6xl px-4 py-8">
	<h1 class="mb-8 text-center text-4xl font-bold">URL Analytics</h1>

	<!-- Stats Overview -->
	<div class="mb-8 grid grid-cols-1 gap-4 sm:grid-cols-2">
		<div class="border-border bg-card text-card-foreground rounded-lg border p-6 shadow-sm">
			<h3 class="text-muted-foreground text-lg font-medium">Total Visits</h3>
			<p class="mt-2 text-4xl font-bold">{data.analytics.totalVisits}</p>
		</div>
		<div class="border-border bg-card text-card-foreground rounded-lg border p-6 shadow-sm">
			<h3 class="text-muted-foreground text-lg font-medium">Unique Visitors</h3>
			<p class="mt-2 text-4xl font-bold">{data.analytics.uniqueVisitors}</p>
		</div>
	</div>

	<!-- Visits Over Time -->
	<div class="border-border bg-card text-card-foreground mb-8 rounded-lg border p-6 shadow-sm">
		<h2 class="mb-4 text-2xl font-semibold">Visits Over Time</h2>
		<div class="h-[300px]">
			<VisXYContainer data={hourlyVisits}>
				<VisLine {x} {y} curve="monotone" color={lineColor} />
				<VisAxis type="x" tickFormat={formatDate} />
				<VisAxis type="y" />
				<VisTooltip
					{tooltipConfig}
					content={(d: ChartDataPoint) => `
						Time: ${formatDate(d.x)}<br>
						Visits: ${d.totalVisits}
					`}
				/>
			</VisXYContainer>
		</div>
	</div>

	<!-- Distribution Charts -->
	<div class="mb-8 grid grid-cols-1 gap-4 sm:grid-cols-2">
		<!-- Browser Distribution -->
		<div class="border-border bg-card text-card-foreground rounded-lg border p-6 shadow-sm">
			<h2 class="mb-4 text-2xl font-semibold">Browser Distribution</h2>
			<div class="h-[300px]">
				<VisSingleContainer data={browserData}>
					<VisDonut
						value={(d: PieDataPoint) => d.value}
						arcLabel={(d: PieDataPoint) => `${d.name} (${d.value})`}
						{colors}
					/>
					<VisTooltip {tooltipConfig} />
				</VisSingleContainer>
			</div>
		</div>

		<!-- Traffic Sources -->
		<div class="border-border bg-card text-card-foreground rounded-lg border p-6 shadow-sm">
			<h2 class="mb-4 text-2xl font-semibold">Traffic Sources</h2>
			<div class="h-[300px]">
				<VisSingleContainer data={referrerData}>
					<VisDonut
						value={(d: PieDataPoint) => d.value}
						arcLabel={(d: PieDataPoint) => `${d.name} (${d.value})`}
						{colors}
					/>
					<VisTooltip {tooltipConfig} />
				</VisSingleContainer>
			</div>
		</div>
	</div>

	<!-- Visit History Table -->
	<div class="border-border bg-card text-card-foreground rounded-lg border p-6 shadow-sm">
		<div class="flex items-center justify-between">
			<h2 class="text-2xl font-semibold">Visit History</h2>
		</div>

		<div class="mt-4 overflow-x-auto">
			<table class="divide-border min-w-full divide-y">
				<thead>
					<tr>
						<th
							class="text-muted-foreground px-6 py-3 text-left text-xs font-medium uppercase tracking-wider"
						>
							Time
						</th>
						<th
							class="text-muted-foreground px-6 py-3 text-left text-xs font-medium uppercase tracking-wider"
						>
							User Agent
						</th>
						<th
							class="text-muted-foreground px-6 py-3 text-left text-xs font-medium uppercase tracking-wider"
						>
							IP Address
						</th>
						<th
							class="text-muted-foreground px-6 py-3 text-left text-xs font-medium uppercase tracking-wider"
						>
							Referer
						</th>
					</tr>
				</thead>
				<tbody class="divide-border divide-y">
					{#each data.analytics.visits as visit}
						<tr class="hover:bg-muted/50">
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
										class="text-primary hover:text-primary/90"
										target="_blank"
										rel="noopener noreferrer"
									>
										{visit.referer}
									</a>
								{:else}
									<span class="text-muted-foreground">Direct</span>
								{/if}
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	</div>
</main>
