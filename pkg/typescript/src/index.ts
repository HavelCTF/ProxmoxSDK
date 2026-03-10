import { WasmLoader } from './wasm/loader';
import { ClusterModule } from './modules/cluster';

export * from './types/cluster';

export class ProxmoxSDK {
    public cluster: ClusterModule;

    private constructor() {
        this.cluster = new ClusterModule();
    }

    public static async create(host: string, token: string, uuid: string, wasmSource: string | Buffer | Uint8Array): Promise<ProxmoxSDK> {
        await WasmLoader.init(wasmSource);

        if (!globalThis.ProxmoxWASM) {
            throw new Error("WASM module failed to initialize");
        }

        const ok = globalThis.ProxmoxWASM.initClient(host, token, uuid);
        if (!ok) {
            throw new Error("Go client initialization failed");
        }

        return new ProxmoxSDK();
    }
}