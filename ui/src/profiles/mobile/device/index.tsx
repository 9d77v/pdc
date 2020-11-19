import React from "react"
import { useHistory } from "react-router-dom";
import { Icon, NavBar, Tabs } from "antd-mobile";
import DeviceCards from "src/profiles/common/device/DeviceCard";
import DeviceCameraList from "./DeviceCameraList";


export default function DeviceIndex() {
    const history = useHistory()

    const tabs = [
        { title: "遥测" },
        { title: "摄像头" },
    ];

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
                    initialPage={0}
                    onChange={(tab, index) => { console.log('onChange', index, tab); }}
                    onTabClick={(tab, index) => { console.log('onTabClick', index, tab); }}
                >
                    <DeviceCards width={200} />
                    <DeviceCameraList />
                </Tabs>
            </div>
        </div>)
}