import gql from 'graphql-tag';

const LIST_MEDIA = gql`
{
   listMedia{
       id
       title
       desc
       cover
       pubDate
       episodes{
        id
       order
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

const ADD_MEDIA = gql`
mutation createMedia($input:NewMedia!){
   createMedia(input:$input)
}
`

const ADD_EPISODE = gql`
mutation createEpisode($input:NewEpisode!){
   createEpisode(input:$input)
}
`

export { LIST_MEDIA, ADD_MEDIA, ADD_EPISODE }