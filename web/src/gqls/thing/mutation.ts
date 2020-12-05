import gql from 'graphql-tag';

const ADD_THING = gql`
mutation createThing($input:NewThing!){
   createThing(input:$input){
     id
   }
}
`
const UPDATE_THING = gql`
mutation updateThing($input:NewUpdateThing!){
   updateThing(input:$input){
     id
   }
}
`

export { ADD_THING, UPDATE_THING }
