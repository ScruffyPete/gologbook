import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './components/App.tsx'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import { projectLoader, ProjectPage } from './components/ProjectPage.tsx'
import { LoginPage } from './components/LoginPage.tsx'
import { SignupPage } from './components/SignupPage.tsx'
import { LayoutWrapper } from './components/Layout.tsx'
import { authLoader, redirectIfAuthedLoader, rootRedirectLoader } from './lib/auth.ts'

const router = createBrowserRouter([
  {
    path: '/',
    Component: App,
    children: [
      {
        index: true,
        loader: rootRedirectLoader,
      },
      {
        path: '/login',
        Component: LoginPage,
        loader: redirectIfAuthedLoader,
      },
      {
        path: '/signup',
        Component: SignupPage,
      },
      {
        path: '/projects',
        Component: LayoutWrapper,
        loader: authLoader,
        children: [
          {
            path: ':projectId',
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
