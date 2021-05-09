import { Modal, Form, Input, Switch } from 'antd';
import React, { FC, useEffect } from 'react'
import { IUpdateDeviceDashboard } from 'src/models/device';

interface IDeviceDashboardUpdateFormProps {
    data: IUpdateDeviceDashboard
    visible: boolean;
    onUpdate: (values: IUpdateDeviceDashboard) => void;
    onCancel: () => void;
}

export const DeviceDashboardUpdateForm: FC<IDeviceDashboardUpdateFormProps> = ({
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
        setTimeout(() => {
            form.setFieldsValue({
                "id": data.id,
                "name": data.name,
                "isVisible": data.isVisible
            })
        }, 3000);
    }, [form, data]);
    return (
        <Modal
            visible={visible}
            title="修改设备仪表盘"
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
                name="deviceDashboardUpdateForm"
                style={{ maxHeight: 600 }}
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
                <Form.Item name="isVisible" label="是否显示" valuePropName='checked'>
                    <Switch />
                </Form.Item>
            </Form>
        </Modal>
    )
}
