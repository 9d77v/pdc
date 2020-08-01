import { Modal, Form, Input } from 'antd';
import React from 'react'

interface Values {
    title: string;
    description: string;
}

interface VideoSeriesCreateFormProps {
    visible: boolean;
    onCreate: (values: Values) => void;
    onCancel: () => void;
}

export const VideoSeriesCreateForm: React.FC<VideoSeriesCreateFormProps> = ({
    visible,
    onCreate,
    onCancel,
}) => {
    const [form] = Form.useForm();

    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }
    return (
        <Modal
            visible={visible}
            title="新增视频系列"
            okText="确定"
            cancelText="取消"
            destroyOnClose
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
                name="videoSeriesCreateForm"
                initialValues={{ isShow: true, subtitle_lang: "简体中文" }}
            >
                <Form.Item
                    name="name"
                    label="名称"
                    rules={[{ required: true, message: '请输入系列名称!' }]}
                >
                    <Input />
                </Form.Item>
            </Form>
        </Modal>
    );
};