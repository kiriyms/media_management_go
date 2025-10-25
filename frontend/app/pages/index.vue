<script lang="ts" setup>
import { ref } from 'vue';
import { useRouter } from '#app';

const router = useRouter();
const passkey = ref('');
const isLoading = ref(false);
const errorMessage = ref('');

const handleSubmit = async (e: Event) => {
    e.preventDefault();
    isLoading.value = true;

    errorMessage.value = ''; // Clear any previous errors
    
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
                router.push('/dash');
            } else {
                errorMessage.value = 'Invalid server response: No token received';
            }
        } else {
            // expected error shape: { error: 'message' }
            errorMessage.value = data?.error || 'Authentication failed';
        }
    } catch (err) {
        errorMessage.value = 'Network error: Unable to connect to server';
    } finally {
        isLoading.value = false;
    }
};
</script>

<template>
    <div
        class="
            flex items-center justify-center
            min-h-screen
        "
    >
        <form
            @submit="handleSubmit"
            class="
                p-8
                rounded
                shadow-md
                w-full max-w-xs
                flex flex-col
                gap-4
            "
        >
            <label for="passkey" class="font-semibold">Passkey</label>
            <input
                id="passkey"
                v-model="passkey"
                type="password"
                class="
                    border
                    rounded
                    px-3 py-2
                    focus:outline-none focus:ring-2 focus:ring-blue-400
                "
                placeholder="Enter passkey"
                required
            />
            <p
                v-if="errorMessage"
                class="
                    text-red-600
                    text-sm
                    mt-1
                "
            >
                {{ errorMessage }}
            </p>
            <button 
                type="submit" 
                :disabled="isLoading"
                class="
                    bg-blue-600 
                    rounded px-4 py-2 
                    hover:bg-blue-700 
                    disabled:opacity-50
                    disabled:cursor-not-allowed
                    transition
                "
            >
                <span class="flex items-center justify-center gap-2">
                    Submit
                    <svg
                        v-if="isLoading"
                        class="animate-spin h-4 w-4"
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                    >
                        <circle
                            class="opacity-25"
                            cx="12"
                            cy="12"
                            r="10"
                            stroke="currentColor"
                            stroke-width="4"
                        ></circle>
                        <path
                            class="opacity-75"
                            fill="currentColor"
                            d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                        ></path>
                    </svg>
                </span>
            </button>
        </form>
    </div>
</template>