import { Modal, Form, Input, Switch, DatePicker, message, Radio, Select } from 'antd';
import React, { useState } from 'react'
import { Uploader } from '../../../../components/Uploader';

const { TextArea } = Input;

interface Values {
    title: string;
    description: string;
}

interface VideoCreateFormProps {
    visible: boolean;
    onCreate: (values: Values) => void;
    onCancel: () => void;
}

export const VideoCreateForm: React.FC<VideoCreateFormProps> = ({
    visible,
    onCreate,
    onCancel,
}) => {
    const [form] = Form.useForm();
    const [url, setUrl] = useState("")
    const [videoURLs, setVideoURLs] = useState([])
    const [subtitles, setSubtitles] = useState([])
    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }
    return (
        <Modal
            visible={visible}
            title="新增视频"
            okText="确定"
            cancelText="取消"
            onCancel={
                () => {
                    onCancel()
                    form.resetFields()
                    setUrl('')
                    setVideoURLs([])
                    setSubtitles([])
                }
            }
            getContainer={false}
            onOk={() => {
                form.setFieldsValue({
                    "cover": url,
                    "videoURLs": videoURLs,
                    "subtitles": subtitles
                })
                if (subtitles.length > 0 && videoURLs.length !== subtitles.length) {
                    message.error(`视频数量与字幕数量不一致,视频数量：${videoURLs.length},字幕数量：${subtitles.length}`);
                    return
                }
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
                setVideoURLs([])
                setSubtitles([])
            }}
        >
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="videoCreateForm"
                initialValues={{ isShow: true, subtitle_lang: "简体中文" }}
            >
                <Form.Item
                    name="title"
                    label="标题"
                    rules={[{ required: true, message: '请输入标题!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item name="desc" label="简介">
                    <TextArea rows={4} />
                </Form.Item>
                <Form.Item name="cover" label="封面">
                    <Uploader
                        fileLimit={1}
                        bucketName="image"
                        validFileTypes={["image/jpeg", "image/png", "image/webp"]}
                        setURL={setUrl}
                    />
                </Form.Item>
                <Form.Item name="pubDate" label="上映时间">
                    <DatePicker />
                </Form.Item>
                <Form.Item name="tags" label="标签">
                    <Select
                        mode="tags"
                        size={"large"}
                        style={{ width: '100%' }}
                    >
                    </Select>
                </Form.Item>
                <Form.Item name="isShow" label="是否显示" valuePropName='checked'>
                    <Switch />
                </Form.Item>
                <Form.Item
                    name="videoURLs"
                    label="视频列表"
                    rules={[{ required: true, message: '请上传视频!' }]}
                >
                    <Uploader
                        fileLimit={0}
                        bucketName="video"
                        validFileTypes={["video/mp4"]}
                        setURL={setVideoURLs}
                    />
                </Form.Item>
                <Form.Item name="subtitle_lang" label="字幕语言">
                    <Radio.Group buttonStyle="solid">
                        <Radio.Button value="简体中文">简体中文</Radio.Button>
                        <Radio.Button value="中日双语">中日双语</Radio.Button>
                    </Radio.Group>
                </Form.Item>
                <Form.Item name="subtitles" label="字幕列表">
                    <Uploader
                        fileLimit={0}
                        bucketName="vtt"
                        validFileTypes={["text/vtt", "text/ass", 'text/srt']}
                        setURL={setSubtitles}
                    />
                </Form.Item>
            </Form>
        </Modal>
    );
};