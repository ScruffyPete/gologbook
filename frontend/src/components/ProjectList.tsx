import { Project } from "@/types/Project";
import { SidebarMenuButton } from "./ui/sidebar";
import { NavLink, useMatch } from "react-router-dom";

export function ProjectItem({ project }: { project: Project }) {
    const match = useMatch('/projects/:projectId')

    return (
        <SidebarMenuButton asChild>
            <NavLink
                className={
                    match?.params.projectId === project.id
                        ? "bg-muted font-medium"
                        : ""
                }
                to={`/projects/${project.id}`}
            >
                {project.title}
            </NavLink>
        </SidebarMenuButton>
    )
}