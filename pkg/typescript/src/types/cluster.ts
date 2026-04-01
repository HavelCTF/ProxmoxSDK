import { z } from 'zod';

export const ClusterNextIdResponseSchema = z.object({ VMID: z.string() });
export type ClusterNextIdResponse = z.infer<typeof ClusterNextIdResponseSchema>;

export const ClusterTasksDataSchema = z.object({ UPID: z.string() });
export type ClusterTasksData = z.infer<typeof ClusterTasksDataSchema>;

export const ClusterTasksResponseSchema = z.object({ Data: z.array(ClusterTasksDataSchema) });
export type ClusterTasksResponse = z.infer<typeof ClusterTasksResponseSchema>;
