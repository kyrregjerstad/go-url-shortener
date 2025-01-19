import { PUBLIC_API_URL } from '$env/static/public';

interface ShortenResponse {
	shortUrl: string;
}

interface StatsResponse {
	visits: number;
	createdAt: string;
	originalUrl: string;
}

export class ApiError extends Error {
	constructor(
		public status: number,
		message: string
	) {
		super(message);
	}
}

export async function shortenUrl(url: string): Promise<string> {
	const response = await fetch(`${PUBLIC_API_URL}/shorten`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({ url })
	});

	if (!response.ok) {
		throw new ApiError(response.status, 'Failed to shorten URL');
	}

	const data: ShortenResponse = await response.json();
	return data.shortUrl;
}

export async function getUrlStats(shortUrl: string): Promise<StatsResponse> {
	const response = await fetch(`${PUBLIC_API_URL}/stats/${shortUrl}`);

	if (!response.ok) {
		throw new ApiError(response.status, 'Failed to get URL stats');
	}

	return response.json();
}
