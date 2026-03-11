import { z } from 'zod';

export const ClusterNextIdResponseSchema = z.object({ id: z.string() });
export type ClusterNextIdResponse = z.infer<typeof ClusterNextIdResponseSchema>;