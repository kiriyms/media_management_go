<script setup lang="ts">
import type { TabsItem } from '@nuxt/ui'
import { ref, onMounted } from 'vue'

interface Note {
    id: string
    title: string
    note: string
    createdAt: string
    updatedAt: string
}

interface TabItem {
    label: string
    description: string
    slot: string
}

const pending = ref(false)
const showModal = ref(false)
const newNoteTitle = ref('')

const items = ref<TabItem[] | undefined>([])

// Using a computed ref to store the auth token
const authToken = computed(() => {
    if (typeof window !== 'undefined') {
        return localStorage.getItem('session_token')
    }
    return null
})

const { data: notesData, refresh } = useFetch('http://localhost:8080/note', {
    headers: computed(() => ({
        'Authorization': `Bearer ${authToken.value}`
    })),
    key: 'notes-data',
    transform: (response: any) => response?.notes as Note[],
    cache: 'no-store'
})

// Wrapper function to handle the loading state
const fetchNotes = async () => {
    if (!authToken.value) {
        console.error('No auth token found')
        return
    }

    pending.value = true
    try {
        await refresh()
        console.log('Fetched notes:', notesData.value)
        // Populate items based on fetched notes
        items.value = notesData.value?.map((note) => ({
            label: note.title,
            description: note.note,
            slot: note.id
        }))
        console.log('Populated items:', items.value)
    } catch (error) {
        console.error('Failed to fetch notes:', error)
    } finally {
        pending.value = false
    }
}

const submitNewNote = async () => {
    if (!authToken.value || !newNoteTitle.value) {
        return
    }

    pending.value = true
    try {
        const response = await $fetch('http://localhost:8080/note', {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${authToken.value}`,
                'Content-Type': 'application/json'
            },
            body: {
                title: newNoteTitle.value,
                note: ' ',
            }
        })

        // Clear form
        newNoteTitle.value = ''
        showModal.value = false

        // Refresh notes list
        await fetchNotes()
    } catch (error) {
        console.error('Failed to create note:', error)
    } finally {
        pending.value = false
    }
}

onMounted(async () => {
    console.log('Component mounted, fetching notes...')
    fetchNotes()
})
</script>

<template>
    <div 
        class="
            flex-1 h-full
            bg-gray-800
            rounded-2xl p-4
            flex
        ">
        <UTabs :items="items" variant="link" class="gap-4 w-full">
            <template #note="{ item }">
                <div class="">
                    <div class="font-bold text-lg mb-2">{{ item.label }}</div>
                    <div>{{ item.description }}</div>
                </div>
            </template>
        </UTabs>
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
                class="hover:text-gray-400 transition-colors h-8"
                color="neutral"
                variant="outline"
            />
            <template #body>
                <form class="flex flex-col gap-4 p-4">
                    <UInput
                        v-model="newNoteTitle"
                        placeholder="Note Title"
                        type="text"
                        required
                    />
                </form>
            </template>

            <template #footer>
                <div class="flex justify-end gap-2">
                    <UButton
                        @click.prevent="submitNewNote"
                        type="submit"
                        color="primary"
                        :loading="pending"
                        :disabled="!newNoteTitle"
                    >
                        Add Note
                    </UButton>
                </div>
            </template>
        </UModal>
    </div>
</template>