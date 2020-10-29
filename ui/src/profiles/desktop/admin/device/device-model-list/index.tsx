import React, { useState } from 'react';
import { DeviceModelList } from './DeviceModelList';
import { DeviceModelTabs } from './DeviceModelTabs';


export default function DeviceModelIndex() {
    const [currentSelectItem, setCurrentSelectItem] = useState({
        id: 0,
        name: "",
        desc: "",
        deviceType: 0,
        createdAt: 0,
        updatedAt: 0,
    })
    return (
        <div style={{ display: "flex", backgroundColor: "#e9f1f1" }}>
            <div style={{ width: 350, padding: 10 }}>
                <DeviceModelList currentSelectID={currentSelectItem.id}
                    setCurrentSelectItem={setCurrentSelectItem} />
            </div>
            <div style={{ flex: "1 1 auto", padding: 10 }}>
                {currentSelectItem.id ? <DeviceModelTabs id={currentSelectItem.id} /> : undefined}
            </div>
        </div>
    )
}
