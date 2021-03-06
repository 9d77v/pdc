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
 query searchVideo($searchParam:SearchParam!) {
   searchVideo(searchParam:$searchParam){
       edges{
            id
            title
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
 query searchVideo($searchParam:SearchParam!) {
   searchVideo(searchParam:$searchParam){
    aggResults{
         key
         value
       }
   }
  }
`;

const VIDEO_RANDOM_TAG_SUGGEST = gql`
    query searchVideo($searchParam:SearchParam!) {
   searchVideo(searchParam:$searchParam){
       edges{
            id
            title
            cover
            totalNum
            episodeID
       }
   }
  }
`;

const VIDEO_COMBO = gql`
 query videos($searchParam:SearchParam!,$filterVideosInVideoSeries:Boolean=true) {
   videos(searchParam:$searchParam,filterVideosInVideoSeries:$filterVideosInVideoSeries){
       edges{
          value:id 
          text:title 
       }
   }
  }
`

const GET_VIDEO_DETAIL = gql`
query videoDetail(
  $similiarVideoParam: SearchParam!
  $searchParam: SearchParam!
  $episodeID: ID!
) {
  histories(sourceType: 1, searchParam: $searchParam, subSourceID: $episodeID) {
    edges {
      subSourceID
      currentTime
      remainingTime
      updatedAt
    }
  }
  videos(searchParam: $searchParam, episodeID: $episodeID) {
    edges {
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
  }
  videoSerieses(searchParam: $searchParam, episodeID: $episodeID) {
    edges {
      id
      name
      items {
        videoID
        videoSeriesID
        alias
        episodeID
      }
    }
  }
  similarVideos(searchParam: $similiarVideoParam, episodeID: $episodeID) {
    edges {
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
  GET_VIDEO_TAGS, VIDEO_RANDOM_TAG_SUGGEST,
  LIST_VIDEO_SERIES
}
