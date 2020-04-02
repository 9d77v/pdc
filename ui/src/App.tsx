import React from 'react';
import './App.css';
import { Spin, ConfigProvider } from 'antd';
import ApolloClient from 'apollo-boost';
import { ApolloProvider } from '@apollo/react-hooks';
import zhCN from 'antd/es/locale/zh_CN';

const DesktopIndex = React.lazy(() => import('./profiles/desktop/index'))

const client = new ApolloClient({
  uri: '/api',
});

function App() {
  return (
    <ApolloProvider client={client}>
      <div className="App">
        <React.Suspense fallback={<Spin />}>
          <ConfigProvider locale={zhCN}>
            <DesktopIndex />
          </ConfigProvider>
        </React.Suspense>
      </div>
    </ApolloProvider>
  );
}

export default App;
