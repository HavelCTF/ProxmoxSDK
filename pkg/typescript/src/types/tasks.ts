import { z } from 'zod';

export const NodeTaskBaseSchema = z.object({
    ID: z.string(),
    Node: z.string(),
    PID: z.number(),
    PStart: z.number(),
    StartTime: z.number(),
    Type: z.string(),
    UPID: z.string(),
    User: z.string(),
});

export const NodeTaskStatusDataSchema = NodeTaskBaseSchema.extend({
    Status: z.string(),
    ExitStatus: z.string().optional(),
});
export type NodeTaskStatusData = z.infer<typeof NodeTaskStatusDataSchema>;

export const NodeTaskStatusResponseSchema = z.object({
    Data: NodeTaskStatusDataSchema,
});
export type NodeTaskStatusResponse = z.infer<typeof NodeTaskStatusResponseSchema>;

export const NodeTaskDeleteResponseSchema = z.object({
    Data: z.string(),
});
export type NodeTaskDeleteResponse = z.infer<typeof NodeTaskDeleteResponseSchema>;
