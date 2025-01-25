import type { PageLoad } from './$types';

function generateMockData() {
	const browsers = ['Chrome', 'Firefox', 'Safari', 'Edge'];
	const referers = [
		'https://google.com',
		'https://twitter.com',
		'https://facebook.com',
		null // Direct traffic
	];
	const now = new Date();
	const sevenDaysAgo = new Date(now.getTime() - 7 * 24 * 60 * 60 * 1000);

	// Generate last 7 days of hourly data with cumulative visits
	let cumulativeVisits = 0;
	const hourlyData = Array.from({ length: 7 * 24 }, (_, i) => {
		const date = new Date(sevenDaysAgo);
		date.setHours(date.getHours() + i); // Start from 7 days ago and move forward

		// Generate new visits for this hour (between 0 and 5)
		const newVisits = Math.floor(Math.random() * 6);
		cumulativeVisits += newVisits;

		return {
			timestamp: date.toISOString(),
			totalVisits: cumulativeVisits,
			uniqueVisits: Math.floor(cumulativeVisits * 0.7) // Maintain roughly 70% unique ratio
		};
	});

	// Generate individual visit records
	const visits = Array.from({ length: cumulativeVisits }, () => {
		const browser = browsers[Math.floor(Math.random() * browsers.length)];
		const version = Math.floor(Math.random() * 15) + 100;
		return {
			timestamp: new Date(
				sevenDaysAgo.getTime() + Math.random() * 7 * 24 * 60 * 60 * 1000
			).toISOString(),
			userAgent: `${browser}/${version}.0.0.0`,
			ipAddress: `${Math.floor(Math.random() * 255)}.${Math.floor(
				Math.random() * 255
			)}.${Math.floor(Math.random() * 255)}.${Math.floor(Math.random() * 255)}`,
			referer: referers[Math.floor(Math.random() * referers.length)]
		};
	}).sort((a, b) => new Date(b.timestamp).getTime() - new Date(a.timestamp).getTime());

	// Calculate browser stats
	const browserStats = visits.reduce(
		(acc, visit) => {
			const browser = visit.userAgent.split('/')[0];
			acc[browser] = (acc[browser] || 0) + 1;
			return acc;
		},
		{} as Record<string, number>
	);

	// Calculate referrer stats
	const referrerStats = visits.reduce(
		(acc, visit) => {
			const source = visit.referer ? new URL(visit.referer).hostname : 'Direct';
			acc[source] = (acc[source] || 0) + 1;
			return acc;
		},
		{} as Record<string, number>
	);

	return {
		visits,
		hourlyData,
		browserStats,
		referrerStats,
		totalVisits: cumulativeVisits,
		uniqueVisitors: Math.floor(cumulativeVisits * 0.7) // Maintain the same unique visitor ratio
	};
}

export const load = (async ({ params }) => {
	const analytics = generateMockData();

	return {
		analytics,
		shortCode: params.shortCode
	};
}) satisfies PageLoad;
