import { Modal, Form, Input, Switch } from 'antd';
import React from 'react';

export interface INewDeviceDashboard {
    name: string
    isVisible: boolean
}

interface DeviceDashboardCreateFormProps {
    visible: boolean;
    onCreate: (values: INewDeviceDashboard) => void;
    onCancel: () => void;
}

export const DeviceDashboardCreateForm: React.FC<DeviceDashboardCreateFormProps> = ({
    visible,
    onCreate,
    onCancel,
}) => {

    const [form] = Form.useForm();
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }

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
                initialValues={{ isVisible: true }}
            >
                <Form.Item
                    name="name"
                    label="名称"
                    rules={[{ required: true, message: '请输入名称!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item name="isVisible" label="是否显示" valuePropName='checked'>
                    <Switch />
                </Form.Item>
            </Form>
        </Modal>
    )
}
