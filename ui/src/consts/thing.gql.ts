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

const THING_SERIES = gql`
 query ThingSeries($dimension:String!,$index1:String!,$index2:String!,$start:Int,$end:Int,$status:[Int!]) {
   Series1:ThingSeries(dimension: $dimension, index: $index1,start:$start,end:$end,status:$status){
      name
      value
   }
    Series2:ThingSeries(dimension: $dimension, index: $index2,start:$start,end:$end,status:$status){
      name
      value
   }
    Series3:ThingSeries(dimension: $dimension, index: $index1,status:$status){
      name
      value
   }
    Series4:ThingSeries(dimension: $dimension, index: $index2,status:$status){
      name
      value
   }
  }
`
export { LIST_THING, ADD_THING, UPDATE_THING, THING_SERIES }