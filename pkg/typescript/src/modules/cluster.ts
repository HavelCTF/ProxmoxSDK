import { ClusterNextIdResponse, ClusterNextIdResponseSchema } from '../types/cluster';

export class ClusterModule {
    public async GetNextId(): Promise<ClusterNextIdResponse> {
        if (!globalThis.ProxmoxWASM?.cluster) {
            throw new Error('ProxmoxSDK not initialized — call ProxmoxSDK.create() first');
        }
        const raw = await globalThis.ProxmoxWASM.cluster.GetNextId();
        return ClusterNextIdResponseSchema.parse(raw);
    }
}