import gql from 'graphql-tag';

const ADD_USER = gql`
mutation createUser($input:NewUser!){
   createUser(input:$input){
     uid
   }
}
`
const UPDATE_USER = gql`
mutation updateUser($input:NewUpdateUser!){
   updateUser(input:$input){
     uid
   }
}
`

const LOGIN = gql`
mutation login($username: String!, $password: String!) {
  login(username: $username, password: $password) {
    accessToken
    refreshToken
  }
}
`

const REFRESH_TOKEN = gql`
mutation refreshToken($refreshToken:String!){
   refreshToken(refreshToken:$refreshToken){
      accessToken
      refreshToken
   }
}
`

const UPDATE_PROFILE = gql`
mutation updateProfile($input:NewUpdateProfile!){
   updateProfile(input:$input){
      uid
   }
}
`
const UPDATE_PASSWORD = gql`
mutation updatePassword($oldPassword:String!,$newPassword:String!){
   updatePassword(oldPassword:$oldPassword,newPassword:$newPassword){
      uid
   }
}
`

export {
    ADD_USER, UPDATE_USER,
    LOGIN, REFRESH_TOKEN,
    UPDATE_PROFILE, UPDATE_PASSWORD
}
