import { redirect } from "react-router-dom"

export function requireAuthLoader() {
    const token = localStorage.getItem('token')
    if (!token) {
        throw redirect('/login')
    }
    return null
}