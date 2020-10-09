import { Form, message } from 'antd';
import React, { forwardRef, Ref, useImperativeHandle, useState } from 'react'
import { Uploader } from '../../../../../../components/Uploader';

interface Values {
    title: string;
    description: string;
}

interface VideoCreateStepTwoFormProps {
    onCreate?: (values: Values) => void;
}

const VideoCreateStepTwoForm  =  (props:VideoCreateStepTwoFormProps,ref:  any)=>{
    const [form] = Form.useForm();
    const onFinish = () => {
        return form;
      }

      useImperativeHandle(ref, () => ({
        onFinish,
      }));
    const [url, setUrl] = useState("")
    const [videoURLs, setVideoURLs] = useState([])
    const [subtitles, setSubtitles] = useState([])
    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }
    const maxNum = 1
    const videoPathPrefix = "desktop/" + maxNum.toString() + "/"

    const create = () => {
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
                // onCreate(values);
            })
            .catch(info => {
                console.log('Validate Failed:', info);
            });
        setUrl('')
        setVideoURLs([])
        setSubtitles([])
    }
    return (
        <Form
            {...layout}
            form={form}
            layout="horizontal"
            name="videoCreateForm"
            initialValues={{ isShow: true, subtitle_lang: "简体中文" }}
        >
            <Form.Item
                name="videoURLs"
                label="视频列表"
                rules={[{ required: true, message: '请上传视频!' }]}
            >
                <Uploader
                    fileLimit={0}
                    bucketName="video"
                    filePathPrefix={videoPathPrefix}
                    validFileTypes={["video/mp4"]}
                    setURL={setVideoURLs}
                />
            </Form.Item>
        </Form>
    )
}

export default forwardRef(VideoCreateStepTwoForm);
