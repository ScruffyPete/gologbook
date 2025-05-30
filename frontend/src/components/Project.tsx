import { useEffect, useState } from "react";
import { Project as ProjectType } from "../types/Project";
import { useNavigate, useParams } from "react-router-dom";
import EntryListWrapper from "./EntryList";


export default function ProjectWrapper() {
    const { projectId } = useParams()
    const [project, setProject] = useState<ProjectType | null>(null)
    const navigate = useNavigate()

    useEffect(() => {
        fetch(`/api/projects/${projectId}`)
            .then(res => res.json())
            .then(setProject)
            .catch(err => console.error('Failed to load project:', err))
    }, [projectId])

    if (!projectId) return <div>Missing project ID</div>
    if (!project) return <div>Loading...</div>

    async function handleDeleteProject() {
        try {
            await fetch(`/api/projects/${projectId}`, { method: 'DELETE' })
            navigate('/')
        } catch (err) {
            console.error('Failed to delete project:', err)
            alert('Could not delete project')
        }
    }

    return <>
        <button onClick={() => navigate(-1)}>← Back</button>
        <button onClick={handleDeleteProject}>Delete</button>
        <Project project={project}></Project>
    </>
}

export function Project({ project }: { project: ProjectType }) {

    return (
        <>
            <header style={{ marginBottom: '1rem' }}>
                <h2>{project.title}</h2>
            </header>
            <EntryListWrapper />
        </>

    )
}
