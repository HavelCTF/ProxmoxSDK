import { ClusterNextIdResponse } from './cluster';

declare global {
    var ProxmoxWASM: {
        initClient: (host: string, token: string, uuid: string) => boolean;
        cluster: {
            GetNextId: () => Promise<ClusterNextIdResponse>;
        }
    };
    
    // Déclare la présence de la variable globale injectée par wasm_exec.js
    var Go: {
        new (): {
            importObject: any;
            run(instance: WebAssembly.Instance): Promise<void>;
        };
    };
}

export {};