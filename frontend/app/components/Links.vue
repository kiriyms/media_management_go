<script setup lang="ts"></script>

<template>
    <div 
        class="
            w-100 h-full shrink-0 
            bg-amber-400 
            rounded-2xl p-4
        ">
        <div 
            class="
                text-xl font-semibold
                border-b-2 border-gray-950
            ">
            LINKS
        </div>
        <ul 
            class="
                pt-2
            ">
            <li class="w-full bg-green-500 p-2 rounded-md flex gap-2">
                <div class="w-10 h-10 bg-red-800"></div>
                <a href="#" class="text-blue-500 hover:underline">Link 1</a>
            </li>
            <li>
                <a href="#" class="text-blue-500 hover:underline">Link 2</a>
            </li>
            <li>
                <a href="#" class="text-blue-500 hover:underline">Link 3</a>
            </li>
        </ul>
    </div>
</template>