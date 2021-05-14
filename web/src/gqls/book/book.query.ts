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

const SEARCH_BOOK = gql`
 query searchBook($searchParam:SearchParam!,$bookshelfsSearchParam:SearchParam!) {
   searchBook(searchParam: $searchParam){
        totalCount
        edges{
            id
            isbn
            name
            cover
            author
            translator
       }
   }
   bookshelfs(searchParam: $bookshelfsSearchParam){
        totalCount
        edges{
            id
            name
            cover
       }
   }
  }
`

const APP_BOOK_DETAIL = gql`
 query searchBook($searchParam:SearchParam!) {
   searchBook(searchParam: $searchParam){
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
       }
   }
  }
`
export {
  LIST_BOOK, BOOK_COMBO, BOOK_DETAIL, SEARCH_BOOK, APP_BOOK_DETAIL
}
