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

const UPDATE_SUBTITLE = gql`
mutation updateSubtitle($input:NewUpdateSubtitles!){
  updateSubtitle(input:$input){
     id
   }
}
`
const LIST_VIDEO = gql`
 query videos($keyword:String, $page: Int, $pageSize: Int, $sorts: [Sort!]) {
   videos(keyword:$keyword,page: $page, pageSize: $pageSize,sorts:$sorts){
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
 query videos( $page: Int, $pageSize: Int, $sorts: [Sort!]) {
   videos(page: $page, pageSize: $pageSize,sorts:$sorts){
       edges{
            id
            title
            cover
            desc
            episodes{
              id
            }
       }
   }
  }
`;

const GET_VIDEO = gql`
 query videos( $ids: [ID!]) {
   videos(ids:$ids){
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
            }
            tags
       }
   }
  }
`;

export {
  LIST_VIDEO, ADD_VIDEO, UPDATE_VIDEO, ADD_EPISODE,
  UPDATE_EPISODE, LIST_VIDEO_CARD, GET_VIDEO, UPDATE_SUBTITLE
}