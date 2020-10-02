import { Icon, NavBar } from 'antd-mobile'
import React from 'react'
import { useHistory } from 'react-router-dom'

interface IScannerResultProps {
    url: string
}

export const ScannerResult = (props: IScannerResultProps) => {
    const history = useHistory()
    return (
        <div style={{ height: "100%" }}>
            <NavBar
                mode="light"
                icon={<Icon type="left" />}
                onLeftClick={() => history.goBack()}
            >
                {props.url}
            </NavBar>
            <iframe src={props.url} title={props.url} width={"100%"} height={"100%"} />
        </div>
    )
}