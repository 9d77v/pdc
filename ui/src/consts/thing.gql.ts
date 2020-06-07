import gql from 'graphql-tag';

const ADD_THING = gql`
mutation createThing($input:NewThing!){
   createThing(input:$input){
     id
   }
}
`
const UPDATE_THING = gql`
mutation updateThing($input:NewUpdateThing!){
   updateThing(input:$input){
     id
   }
}
`
const LIST_THING = gql`
 query Things( $page: Int, $pageSize: Int, $sorts: [Sort!]) {
   Things(page: $page, pageSize: $pageSize,sorts:$sorts){
        totalCount
        edges{
            id
            uid
            name
            num
            brandName
            pics
            unitPrice
            unit
            specifications
            category
            location
            status
            purchaseDate
            purchasePlatform
            refOrderID
            rubbishCategory
            createdAt
            updatedAt
       }
   }
  }
`;


export { LIST_THING, ADD_THING, UPDATE_THING }