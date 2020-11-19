import { Modal, Form, Input, InputNumber } from 'antd';
import React, { FC } from 'react'
import { INewTelemetryModel } from 'src/models/device';


interface ITelemetryModelCreateFormProps {
    visible: boolean;
    onCreate: (values: INewTelemetryModel) => void;
    onCancel: () => void;
}

export const TelemetryModelCreateForm: FC<ITelemetryModelCreateFormProps> = ({
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
            title="新增遥测"
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
                name="telemetryModelCreateForm"
                style={{ maxHeight: 600 }}
                initialValues={{ factor: 1, scale: 2 }}
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
                <Form.Item
                    name="factor"
                    label="系数"
                    rules={[{ required: true, message: '请输入系数!' }]}
                >
                    <InputNumber />
                </Form.Item>
                <Form.Item
                    name="unit"
                    label="单位"
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="unitName"
                    label="单位名称"
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="scale"
                    label="小数位数"
                    rules={[{ required: true, message: '请输入小数位数!' }]}
                >
                    <InputNumber min={0} max={10} />
                </Form.Item>
            </Form>
        </Modal>
    );
};