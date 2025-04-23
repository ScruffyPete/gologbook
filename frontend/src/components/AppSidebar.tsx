import {
    Sidebar,
    SidebarContent,
    SidebarGroup,
    SidebarGroupAction,
    SidebarGroupContent,
    SidebarGroupLabel,
    SidebarMenu,
    SidebarMenuButton,
    SidebarMenuItem,
} from "@/components/ui/sidebar"
import { Project } from "@/types/Project"
import { Plus } from "lucide-react"
import React, { useEffect, useRef, useState } from "react"
import { Link } from "react-router-dom"
import { Input } from "./ui/input"

export function AppSidebar() {
    const [projects, setProjects] = useState<Project[]>([])
    const [newProjectTitle, setNewProjectTitle] = useState('')
    const [isCreating, setIsCreating] = useState(false)

    const inputRef = useRef<HTMLInputElement>(null)

    useEffect(() => { loadProjects() }, [])

    useEffect(() => {
        if (isCreating && inputRef.current) {
            inputRef.current.focus()
        }
    }, [isCreating])

    async function loadProjects() {
        try {
            const res = await fetch("/api/projects")
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
            const res = await fetch('/api/projects', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ title: newProjectTitle }),
            })

            if (!res.ok) throw new Error('Failed to cretae project')

            const newProject = await res.json()
            setProjects(prev => [...prev, newProject])

            await loadProjects()

            setNewProjectTitle('')
            setIsCreating(false)
        } catch (err) {
            console.error(err)
            alert('Could not create project')
        }
    }

    return (
        <Sidebar>
            <SidebarContent>
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
                                <SidebarMenuItem key={project.id}>
                                    <SidebarMenuButton asChild>
                                        <Link to={`/projects/${project.id}`}>{project.title}</Link>
                                    </SidebarMenuButton>
                                </SidebarMenuItem>
                            ))}
                        </SidebarMenu>
                    </SidebarGroupContent>
                </SidebarGroup>
            </SidebarContent>
        </Sidebar>
    )
}
