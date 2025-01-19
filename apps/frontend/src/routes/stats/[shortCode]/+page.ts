import type { PageLoad } from './$types';
import { error } from '@sveltejs/kit';
import { api } from '$lib/api';

export const load: PageLoad = async ({ params }) => {
	try {
		console.log('params', params);
		const stats = await api.getStats(params.shortCode);
		console.log('stats', stats);
		return { stats };
	} catch (err) {
		if (err instanceof Error) {
			throw error(500, err.message);
		}
		throw error(500, 'Failed to fetch stats');
	}
};
