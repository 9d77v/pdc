import { Modal, Form, Input, InputNumber } from 'antd';
import React, { useEffect } from 'react'

export interface IUpdateDevice {
    id: number
    name: string
    ip: string
    port: number
}

interface DeviceUpdateFormProps {
    data: IUpdateDevice
    visible: boolean;
    onUpdate: (values: IUpdateDevice) => void;
    onCancel: () => void;
}

export const DeviceUpdateForm: React.FC<DeviceUpdateFormProps> = ({
    data,
    visible,
    onUpdate,
    onCancel,
}) => {
    const [form] = Form.useForm();
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    useEffect(() => {
        form.setFieldsValue({
            "id": data.id,
            "name": data.name,
            "ip": data.ip,
            "port": data.port
        })
    }, [form, data]);
    return (
        <Modal
            visible={visible}
            title="修改设备"
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
                        onUpdate(values);
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
                name="deviceUpdateForm"
                style={{ maxHeight: 600 }}
                initialValues={{ deviceType: 0 }}
            >
                <Form.Item
                    name="id"
                    noStyle
                >
                    <Input hidden />
                </Form.Item>
                <Form.Item
                    name="name"
                    label="名称"
                    rules={[{ required: true, message: '请输入名称!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="ip"
                    label="IP地址"
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="port"
                    label="端口"
                >
                    <InputNumber />
                </Form.Item>
            </Form>
        </Modal>
    )
}
