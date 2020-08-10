import React, { useEffect, useState } from "react"
import { useHistory } from 'react-router-dom';
import { message } from "antd"
import { useQuery } from "@apollo/react-hooks";
import { LIST_VIDEO_CARD } from '../../../../consts/video.gql';
import { Img } from "../../../../components/Img";

export default function VideoList() {

    const [cards, setCards] = useState(<div />)
    const { error, data } = useQuery(LIST_VIDEO_CARD,
        {
            variables: {
                page: 1,
                pageSize: -1,
                sorts: [{
                    field: "title",
                    isAsc: true
                }]
            }
        });

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const history = useHistory()
    useEffect(() => {
        if (data && data.videos.edges) {
            const videos = data.videos.edges
            setCards(videos.map((item: any) =>
                <div key={item.id}
                    onClick={() => history.push('/app/media/videos/' + item.id)}
                    style={{
                        width: "30%",
                        margin: "2.5% 0 0 2.5%",
                        height: 210,
                        display: "flex",
                        float: "left",
                        flexDirection: "column"
                    }}
                >
                    <Img src={item.cover} width={"100%"} height={"70%"} />
                    <div style={{
                        fontSize: 12,
                        height: 36,
                        overflow: "hidden",
                        textOverflow: "ellipsis"
                    }}>{item.title}</div>
                    <div style={{ fontSize: 10 }}>全{item.episodes ? item.episodes.length : 0}话</div>
                </div >
            ))
        }
    }, [data, history])
    return (
        <div >
            {cards}
        </div>
    )
}