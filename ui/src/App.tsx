import React, { Suspense, useEffect } from 'react';
import { ApolloProvider } from '@apollo/react-hooks';
import { apolloClient } from './utils/apollo_client';
import { isMobile } from './utils/util';
import {
  BrowserRouter as Router,
  Switch, Route, Redirect, useHistory
} from 'react-router-dom';
import { Spin } from 'antd';
const DesktopIndex = React.lazy(() => import('./profiles/desktop/index'))
const MobileIndex = React.lazy(() => import('./profiles/mobile/index'))
const Login = React.lazy(() => import('./profiles/login/index'))
export default function App() {
  const history = useHistory();
  useEffect(() => {
    const token = localStorage.getItem('accessToken');
    if (!token) {
      if (history) {
        history.push('/login')
      }
    }
  }, [history])
  document.documentElement.style.setProperty('--theme-primary', '#108ee9')
  return (
    <Suspense fallback={<Spin />}>
      <ApolloProvider client={apolloClient}>
        <Router>
          <Switch>
            <Route exact path="/">
              <Redirect to="/login" />
            </Route>
            <Route exact path="/login" component={Login} />
            <Route path="/app" component={isMobile() ? MobileIndex : DesktopIndex} />
            <Route path="/admin" component={DesktopIndex} />
          </Switch>
        </Router >
      </ApolloProvider >
    </Suspense >
  );
}
