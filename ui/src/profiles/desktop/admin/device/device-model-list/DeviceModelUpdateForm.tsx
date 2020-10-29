import { Modal, Form, Input } from 'antd';
import React, { useEffect } from 'react'
import TextArea from 'antd/lib/input/TextArea';


export interface IUpdateDeviceModel {
    id: number
    name: string
    desc: string
}

interface DeviceModelUpdateFormProps {
    data: IUpdateDeviceModel
    visible: boolean;
    onUpdate: (values: IUpdateDeviceModel) => void;
    onCancel: () => void;
}

export const DeviceModelUpdateForm: React.FC<DeviceModelUpdateFormProps> = ({
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
            "desc": data.desc
        })
    }, [form, data]);
    return (
        <Modal
            visible={visible}
            title="修改设备模型"
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
                name="deviceModelUpdateForm"
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
                <Form.Item name="desc" label="简介">
                    <TextArea rows={4} />
                </Form.Item>
            </Form>
        </Modal>
    );
};