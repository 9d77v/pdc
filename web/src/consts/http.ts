import { isMobile } from "src/utils/util";
import jwt_decode from "jwt-decode";
import dayjs from "dayjs";

export const getServerURL=()=>{
    if (document.location.protocol === "http:") {
        return `http://${document.location.host}/api`
    }
    return process.env.REACT_APP_SERVER_URL+"/api"
}

//getUploadURL 获取minio上传地址
export const getUploadURL = async (bucketName: String, fileName: String) => {
    const body = JSON.stringify({
        operationName: "presignedUrl",
        query: ` query presignedUrl($bucketName: String!,$objectName:String!) {
 \n    presignedUrl(bucketName: $bucketName, objectName: $objectName){
    \n               ok
    \n               url
    \n            }
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
    remainingTime: number,
    duration: number,
    clientTs: number) => {
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
                "remainingTime": remainingTime,
                "duration": duration,
                "clientTs": clientTs
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
        redirect: 'follow' as const,
        mode: "cors" as const,
        keepalive: true
    };
    const data = await fetch(getServerURL(), requestOptions)
    return data.json()
}
