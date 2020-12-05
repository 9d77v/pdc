import gql from 'graphql-tag';


const LIST_USER = gql`
 query users($searchParam:SearchParam!) {
   users(searchParam: $searchParam){
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

const GET_USER = gql`
 query userInfo($uid:ID) {
   userInfo(uid:$uid){
      uid
      name
      avatar
      roleID
      gender
      color
      birthDate
      ip 
   }
  }
`;


export {
  LIST_USER, GET_USER
}
