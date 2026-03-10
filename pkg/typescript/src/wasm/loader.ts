export class WasmLoader {
    private static isLoaded = false;

    public static async init(wasmSource: string | Buffer | Uint8Array): Promise<void> {
        if (this.isLoaded) return;

        // Bypass TypeScript pour vérifier l'existence sur globalThis
        if (typeof (globalThis as any).Go === 'undefined') {
            throw new Error("[ProxmoxSDK] L'objet 'Go' est introuvable. Importez 'wasm_exec.js' avant d'initialiser le SDK.");
        }

        // On peut maintenant instancier Go sereinement
        const go = new globalThis.Go();
        let instance: WebAssembly.Instance;

        try {
            console.debug(`[WasmLoader] Chargement du WASM depuis source : ${typeof wasmSource === 'string' ? wasmSource : 'Buffer/Uint8Array'}`);
            if (typeof window !== 'undefined' && typeof wasmSource === 'string') {
                console.log(`[WasmLoader] Chargement du WASM depuis URL : ${wasmSource}`);
                const response = await fetch(wasmSource);
                const wasmObj = await WebAssembly.instantiateStreaming(response, go.importObject);
                instance = wasmObj.instance;
                go.run(instance);
                this.isLoaded = true;
            } else {
                console.log(`[WasmLoader] Chargement du WASM depuis Buffer/Uint8Array...`);
                const wasmObj = await WebAssembly.instantiate(wasmSource as ArrayBuffer | Uint8Array, go.importObject) as any;
                instance = wasmObj.instance;
                go.run(instance);
                this.isLoaded = true;
            }
        } catch (error) {
            throw new Error(`[ProxmoxSDK] Échec du chargement WASM : ${error}`);
        }
    }
}