import { Modal, Form, Input, Switch, Select } from 'antd';
import React, { FC } from 'react';
import { DeviceTypeMap } from 'src/consts/consts';
import { INewDeviceDashboard } from 'src/models/device';

interface IDeviceDashboardCreateFormProps {
    visible: boolean;
    onCreate: (values: INewDeviceDashboard) => void;
    onCancel: () => void;
}

export const DeviceDashboardCreateForm: FC<IDeviceDashboardCreateFormProps> = ({
    visible,
    onCreate,
    onCancel,
}) => {

    const [form] = Form.useForm();
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    let deviceTypeOptions: any[] = []
    DeviceTypeMap.forEach((value: string, key: number) => {
        deviceTypeOptions.push(<Select.Option
            value={key}
            key={'deviceType_options_' + key}>{value}</Select.Option>)
    })
    return (
        <Modal
            visible={visible}
            title="新增设备仪表盘"
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
                name="deviceDashboardCreateForm"
                style={{ maxHeight: 600 }}
                initialValues={{ isVisible: true, deviceType: 0 }}
            >
                <Form.Item
                    name="name"
                    label="名称"
                    rules={[{ required: true, message: '请输入名称!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="deviceType"
                    label="设备类型"
                    hasFeedback
                    rules={[{ required: true, message: '请选择设备类型!' }]}
                >
                    <Select >
                        {deviceTypeOptions}
                    </Select>
                </Form.Item>
                <Form.Item name="isVisible" label="是否显示" valuePropName='checked'>
                    <Switch />
                </Form.Item>
            </Form>
        </Modal>
    )
}
