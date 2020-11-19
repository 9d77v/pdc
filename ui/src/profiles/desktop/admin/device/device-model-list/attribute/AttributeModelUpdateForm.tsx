import { Modal, Form, Input } from 'antd';
import React, { FC, useEffect } from 'react'
import { IUpdateAttributeModel } from 'src/models/device';


interface IAttributeModelCreateFormProps {
    data: IUpdateAttributeModel
    visible: boolean;
    onUpdate: (values: IUpdateAttributeModel) => void;
    onCancel: () => void;
}

export const AttributeModelUpdateForm: FC<IAttributeModelCreateFormProps> = ({
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
            "key": data.key,
            "name": data.name,
        })
    }, [form, data]);

    return (
        <Modal
            visible={visible}
            title="修改属性"
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
                name="attributeModelUpdateForm"
                style={{ maxHeight: 600 }}
            >
                <Form.Item
                    name="id"
                    noStyle
                >
                    <Input hidden />
                </Form.Item>
                <Form.Item
                    name="key"
                    label="键"
                    rules={[{ required: true, message: '请输入key!' }]}
                >
                    <Input disabled />
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