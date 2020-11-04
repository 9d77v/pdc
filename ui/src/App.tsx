import React, { Suspense, useEffect } from 'react';
import { ApolloProvider } from '@apollo/react-hooks';
import { apolloClient } from './utils/apollo_client';
import { isMobile } from './utils/util';
import {
  BrowserRouter as Router,
  Switch, Route, Redirect
} from 'react-router-dom';
import { Spin } from 'antd';
import { GesturePasswordKey } from './consts/consts';
const DesktopIndex = React.lazy(() => import('./profiles/desktop/index'))
const MobileIndex = React.lazy(() => import('./profiles/mobile/index'))
const Login = React.lazy(() => import('./profiles/login/index'))
const GestureLogin = React.lazy(() => import("./profiles/login/GestureLogin"))
const App = () => {
  useEffect(() => {
    if (document.location.pathname !== "/login") {
      if (!localStorage.getItem("accessToken")) {
        window.location.href = '/login';
      }
      if (document.location.pathname !== "/gesture_login" &&
        document.location.pathname !== "/app/user/gesture_password") {
        if (localStorage.getItem(GesturePasswordKey)) {
          window.location.href = '/gesture_login';
        }
      }
    }
  }, [])
  return (
    <Suspense fallback={<Spin />}>
      <ApolloProvider client={apolloClient}>
        <Router>
          <Switch>
            <Route exact path="/">
              <Redirect to="/login" />
            </Route>
            <Route exact path="/gesture_login" component={GestureLogin} />
            <Route exact path="/login" component={Login} />
            <Route path="/app" component={isMobile() ? MobileIndex : DesktopIndex} />
            <Route path="/admin" component={DesktopIndex} />
          </Switch>
        </Router >
      </ApolloProvider >
    </Suspense >
  );
}

export default App
