import { beforeEach, describe, expect, test, vi, type Mock } from "vitest";
import { ProxmoxClient } from "../client.js";
import "./version.js";
import axios from "axios";
import { Effect } from "effect";

vi.mock("axios");

describe("Version Module", () => {
  let client: ProxmoxClient;
  const baseUrl = "https://proxmox.example.com";
  const api = "my-api-token";
  const uuid = "my-uuid";

  beforeEach(() => {
    client = new ProxmoxClient(baseUrl, api, uuid);
  });

  test("should initialize axios instance with correct configuration", () => {
    expect(axios.create).toHaveBeenCalledWith({
      baseURL: `${baseUrl}/api2/json`,
      headers: {
        Authorization: api,
      },
      httpsAgent: expect.any(Object),
    });
  });

  test("should have version method", () => {
    expect(typeof client.version).toBe("function");
  });

  test("version method should make a GET request to /version", async () => {
    const mockGet = vi.fn().mockResolvedValue({ data: { version: "1.0.0" } });
    (axios.create as Mock).mockReturnValue({ get: mockGet });

    // Create a new client instance after setting up the mock
    const testClient = new ProxmoxClient(baseUrl, api, uuid);

    const result = await testClient.version().pipe(Effect.runPromise);

    expect(mockGet).toHaveBeenCalledWith("/version");
    expect(result).toEqual({ version: "1.0.0" });
  });
});
