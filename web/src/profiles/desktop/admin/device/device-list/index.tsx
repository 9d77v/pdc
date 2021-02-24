import React, { useState } from 'react';
import { DeviceList } from './DeviceList';
import { DeviceTabs } from './DeviceTabs';


export default function DeviceModelIndex() {
    const [currentSelectItem, setCurrentSelectItem] = useState({
        id: 0,
        name: "",
        deviceModelID: 0,
        deviceModel: {
            name: "",
            desc: ""
        },
        createdAt: 0,
        updatedAt: 0,
    })

    return (
        <div style={{ display: "flex" }}>
            <div style={{ width: 380, padding: 10 }}>
                <DeviceList currentSelectID={currentSelectItem.id}
                    setCurrentSelectItem={setCurrentSelectItem} />
            </div>
            <div style={{ flex: "1 1 auto", padding: 10 }}>
                {currentSelectItem.id ? <DeviceTabs id={currentSelectItem.id} /> : undefined}
            </div>
        </div>
    )
}
