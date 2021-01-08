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
`

const HISTORY_STATISTIC = gql`
query historyStatistic($sourceType: Int) {
  historyStatistic(sourceType: $sourceType) {
    data
  }
}
`
const APP_HISTORY_STATISTIC = gql`
query appHistoryStatistic($sourceType: Int) {
  appHistoryStatistic(sourceType: $sourceType) {
    data
  }
}
`
export { LIST_HISTORY, HISTORY_STATISTIC, APP_HISTORY_STATISTIC }
