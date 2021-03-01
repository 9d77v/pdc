import { Button } from 'antd';
import React, { FC } from 'react';

export interface ICircleButtonProps {
    right: number,
    bottom: number,
    radius: number,
    display: string,
    icon: JSX.Element,
    onClick: () => void,
}

const CircleButton: FC<ICircleButtonProps> = ({
    bottom, radius, right, display, icon, onClick
}) => {

    return (<div>
        <Button type="primary" style={{
            borderRadius: '50%',
            position: 'fixed',
            zIndex: 1,
            right, bottom,
            width: radius,
            height: radius,
            justifyContent: 'center',
            alignItems: 'center',
            display,
            fontSize: 'xx-large',
            boxShadow: '5px 5px 5px darkgrey',

        }}
            icon={icon}
            onClick={onClick}
        />
    </div>);
}

export default CircleButton
