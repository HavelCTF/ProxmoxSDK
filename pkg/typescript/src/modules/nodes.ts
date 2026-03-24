import { type NodesResponse, NodesResponseSchema } from '../types/nodes';

export class NodesModule {
    async getNodes(): Promise<NodesResponse> {
        return NodesResponseSchema.parse(await globalThis.ProxmoxWASM.nodes.GetNodes());
    }
}
