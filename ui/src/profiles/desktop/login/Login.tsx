import React from 'react';
import "./Login.less"
import { LoginForm } from '../../../components/LoginForm';
import { useHistory } from 'react-router-dom';

export const Login = () => {

    const token = localStorage.getItem('accessToken');
    const history = useHistory();
    if (token) {
        history.push('/app/home')
    }
    return (
        <div className="login-background">
            <div className="login-form">
                <div className={'title'}>个人数据中心</div>
                <LoginForm />
            </div>
        </div>
    );
};