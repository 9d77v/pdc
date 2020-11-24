import { useQuery } from '@apollo/react-hooks';
import { Modal, Form, Select } from 'antd';
import React, { FC, useMemo } from 'react';
import { LIST_DEVICE_SELECTOR } from 'src/consts/device.gql';
import { IDeviceDashboardCamera, INewDeviceDashboardCamera } from 'src/models/device';

const { Option } = Select;

interface IDeviceDashboardCameraAddFormProps {
    existData: IDeviceDashboardCamera[]
    visible: boolean
    onCreate: (values: INewDeviceDashboardCamera) => void
    onCancel: () => void
}

export const DeviceDashboardCameraAddForm: FC<IDeviceDashboardCameraAddFormProps> = ({
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
                sorts: [{
                    field: 'id',
                    isAsc: true
                }],
                deviceType: 1
            },
            fetchPolicy: "cache-and-network"
        })
    function handleChange(value: any) {
        // console.log(`selected ${value}`);
    }

    const existDataMap = useMemo(() => {
        const m = new Map<number, boolean>()
        for (const t of existData || []) {
            m.set(t.deviceID, true)
        }
        return m
    }, [existData])

    const optGroups = useMemo(() => {
        const d = data ? data.devices.edges : []
        return d.map((value: any) => {
            return <Option key={value.id} value={value.id} disabled={existDataMap.get(value.id)}>{value.name}</Option>
        })
    }, [existDataMap, data])
    return (
        <Modal
            visible={visible}
            title="添加摄像头"
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
                name="deviceDashboardCameraAddForm"
                style={{ maxHeight: 600 }}
                initialValues={{ isVisible: true }}
            >
                <Form.Item name="deviceIDs" label="摄像头" valuePropName='checked'
                    rules={[{ required: true, message: '请选择摄像头!', type: 'array' }]}
                >
                    <Select
                        mode="multiple"
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
