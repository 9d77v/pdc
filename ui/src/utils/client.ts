import { ApolloClient } from 'apollo-client';
import { HttpLink } from 'apollo-link-http';
import { InMemoryCache } from 'apollo-cache-inmemory';
import { ApolloLink, from } from 'apollo-link';
import { onError } from 'apollo-link-error';

const httpLink = new HttpLink({ uri: '/api' });


const authLink = new ApolloLink((operation, forward) => {
    // add the authorization to the headers
    const token = localStorage.getItem('accessToken');
    operation.setContext({
        headers: {
            Authorization: token ? `Bearer ${token}` : "",
        }
    });

    return forward(operation);
})

const errorLink = onError(({ graphQLErrors, networkError }) => {
    if (graphQLErrors) {
        graphQLErrors.map(({ message, locations, path }) =>
            console.log(
                `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`,
            ),
        )
    }
    if (networkError) {
        const err: any = networkError
        if (err.statusCode === 401) {
            localStorage.clear()
        }
    };
})

export const client = new ApolloClient({
    link: from([
        errorLink,
        authLink,
        httpLink,

    ]),
    cache: new InMemoryCache()
});