import {
    type CloneLXCData,
    type TaskBaseResponse,
    TaskBaseResponseSchema,
    type LXCsResponse,
    LXCsResponseSchema,
    type CreateLXCData
} from '../types/lxc';

export class LxcModule {
    async getLXCs(node: string): Promise<LXCsResponse> {
        return LXCsResponseSchema.parse(
            await globalThis.ProxmoxWASM.lxc.GetLXCs(node)
        );
    }

    async postLXC(node: string, data: CreateLXCData): Promise<TaskBaseResponse> {
        return TaskBaseResponseSchema.parse(
            await globalThis.ProxmoxWASM.lxc.PostLXC(node, data)
        );
    }

    async startLXC(node: string, vmid: number): Promise<TaskBaseResponse> {
        return TaskBaseResponseSchema.parse(
            await globalThis.ProxmoxWASM.lxc.StartLXC(node, vmid)
        );
    }

    async stopLXC(node: string, vmid: number): Promise<TaskBaseResponse> {
        return TaskBaseResponseSchema.parse(
            await globalThis.ProxmoxWASM.lxc.StopLXC(node, vmid)
        );
    }

    async deleteLXC(node: string, vmid: number): Promise<TaskBaseResponse> {
        return TaskBaseResponseSchema.parse(
            await globalThis.ProxmoxWASM.lxc.DeleteLXC(node, vmid)
        );
    }

    async cloneLXC(node: string, vmid: number, data: CloneLXCData): Promise<TaskBaseResponse> {
        return TaskBaseResponseSchema.parse(
            await globalThis.ProxmoxWASM.lxc.CloneLXC(node, vmid, data)
        );
    }
}