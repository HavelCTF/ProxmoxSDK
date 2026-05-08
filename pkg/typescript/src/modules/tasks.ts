import { type NodeTaskStatusResponse, NodeTaskStatusResponseSchema } from '../types/tasks';

export interface WaitOptions {
    pollIntervalMs?: number;
    timeoutMs?: number;
}

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

    async wait(opts?: WaitOptions): Promise<NodeTaskStatusResponse> {
        const wasmOpts = {
            PollIntervalMs: opts?.pollIntervalMs ?? 1000,
            TimeoutMs: opts?.timeoutMs ?? 0,
        };
        return NodeTaskStatusResponseSchema.parse(
            await globalThis.ProxmoxWASM.tasks.Wait(this.nodeName, this.upid, wasmOpts),
        );
    }
}
