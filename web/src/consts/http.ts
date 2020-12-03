import { isMobile } from "src/utils/util";
import jwt_decode from "jwt-decode";
import dayjs from "dayjs";

//getUploadURL 获取minio上传地址
export const getUploadURL = async (bucketName: String, fileName: String) => {
    const body = JSON.stringify({
        operationName: "presignedUrl",
        query: ` query presignedUrl($bucketName: String!,$objectName:String!) {
 \n    presignedUrl(bucketName: $bucketName, objectName: $objectName)
 \n  }`,
        variables: {
            "bucketName": bucketName,
            "objectName": fileName
        }
    })
    return await request(body, true)
}

export const getRefreshToken = async (refreshToken: String) => {
    const body = JSON.stringify({
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
    return await request(body, false)
}

export const recordHistory = async (
    sourceType: number,
    sourceID: number,
    subSourceID: number,
    currentTime: number,
    remainingTime: number) => {
    const body = JSON.stringify({
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
                "platform": isMobile() ? "mobile" : "desktop",
                "currentTime": currentTime,
                "remainingTime": remainingTime
            }
        }
    })
    return await request(body, true)
}


const request = async (body: string, needToken?: boolean): Promise<any> => {
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    if (needToken) {
        let token = localStorage.getItem('accessToken') || "";
        if (token !== "") {
            const accessToken: any = jwt_decode(token)
            if (Number(accessToken.exp) - dayjs().unix() <= 0) {
                const refreshToken = localStorage.getItem('refreshToken') || "";
                const data = await getRefreshToken(refreshToken)
                if (data.data) {
                    const refreshData: any = data.data.refreshToken
                    localStorage.setItem("accessToken", refreshData.accessToken)
                    localStorage.setItem("refreshToken", refreshData.refreshToken)
                    token = refreshData.accessToken
                }
            }
        }
        myHeaders.append("Authorization", token ? `Bearer ${token}` : "");
    }
    const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: body,
        redirect: 'follow' as const
    };
    const data = await fetch("/api", requestOptions)
    return data.json()
}