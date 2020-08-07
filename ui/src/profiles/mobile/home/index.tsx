import React from "react"
import { Route, useLocation } from "react-router-dom";

export default function HomeIndex() {
    const location = useLocation();

    switch (location.pathname) {
        case "/app/home":
            break
    }
    return (
        <Route exact path="/app/home">
            <div style={{ display: 'flex', alignItems: 'center', height: "100%", justifyContent: 'center', backgroundColor: '#eee' }}>
                欢迎使用个人数据 中心
        </div>
        </Route>)
}