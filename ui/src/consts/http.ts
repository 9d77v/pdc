
//getUploadURL 获取minio上传地址
export const getUploadURL = async (bucketName: String, fileName: String) => {
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    const token = localStorage.getItem('accessToken') || "";
    myHeaders.append("Authorization", token ? `Bearer ${token}` : "");

    const graphql = JSON.stringify({
        operationName: "presignedUrl",
        query: ` query presignedUrl($bucketName: String!,$objectName:String!) {
 \n    presignedUrl(bucketName: $bucketName, objectName: $objectName)
 \n  }`,
        variables: {
            "bucketName": bucketName,
            "objectName": fileName
        }
    })
    const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: graphql,
        redirect: 'follow' as const
    };
    const data = await fetch("/api", requestOptions)
    return data.json()
}

export const getRefreshToken = async (refreshToken: String) => {
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    const graphql = JSON.stringify({
        operationName: "refreshToken",
        query: `mutation refreshToken($refreshToken:String!){
\n            refreshToken(refreshToken:$refreshToken){
\n               accessToken
\n               refreshToken
\n            }
\n         }`,
        variables: {
            "refreshToken": refreshToken
        }
    })
    const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: graphql,
        redirect: 'follow' as const
    };
    const data = await fetch("/api", requestOptions)
    return data.json()
}

export const recordHistory = async (
    sourceType: number,
    sourceID: number,
    subSourceID: number,
    currentTime: number,
    remainingTime: number) => {
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    const token = localStorage.getItem('accessToken') || "";
    myHeaders.append("Authorization", token ? `Bearer ${token}` : "");

    const graphql = JSON.stringify({
        operationName: "recordHistory",
        query: `mutation recordHistory($input:NewHistoryInput!){
\n            recordHistory(input:$input){
\n               subSourceID
\n            }
\n         }`,
        variables: {
            "input": {
                "sourceType": sourceType,
                "sourceID": sourceID,
                "subSourceID": subSourceID,
                "currentTime": currentTime,
                "remainingTime": remainingTime
            }
        }
    })
    const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: graphql,
        redirect: 'follow' as const
    };
    const data = await fetch("/api", requestOptions)
    return data.json()
}