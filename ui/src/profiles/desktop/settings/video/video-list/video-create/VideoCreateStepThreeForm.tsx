import { Form, message, Radio } from 'antd';
import React, { forwardRef, Ref, useImperativeHandle, useState } from 'react'
import { Uploader } from '../../../../../../components/Uploader';

interface Values {
    title: string;
    description: string;
}

interface VideoCreateStepThreeFormProps {
    onCreate?: (values: Values) => void;
}

const VideoCreateStepThreeForm  =  (props:VideoCreateStepThreeFormProps,ref:  any)=>{
    const [form] = Form.useForm();
    const onFinish = () => {
        return form;
      }

      useImperativeHandle(ref, () => ({
        onFinish,
      }))
    const [url, setUrl] = useState("")
    const [videoURLs, setVideoURLs] = useState([])
    const [subtitles, setSubtitles] = useState([])
    const [maxNum, setMaxNum] = useState(0)
    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }
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
    )
}

export default forwardRef(VideoCreateStepThreeForm);
