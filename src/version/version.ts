import axios from 'axios'
import type { AxiosResponse, AxiosInstance } from 'axios'

export class ProxmoxVersion {
    private versionInstance : AxiosInstance
    private uuid : string

    constructor(baseInstance : AxiosInstance, id : string) {
        this.versionInstance = baseInstance
        this.uuid = id
    }

    public async version() : Promise<AxiosResponse<any, any> | undefined> {
        try {
            const response = await this.versionInstance.get('/version')
            return response.data
        } catch (error) {
            if (axios.isAxiosError(error)) {
                console.error('[%s] Version request failed: %s', this.uuid, error)
            } else {
                console.error()
            }
            return undefined
        }
    }
}
