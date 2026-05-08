import { z } from 'zod';

export const CloneLXCDataSchema = z.object({
    NewId: z.number(),
    Target: z.string().optional(),
});
export type CloneLXCData = z.infer<typeof CloneLXCDataSchema>;

export const TaskBaseResponseSchema = z.object({
    Data: z.string(),
});
export type TaskBaseResponse = z.infer<typeof TaskBaseResponseSchema>;

export type DeleteLXCResponse = TaskBaseResponse;
export type CloneLXCResponse = TaskBaseResponse;
export type StartLXCResponse = TaskBaseResponse;
export type StopLXCResponse = TaskBaseResponse;
export type CreateLXCResponse = TaskBaseResponse; // Ajout pour la création

export const LXCFeaturesSchema = z.object({
    Nesting: z.boolean(),
});
export type LXCFeatures = z.infer<typeof LXCFeaturesSchema>;

export const CreateLXCDataSchema = z.object({
    OSTemplate: z.string(),
    VMID: z.number(),
    Features: LXCFeaturesSchema.optional(),
});
export type CreateLXCData = z.infer<typeof CreateLXCDataSchema>;

export const LXCsDataSchema = z.object({
    Status: z.string(),
    VMID: z.number(),
    CPU: z.number().optional(),
    CPUS: z.number().optional(),
    Disk: z.number().optional(),
    DiskRead: z.number().optional(),
    DiskWrite: z.number().optional(),
    Lock: z.string().optional(),
    MaxDisk: z.number().optional(),
    MaxMem: z.number().optional(),
    MaxSwap: z.number().optional(),
    Mem: z.number().optional(),
    Name: z.string().optional(),
    NetIn: z.number().optional(),
    NetOut: z.number().optional(),
    PressureCPUSome: z.string().optional(),
    PressureIOFull: z.string().optional(),
    PressureIOSome: z.string().optional(),
    PressureMemoryFull: z.string().optional(),
    PressureMemorySome: z.string().optional(),
    Tags: z.string().optional(),
    Template: z.boolean().optional(),
    Uptime: z.number().optional(),
});
export type LXCsData = z.infer<typeof LXCsDataSchema>;

export const LXCsResponseSchema = z.object({
    LXCs: z.array(LXCsDataSchema),
});
export type LXCsResponse = z.infer<typeof LXCsResponseSchema>;

export const LXCStatusDataSchema = z.object({
    Name: z.string().optional(),
    Status: z.string(),
    VMID: z.number().optional(),
    Uptime: z.number().optional(),
    CPUs: z.number().optional(),
    CPU: z.number().optional(),
    Mem: z.number().optional(),
    MaxMem: z.number().optional(),
    Disk: z.number().optional(),
    MaxDisk: z.number().optional(),
    Swap: z.number().optional(),
    MaxSwap: z.number().optional(),
    NetIn: z.number().optional(),
    NetOut: z.number().optional(),
    DiskRead: z.number().optional(),
    DiskWrite: z.number().optional(),
    Lock: z.string().optional(),
    Tags: z.string().optional(),
    Type: z.string().optional(),
});
export type LXCStatusData = z.infer<typeof LXCStatusDataSchema>;

export const LXCStatusResponseSchema = z.object({
    Data: LXCStatusDataSchema,
});
export type LXCStatusResponse = z.infer<typeof LXCStatusResponseSchema>;
