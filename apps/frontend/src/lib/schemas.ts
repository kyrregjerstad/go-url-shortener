import { z } from 'zod';

export const urlSchema = z.object({
	url: z.string().min(3, 'Please enter a valid URL')
});

export type UrlSchema = typeof urlSchema;
