import { useQuery } from '@apollo/react-hooks';
import { Modal, Form, Select } from 'antd';
import React from 'react';
import { LIST_DEVICE_SELECTOR } from 'src/consts/device.gql';

const { Option, OptGroup } = Select;

export interface INewDeviceDashboardTelemetry {
    deviceDashboardID: number
    telemetryIDs: number[]
}

interface DeviceDashboardTelemetryAddFormProps {
    visible: boolean
    onCreate: (values: INewDeviceDashboardTelemetry) => void
    onCancel: () => void
}

export const DeviceDashboardTelemetryAddForm: React.FC<DeviceDashboardTelemetryAddFormProps> = ({
    visible,
    onCreate,
    onCancel,
}) => {
    const [form] = Form.useForm();
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }

    const { data } = useQuery(LIST_DEVICE_SELECTOR,
        {
            variables: {
                sorts: [{
                    field: 'id',
                    isAsc: true
                }]
            },
            fetchPolicy: "cache-and-network"
        })
    function handleChange(value: any) {
        // console.log(`selected ${value}`);
    }

    const getOptGroups = (data: any[]) => {
        return data.map((value: any) => {
            const options = value.telemetries.map((t: any) => {
                return <Option key={t.id} value={t.id}>{t.name}({t.key})</Option>
            })
            return <OptGroup label={value.name} key={value.id}>
                {options}
            </OptGroup>
        })
    }
    return (
        <Modal
            visible={visible}
            title="添加遥测"
            okText="确定"
            cancelText="取消"
            onCancel={
                () => {
                    onCancel()
                    form.resetFields()
                }
            }
            getContainer={false}
            onOk={() => {
                form
                    .validateFields()
                    .then((values: any) => {
                        form.resetFields();
                        onCreate(values);
                    })
                    .catch(info => {
                        console.log('Validate Failed:', info);
                    });
            }}
            maskClosable={false}
        >
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="deviceDashboardTelemetryAddForm"
                style={{ maxHeight: 600 }}
                initialValues={{ isVisible: true }}
            >
                <Form.Item name="telemetryIDs" label="遥测" valuePropName='checked'
                    rules={[{ required: true, message: 'Please select your favourite colors!', type: 'array' }]}
                >
                    <Select mode="multiple"
                        style={{ width: 200 }}
                        listItemHeight={10}
                        onChange={handleChange}>
                        {getOptGroups(data ? data.devices.edges : [])}
                    </Select>
                </Form.Item>
            </Form>
        </Modal>
    )
}
