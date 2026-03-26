import {
    type ClusterNextIdResponse,
    ClusterNextIdResponseSchema,
    type ClusterTasksResponse,
    ClusterTasksResponseSchema,
} from '../types/cluster';

export class ClusterModule {
    async getNextId(): Promise<ClusterNextIdResponse> {
        return ClusterNextIdResponseSchema.parse(await globalThis.ProxmoxWASM.cluster.GetNextId());
    }

    async getTasks(): Promise<ClusterTasksResponse> {
        return ClusterTasksResponseSchema.parse(await globalThis.ProxmoxWASM.cluster.GetTasks());
    }
}
