export async function loadWasm(source: string | Buffer | Uint8Array): Promise<void> {
    if (typeof (globalThis as any).Go === 'undefined') {
        throw new Error('wasm_exec.js must be loaded before calling loadWasm()');
    }

    const go = new globalThis.Go();
    const isBrowser = typeof window !== 'undefined' && typeof source === 'string';

    const { instance } = isBrowser
        ? await WebAssembly.instantiateStreaming(await fetch(source), go.importObject)
        : await WebAssembly.instantiate(source as BufferSource, go.importObject);

    go.run(instance);
}
