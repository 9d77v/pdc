import gql from 'graphql-tag';

const ADD_VIDEO = gql`
mutation createVideo($input:NewVideo!){
   createVideo(input:$input){
     id
   }
}
`

const ADD_VIDEO_RESOURCE = gql`
mutation addVideoResource($input:NewVideoResource!){
   addVideoResource(input:$input){
     id
   }
}
`

const SAVE_SUBTITLES = gql`
mutation saveSubtitles($input:NewSaveSubtitles!){
  saveSubtitles(input:$input){
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
            theme
            tags
            createdAt
            updatedAt
       }
   }
  }
`;

const LIST_VIDEO_CARD = gql`
 query searchVideo($keyword:String,$tags:[String!],  $page: Int, $pageSize: Int) {
   searchVideo(keyword:$keyword,tags:$tags,page: $page, pageSize: $pageSize){
       edges{
            id
            title
            desc
            cover
            totalNum
       }
       totalCount
       aggResults{
         key
         value
       }
   }
  }
`;

const GET_VIDEO_TAGS = gql`
 query searchVideo($page:Int=1) {
   searchVideo(page: $page){
    aggResults{
         key
         value
       }
   }
  }
`;

const VIDEO_RANDOM_TAG_SUGGEST = gql`
 query searchVideo($tag:String!,$pageSize:Int!) {
   searchVideo(tags:[$tag],page: 1, pageSize:$pageSize,isRandom:true){
       edges{
            id
            title
            desc
            cover
            totalNum
       }
   }
  }
`;

const SIMILAR_VIDEOS = gql`
 query similarVideos($videoID:ID!,$pageSize:Int!) {
   similarVideos(videoID:$videoID, pageSize:$pageSize,){
       edges{
            id
            title
            desc
            cover
            totalNum
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
 query videos($videoID:ID!,$sourceType:Int!=1) {
   videos(ids:[$videoID]){
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
            theme
       }
   }
   videoSerieses(videoID:$videoID){
       edges{
            id
            name
            items{
              videoID
              videoSeriesID
              alias
            }
       }
   }
   historyInfo(sourceType:$sourceType,sourceID:$videoID){
      subSourceID
      currentTime
      remainingTime
      updatedAt
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
  ADD_VIDEO, ADD_VIDEO_RESOURCE, SAVE_SUBTITLES,
  LIST_VIDEO, VIDEO_COMBO, UPDATE_VIDEO, ADD_EPISODE,
  UPDATE_EPISODE, LIST_VIDEO_CARD, GET_VIDEO, GET_VIDEO_TAGS,
  UPDATE_MOBILE_VIDEO, VIDEO_RANDOM_TAG_SUGGEST, SIMILAR_VIDEOS,
  LIST_VIDEO_SERIES, ADD_VIDEO_SERIES, UPDATE_VIDEO_SERIES,
  ADD_VIDEO_SERIES_ITEM, UPDATE_VIDEO_SERIES_ITEM
}