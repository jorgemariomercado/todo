import MainLayout from '../layout';
import {ErrorBoundary} from "../components/ErrorBoundary";

const MainRoutes = {
    path: '/',
    element: <MainLayout/>,
    errorElement: <ErrorBoundary/>,

    children: [
        {
            path: "ui/tasks",
            lazy: () => import('../components/Tasks'),
        },
        {
            path: "ui/reports",
            lazy: () => import('../components/Reports'),
        },
    ]
};

export default MainRoutes;