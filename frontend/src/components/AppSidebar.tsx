import {
    Sidebar,
    SidebarContent,
    SidebarGroup,
    SidebarGroupAction,
    SidebarGroupContent,
    SidebarGroupLabel,
    SidebarMenu,
    SidebarMenuItem,
} from "@/components/ui/sidebar"
import { Project } from "@/types/Project"
import { Plus } from "lucide-react"
import React, { useEffect, useRef, useState } from "react"
import { Input } from "./ui/input"
import { ProjectItem } from "./ProjectList"
import { useNavigate } from "react-router-dom"
import { Button } from "./ui/button"

export function AppSidebar() {
    const [projects, setProjects] = useState<Project[]>([])
    const [newProjectTitle, setNewProjectTitle] = useState('')
    const [isCreating, setIsCreating] = useState(false)

    const navigate = useNavigate()

    const inputRef = useRef<HTMLInputElement>(null)

    useEffect(() => { loadProjects() }, [])

    useEffect(() => {
        if (isCreating && inputRef.current) {
            inputRef.current.focus()
        }
    }, [isCreating])

    async function loadProjects() {
        try {
            const res = await fetch("/api/projects/", {
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                }
            })
            if (!res.ok) throw new Error("Failed to fetch projects")

            const data = await res.json()
            setProjects(data)
        } catch (err) {
            console.error("Error loading projects:", err)
            alert("Could not load projects")
        }
    }

    async function handleCreateProject(e: React.FormEvent) {
        e.preventDefault()

        try {
            const res = await fetch('/api/projects/', {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem('token')}`,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ title: newProjectTitle }),
            })

            if (!res.ok) throw new Error('Failed to cretae project')

            const newProject = await res.json()
            setProjects(prev => [...prev, newProject])

            navigate(`/projects/${newProject.id}`)

            await loadProjects()

            setNewProjectTitle('')
            setIsCreating(false)
        } catch (err) {
            console.error(err)
            alert('Could not create project')
        }
    }

    async function handleLogout() {
        localStorage.removeItem('token')
        navigate('/login')
    }

    return (
        <Sidebar>
            <SidebarContent>
                <SidebarGroup>
                    <SidebarGroupLabel>User</SidebarGroupLabel>
                    <SidebarGroupContent>
                        <SidebarMenu>
                            <SidebarMenuItem>
                                <Button variant="outline" onClick={handleLogout}>Logout</Button>
                            </SidebarMenuItem>
                        </SidebarMenu>
                    </SidebarGroupContent>
                </SidebarGroup>
                <SidebarGroup>
                    <SidebarGroupLabel>Projects</SidebarGroupLabel>
                    <SidebarGroupAction onClick={() => setIsCreating(true)}>
                        <Plus /> <span className="sr-only">Add Project</span>
                    </SidebarGroupAction>
                    <SidebarGroupContent>
                        <SidebarMenu>
                            {isCreating && (
                                <SidebarMenuItem>
                                    <form onSubmit={handleCreateProject}>
                                        <Input
                                            ref={inputRef}
                                            value={newProjectTitle}
                                            onChange={(e) => setNewProjectTitle(e.target.value)}
                                        />
                                    </form>
                                </SidebarMenuItem>
                            )}
                            {projects.map((project) => (
                                <SidebarMenuItem className="h-9" key={project.id}>
                                    <ProjectItem project={project} />
                                </SidebarMenuItem>
                            ))}
                        </SidebarMenu>
                    </SidebarGroupContent>
                </SidebarGroup>
            </SidebarContent>
        </Sidebar>
    )
}
