// SPDX-License-Identifier: MIT
// Copyright (c) 2026 HavelCTF

import { type VersionResponse, VersionResponseSchema } from '../types/version';

export async function getVersion(): Promise<VersionResponse> {
    return VersionResponseSchema.parse(await globalThis.ProxmoxWASM.version.GetVersion());
}
