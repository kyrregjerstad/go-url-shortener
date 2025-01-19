import { fail } from '@sveltejs/kit';
import { message, superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';

import { shortenUrl } from '$lib/api';
import { urlSchema } from '$lib/schemas';
import type { PageServerLoad, Actions } from './$types';

export const load = (async () => {
	const form = await superValidate(zod(urlSchema));
	return { form };
}) satisfies PageServerLoad;

export const actions = {
	default: async ({ request }) => {
		const form = await superValidate(request, zod(urlSchema));

		if (!form.valid) {
			return fail(400, { form });
		}

		try {
			const shortUrl = await shortenUrl(form.data.url);
			form.message = 'URL shortened successfully';
			return { form, shortUrl };
		} catch (error) {
			return message(form, error instanceof Error ? error.message : 'Something went wrong', {
				status: 400
			});
		}
	}
} satisfies Actions;
