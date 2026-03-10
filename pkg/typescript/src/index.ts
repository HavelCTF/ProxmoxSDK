import { WasmLoader } from './wasm/loader';
import { ClusterModule } from './modules/cluster';

export * from './types/cluster';

export class ProxmoxSDK {
    public cluster: ClusterModule;
    private host: string;
    private token: string;

    private constructor(host: string, token: string) {
        this.host = host;
        this.token = token;
        this.cluster = new ClusterModule(host, token);
    }

    /**
     * @param host L'URL de l'API (ex: https://10.0.0.1:8006)
     * @param token Le token d'API
     * @param uuid L'identifiant unique
     * @param wasmSource Le Buffer (Node.js) ou l'URL (Navigateur) du fichier main_wasm.wasm
     */
    public static async create(host: string, token: string, uuid: string, wasmSource: string | Buffer | Uint8Array): Promise<ProxmoxSDK> {
        // 1. Charge WASM
        await WasmLoader.init(wasmSource);

        // 2. Initialise le client Go
        if (!globalThis.ProxmoxWASM) {
            throw new Error("[ProxmoxSDK] Le client interne Go n'est pas disponible.");
        }
        const isInit = globalThis.ProxmoxWASM.initClient(host, token, uuid);
        if (!isInit) {
            throw new Error("[ProxmoxSDK] Échec de l'initialisation du client interne Go.");
        }

        // 3. Retourne l'instance du SDK
        return new ProxmoxSDK(host, token);
    }
}