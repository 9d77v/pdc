import { Tabs, Layout } from 'antd';
import React from 'react';
import { DeviceModelDetailDescriptions } from './DeviceModelDetailDescriptions';
import AttributeModelTable from './attribute/AttributeModelTable';
import { useQuery } from '@apollo/react-hooks';
import TelemetryModelTable from './telemetry/TelemetryModelTable';
import { GET_DEVICE_MODEL } from 'src/gqls/device/query';

const { TabPane } = Tabs;
const { Header } = Layout;

interface IDeviceModelTabsProps {
    id: number
}
export const DeviceModelTabs = (props: IDeviceModelTabsProps) => {
    const { id } = props
    const { data, refetch } = useQuery(GET_DEVICE_MODEL,
        {
            variables: {
                searchParam: {
                    ids: [id]
                }
            },
            fetchPolicy: "cache-and-network"
        })
    const deviceModel = data?.deviceModels.edges[0]
    return (
        <div style={{ width: "100%" }}>
            <Header style={{ color: "white", textAlign: "left" }}>
                {deviceModel?.name}
            </Header>
            <Tabs type="card" style={{ backgroundColor: "#fff" }}>
                <TabPane tab="详细信息" key="1">
                    <DeviceModelDetailDescriptions {...deviceModel} />
                </TabPane>
                <TabPane tab="属性" key="2">
                    <AttributeModelTable
                        id={id}
                        data={deviceModel?.attributeModels || []}
                        refetch={refetch} />
                </TabPane>
                {
                    deviceModel?.deviceType !== 1 ?
                        <TabPane tab="遥测" key="3">
                            <TelemetryModelTable
                                id={id}
                                data={deviceModel?.telemetryModels || []}
                                refetch={refetch} />
                        </TabPane> : null
                }

            </Tabs>
        </div>
    )
}