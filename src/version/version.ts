import type { AxiosInstance } from 'axios'
import { Effect } from 'effect'


export class ProxmoxVersion {
    private versionInstance : AxiosInstance
    private uuid : string

    constructor(baseInstance : AxiosInstance, id : string) {
        this.versionInstance = baseInstance
        this.uuid = id
    }

    public version = () => Effect.tryPromise({
         try: async () => {
            const response = await this.versionInstance.get('/version')
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
