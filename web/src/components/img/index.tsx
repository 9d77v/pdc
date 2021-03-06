import React, { useState, useRef } from "react"

import "./index.less"
import { Modal, Progress } from "antd"
import useIntersectionObserver from "src/hooks/use-intersection-observer"
import { formatTimeLength } from "src/utils/util"
interface ImageProps {
    src: string
    width?: number | string
    height?: number | string
    currentTime?: number,
    remainingTime?: number,
    showModal?: boolean
}

const Img: React.FC<ImageProps> = ({
    src,
    width,
    height,
    currentTime,
    remainingTime,
    showModal = false
}) => {
    const [visible, setVisible] = useState(false)
    const ref: any = useRef();
    const [isVisible, setIsVisible] = useState(false);

    useIntersectionObserver({
        target: ref,
        onIntersect: ([{ isIntersecting }]: any, observerElement: any) => {
            if (isIntersecting) {
                setIsVisible(true);
                observerElement.unobserve(ref.current);
            }
        }
    })
    return (<div
        ref={ref}
        className={"img-box"}
        style={{
            maxHeight: height ? height : 214,
            width: width ? width : 160,
            position: "relative",
            overflowY: "hidden"
        }}>
        {showModal ? <Modal
            title="查看图片"
            visible={visible}
            destroyOnClose={true}
            maskClosable={true}
            onCancel={() => setVisible(false)}
            footer={null}
        >
            {src ? <img src={src} alt="图片不存在" /> : "暂无图片"}
        </Modal> : null}
        {src ? isVisible && (<div style={{ width: "100%", height: "100%" }}>
            <img src={src}
                alt="图片不存在"
                onClick={() => setVisible(true)} />
            {currentTime ?
                <div>
                    <Progress percent={currentTime / (currentTime + (remainingTime || 0)) * 100} showInfo={false} />
                    <div style={{
                        position: "absolute",
                        bottom: 10,
                        right: 10,
                        color: "white",
                        opacity: 0.5,
                        backgroundColor: "black"
                    }}>{formatTimeLength(currentTime)}/{formatTimeLength(currentTime + (remainingTime || 0))}</div>
                </div> : ''}

        </div>
        ) : "暂无图片"}
    </div>)
}
export default Img
