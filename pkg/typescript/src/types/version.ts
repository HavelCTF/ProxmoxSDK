import { z } from 'zod';

export const VersionDataSchema = z.object({
    Version: z.string(),
    Release: z.string(),
    RepoID: z.string(),
    Console: z.string(),
});
export type VersionData = z.infer<typeof VersionDataSchema>;

export const VersionResponseSchema = z.object({
    Data: VersionDataSchema,
});
export type VersionResponse = z.infer<typeof VersionResponseSchema>;
