import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { RouterProvider } from 'react-router-dom';
import { CssBaseline, CssVarsProvider } from '@mui/joy';
import { theme } from './utils/theme'
import { router } from './utils/router';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

root.render(
  <React.StrictMode>
    <CssVarsProvider defaultMode="dark" theme={theme}>
      <CssBaseline/>
    <RouterProvider router={router}/>
    </CssVarsProvider>
  </React.StrictMode>
);

