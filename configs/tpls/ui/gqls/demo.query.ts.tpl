import gql from 'graphql-tag';


const LIST_[[.TitleName]] = gql`
 query [[.LowerName]]s($searchParam:SearchParam!) {
   [[.LowerName]]s(searchParam: $searchParam){
        totalCount
        edges{
            id[[range .Columns]]
            [[.Name]][[end]]
            createdAt
            updatedAt
       }
   }
  }
`

export {
  LIST_[[.TitleName]]
}
