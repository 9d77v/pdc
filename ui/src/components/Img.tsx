import React, { useState, useRef } from "react"

import "./img.less"
import { Modal } from "antd"
import useIntersectionObserver from "../hooks/use-intersection-observer"
interface ImageProps {
    src: string
    width?: number | string
    height?: number | string
}

export const Img: React.FC<ImageProps> = ({
    src,
    width,
    height
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
            height: height ? height : 214,
            width: width ? width : 160
        }}>
        <Modal
            title="查看图片"
            visible={visible}
            destroyOnClose={true}
            maskClosable={true}
            onCancel={() => setVisible(false)}
            footer={null}
        >
            {src ? <img src={src} alt="图片不存在" /> : "暂无图片"}
        </Modal>
        {src ? isVisible && (<img src={src} alt="图片不存在" onClick={() => setVisible(true)} />) : "暂无图片"}
    </div>)
}
