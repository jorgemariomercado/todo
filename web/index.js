import * as React from 'react';
import * as ReactDOM from 'react-dom/client';

import App from './App';

const rootElement = document.getElementById('root');
const root = ReactDOM.createRoot(rootElement);

if (window.location.host.includes('localhost') ) {
    new EventSource('/esbuild').addEventListener(
        'change',
        () => location.reload())
}

root.render(
    <App/>
);
