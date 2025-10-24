<script lang="ts" setup>
import { useRouter } from '#app';

const router = useRouter();

definePageMeta({
    middleware: async (to) => {
        try {
            const token = localStorage.getItem('session_token');
            if (!token) {
                return navigateTo('/');
            }

            const res = await fetch('http://localhost:8080/login', {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Accept': 'application/json'
                },
                // Don't automatically follow redirects
                redirect: 'manual'
            });

            // Check specific status codes
            if (res.status === 200) {
                // Valid session, allow access to dashboard
                return;
            } else if (res.status === 401 || res.status === 403) {
                // Unauthorized or forbidden
                console.error('Authentication failed:', res.status);
                localStorage.removeItem('session_token');
                return navigateTo('/', { replace: true });
            } else if (res.status === 302) {
                // Handle redirect manually - likely due to missing or invalid auth
                console.error('Redirect detected - invalid session');
                localStorage.removeItem('session_token');
                return navigateTo('/', { replace: true });
            }

            // Any other error
            console.error('Unexpected response:', res.status);
            return navigateTo('/', { replace: true });
        } catch (err) {
            console.error('Session validation failed:', err);
            localStorage.removeItem('session_token');
            return navigateTo('/', { replace: true });
        }
    }
});
</script>

<template>
    <div class="p-8">
        <h1 class="text-2xl font-bold mb-4">Dashboard</h1>
    </div>
</template>