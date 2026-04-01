import {
    type ClusterNextIdResponse,
    ClusterNextIdResponseSchema,
    type ClusterTasksResponse,
    ClusterTasksResponseSchema,
} from '../types/cluster';

export class ClusterService {
    async getTasks(): Promise<ClusterTasksResponse> {
        return ClusterTasksResponseSchema.parse(await globalThis.ProxmoxWASM.cluster.GetTasks());
    }

    async getNextId(): Promise<ClusterNextIdResponse> {
        return ClusterNextIdResponseSchema.parse(await globalThis.ProxmoxWASM.cluster.GetNextId());
    }
}
