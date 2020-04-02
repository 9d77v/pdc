import { Modal, Form, Input, InputNumber } from 'antd';
import React from 'react'
interface Values {
    title: string;
    description: string;
}

interface EpisodeCreateFormProps {
    visible: boolean;
    onCreate: (values: Values) => void;
    onCancel: () => void;
}

export const EpisodeCreateForm: React.FC<EpisodeCreateFormProps> = ({
    visible,
    onCreate,
    onCancel,
}) => {
    const [form] = Form.useForm();
    return (
        <Modal
            visible={visible}
            title="新增分集"
            okText="确定"
            cancelText="取消"
            onCancel={onCancel}
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
        >
            <Form
                form={form}
                layout="vertical"
                name="form_in_modal"
                initialValues={{ modifier: 'public' }}
            >
                <Form.Item
                    name="title"
                    label="标题"
                    rules={[{ required: true, message: '请输入标题!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item name="desc" label="简介">
                    <Input type="textarea" />
                </Form.Item>
                <Form.Item
                    name="order"
                    label="第几话"
                    rules={[{ required: true, message: '请输入第几话!' }]}
                >
                    <InputNumber min={0} max={50} />
                </Form.Item>
                <Form.Item
                    name="url"
                    label="视频地址"
                    rules={[{ required: true, message: '请输入视频地址!' }]}
                >
                    <Input />
                </Form.Item>

            </Form>
        </Modal>
    );
};