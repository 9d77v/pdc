import gql from 'graphql-tag';

const ADD_BOOK = gql`
mutation createBook($input:NewBook!){
   createBook(input:$input){
     id
   }
}
`
const UPDATE_BOOK = gql`
mutation updateBook($input:NewUpdateBook!){
   updateBook(input:$input){
     id
   }
}
`

export {
   ADD_BOOK, UPDATE_BOOK
}
