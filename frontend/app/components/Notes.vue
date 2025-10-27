<script setup lang="ts">
import type { TabsItem } from '@nuxt/ui'
import { ref } from 'vue'

const items = ref<TabsItem[]>([
{
  label: 'Account',
  description: 'Make changes to your account here. Click save when you\'re done.',
  icon: 'i-lucide-user',
  slot: 'account' as const
},
{
  label: 'Password',
  description: 'Change your password here. After saving, you\'ll be logged out.',
  icon: 'i-lucide-lock',
  slot: 'password' as const
}
] satisfies TabsItem[])
</script>

<template>
    <div 
        class="
            flex-1 h-full
            bg-gray-800
            rounded-2xl p-4
        ">
        <UTabs :items="items" variant="link" class="gap-4 w-full">
            <template #account="{ item }">
                <div class="">
                    ACCOUNT SETTINGS {{ item.description }}
                </div>
            </template>

            <template #password="{ item }">
                <div class="">
                    PASSWORD SETTINGS {{ item.description }}
                </div>
            </template>
        </UTabs>
    </div>
</template>