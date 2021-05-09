import React from "react"

interface ICardProps {
    title: string
    width: string | number
    height?: string | number
    cardItems: JSX.Element[]
    onClick?: ((event: React.MouseEvent<HTMLDivElement, MouseEvent>) => void) | undefined
}

const Card: React.FC<ICardProps> = ({
    title,
    width,
    height,
    cardItems,
    onClick
}) => {
    return (
        <div
            className="pdc-card-default"
            style={{
                height: height,
                width: width,
                float: "left",
                paddingTop: 10
            }}
            onClick={onClick}
        >
            <div style={{ fontSize: 22 }}>
                {title}
            </div>
            <div style={{
                textAlign: "left",
                padding: 10,
                paddingLeft: 20
            }}>
                {cardItems}
            </div>
        </div >
    )
}
export default Card
