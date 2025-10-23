<script lang="ts" setup>
import { ref } from 'vue';

const passkey = ref('');
const handleSubmit = async (e: Event) => {
    e.preventDefault();

    try {
        const res = await fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ key: passkey.value }),
        });

        const data = await res.json();

        if (res.ok) {
            if (data && data.token) {
                localStorage.setItem('session_token', data.token);
                // you may want to redirect or update UI here
            } else {
                console.error('No token in response', data);
            }
        } else {
            // expected error shape: { error: 'message' }
            console.error('Login error:', data?.error ?? data);
        }
    } catch (err) {
        console.error('Network or unexpected error:', err);
    }
};
</script>

<template>
    <div
        class="
            flex items-center justify-center
            min-h-screen
            bg-gray-50
        "
    >
        <form
            @submit="handleSubmit"
            class="
                bg-white
                p-8
                rounded
                shadow-md
                w-full max-w-xs
                flex flex-col
                gap-4
            "
        >
            <label for="passkey" class="text-gray-700 font-semibold">Passkey</label>
            <input
                id="passkey"
                v-model="passkey"
                type="password"
                class="
                    border border-gray-300
                    rounded
                    px-3 py-2
                    focus:outline-none focus:ring-2 focus:ring-blue-400
                "
                placeholder="Enter passkey"
                required
            />
            <button 
                type="submit" 
                class="
                    bg-blue-600 
                    text-white 
                    rounded px-4 py-2 
                    hover:bg-blue-700 
                    transition
                "
            >
                Submit
            </button>
        </form>
    </div>
</template>