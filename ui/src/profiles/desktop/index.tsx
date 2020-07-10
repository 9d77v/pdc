import React from 'react';
import "./index.less"
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Redirect
} from "react-router-dom";
import { Login } from './login/Login';
import App from './App';

export default function DesktopIndex() {
    return (
        <Router>
            <Switch>
                <Route exact path="/">
                    <Redirect to="/login" />
                </Route>
                <Route exact path="/login" component={Login} />
                <Route path="/app" component={App} />
            </Switch>
        </Router >
    )
}
