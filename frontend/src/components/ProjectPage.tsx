import { Entry } from "@/types/Entry";
import { Project } from "@/types/Project";
import { LoaderFunctionArgs, useLoaderData } from "react-router-dom";
import { useEffect, useRef, useState } from "react";
import { EntryFeed } from "./EntryFeed";
import { ProjectLayout } from "./ProjectLayout";
import { OutputStream } from "./OutputStream";


export async function projectLoader({ params }: LoaderFunctionArgs) {
    try {
        const [projectRes, entriesRes] = await Promise.all([
            fetch(`/api/projects/${params.projectId}`, {
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                }
            }),
            fetch(`/api/entries/?project_id=${params.projectId}`, {
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                }
            }),
        ]);

        if (!projectRes.ok || !entriesRes.ok) {
            throw new Error('Failed to load data');
        }

        const [loaderProject, loaderEntries] = await Promise.all([
            projectRes.json(),
            entriesRes.json()
        ]);

        return { loaderProject, loaderEntries };
    } catch (error) {
        console.error(error);
        throw new Response('Failed to load project', { status: 500 });
    }
}

export function ProjectPage() {
    const { loaderProject, loaderEntries } = useLoaderData<{ loaderProject: Project, loaderEntries: Entry[] }>()

    const [project, setProject] = useState<Project>(loaderProject)
    const [entries, setEntries] = useState<Entry[]>(loaderEntries)
    const [newEntryBody, setNewEntryBody] = useState("")
    const [streamKey, setStreamKey] = useState(0)

    const scrollAnchorRef = useRef<HTMLDivElement | null>(null)
    const textareaRef = useRef<HTMLTextAreaElement>(null)

    useEffect(() => {
        setProject(loaderProject)
    }, [loaderProject])

    useEffect(() => {
        setEntries(loaderEntries)
        setStreamKey(0)
    }, [loaderProject.id, loaderEntries])

    useEffect(() => {
        scrollAnchorRef.current?.scrollIntoView({ behavior: "smooth" })
    }, [entries])

    const handleSubmit = async () => {
        const entry = newEntryBody.trim()
        try {
            const res = await fetch(`/api/entries/`, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem('token')}`,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ body: entry, project_id: project.id })
            })
            const newEntry = await res.json()
            setEntries((prev) => [...prev, newEntry])
            setNewEntryBody("")
            setStreamKey(prev => prev + 1)
            if (textareaRef.current) {
                textareaRef.current.style.height = "auto"
            }
        } catch (err) {
            console.error(err)
            alert('Error submitting an entry')
        }
    }

    return <ProjectLayout
        outputStream={<OutputStream key={project.id} projectID={project.id} streamKey={streamKey} />}
    >
        <EntryFeed
            entries={entries}
            newEntryBody={newEntryBody}
            setNewEntryBody={setNewEntryBody}
            onSubmit={handleSubmit}
        />
    </ProjectLayout>
}
