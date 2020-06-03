import React, { useEffect, useState } from "react"
import { message } from "antd"
import "./video.less"
import { useQuery } from "@apollo/react-hooks";
import { LIST_VIDEO_CARD } from '../../../consts/video.gql';
import { Img } from "../../../components/Img";

export const VideoPage = () => {

    const [cards, setCards] = useState(<div />)
    const { error, data } = useQuery(LIST_VIDEO_CARD,
        {
            variables: {
                page: 1,
                pageSize: 10
            }
        });

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    useEffect(() => {
        if (data && data.Videos.edges) {
            const videos = data.Videos.edges
            setCards(videos.map((item: any) =>
                <div key={item.id}
                    onClick={() => console.log('clicked')}
                    className={"card"}
                >
                    <Img src={item.cover} />
                    <div>{item.title}</div>
                    <div style={{ color: "#99a2aa" }}>全{item.episodes ? item.episodes.length : 0}话</div>
                </div >
            ))
        }
    }, [data])
    return (<div>{cards}</div>)
}