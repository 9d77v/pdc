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
const UPDATE_VIDEO_RESOURCE = gql`
mutation updateVideoResource($input:NewVideoResource!){
   updateVideoResource(input:$input){
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
  ADD_VIDEO, ADD_VIDEO_RESOURCE, UPDATE_VIDEO_RESOURCE,
  SAVE_SUBTITLES, UPDATE_VIDEO, ADD_EPISODE,
  UPDATE_EPISODE, ADD_VIDEO_SERIES, UPDATE_VIDEO_SERIES,
  ADD_VIDEO_SERIES_ITEM, UPDATE_VIDEO_SERIES_ITEM
}
