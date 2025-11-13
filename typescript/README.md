# TypeScript SDK

TypeScript implementation of the ProxmoxSDK.

## Installation

```bash
npm install
```

## Building

```bash
npm run build
```

## Testing

```bash
npm test
```

## Usage

```typescript
import { ProxmoxClient } from 'proxmoxsdk';
import 'proxmoxsdk/version';

const client = new ProxmoxClient(
  'https://proxmox.example.com',
  'PVEAPIToken=user@pam!token=...',
  'unique-client-id'
);

const version = await client.version();
console.log(version);
```

## Development

```bash
# Run in development mode
npm run dev

# Clean build artifacts
npm run clean

# Lint code
npx @biomejs/biome check .

# Format code
npx @biomejs/biome format --write .
```
