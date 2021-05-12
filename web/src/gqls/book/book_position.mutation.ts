import gql from 'graphql-tag';


const CREATE_BOOK_POSITION = gql`
mutation createBookPosition($input:NewBookPosition!){
   createBookPosition(input:$input){
     id
   }
}
`

const UPDATE_BOOK_POSITION = gql`
mutation updateBookPosition($input:NewUpdateBookPosition!){
   updateBookPosition(input:$input){
     id
   }
}
`
const REMOVE_BOOK_POSITION = gql`
mutation removeBookPosition($id:Int!){
   removeBookPosition(id:$id){
     id
   }
}
`

export {
  CREATE_BOOK_POSITION, UPDATE_BOOK_POSITION, REMOVE_BOOK_POSITION
}
