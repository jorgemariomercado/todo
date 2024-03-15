import {
  createBrowserRouter
} from "react-router-dom";

import MainRoutes from './MainRoutes';

export default function generateRoutes() {
  return createBrowserRouter([MainRoutes]);
}