<script setup lang="ts"></script>

<template>
    <div class="flex flex-col min-h-screen">
        <nav>
            <h1 class="text-2xl font-bold p-4">My Nuxt 3 App</h1>
        </nav>
        <main class="flex-1 flex flex-col">
            <slot />
        </main>
        <nav>
            <ul>
                <li><NuxtLink to="/">Home</NuxtLink></li>
                <li><NuxtLink to="/about">About</NuxtLink></li>
                <li><NuxtLink to="/contact">Contact</NuxtLink></li>
            </ul>
        </nav>
    </div>
</template>