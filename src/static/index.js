const { createApp, ref, onMounted } = Vue

createApp({
    setup() {
        const ideas = ref([])
        const error = ref(null)
        const showModal = ref(false)
        const newIdea = ref({
            id: null,
            timestamp: '',
            title: '',
            description: '',
            writer: '',
            tags: []
        })

        onMounted(async () => {
            try {
                const response = await axios.get('/ideas')
                if (response.data) {
                    ideas.value = response.data
                }
            } catch (err) {
                error.value = err
            }
        })

        const submitIdea = async () => {
            try {
                // Generate a unique ID and timestamp
                newIdea.value.id = ideas.value.length + 1
                newIdea.value.timestamp = new Date().toISOString()
                newIdea.value.tags = newIdea.value.tags.split(',').map(tag => tag.trim())

                const response = await axios.post('/idea', newIdea.value)
                if (response.status === 201) {
                    ideas.value.push(newIdea.value)
                    showModal.value = false
                    newIdea.value = {
                        id: null,
                        timestamp: '',
                        title: '',
                        description: '',
                        writer: '',
                        tags: []
                    }
                }
            } catch (err) {
                console.error('Failed to add idea:', err)
            }
        }

        return {
            ideas,
            error,
            showModal,
            newIdea,
            submitIdea
        }
    }
}).mount('#app')