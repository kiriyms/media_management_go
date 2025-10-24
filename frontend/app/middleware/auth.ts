export default defineNuxtRouteMiddleware(async (to, from) => {
    if (import.meta.server) return

    // const token = useCookie('session_token').value
    const token = localStorage.getItem('session_token')
    
    if (!token) {
      console.error('No session_token found')
      return navigateTo('/')
    }

    const res = await fetch('http://localhost:8080/login', {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Accept': 'application/json',
      },
      // don't automatically follow redirects -- let us handle non-200s
      redirect: 'manual',
    })

    if (res.status === 200) {
        return
    }

    return navigateTo('/')
})