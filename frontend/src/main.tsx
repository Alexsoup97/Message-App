import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { createBrowserRouter,RouterProvider } from 'react-router-dom';
import  { SignIn } from './pages/SignIn'
import { CssBaseline, CssVarsProvider } from '@mui/joy';
import { theme } from './utils/theme'
const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);


const router = createBrowserRouter([
    {
      path: "/",
      element: <SignIn/>
    }])

root.render(
  <React.StrictMode>
    <CssVarsProvider defaultMode="dark" theme={theme}>
      <CssBaseline/>
    <RouterProvider router={router}/>
    </CssVarsProvider>
  </React.StrictMode>
);

