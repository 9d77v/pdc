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
 query things( $page: Int, $pageSize: Int, $sorts: [Sort!]) {
   things(page: $page, pageSize: $pageSize,sorts:$sorts){
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
            consumerExpenditure
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
 query thingSeries($dimension:String!,$index1:String!,$index2:String!,$status:[Int!]) {
    series3:thingSeries(dimension: $dimension, index: $index1,status:$status){
      name
      value
   }
   series4:thingSeries(dimension: $dimension, index: $index2,status:$status){
      name
      value
   }
  }
`

const THING_ANALYZE = gql`
 query thingAnalyze($dimension:String!,$index1:String!,$index2:String!,$start:Int,$group:String!) {
   series1:thingAnalyze(dimension: $dimension, index: $index1,start:$start,group:$group){
      x1
      x2
      y
   }
    series2:thingAnalyze(dimension: $dimension, index: $index2,start:$start,group:$group){
      x1
      x2
      y
   }
  }
`
export { LIST_THING, ADD_THING, UPDATE_THING, THING_SERIES, THING_ANALYZE }
