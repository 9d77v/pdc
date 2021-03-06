import { useQuery } from '@apollo/react-hooks';
import { Modal, Form, Select } from 'antd';
import React, { FC, useMemo } from 'react';
import { LIST_DEVICE_SELECTOR } from 'src/gqls/device/query';
import { IDeviceDashboardTelemetry, INewDeviceDashboardTelemetry } from 'src/models/device';

const { Option, OptGroup } = Select;

interface IDeviceDashboardTelemetryAddFormProps {
    existData: IDeviceDashboardTelemetry[]
    visible: boolean
    onCreate: (values: INewDeviceDashboardTelemetry) => void
    onCancel: () => void
}

export const DeviceDashboardTelemetryAddForm: FC<IDeviceDashboardTelemetryAddFormProps> = ({
    existData,
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
                searchParam: {
                    sorts: [{
                        field: 'id',
                        isAsc: true
                    }],
                },
                deviceType: 0
            },
            fetchPolicy: "cache-and-network"
        })
    function handleChange(value: any) {
        // console.log(`selected ${value}`);
    }

    const existDataMap = useMemo(() => {
        const m = new Map<number, boolean>()
        for (const t of existData || []) {
            m.set(t.telemetryID, true)
        }
        return m
    }, [existData])

    const optGroups = useMemo(() => {
        const d = data ? data.devices.edges : []
        return d.map((value: any) => {
            const options = value.telemetries.map((t: any) => {
                return <Option key={t.id} value={t.id} disabled={existDataMap.get(t.id)}>{t.name}({t.key})</Option>
            })
            return <OptGroup label={value.name} key={value.id}>
                {options}
            </OptGroup>
        })
    }, [existDataMap, data])
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
                        {optGroups}
                    </Select>
                </Form.Item>
            </Form>
        </Modal>
    )
}
