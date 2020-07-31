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

const UPDATE_MOBILE_VIDEO = gql`
mutation updateMobileVideo($input:NewUpdateMobileVideos!){
  updateMobileVideo(input:$input){
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
 query videos(  $page: Int, $pageSize: Int, $sorts: [Sort!]) {
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

const VIDEO_COMBO = gql`
 query videos($keyword:String, $page: Int, $pageSize: Int, $sorts: [Sort!],$isFilterVideoSeries:Boolean=true) {
   videos(keyword:$keyword,page: $page, pageSize: $pageSize,sorts:$sorts,isFilterVideoSeries:$isFilterVideoSeries){
       edges{
          value:id 
          text:title 
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
              mobileURL
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

const LIST_VIDEO_SERIES = gql`
 query videoSerieses($keyword:String, $page: Int, $pageSize: Int, $sorts: [Sort!]) {
  videoSerieses(keyword:$keyword,page: $page, pageSize: $pageSize,sorts:$sorts){
       totalCount
       edges{
            id
            name
            items{
              videoSeriesID
              videoID
              title
              alias
              num
            }
            createdAt
            updatedAt
       }
   }
  }
`;

const ADD_VIDEO_SERIES = gql`
mutation createVideoSeries($input:NewVideoSeries!){
   createVideoSeries(input:$input){
     id
   }
}
`
const UPDATE_VIDEO_SERIES = gql`
mutation updateVideoSeries($input:NewUpdateVideoSeries!){
   updateVideoSeries(input:$input){
     id
   }
}
`

const ADD_VIDEO_SERIES_ITEM = gql`
mutation createVideoSeriesItem($input:NewVideoSeriesItem!){
  createVideoSeriesItem(input:$input){
     videoID
     videoSeriesID
   }
}
`

const UPDATE_VIDEO_SERIES_ITEM = gql`
mutation updateVideoSeriesItem($input:NewUpdateVideoSeriesItem!){
  updateVideoSeriesItem(input:$input){
     videoID
     videoSeriesID
   }
}
`

export {
  LIST_VIDEO, VIDEO_COMBO, ADD_VIDEO, UPDATE_VIDEO, ADD_EPISODE,
  UPDATE_EPISODE, LIST_VIDEO_CARD, GET_VIDEO, UPDATE_SUBTITLE,
  UPDATE_MOBILE_VIDEO,
  LIST_VIDEO_SERIES, ADD_VIDEO_SERIES, UPDATE_VIDEO_SERIES,
  ADD_VIDEO_SERIES_ITEM, UPDATE_VIDEO_SERIES_ITEM
}