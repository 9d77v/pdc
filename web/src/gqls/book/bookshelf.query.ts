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

const BOOKSHELF_DETAIL = gql`
 query bookshelfDetail($searchParam:SearchParam!, $bookshelfID:Int!) {
   bookshelfs(searchParam: $searchParam){
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
   bookPositions(bookshelfID: $bookshelfID){
       edges{
        id
        bookID
        book{
          name 
          cover
        }
        layer
        partition
        prevID
       }
   }
  }
`

const APP_BOOKSHELF_DETAIL = gql`
 query bookshelfDetail($searchParam:SearchParam!,$bookshelfID:Int!) {
   bookshelfs(searchParam: $searchParam){
        edges{
            id
            name
            cover
            layerNum
            partitionNum
       }
   }
   bookPositions(bookshelfID: $bookshelfID){
       edges{
        id
        bookID
        book{
          name 
          cover
        }
        layer
        partition
        prevID
       }
   }
  }
`

export {
  LIST_BOOKSHELF, BOOKSHELF_DETAIL, APP_BOOKSHELF_DETAIL
}
