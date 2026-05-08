import {
    type CloneLXCData,
    type LXCStatusResponse,
    LXCStatusResponseSchema,
    type TaskBaseResponse,
    TaskBaseResponseSchema,
} from '../types/lxc';

export class LXCStatusService {
    constructor(
        private readonly nodeName: string,
        private readonly vmid: number,
    ) {}

    async startLXC(): Promise<TaskBaseResponse> {
        return TaskBaseResponseSchema.parse(await globalThis.ProxmoxWASM.lxc.StartLXC(this.nodeName, this.vmid));
    }

    async stopLXC(): Promise<TaskBaseResponse> {
        return TaskBaseResponseSchema.parse(await globalThis.ProxmoxWASM.lxc.StopLXC(this.nodeName, this.vmid));
    }
}

export class LXCService {
    constructor(
        private readonly nodeName: string,
        private readonly vmid: number,
    ) {}

    async deleteLXC(): Promise<TaskBaseResponse> {
        return TaskBaseResponseSchema.parse(await globalThis.ProxmoxWASM.lxc.DeleteLXC(this.nodeName, this.vmid));
    }

    async cloneLXC(data: CloneLXCData): Promise<TaskBaseResponse> {
        return TaskBaseResponseSchema.parse(await globalThis.ProxmoxWASM.lxc.CloneLXC(this.nodeName, this.vmid, data));
    }

    async get(): Promise<LXCStatusResponse> {
        return LXCStatusResponseSchema.parse(await globalThis.ProxmoxWASM.lxc.Get(this.nodeName, this.vmid));
    }

    status(): LXCStatusService {
        return new LXCStatusService(this.nodeName, this.vmid);
    }
}
