import { z } from 'zod';

export const StorageContentDataSchema = z.object({
    VolID: z.string(),
    Content: z.string(),
    Format: z.string().optional(),
    Size: z.number().optional(),
    Used: z.number().optional(),
    CTime: z.number().optional(),
    VMID: z.number().optional(),
    Notes: z.string().optional(),
    Parent: z.string().optional(),
});
export type StorageContentData = z.infer<typeof StorageContentDataSchema>;

export const StorageContentResponseSchema = z.object({
    Data: z.array(StorageContentDataSchema),
});
export type StorageContentResponse = z.infer<typeof StorageContentResponseSchema>;

/**
 * Content type filter for `node().storage(name).getContent(...)`.
 *
 * Proxmox storage volumes are categorised by their content type. The four
 * we currently care about are templates, ISOs, container backups and
 * container root volumes. Pass an empty string to return everything.
 */
export type StorageContentType = 'vztmpl' | 'iso' | 'backup' | 'rootdir' | '';

export const UploadResponseSchema = z.object({
    Data: z.string(),
});
export type UploadResponse = z.infer<typeof UploadResponseSchema>;

/**
 * Body accepted by `uploadTemplate`. We accept either a raw byte buffer
 * (Node.js + browser-side `Uint8Array`) or a `Blob` (browser File picker).
 */
export type UploadTemplateBody = Uint8Array | Blob;
