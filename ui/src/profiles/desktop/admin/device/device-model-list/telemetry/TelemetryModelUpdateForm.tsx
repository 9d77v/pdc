import { Modal, Form, Input, InputNumber } from 'antd';
import React, { FC, useEffect } from 'react'
import { IUpdateTelemetryModel } from 'src/models/device';


interface ITelemetryModelUpdateFormProps {
    data: IUpdateTelemetryModel
    visible: boolean;
    onUpdate: (values: IUpdateTelemetryModel) => void;
    onCancel: () => void;
}

export const TelemetryModelUpdateForm: FC<ITelemetryModelUpdateFormProps> = ({
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
            "factor": data.factor,
            "unit": data.unit,
            "unitName": data.unitName,
            "scale": data.scale
        })
    }, [form, data]);

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
                name="telemetryModelUpdateForm"
                style={{ maxHeight: 600 }}
                initialValues={{ factor: 1, scale: 2 }}
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