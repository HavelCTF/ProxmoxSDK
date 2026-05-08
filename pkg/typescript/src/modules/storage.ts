import {
    type StorageContentResponse,
    StorageContentResponseSchema,
    type StorageContentType,
    type UploadResponse,
    UploadResponseSchema,
    type UploadTemplateBody,
} from '../types/storage';

async function toUint8Array(body: UploadTemplateBody): Promise<Uint8Array> {
    if (body instanceof Uint8Array) {
        return body;
    }
    // Blob (browser File picker, fetch responses, ...). Avoid relying on
    // Node-only Buffer APIs so the SDK stays isomorphic.
    const buf = await body.arrayBuffer();
    return new Uint8Array(buf);
}

export class StorageService {
    constructor(
        private readonly nodeName: string,
        private readonly storageName: string,
    ) {}

    async getContent(contentType: StorageContentType = ''): Promise<StorageContentResponse> {
        return StorageContentResponseSchema.parse(
            await globalThis.ProxmoxWASM.storage.GetContent(this.nodeName, this.storageName, contentType),
        );
    }

    async uploadTemplate(filename: string, body: UploadTemplateBody): Promise<UploadResponse> {
        const payload = await toUint8Array(body);
        return UploadResponseSchema.parse(
            await globalThis.ProxmoxWASM.storage.UploadTemplate(this.nodeName, this.storageName, filename, payload),
        );
    }
}
