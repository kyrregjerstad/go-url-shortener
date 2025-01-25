import type { PageLoad } from './$types';
import { error } from '@sveltejs/kit';
import { api } from '$lib/api';

export const load: PageLoad = async ({ params, fetch }) => {
	try {
		const analytics = await api.getAnalytics(params.shortCode, fetch);
		return { analytics };
	} catch (err) {
		if (err instanceof Error) {
			throw error(500, err.message);
		}
		throw error(500, 'Failed to fetch analytics');
	}
};
