output: {
    type: "raw"
    properties: {
        apiVersion: "v1"
        kind:       "ConfigMap"
        metadata: {
            name: "redisconfig"
            namespace: "redis"
        }
        data: input: parameter.redisconfig
    }
}