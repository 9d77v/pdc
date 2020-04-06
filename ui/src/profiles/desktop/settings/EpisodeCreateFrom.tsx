import { Modal, Form, Input, InputNumber } from 'antd';
import React, { useState } from 'react'
import { SingleUploader } from '../../components/Uploader';
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
    const [url, setUrl] = useState('')

    return (
        <Modal
            visible={visible}
            title="新增分集"
            okText="确定"
            cancelText="取消"
            onCancel={onCancel}
            onOk={() => {
                form.setFieldsValue({
                    "url": url
                })
                form
                    .validateFields()
                    .then((values: any) => {
                        form.resetFields();
                        onCreate(values);
                    })
                    .catch(info => {
                        console.log('Validate Failed:', info);
                    });
                setUrl('')
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
                >
                    <Input />
                </Form.Item>
                <Form.Item name="num" label="集数"
                    rules={[{ required: true, message: '请输入集数!' }]}
                >
                    <InputNumber />
                </Form.Item>
                <Form.Item name="desc" label="简介">
                    <Input type="textarea" />
                </Form.Item>
                <Form.Item
                    name="url"
                    label="上传视频"
                    rules={[{ required: true, message: '请上传视频!' }]}
                >
                    <SingleUploader
                        bucketName="video"
                        validFileTypes={["video/mp4"]}
                        setURL={setUrl}
                    />
                </Form.Item>
            </Form>
        </Modal >
    );
};