import {
    type NodeTaskDeleteResponse,
    NodeTaskDeleteResponseSchema,
    type NodeTaskStatusResponse,
    NodeTaskStatusResponseSchema,
} from '../types/tasks';

export class TasksModule {
    async getTaskStatus(node: string, upid: string): Promise<NodeTaskStatusResponse> {
        return NodeTaskStatusResponseSchema.parse(await globalThis.ProxmoxWASM.tasks.GetTaskStatus(node, upid));
    }

    async deleteTask(node: string, upid: string): Promise<NodeTaskDeleteResponse> {
        return NodeTaskDeleteResponseSchema.parse(await globalThis.ProxmoxWASM.tasks.DeleteTask(node, upid));
    }
}
