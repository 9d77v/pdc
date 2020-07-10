import React from 'react';
import './App.css';
import { Spin, ConfigProvider } from 'antd';
import { ApolloProvider } from '@apollo/react-hooks';
import zhCN from 'antd/es/locale/zh_CN';
import { client } from './utils/client';

const DesktopIndex = React.lazy(() => import('./profiles/desktop/index'))

function App() {
  return (
    <ApolloProvider client={client}>
      <React.Suspense fallback={<Spin />}>
        <ConfigProvider locale={zhCN}>
          <DesktopIndex />
        </ConfigProvider>
      </React.Suspense>
    </ApolloProvider>
  );
}

export default App;
