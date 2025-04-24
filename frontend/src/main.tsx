import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './components/App.tsx'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import { loader as projectLoader, ProjectPage } from './components/ProjectPage.tsx'

const router = createBrowserRouter([
  {
    path: '/',
    Component: App,
    children: [
      {
        path: '/projects/:projectId',
        Component: ProjectPage,
        loader: projectLoader,
      }
    ]
  },

]);

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
