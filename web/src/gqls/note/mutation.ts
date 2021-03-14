import gql from 'graphql-tag';


const SYNC_NOTES = gql`
mutation syncNotes($input: SyncNotesInput!) {
  syncNotes(input: $input) {
    last_update_time:lastUpdateTime
    list {
      id
      parent_id:parentID
      uid
      note_type:noteType
      level
      title
      color
      state
      version
      created_at:createdAt
      updated_at:updatedAt
      content
      tags
      sha1
    }
  }
}
`

export {
  SYNC_NOTES
}
