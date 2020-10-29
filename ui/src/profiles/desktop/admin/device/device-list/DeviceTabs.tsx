import { Tabs, Layout } from 'antd';
import React from 'react';
import { DeviceDetailDescriptions } from './DeviceDetailDescriptions';
import { useQuery } from '@apollo/react-hooks';
import { GET_DEVICE } from 'src/consts/device.gql';
import TelemetryTable from './TelemetryTable';
import AttributeTable from './AttributeTable';

const { TabPane } = Tabs;
const { Header } = Layout;

interface IDeviceTabsProps {
    id: number
}
export const DeviceTabs = (props: IDeviceTabsProps) => {
    const { id } = props
    const { data } = useQuery(GET_DEVICE,
        {
            variables: {
                ids: [id]
            },
            fetchPolicy: "cache-and-network"
        })
    const device = data?.devices.edges[0]
    return (
        <div style={{ width: "100%" }}>
            <Header style={{ color: "white", textAlign: "left" }}>
                {device?.name}
            </Header>
            <Tabs type="card" style={{ backgroundColor: "#fff" }}>
                <TabPane tab="详细信息" key="1">
                    <DeviceDetailDescriptions {...device} />
                </TabPane>
                <TabPane tab="属性" key="2">
                    <AttributeTable
                        id={id}
                        data={device?.attributes || []} />
                </TabPane>
                <TabPane tab="遥测" key="3">
                    <TelemetryTable
                        id={id}
                        data={device?.telemetries || []} />
                </TabPane>
            </Tabs>
        </div>
    )
}
