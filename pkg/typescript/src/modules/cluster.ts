import { type ClusterNextIdResponse, ClusterNextIdResponseSchema } from '../types/cluster';

export class ClusterModule {
    async getNextId(): Promise<ClusterNextIdResponse> {
        return ClusterNextIdResponseSchema.parse(await globalThis.ProxmoxWASM.cluster.GetNextId());
    }
}
