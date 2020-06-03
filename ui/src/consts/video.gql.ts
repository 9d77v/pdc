import gql from 'graphql-tag';

const ADD_VIDEO = gql`
mutation createVideo($input:NewVideo!){
   createVideo(input:$input){
     id
   }
}
`
const UPDATE_VIDEO = gql`
mutation updateVideo($input:NewUpdateVideo!){
   updateVideo(input:$input){
     id
   }
}
`

const ADD_EPISODE = gql`
mutation createEpisode($input:NewEpisode!){
   createEpisode(input:$input){
     id
   }
}
`

const UPDATE_EPISODE = gql`
mutation updateEpisode($input:NewUpdateEpisode!){
   updateEpisode(input:$input){
     id
   }
}
`
const LIST_VIDEO = gql`
 query Videos( $page: Int, $pageSize: Int) {
   Videos(page: $page, pageSize: $pageSize){
       totalCount
       edges{
             id
            title
            desc
            cover
            pubDate
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
            isShow
            tags
            createdAt
            updatedAt
       }
   }
  }
`;

const LIST_VIDEO_CARD = gql`
 query Videos( $page: Int, $pageSize: Int) {
   Videos(page: $page, pageSize: $pageSize){
       edges{
            id
            title
            cover
            episodes{
              id
            }
       }
   }
  }
`;

export { LIST_VIDEO, ADD_VIDEO, UPDATE_VIDEO, ADD_EPISODE, UPDATE_EPISODE, LIST_VIDEO_CARD }