import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './components/App.tsx'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import { loader as projectLoader, ProjectPage } from './components/ProjectPage.tsx'
import { LoginPage } from './components/LoginPage.tsx'
import { SignupPage } from './components/SignupPage.tsx'
import { LayoutWrapper } from './components/Layout.tsx'
import { RootPage } from './components/RootPage.tsx'

const router = createBrowserRouter([
  {
    path: '/',
    Component: App,
    children: [
      {
        index: true,
        Component: RootPage,
      },
      {
        path: '/login',
        Component: LoginPage,
      },
      {
        path: '/signup',
        Component: SignupPage,
      },
      {
        path: '/projects',
        Component: LayoutWrapper,
        children: [
          {
            path: '/projects/:projectId',
            Component: ProjectPage,
            loader: projectLoader,
          },
        ]
      }
    ]
  }
])

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
