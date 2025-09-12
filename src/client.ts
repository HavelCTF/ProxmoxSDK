import axios from 'axios'
import type { AxiosInstance } from 'axios'
import * as https from 'https'
import { Effect } from 'effect'

export class ProxmoxClient {
    private axiosInstance : AxiosInstance
    private uuid : string

    constructor (baseUrl : string, apiToken : string, uuid : string) {
        this.axiosInstance = axios.create({
            baseURL: baseUrl + '/api2/json',
            headers: {
                'Authorization': apiToken
            },
            httpsAgent: new https.Agent({ rejectUnauthorized: false })
        })
        this.uuid = uuid
    }

    public version = () => Effect.tryPromise({
         try: async () => {
            const response = await this.axiosInstance.get('/version')
            return response.data
         },
         catch: (error) => {
            return new Error(`[${this.uuid}] Version request failed: ${error}`)
         }
    }).pipe(
        Effect.retry({ times: 3 }),
        Effect.timeout(10000),

        Effect.tapError((error) => 
            Effect.sync(() =>
                console.error(error.message)
        ))
    )
}
