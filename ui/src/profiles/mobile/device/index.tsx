import React from "react"
import { useHistory } from "react-router-dom";
import { Icon, NavBar } from "antd-mobile";
import DeviceCards from "src/profiles/common/device/DeviceCard";

export default function DeviceIndex() {
    const history = useHistory()

    return (
        <div style={{ height: "100%", textAlign: "center" }}>
            <NavBar
                mode="light"
                style={{ position: "fixed", width: "100%", zIndex: 10, top: 0 }}
                icon={<Icon type="left" />}
                onLeftClick={() => history.goBack()}
            >设备</NavBar>
            <div style={{ marginTop: 45 }}>
                <DeviceCards width={"42%"} />
            </div>
        </div>)
}