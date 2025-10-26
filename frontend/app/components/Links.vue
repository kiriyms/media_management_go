<script setup lang="ts">
import { ref } from 'vue'

interface Link {
    ID: string;
    Link: string;
    ImgPath: string;
    CreatedAt: string;
    UpdatedAt: string;
}

const pending = ref(false)
const showModal = ref(false)
const newLinkUrl = ref('')
const selectedIcon = ref('')

const iconOptions = [
    { value: 'youtube.png', label: 'YouTube', icon: 'i-logos-youtube-icon' },
    { value: 'github.png', label: 'GitHub', icon: 'i-logos-github-icon' },
    { value: 'twitter.png', label: 'Twitter', icon: 'i-logos-twitter' },
    { value: 'linkedin.png', label: 'LinkedIn', icon: 'i-logos-linkedin-icon' },
    { value: 'instagram.png', label: 'Instagram', icon: 'i-logos-instagram-icon' }
]

// Using a computed ref to store the auth token
const authToken = computed(() => {
    if (typeof window !== 'undefined') {
        return localStorage.getItem('session_token')
    }
    return null
})

// Use Nuxt's built-in state management with useFetch
const { data: linksData, refresh } = useFetch('http://localhost:8080/link', {
    headers: computed(() => ({
        'Authorization': `Bearer ${authToken.value}`
    })),
    key: 'links-data',
    transform: (response: any) => response?.links as Link[],
    cache: 'no-store'
})

const links = computed(() => linksData.value || [])

// Wrapper function to handle the loading state
const fetchLinks = async () => {
    if (!authToken.value) {
        console.error('No auth token found')
        return
    }

    pending.value = true
    try {
        await refresh()
        console.log('Fetched links:', links.value)
    } catch (error) {
        console.error('Failed to fetch links:', error)
    } finally {
        pending.value = false
    }
}

const submitNewLink = async () => {
    if (!authToken.value || !newLinkUrl.value || !selectedIcon.value) {
        return
    }

    pending.value = true
    try {
        const response = await $fetch('http://localhost:8080/link', {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${authToken.value}`,
                'Content-Type': 'application/json'
            },
            body: {
                link: newLinkUrl.value,
                img_path: selectedIcon.value
            }
        })

        // Clear form
        newLinkUrl.value = ''
        selectedIcon.value = ''
        showModal.value = false

        // Refresh links list
        await fetchLinks()
    } catch (error) {
        console.error('Failed to create link:', error)
    } finally {
        pending.value = false
    }
}

// Fetch links when component mounts
onMounted(() => {
    console.log('Component mounted, fetching links...')
    fetchLinks()
})
</script>

<template>
    <div 
        class="
            w-100 h-full shrink-0 
            bg-gray-800
            rounded-2xl p-4
        ">
        <div 
            class="
                text-xl font-semibold
                border-b-2 border-gray-950/50
                flex items-center gap-2
                text-white
            ">
            LINKS
            <div v-if="pending" class="w-4 h-4 border-2 border-gray-950/80 border-t-transparent rounded-full animate-spin"></div>
            <UModal
                v-model="showModal"
                :close="{
                    color: 'neutral',
                    variant: 'outline',
                }"
                title="Add New Link"
            >
                <UButton 
                    icon="material-symbols:add-2-rounded"
                    class="hover:text-gray-400 transition-colors"
                    color="neutral"
                    variant="outline"
                />
                <template #body>
                    <form class="flex flex-col gap-4 p-4">
                        <UFormGroup label="Link URL">
                            <UInput
                                v-model="newLinkUrl"
                                placeholder="https://example.com"
                                type="url"
                                required
                            />
                        </UFormGroup>

                        <UFormGroup label="Choose Icon">
                            <div class="grid grid-cols-5 gap-2">
                                <UButton
                                    v-for="option in iconOptions"
                                    :key="option.value"
                                    :class="{
                                        'ring-2 ring-primary-500': selectedIcon === option.icon
                                    }"
                                    @click="selectedIcon = option.icon"
                                    variant="ghost"
                                    :title="option.label"
                                >
                                    <UIcon
                                        :name="option.icon"
                                        class="text-2xl"
                                    />
                                </UButton>
                            </div>
                        </UFormGroup>
                    </form>
                </template>

                <template #footer>
                    <div class="flex justify-end gap-2">
                        <UButton
                            @click.prevent="submitNewLink"
                            type="submit"
                            color="primary"
                            :loading="pending"
                            :disabled="!newLinkUrl || !selectedIcon"
                        >
                            Add Link
                        </UButton>
                    </div>
                </template>
            </UModal>
            <button 
                @click="fetchLinks"
                class="ml-auto hover:text-gray-400 transition-colors"
                title="Refresh links"
            >
                <svg 
                    xmlns="http://www.w3.org/2000/svg" 
                    class="h-5 w-5"
                    :class="{ 'animate-spin': pending }"
                    fill="none" 
                    viewBox="0 0 24 24" 
                    stroke="currentColor"
                >
                    <path 
                        stroke-linecap="round" 
                        stroke-linejoin="round" 
                        stroke-width="2" 
                        d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" 
                    />
                </svg>
            </button>
        </div>
        <ul 
            class="
                pt-2 flex flex-col gap-2 overflow-auto h-[90%]
            ">
            <li 
                v-for="link in links" 
                :key="link.ID"
                class="w-full bg-gray-700 p-2 rounded-md flex items-center gap-2"
            >
                <UIcon 
                    v-if="link.ImgPath" 
                    :name="link.ImgPath"
                    class="size-10"
                />
                <div v-else class="w-10 h-10 bg-red-800 rounded-md" />
                <a 
                    :href="link.Link" 
                    target="_blank"
                    rel="noopener noreferrer"
                    class="flex-1 text-blue-500 hover:underline overflow-auto"
                >
                    {{ link.Link }}
                </a>
                <div class="text-xs text-white light:text-black">
                    {{ new Date(link.CreatedAt).toLocaleDateString() }}
                </div>
            </li>
        </ul>
    </div>
</template>