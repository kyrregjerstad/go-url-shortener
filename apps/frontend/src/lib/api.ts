import { PUBLIC_API_URL } from '$env/static/public';

interface ShortenResponse {
	shortUrl: string;
}

interface StatsResponse {
	longUrl: string;
	createdAt: string;
	visits: number;
	lastVisit?: string;
}

interface Visit {
	timestamp: string;
	userAgent: string;
	ipAddress: string;
	referer?: string;
}

interface AnalyticsResponse {
	short_code: string;
	visits: Visit[];
}

export class ApiError extends Error {
	constructor(
		public status: number,
		message: string
	) {
		super(message);
	}
}

abstract class BaseApi {
	constructor(protected baseUrl: string = PUBLIC_API_URL) {}

	protected async get<T>(path: string, fetcher: typeof fetch): Promise<T> {
		try {
			const response = await fetcher(`${this.baseUrl}${path}`, {
				headers: {
					'Content-Type': 'application/json'
				}
			});

			if (!response.ok) {
				throw new ApiError(response.status, await response.text());
			}

			return response.json();
		} catch (error) {
			console.error('Error fetching data:', error);
			throw error;
		}
	}

	protected async post<T>(path: string, data: unknown): Promise<T> {
		try {
			const response = await fetch(`${this.baseUrl}${path}`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(data)
			});

			if (!response.ok) {
				throw new ApiError(response.status, await response.text());
			}

			return response.json();
		} catch (error) {
			console.error('Error posting data:', error);
			throw error;
		}
	}
}

class UrlShortenerApi extends BaseApi {
	async shorten(url: string): Promise<string> {
		const data = await this.post<ShortenResponse>('/shorten', { url });
		return data.shortUrl;
	}

	async getStats(shortCode: string, fetcher: typeof fetch): Promise<StatsResponse> {
		return this.get<StatsResponse>(`/stats/${shortCode}`, fetcher);
	}

	async getAnalytics(shortCode: string, fetcher: typeof fetch): Promise<AnalyticsResponse> {
		return this.get<AnalyticsResponse>(`/analytics/${shortCode}`, fetcher);
	}
}

// Create a singleton instance
export const api = new UrlShortenerApi();

// Export type-only interfaces for use in components
export type { ShortenResponse, StatsResponse, Visit, AnalyticsResponse };
