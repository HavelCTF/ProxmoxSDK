import { z } from 'zod';

export const NodesDataSchema = z.object({
    Node: z.string(),
    Status: z.string(),
    Uptime: z.number().optional(),
    SslFingerprint: z.string().optional(),
    CPU: z.number().optional(),
    Level: z.string().optional(),
    MaxCPU: z.number().optional(),
    MaxMEM: z.number().optional(),
    MEM: z.number().optional(),
});
export type NodesData = z.infer<typeof NodesDataSchema>;

export const NodesResponseSchema = z.object({
    Data: z.array(NodesDataSchema),
});
export type NodesResponse = z.infer<typeof NodesResponseSchema>;

export const NodeTaskDataSchema = z.object({
    UPID: z.string(),
    Node: z.string().optional(),
    PID: z.number().optional(),
    PStart: z.number().optional(),
    StartTime: z.number().optional(),
    Type: z.string().optional(),
    User: z.string().optional(),
    Status: z.string().optional(),
});

export const NodeTasksResponseSchema = z.object({
    Data: z.array(NodeTaskDataSchema),
});
export type NodeTasksResponse = z.infer<typeof NodeTasksResponseSchema>;
