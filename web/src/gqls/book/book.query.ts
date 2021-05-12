import gql from 'graphql-tag';


const LIST_BOOK = gql`
 query books($searchParam:SearchParam!) {
   books(searchParam: $searchParam){
        totalCount
        edges{
            id
            isbn
            name
            desc
            cover
            author
            translator
            publishingHouse
            createdAt
            updatedAt
       }
   }
  }
`

const BOOK_COMBO = gql`
 query books($searchParam:SearchParam!,$filterBooksInBookPositions:Boolean=true) {
   books(searchParam:$searchParam,filterBooksInBookPositions:$filterBooksInBookPositions){
       edges{
          value:id 
          text:name
          cover 
       }
   }
  }
`

const BOOK_DETAIL = gql`
 query books($searchParam:SearchParam!) {
   books(searchParam: $searchParam){
        edges{
            id
            isbn
            name
            desc
            cover
            author
            translator
            publishingHouse
            edition
            printedTimes
            printedSheets
            format
            wordCount
            pricing
            purchasePrice
            purchaseTime
            purchaseSource
            bookBorrowUID
            createdAt
            updatedAt
       }
   }
  }
`
export {
  LIST_BOOK, BOOK_COMBO, BOOK_DETAIL
}
