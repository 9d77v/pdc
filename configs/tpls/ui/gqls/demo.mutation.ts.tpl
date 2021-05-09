import gql from 'graphql-tag';

const ADD_[[.TitleName]] = gql`
mutation create[[.Name]]($input:New[[.Name]]!){
   create[[.Name]](input:$input){
     id
   }
}
`
const UPDATE_[[.TitleName]] = gql`
mutation update[[.Name]]($input:NewUpdate[[.Name]]!){
   update[[.Name]](input:$input){
     id
   }
}
`

export {
   ADD_[[.TitleName]], UPDATE_[[.TitleName]]
}
