import axios from 'axios'
import type { AxiosInstance } from 'axios'
import * as https from 'https'
import { ProxmoxVersion } from "./version/version.js"

export class ProxmoxClient {
    private axiosInstance : AxiosInstance
    private uuid : string

    public version : ProxmoxVersion

    constructor (baseUrl : string, apiToken : string, uuid : string) {
        this.axiosInstance = axios.create({
            baseURL: baseUrl + '/api2/json',
            headers: {
                'Authorization': apiToken
            },
            httpsAgent: new https.Agent({ rejectUnauthorized: false })
        })
        this.uuid = uuid
        this.version = new ProxmoxVersion(this.axiosInstance, this.uuid)
    }
}
