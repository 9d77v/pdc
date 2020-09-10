import React, { useEffect } from 'react';
import { Spin, ConfigProvider } from 'antd';
import { ApolloProvider } from '@apollo/react-hooks';
import zhCN from 'antd/es/locale/zh_CN';
import { apolloClient } from './utils/apollo_client';
import { isMobile } from './utils/util';
import {
  BrowserRouter as Router,
  Switch, Route, Redirect, useHistory
} from 'react-router-dom';
import { Login } from './profiles/login/Login';
const DesktopIndex = React.lazy(() => import('./profiles/desktop/index'))
const MobileIndex = React.lazy(() => import('./profiles/mobile/index'))

export default function App() {
  const history = useHistory();
  useEffect(() => {
    const token = localStorage.getItem('accessToken');
    if (!token) {
      history.push('/login')
    }
  }, [history])
  return (
    <ApolloProvider client={apolloClient}>
      <React.Suspense fallback={<Spin />}>
        <ConfigProvider locale={zhCN}>
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
        </ConfigProvider>
      </React.Suspense>
    </ApolloProvider>
  );
}
