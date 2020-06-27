import gql from 'graphql-tag';

const ADD_USER = gql`
mutation createUser($input:NewUser!){
   createUser(input:$input){
     id
   }
}
`
const UPDATE_USER = gql`
mutation updateUser($input:NewUpdateUser!){
   updateUser(input:$input){
     id
   }
}
`
const LIST_USER = gql`
 query Users( $page: Int, $pageSize: Int, $sorts: [Sort!]) {
   Users(page: $page, pageSize: $pageSize,sorts:$sorts){
        totalCount
        edges{
            id
            name
            avatar
            roleID
            gender
            color
            birthDate
            ip 
            createdAt
            updatedAt
       }
   }
  }
`;


export { LIST_USER, ADD_USER, UPDATE_USER }
