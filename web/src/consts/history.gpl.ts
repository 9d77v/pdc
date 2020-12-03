import gql from 'graphql-tag';

const RECORD_HISTORY = gql`
mutation recordHistory($input:NewHistoryInput!){
    recordHistory(input:$input){
        subSourceID
   }
}
`

const LIST_HISTORY = gql`
 query histories($sourceType:Int, $page: Int, $pageSize: Int) {
   histories(sourceType:$sourceType,page: $page, pageSize: $pageSize){
       totalCount
       edges{
            sourceType
            sourceID
            title
            num
            subTitle
            cover
            subSourceID  
            platform
            currentTime
            remainingTime
            updatedAt
       }
   }
  }
`;
export { RECORD_HISTORY, LIST_HISTORY }
