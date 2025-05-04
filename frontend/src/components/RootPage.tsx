import { useNavigate } from "react-router-dom"
import { useEffect } from "react"

export function RootPage() {
    const navigate = useNavigate()

    useEffect(() => {
        if (localStorage.getItem('token')) {
            navigate('/projects')
        } else {
            navigate('/login')
        }
    }, [navigate])

    return null
}