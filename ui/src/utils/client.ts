import {
    from,
    ApolloClient,
    InMemoryCache,
    HttpLink
} from '@apollo/client';
import { setContext } from '@apollo/client/link/context';
import { onError } from "@apollo/client/link/error";
import jwt_decode from 'jwt-decode';
import moment from 'moment';
import { getRefreshToken } from '../consts/http';
import { message as msg, message } from 'antd'
const httpLink = new HttpLink({ uri: '/api' });

const authLink = setContext(
    () => {
        let token = localStorage.getItem('accessToken') || "";
        if (token === "") {
            return {
                headers: {
                    Authorization: token ? `Bearer ${token}` : "",
                }
            }
        }
        const accessToken: any = jwt_decode(token)
        if (Number(accessToken.exp) - moment().unix() > 0) {
            return {
                headers: {
                    Authorization: token ? `Bearer ${token}` : "",
                }
            }
        }
        const refreshToken = localStorage.getItem('refreshToken') || "";
        return new Promise(async (success, fail) => {
            //refresh token
            const data = await getRefreshToken(refreshToken)
            if (!data.data) {
                success({
                    headers: {
                        Authorization: token ? `Bearer ${token}` : "",
                    }
                })
                return
            }
            const refreshData: any = data.data.refreshToken
            localStorage.setItem("accessToken", refreshData.accessToken)
            localStorage.setItem("refreshToken", refreshData.refreshToken)
            token = refreshData.accessToken
            success({
                headers: {
                    Authorization: token ? `Bearer ${token}` : "",
                }
            })
        })
    })
const errorLink = onError(({ graphQLErrors, networkError }) => {
    if (graphQLErrors) {
        graphQLErrors.map(({ message, locations, path }) =>
            msg.error(
                `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`,
            ),
        )
    }
    if (networkError) {
        const err: any = networkError
        if (err.statusCode === 401) {
            localStorage.clear()
            client.resetStore()
            message.error("token失效，请刷新页面")
        } else if (err.statusCode === 403) {
            message.error("无操作权限")
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
