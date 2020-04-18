
//getUploadURL 获取minio上传地址
export const getUploadURL = async (bucketName: String, fileName: String) => {
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    const graphql = JSON.stringify({
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
