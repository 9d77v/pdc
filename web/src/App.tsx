import React, { Suspense, useEffect } from 'react'
import {
  RecoilRoot,
} from 'recoil';
import { ApolloProvider } from '@apollo/react-hooks'
import { apolloClient } from './utils/apollo_client'
import { isMobile } from './utils/util'
import {
  BrowserRouter as Router,
  Switch, Route, Redirect
} from 'react-router-dom'
import { Spin } from 'antd'
import { GesturePasswordKey } from './consts/consts'
import { AppPath } from './consts/path'
const DesktopIndex = React.lazy(() => import('./profiles/desktop/index'))
const MobileIndex = React.lazy(() => import('./profiles/mobile/index'))
const Login = React.lazy(() => import('./profiles/login/index'))
const GestureLogin = React.lazy(() => import("./profiles/login/GestureLogin"))
const App = () => {
  useEffect(() => {
    if (document.location.pathname !== AppPath.LOGIN) {
      if (!localStorage.getItem("accessToken")) {
        window.location.href = AppPath.LOGIN
      }
      if (document.location.pathname !== AppPath.GESTURE_LOGIN &&
        document.location.pathname !== AppPath.USER_GESTURE_PASSWORD) {
        if (localStorage.getItem(GesturePasswordKey)) {
          window.location.href = AppPath.GESTURE_LOGIN
        }
      }
    }
  }, [])
  return (
    <RecoilRoot>
      <Suspense fallback={<Spin />}>
        <ApolloProvider client={apolloClient}>
          <Router>
            <Switch>
              <Route exact path="/">
                <Redirect to={AppPath.LOGIN} />
              </Route>
              <Route exact path={AppPath.GESTURE_LOGIN} component={GestureLogin} />
              <Route exact path={AppPath.LOGIN} component={Login} />
              <Route path="/app" component={isMobile() ? MobileIndex : DesktopIndex} />
              <Route path="/admin" component={DesktopIndex} />
            </Switch>
          </Router >
        </ApolloProvider >
      </Suspense >
    </RecoilRoot>
  )
}

export default App
