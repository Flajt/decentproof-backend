//Based on: https://www.scaleway.com/en/docs/compute/functions/reference-content/code-examples/#using-event-components
export type FunctionEvent = {
    pathParameters: Object
    queryStringParameters: Object
    body: string | Buffer // Might not be buffer but thats my best guess
    headers: any,
    method: string,
    isBase64Encoded: boolean
}