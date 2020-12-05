import gql from 'graphql-tag';

const LIST_HISTORY = gql`
query histories($sourceType: Int, $searchParam: SearchParam!) {
  histories(sourceType: $sourceType, searchParam: $searchParam) {
    totalCount
    edges {
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
export { LIST_HISTORY }
