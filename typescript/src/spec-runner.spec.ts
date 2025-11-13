import { readFileSync, readdirSync } from "node:fs";
import { join, dirname, basename } from "node:path";
import { fileURLToPath } from "node:url";
import { expect, test, describe, beforeEach, afterEach } from "vitest";
import { Effect } from "effect";
import axios from "axios";
import MockAdapter from "axios-mock-adapter";
import { ProxmoxClient } from "./client.js";
import "./version/version.js";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

interface BaseTestCase {
	name: string;
	description: string;
	input: {
		baseURL: string;
		apiToken: string;
		uuid: string;
	};
}

interface ClientTestCase extends BaseTestCase {
	expected: {
		baseURL: string;
		uuid: string;
	};
}

interface MockConfig {
	statusCode?: number;
	body?: Record<string, unknown> | string;
	retries?: Array<{
		attempt: number;
		statusCode: number;
		body?: Record<string, unknown>;
	}>;
}

interface VersionTestCase extends BaseTestCase {
	mock: MockConfig;
	expected?: {
		version: string;
		release: string;
		repoid: string;
	};
	expectError?: boolean;
	errorContains?: string;
	expectedAttempts?: number;
}

function loadAllSpecs(): Map<string, unknown[]> {
	const specsDir = join(__dirname, "../../test-specs");
	const specs = new Map<string, unknown[]>();

	const files = readdirSync(specsDir).filter((f) => f.endsWith(".json"));

	for (const file of files) {
		const specName = basename(file, ".json");
		const specPath = join(specsDir, file);
		const raw = readFileSync(specPath, "utf8");
		specs.set(specName, JSON.parse(raw));
	}

	return specs;
}

const allSpecs = loadAllSpecs();

// Run client specs
if (allSpecs.has("client")) {
	describe("Client Spec Tests", () => {
		const cases = allSpecs.get("client") as ClientTestCase[];

		for (const testCase of cases) {
			test(testCase.name, () => {
				const client = new ProxmoxClient(
					testCase.input.baseURL,
					testCase.input.apiToken,
					testCase.input.uuid,
				);

				expect(client).toBeDefined();

				// biome-ignore lint/suspicious/noExplicitAny: needed for testing internal properties
				const clientAny = client as any;

				expect(clientAny.uuid).toBe(testCase.expected.uuid);
				expect(clientAny.axiosInstance.defaults.baseURL).toBe(
					testCase.expected.baseURL,
				);
			});
		}
	});
}

// Run version specs
if (allSpecs.has("version")) {
	describe("Version Spec Tests", () => {
		let mock: MockAdapter;

		beforeEach(() => {
			mock = new MockAdapter(axios);
		});

		afterEach(() => {
			mock.restore();
		});

		const cases = allSpecs.get("version") as VersionTestCase[];

		for (const testCase of cases) {
			test(testCase.name, async () => {
				const client = new ProxmoxClient(
					testCase.input.baseURL,
					testCase.input.apiToken,
					testCase.input.uuid,
				);

				// Handle retry scenarios
				if (testCase.mock.retries) {
					let callCount = 0;
					mock.onGet(/\/api2\/json\/version/).reply(() => {
						callCount++;
						const response = testCase.mock.retries?.find(
							(r) => r.attempt === callCount,
						);
						if (response) {
							return [
								response.statusCode,
								response.body || "",
								{ "Content-Type": "application/json" },
							];
						}
						return [500, "Unexpected call"];
					});

					if (testCase.expectError) {
						try {
							await Effect.runPromise(client.version());
							throw new Error("Expected version() to throw but it succeeded");
						} catch (error: unknown) {
							const err = error as Error;
							if (testCase.errorContains) {
								expect(err.message).toContain(testCase.errorContains);
							}
						}
					} else {
						const result = await Effect.runPromise(client.version());
						expect(result).toEqual(testCase.expected);
					}

					if (testCase.expectedAttempts) {
						expect(callCount).toBe(testCase.expectedAttempts);
					}
				} else {
					// Single response scenario
					mock
						.onGet(/\/api2\/json\/version/)
						.reply(testCase.mock.statusCode || 200, testCase.mock.body, {
							"Content-Type": "application/json",
						});

					if (testCase.expectError) {
						try {
							await Effect.runPromise(client.version());
							throw new Error("Expected version() to throw but it succeeded");
						} catch (error: unknown) {
							const err = error as Error;
							if (testCase.errorContains) {
								expect(err.message).toContain(testCase.errorContains);
							}
						}
					} else {
						const result = await Effect.runPromise(client.version());
						expect(result).toEqual(testCase.expected);
					}
				}
			});
		}
	});
}
