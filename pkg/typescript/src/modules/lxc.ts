import { type CloneLXCData, type TaskBaseResponse, TaskBaseResponseSchema } from '../types/lxc';

export class LxcStatusService {
    constructor(
        private readonly nodeName: string,
        private readonly vmid: number,
    ) {}

    async start(): Promise<TaskBaseResponse> {
        return TaskBaseResponseSchema.parse(await globalThis.ProxmoxWASM.lxc.StartLXC(this.nodeName, this.vmid));
    }

    async stop(): Promise<TaskBaseResponse> {
        return TaskBaseResponseSchema.parse(await globalThis.ProxmoxWASM.lxc.StopLXC(this.nodeName, this.vmid));
    }
}

export class LxcService {
    constructor(
        private readonly nodeName: string,
        private readonly vmid: number,
    ) {}

    async delete(): Promise<TaskBaseResponse> {
        return TaskBaseResponseSchema.parse(await globalThis.ProxmoxWASM.lxc.DeleteLXC(this.nodeName, this.vmid));
    }

    async clone(data: CloneLXCData): Promise<TaskBaseResponse> {
        return TaskBaseResponseSchema.parse(await globalThis.ProxmoxWASM.lxc.CloneLXC(this.nodeName, this.vmid, data));
    }

    status(): LxcStatusService {
        return new LxcStatusService(this.nodeName, this.vmid);
    }
}
