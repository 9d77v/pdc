import { Modal, Form, Input, InputNumber } from 'antd';
import React, { FC, useEffect } from 'react'
import { IUpdateDevice } from 'src/models/device';

interface IDeviceUpdateFormProps {
    data: IUpdateDevice
    visible: boolean;
    onUpdate: (values: IUpdateDevice) => void;
    onCancel: () => void;
}

export const DeviceUpdateForm: FC<IDeviceUpdateFormProps> = ({
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
            "port": data.port,
            "username": data.username,
            "password": data.password
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
                <Form.Item
                    name="username"
                    label="用户名"
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="password"
                    label="密码"
                >
                    <Input.Password />
                </Form.Item>
            </Form>
        </Modal>
    )
}
