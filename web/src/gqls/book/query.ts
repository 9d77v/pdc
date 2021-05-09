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
  LIST_BOOK
}
