import { Outlet } from 'react-router-dom'
import './App.css'
import { AppSidebar } from './AppSidebar'
import { SidebarInset, SidebarProvider, SidebarTrigger } from './ui/sidebar'
import { ThemeProvider } from './ui/theme-provider'
import { Separator } from '@radix-ui/react-separator'

function App() {

  return (
    <ThemeProvider defaultTheme='dark' storageKey="vite-ui-theme">
      <SidebarProvider>
        <AppSidebar />
        {/* <SidebarTrigger /> */}
        <SidebarInset>
          <main className="flex-1 flex flex-col overflow-hidden">
            <header className="flex h-16 shrink-0 items-center gap-2 border-b px-4">
              <SidebarTrigger className="-ml-1" />
              <Separator orientation="vertical" className="mr-2 h-4" />
              {/* <Breadcrumb>
                <BreadcrumbList>
                  <BreadcrumbItem className="hidden md:block">
                    <BreadcrumbLink href="#">
                      Building Your Application
                    </BreadcrumbLink>
                  </BreadcrumbItem>
                  <BreadcrumbSeparator className="hidden md:block" />
                  <BreadcrumbItem>
                    <BreadcrumbPage>Data Fetching</BreadcrumbPage>
                  </BreadcrumbItem>
                </BreadcrumbList>
              </Breadcrumb> */}
            </header>

            <Outlet />
          </main>
        </SidebarInset>
      </SidebarProvider>
    </ThemeProvider>
  )
}

export default App
