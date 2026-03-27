import {
    type CreateLXCData,
    type LXCsResponse,
    LXCsResponseSchema,
    type TaskBaseResponse,
    TaskBaseResponseSchema,
} from '../types/lxc';
import {
    type NodesResponse,
    NodesResponseSchema,
    type NodeTasksResponse,
    NodeTasksResponseSchema,
} from '../types/nodes';
import { LxcService } from './lxc';
import { TaskService } from './tasks';

export class NodeService {
    constructor(private readonly nodeName: string) {}

    async getTasks(): Promise<NodeTasksResponse> {
        return NodeTasksResponseSchema.parse(await globalThis.ProxmoxWASM.tasks.GetTasks(this.nodeName));
    }

    tasks(upid: string): TaskService {
        return new TaskService(this.nodeName, upid);
    }

    async getLXCs(): Promise<LXCsResponse> {
        return LXCsResponseSchema.parse(await globalThis.ProxmoxWASM.lxc.GetLXCs(this.nodeName));
    }

    async postLXC(data: CreateLXCData): Promise<TaskBaseResponse> {
        return TaskBaseResponseSchema.parse(await globalThis.ProxmoxWASM.lxc.PostLXC(this.nodeName, data));
    }

    lxc(vmid: number): LxcService {
        return new LxcService(this.nodeName, vmid);
    }
}

export async function getNodes(): Promise<NodesResponse> {
    return NodesResponseSchema.parse(await globalThis.ProxmoxWASM.nodes.GetNodes());
}
