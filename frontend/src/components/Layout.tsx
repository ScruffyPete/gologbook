import { Navigate, Outlet } from 'react-router-dom'
import { AppSidebar } from './AppSidebar'
import { SidebarInset, SidebarProvider, SidebarTrigger } from './ui/sidebar'
import { Separator } from '@radix-ui/react-separator'


export function LayoutWrapper() {
    const token = localStorage.getItem('token')
    if (!token) {
        return <Navigate to="/login" />
    }
    return <Layout />
}

export function Layout() {

    return (
        <SidebarProvider>
            <AppSidebar />
            <SidebarInset>
                <main className="flex-1 flex flex-col overflow-hidden">
                    <header className="flex h-16 shrink-0 items-center gap-2 border-b px-4">
                        <SidebarTrigger className="-ml-1" />
                        <Separator orientation="vertical" className="mr-2 h-4" />
                    </header>

                    <Outlet />
                </main>
            </SidebarInset>
        </SidebarProvider>
    )
}