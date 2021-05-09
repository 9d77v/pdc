import gql from 'graphql-tag';


const LIST_BOOKSHELF = gql`
 query bookshelfs($searchParam:SearchParam!) {
   bookshelfs(searchParam: $searchParam){
        totalCount
        edges{
            id
            name
            cover
            layerNum
            partitionNum
            createdAt
            updatedAt
       }
   }
  }
`

export {
  LIST_BOOKSHELF
}
