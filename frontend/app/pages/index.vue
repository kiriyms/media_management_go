<script lang="ts" setup>
import { ref } from 'vue';

const passkey = ref('');
const handleSubmit = (e: Event) => {
    e.preventDefault();
    // Handle passkey submit logic here
    alert(`Passkey submitted: ${passkey.value}`);
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