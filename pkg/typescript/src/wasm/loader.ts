export class WasmLoader {
    private static loaded = false;
    private static goExited: Promise<void> | null = null;

    public static async init(source: string | Buffer | Uint8Array): Promise<void> {
        if (this.loaded) return;

        if (typeof (globalThis as any).Go === 'undefined') {
            throw new Error('wasm_exec.js must be loaded before initializing the SDK');
        }

        const go = new globalThis.Go();
        const isBrowser = typeof window !== 'undefined' && typeof source === 'string';

        let instance: WebAssembly.Instance;
        try {
            if (isBrowser) {
                const result = await WebAssembly.instantiateStreaming(await fetch(source), go.importObject);
                instance = result.instance;
            } else {
                const result = await WebAssembly.instantiate(source as BufferSource, go.importObject);
                instance = (result as WebAssembly.WebAssemblyInstantiatedSource).instance;
            }
        } catch (err) {
            throw new Error(`Failed to load WASM binary: ${err instanceof Error ? err.message : err}`);
        }

        this.goExited = go.run(instance).catch((err: unknown) => {
            console.error('[ProxmoxSDK] Go runtime exited unexpectedly:', err);
        });

        this.loaded = true;
    }
}