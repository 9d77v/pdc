import React from "react"
import { Route, useHistory, useLocation } from "react-router-dom";
import { Icon, NavBar, Tabs } from "antd-mobile";
import DeviceCards from "src/profiles/common/device/DeviceCard";
import DeviceCameraList from "./DeviceCameraList";
import { AppPath } from "src/consts/path";


export default function DeviceIndex() {
    const history = useHistory()
    const tabs = [
        { title: "遥测" },
        { title: "摄像头" },
    ];
    const location = useLocation()
    let initialPage = 0
    if (location.pathname === AppPath.DEVICE_CAMERA) {
        initialPage = 1
    }

    return (
        <div style={{ height: "100%", textAlign: "center" }}>
            <NavBar
                mode="light"
                style={{ position: "fixed", width: "100%", zIndex: 10, top: 0 }}
                icon={<Icon type="left" />}
                onLeftClick={() => history.goBack()}
            >设备</NavBar>
            <div style={{ marginTop: 45 }}>
                <Tabs tabs={tabs}
                    initialPage={initialPage}
                    onChange={(tab, index) => {
                    }}
                    onTabClick={(tab, index) => {
                        if (index !== initialPage) {
                            if (index === 1) {
                                window.location.replace(AppPath.DEVICE_CAMERA)
                            } else {
                                history.replace(AppPath.DEVICE_TELEMETRY)
                            }
                        }
                    }}
                >
                    <Route exact path={AppPath.DEVICE_TELEMETRY}  >
                        <DeviceCards width={160} />
                    </Route>
                    <Route exact path={AppPath.DEVICE_CAMERA}  >
                        <DeviceCameraList />
                    </Route>
                </Tabs>
            </div>
        </div>)
}
