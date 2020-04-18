import gql from 'graphql-tag';

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
       createdAt
       updatedAt
       }
   }
  }
`;

export { LIST_VIDEO, ADD_VIDEO, ADD_EPISODE }