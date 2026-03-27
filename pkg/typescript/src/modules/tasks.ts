import { type NodeTaskStatusResponse, NodeTaskStatusResponseSchema } from '../types/tasks';

export class TaskService {
    constructor(
        private readonly nodeName: string,
        private readonly upid: string,
    ) {}

    async getTaskStatus(): Promise<NodeTaskStatusResponse> {
        return NodeTaskStatusResponseSchema.parse(
            await globalThis.ProxmoxWASM.tasks.GetTaskStatus(this.nodeName, this.upid),
        );
    }

    async deleteTask(): Promise<void> {
        await globalThis.ProxmoxWASM.tasks.DeleteTask(this.nodeName, this.upid);
    }
}
