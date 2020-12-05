import { useMutation } from '@apollo/react-hooks'
import { Button, message, Steps } from 'antd'
import React, { useRef, useState } from 'react'
import { useHistory } from 'react-router-dom'
import { ADD_VIDEO, ADD_VIDEO_RESOURCE, SAVE_SUBTITLES } from 'src/gqls/video/mutation'
import "./index.less"
import VideoCreateStepOneForm from './VideoCreateStepOneForm'
import VideoCreateStepThreeForm from './VideoCreateStepThreeForm'
import VideoCreateStepTwoForm from './VideoCreateStepTwoForm'
const { Step } = Steps

export default function VideoCreateIndex() {
    const [addVideo] = useMutation(ADD_VIDEO)
    const [addVideoResource] = useMutation(ADD_VIDEO_RESOURCE)
    const [saveSubtitles] = useMutation(SAVE_SUBTITLES)

    const [current, setCurrent] = useState(0)
    const [currentVideoID, setCurrentVideoID] = useState(0)
    const [videoURLs, setVideoURLs] = useState([])
    const history = useHistory()
    const stepOneRef = useRef()
    const stepTwoRef = useRef()
    const stepThreeRef = useRef()

    const steps = [
        {
            title: '基本信息',
            content: <VideoCreateStepOneForm ref={stepOneRef} />,
        },
        {
            title: '上传视频',
            content: <VideoCreateStepTwoForm id={currentVideoID} ref={stepTwoRef} />,
        },
        {
            title: '上传字幕',
            content: <VideoCreateStepThreeForm ref={stepThreeRef} />,
        }
    ]

    const handleStepOne = () => {
        const stepOne: any = stepOneRef.current
        const form = stepOne.getForm()
        form.setFieldsValue({
            "cover": stepOne.getURL(),
        })
        form.validateFields()
            .then(async (values: any) => {
                const data = await addVideo({
                    variables: {
                        "input": {
                            "title": values.title,
                            "desc": values.desc,
                            "cover": values.cover,
                            "pubDate": values.pubDate ? values.pubDate.unix() : 0,
                            "tags": values.tags || [],
                            "isShow": values.isShow,
                            "isHideOnMobile": values.isHideOnMobile,
                            "theme": values.theme
                        }
                    }
                })
                if (data) {
                    const id: number = data.data?.createVideo.id
                    setCurrentVideoID(id)
                    setCurrent(current + 1)
                } else {
                    message.error("新建视频出错")
                }
            })
            .catch((info: any) => {
                console.log('Validate Failed:', info)
            })
    }

    const handleStepTwo = () => {
        const stepTwo: any = stepTwoRef.current
        const form = stepTwo.getForm()
        form.setFieldsValue({
            "videoURLs": stepTwo.getVideoURLs(),
        })
        form
            .validateFields()
            .then(async (values: any) => {
                const data = await addVideoResource({
                    variables: {
                        "input": {
                            "id": currentVideoID,
                            "videoURLs": values.videoURLs,
                        }
                    }
                })
                if (data) {
                    setVideoURLs(stepTwo.getVideoURLs)
                    setCurrent(current + 1)
                } else {
                    message.error("添加视频资源出错")
                }
            })
            .catch((info: any) => {
                console.log('Validate Failed:', info)
            })
        stepTwo.resetVideoURLS()
    }

    const handleStepThree = () => {
        const stepThree: any = stepThreeRef.current
        const form = stepThree.getForm()
        const subtitles = stepThree.getSubtitles()
        form.setFieldsValue({
            "subtitles": subtitles
        })
        if (subtitles.length > 0 && videoURLs.length !== subtitles.length) {
            message.error(`字幕数量与视频数量不一致,视频数量：${videoURLs.length},字幕数量：${subtitles.length}`)
            return
        }
        form
            .validateFields()
            .then(async (values: any) => {
                let subtitles = undefined
                if (values.subtitles && values.subtitles.length > 0) {
                    subtitles = {
                        "name": values.subtitle_lang,
                        "urls": values.subtitles
                    }
                    const data = await saveSubtitles({
                        variables: {
                            "input": {
                                "id": currentVideoID,
                                "subtitles": subtitles
                            }
                        }
                    })
                    if (data) {
                        message.success('视频新建完成!')
                        history.goBack()
                    } else {
                        message.error("添加视频字幕出错")
                    }
                } else {
                    message.success('视频新建完成!')
                    history.goBack()
                }

            })
            .catch((info: any) => {
                console.log('Validate Failed:', info)
            })
        stepThree.resetSubtitles()
    }

    const nextFunc = new Map<Number, any>([
        [0, handleStepOne],
        [1, handleStepTwo],
    ])

    return (
        <div style={{ display: "flex", flexDirection: "column", padding: 16 }}>
            <Button type="primary"
                style={{ width: 100, margin: 6 }}
                onClick={() => history.goBack()}>
                返回
                    </Button>
            < Steps current={current} >
                {
                    steps.map(item => (
                        <Step key={item.title} title={item.title} />
                    ))
                }
            </Steps >
            <div className="steps-content">{steps[current].content}</div>
            <div className="steps-action">
                {current < steps.length - 1 && (
                    <Button type="primary" style={{ float: "right" }}
                        onClick={nextFunc.get(current)}>
                        下一步
                    </Button>
                )}
                {current === steps.length - 1 && (
                    <Button type="primary" style={{ float: "right" }}
                        onClick={handleStepThree}>
                        完成
                    </Button>
                )}
            </div>
        </div >
    )
}
