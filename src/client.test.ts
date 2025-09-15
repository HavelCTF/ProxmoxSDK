import { beforeEach, describe, expect, test, vi } from "vitest";
import { ProxmoxClient } from "./client.js";
import axios from "axios";

vi.mock("axios");

describe("ProxmoxClient", () => {
  let client: ProxmoxClient;
  const baseUrl = "https://proxmox.example.com";
  const apiToken = "my-api-token";
  const uuid = "my-uuid";

  beforeEach(() => {
    client = new ProxmoxClient(baseUrl, apiToken, uuid);
  });

  test("should initialize axios instance with correct configuration", () => {
    expect(axios.create).toHaveBeenCalledWith({
      baseURL: `${baseUrl}/api2/json`,
      headers: {
        Authorization: apiToken,
      },
      httpsAgent: expect.any(Object),
    });
  });

  test("should set the uuid correctly", () => {
    // biome-ignore lint/suspicious/noExplicitAny: because we cannot test it properly otherwise
    expect((client as unknown as any).uuid).toBe(uuid);
  });
});
