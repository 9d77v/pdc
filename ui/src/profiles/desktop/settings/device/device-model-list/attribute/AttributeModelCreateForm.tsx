import { Modal, Form, Input } from 'antd';
import React from 'react'

export interface INewAttributeModel {
    key: string
    name: string
}

interface AttributeModelCreateFormProps {
    visible: boolean;
    onCreate: (values: INewAttributeModel) => void;
    onCancel: () => void;
}

export const AttributeModelCreateForm: React.FC<AttributeModelCreateFormProps> = ({
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
            title="新增设备模型"
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
                name="attributeModelCreateForm"
                style={{ maxHeight: 600}}
                initialValues={{ deviceType: 0 }}
            >
                <Form.Item
                    name="key"
                    label="键"
                    rules={[{ required: true, message: '请输入key!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="name"
                    label="名称"
                    rules={[{ required: true, message: '请输入名称!' }]}
                >
                    <Input />
                </Form.Item>
            </Form>
        </Modal>
    );
};