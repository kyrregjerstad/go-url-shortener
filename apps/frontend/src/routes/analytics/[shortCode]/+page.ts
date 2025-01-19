import type { PageLoad } from './$types';
import { error } from '@sveltejs/kit';
import { api } from '$lib/api';

export const load: PageLoad = async ({ params }) => {
	try {
		const analytics = await api.getAnalytics(params.shortCode);
		return { analytics };
	} catch (err) {
		if (err instanceof Error) {
			throw error(500, err.message);
		}
		throw error(500, 'Failed to fetch analytics');
	}
};
