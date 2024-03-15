import * as React from 'react';

import {
    RouterProvider,
} from "react-router-dom";

import generateRoutes from './routes';

export default function App() {
    return (
        <RouterProvider router={generateRoutes()}/>
    );
}
