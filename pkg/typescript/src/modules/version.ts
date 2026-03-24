import { type VersionResponse, VersionResponseSchema } from '../types/version';

export class VersionModule {
    async getVersion(): Promise<VersionResponse> {
        return VersionResponseSchema.parse(await globalThis.ProxmoxWASM.version.GetVersion());
    }
}
