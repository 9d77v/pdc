import gql from 'graphql-tag';

const LIST_VIDEO = gql`
{
   listVideo{
       id
       title
       desc
       cover
       pubDate
       isShow
       episodes{
        id
       num
       title
       desc
       cover
       url
       subtitles{
          name
          url
       }
       createdAt
       updatedAt
       }
       createdAt
       updatedAt
   }
}
`;

const ADD_VIDEO = gql`
mutation createVideo($input:NewVideo!){
   createVideo(input:$input)
}
`

const ADD_EPISODE = gql`
mutation createEpisode($input:NewEpisode!){
   createEpisode(input:$input)
}
`

//getUploadURL 获取minio上传地址
const getUploadURL = async (bucketName: String, fileName: String) => {
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

export { LIST_VIDEO, ADD_VIDEO, ADD_EPISODE, getUploadURL }