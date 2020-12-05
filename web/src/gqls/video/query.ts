import gql from 'graphql-tag';

const LIST_VIDEO = gql`
 query videos($searchParam:SearchParam!) {
   videos(searchParam: $searchParam){
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
            isHideOnMobile
            theme
            tags
            createdAt
            updatedAt
       }
   }
  }
`;

const LIST_VIDEO_CARD = gql`
 query searchVideo($searchParam:VideoSearchParam!) {
   searchVideo(searchParam:$searchParam){
       edges{
            id
            title
            desc
            cover
            totalNum
            episodeID
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
 query searchVideo($searchParam:VideoSearchParam!) {
   searchVideo(searchParam:$searchParam){
    aggResults{
         key
         value
       }
   }
  }
`;

const VIDEO_RANDOM_TAG_SUGGEST = gql`
    query searchVideo($searchParam:VideoSearchParam!) {
   searchVideo(searchParam:$searchParam){
       edges{
            id
            title
            desc
            cover
            totalNum
            episodeID
       }
   }
  }
`;

const SIMILAR_VIDEOS = gql`
 query similarVideos($searchParam:VideoSimilarParam!) {
   similarVideos(searchParam:$searchParam){
       edges{
            id
            title
            desc
            cover
            totalNum
            episodeID
       }
   }
  }
`;

const VIDEO_COMBO = gql`
 query videos($searchParam:SearchParam!,$isFilterVideoSeries:Boolean=true) {
   videos(searchParam:$searchParam,isFilterVideoSeries:$isFilterVideoSeries){
       edges{
          value:id 
          text:title 
       }
   }
  }
`;

const GET_VIDEO_DETAIL = gql`
query videoDetail($episodeID: ID!) {
  videoDetail(episodeID: $episodeID) {
    video {
      id
      title
      desc
      cover
      pubDate
      episodes {
        id
        num
        title
        desc
        cover
        url
        subtitles {
          name
          url
        }
      }
      tags
      theme
    }
    videoSerieses {
      id
      name
      items {
        videoID
        videoSeriesID
        alias
        episodeID
      }
    }
    historyInfo {
      subSourceID
      currentTime
      remainingTime
      updatedAt
    }
  }
}
`;

const LIST_VIDEO_SERIES = gql`
 query videoSerieses($searchParam:SearchParam!) {
  videoSerieses(searchParam:$searchParam){
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

export {
  LIST_VIDEO, VIDEO_COMBO, LIST_VIDEO_CARD, GET_VIDEO_DETAIL,
  GET_VIDEO_TAGS, VIDEO_RANDOM_TAG_SUGGEST, SIMILAR_VIDEOS,
  LIST_VIDEO_SERIES
}