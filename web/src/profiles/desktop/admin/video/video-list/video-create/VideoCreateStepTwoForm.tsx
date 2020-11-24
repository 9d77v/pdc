import { Form } from 'antd'
import React, { forwardRef, Ref, useImperativeHandle, useState } from 'react'
import { Uploader } from 'src/components/Uploader'


interface Values {
    title: string
    description: string
}

interface IVideoCreateStepTwoFormProps {
    id: number
    onCreate?: (values: Values) => void
}

const VideoCreateStepTwoForm = (props: IVideoCreateStepTwoFormProps, ref: Ref<any>) => {
    const [form] = Form.useForm()
    const getForm = () => {
        return form
    }
    const getVideoURLs = () => {
        return videoURLs
    }
    const resetVideoURLS = () => {
        setVideoURLs([])
    }
    useImperativeHandle(ref, () => ({
        getForm,
        getVideoURLs,
        resetVideoURLS
    }))

    const [videoURLs, setVideoURLs] = useState([])
    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }

    const videoPathPrefix = props.id.toString() + "/"
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

export default forwardRef(VideoCreateStepTwoForm)