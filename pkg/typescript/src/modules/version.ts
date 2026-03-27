import { type VersionResponse, VersionResponseSchema } from '../types/version';

export async function getVersion(): Promise<VersionResponse> {
    return VersionResponseSchema.parse(await globalThis.ProxmoxWASM.version.GetVersion());
}
