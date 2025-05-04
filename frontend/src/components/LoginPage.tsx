'use client'

import { useForm } from 'react-hook-form'

import {
    Form,
    FormField,
    FormItem,
    FormLabel,
    FormControl,
} from './ui/form'
import { Input } from './ui/input'
import { Button } from './ui/button'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { Link, useNavigate } from 'react-router-dom'


const loginSchema = z.object({
    email: z.string().email(),
    password: z.string().min(1),
})

export function LoginPage() {
    const navigate = useNavigate()
    const form = useForm<z.infer<typeof loginSchema>>({
        resolver: zodResolver(loginSchema),
        defaultValues: {
            email: '',
            password: '',
        },
    })

    const onSubmit = async (values: z.infer<typeof loginSchema>) => {
        try {
            const res = await fetch('/api/login', {
                method: 'POST',
                body: JSON.stringify(values),
            })
            const data = await res.json()
            if (res.ok) {
                localStorage.setItem('token', data.token)
                navigate('/projects')
            } else {
                alert(data.message)
            }
            console.log(data)
        } catch (err) {
            console.error(err)
            alert('Error logging in')
        }
    }

    return (
        <div className="flex flex-col items-center justify-center h-screen">
            <h1 className="text-2xl font-bold">Login</h1>
            <Form {...form}>
                <form
                    onSubmit={form.handleSubmit(onSubmit)}
                    className="space-y-4 w-80"
                >
                    <FormField
                        control={form.control}
                        name="email"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Email</FormLabel>
                                <FormControl>
                                    <Input
                                        id="email"
                                        type="email"
                                        placeholder="email@example.com"
                                        required
                                        {...field}
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="password"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Password</FormLabel>
                                <FormControl>
                                    <Input
                                        id="password"
                                        type="password"
                                        placeholder="Password"
                                        required
                                        {...field}
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <Button type="submit" className="w-full">Login</Button>
                    <div className="mt-4 text-center text-sm">
                        Don&apos;t have an account?{" "}
                        <Link to="/signup" className="underline underline-offset-4">
                            Sign up
                        </Link>
                    </div>
                </form>
            </Form>
        </div>
    )
}