import gql from 'graphql-tag';

const ADD_BOOKSHELF = gql`
mutation createBookshelf($input:NewBookshelf!){
   createBookshelf(input:$input){
     id
   }
}
`
const UPDATE_BOOKSHELF = gql`
mutation updateBookshelf($input:NewUpdateBookshelf!){
   updateBookshelf(input:$input){
     id
   }
}
`

export {
   ADD_BOOKSHELF, UPDATE_BOOKSHELF
}
