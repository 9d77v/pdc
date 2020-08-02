import gql from 'graphql-tag';

const RECORD_HISTORY = gql`
mutation recordHistory($input:NewHistoryInput!){
    recordHistory(input:$input){
        subSourceID
   }
}
`

export { RECORD_HISTORY }
