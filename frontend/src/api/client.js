// =========================================================
// Pure Static Client with LocalStorage Persistence
// 100% Standalone - No Backend / Database Required
// =========================================================

const STORAGE_KEY = 'me_too_static_store_v1'

const initialData = {
  user: {
    id: 1,
    username: 'lavender_nation',
    nickname: '薰衣草國度',
    bio: '把靈感變成行動，把行動留下成果',
    weekly_info: '★ 週進度：4 個任務，已達成連續 3 天記錄'
  },
  dailyTasks: [
    { id: 1, title: '完成專題使用流程圖', is_completed: false },
    { id: 2, title: '拍一張今天的天空', is_completed: false },
    { id: 3, title: '整理書桌 10 分鐘', is_completed: false }
  ],
  posts: [
    {
      id: 1,
      author_name: 'bbb',
      post_type: 'task',
      title: '拍一張今天的天空',
      content: '用一張照片記錄今天的狀態，完成後分享一句話。',
      expires_info: '任務貼文 · 剩 18 小時',
      result_count: 3,
      likes_count: 720,
      is_joined: true,
      submissions: [
        { id: 1, author_name: 'kb_q_p', time_ago: '4小時', content: '應該長這樣', image_url: 'placeholder-sky-1' },
        { id: 2, author_name: '同學A', time_ago: '2小時', content: '今天的天空是藍灰色，很適合當背景。' },
        { id: 3, author_name: '同學B', time_ago: '2小時', content: '我看倒像綠豆糕。' }
      ]
    },
    {
      id: 2,
      author_name: 'dafuu0000',
      post_type: 'inspiration',
      title: '把通勤路上的顏色做成色票',
      content: '看到有趣的配色先收藏，之後可以轉成自己的創作。',
      expires_info: '靈感貼文 · 收藏保留 30 天',
      image_url: 'placeholder-palette',
      likes_count: 86,
      saves_count: 12,
      is_joined: false,
      submissions: []
    },
    {
      id: 3,
      author_name: 'bbb',
      post_type: 'task',
      title: '拍一張今天的天空',
      content: '用一張照片記錄今天的狀態，完成後分享一句話。',
      expires_info: '任務貼文 · 剩 18 小時',
      result_count: 3,
      likes_count: 720,
      is_joined: false,
      submissions: []
    },
    {
      id: 4,
      author_name: 'dafuu0000',
      post_type: 'inspiration',
      title: '把通勤路上的顏色做成色票',
      content: '看到有趣的配色先收藏，之後可以轉成自己的創作。',
      expires_info: '靈感貼文 · 收藏保留 30 天',
      likes_count: 86,
      saves_count: 12,
      is_joined: false,
      submissions: []
    }
  ],
  myTasks: {
    private: [
      {
        id: 101,
        title: '完成企劃草稿',
        description: '私人任務，永久留在自己的清單',
        is_completed: false,
        progress_note: ''
      }
    ],
    public: [
      {
        id: 201,
        post_id: 1,
        title: '公共任務：拍一張今天的天空',
        description: '來自任務貼文，完成後需分享成果',
        is_completed: false
      },
      {
        id: 202,
        post_id: 3,
        title: '公共任務：整理書桌 10 分鐘',
        description: '來自任務貼文，完成後需分享成果',
        is_completed: false
      }
    ],
    completed: []
  },
  inspirations: [
    {
      id: 1,
      title: '去租書店把柯南漫畫裡的壞人都圈出來',
      date_group: '2026/06/24',
      is_pinned: false
    },
    {
      id: 2,
      title: '蛇的尿道跟肛門是同一個，尿液結晶有時會堵塞',
      preview_content: '貼文連結 / 照片預覽',
      date_group: '2026/06/24',
      is_pinned: false
    }
  ],
  profile: {
    user: {
      username: 'lavender_nation',
      bio: '把靈感變成行動，把行動留下成果',
      weekly_info: '★ 週進度：4 個任務，已達成連續 3 天記錄'
    },
    stats: {
      outcomes: 12,
      tasks: 28,
      inspirations: 46,
      launches: 4
    }
  }
}

function getStore() {
  try {
    const raw = localStorage.getItem(STORAGE_KEY)
    if (!raw) {
      localStorage.setItem(STORAGE_KEY, JSON.stringify(initialData))
      return JSON.parse(JSON.stringify(initialData))
    }
    return JSON.parse(raw)
  } catch (e) {
    return JSON.parse(JSON.stringify(initialData))
  }
}

function saveStore(data) {
  try {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(data))
  } catch (e) {
    console.error('LocalStorage write failed:', e)
  }
}

// ---------------------------------------------------------
// Authentication
// ---------------------------------------------------------
export async function login(username, password) {
  const store = getStore()
  return {
    token: 'static-jwt-token',
    user: store.user
  }
}

export async function register(username, password, nickname) {
  const store = getStore()
  store.user.username = username
  store.user.nickname = nickname || username
  saveStore(store)
  return { token: 'static-jwt-token', user: store.user }
}

export async function getMe() {
  const store = getStore()
  return store.user
}

// ---------------------------------------------------------
// Daily Tasks
// ---------------------------------------------------------
export async function fetchDailyTasks() {
  const store = getStore()
  return store.dailyTasks || []
}

export async function createDailyTask(title) {
  const store = getStore()
  const newTask = {
    id: Date.now(),
    title,
    is_completed: false
  }
  if (!store.dailyTasks) store.dailyTasks = []
  store.dailyTasks.push(newTask)
  saveStore(store)
  return newTask
}

export async function deleteDailyTask(id) {
  const store = getStore()
  store.dailyTasks = (store.dailyTasks || []).filter(t => t.id !== id)
  saveStore(store)
}

export async function toggleDailyTask(id) {
  const store = getStore()
  const task = (store.dailyTasks || []).find(t => t.id === id)
  if (task) {
    task.is_completed = !task.is_completed
    saveStore(store)
    return task
  }
  return null
}

// ---------------------------------------------------------
// Posts & Submissions
// ---------------------------------------------------------
export async function fetchPosts(type = 'all') {
  const store = getStore()
  let list = store.posts || []
  if (type === 'task') list = list.filter(p => p.post_type === 'task')
  if (type === 'inspiration') list = list.filter(p => p.post_type === 'inspiration')
  return list
}

export async function createPost(payload) {
  const store = getStore()
  const isTask = payload.post_type === 'task'
  const newPost = {
    id: Date.now(),
    ...payload,
    author_name: store.user?.username || 'lavender_nation',
    expires_info: isTask ? '任務貼文 · 剩 24 小時' : '靈感貼文 · 收藏保留 30 天',
    likes_count: 0,
    result_count: 0,
    is_joined: isTask,
    submissions: []
  }
  if (!store.posts) store.posts = []
  store.posts.unshift(newPost)
  store.profile.stats.launches++
  saveStore(store)
  return newPost
}

export async function joinPost(id) {
  const store = getStore()
  const post = (store.posts || []).find(p => p.id === id)
  if (post) {
    post.is_joined = true
    saveStore(store)
  }
  return { success: true }
}

export async function fetchPostThread(id) {
  const store = getStore()
  const post = (store.posts || []).find(p => p.id === id)
  return post ? (post.submissions || []) : []
}

export async function addSubmission(postId, payload) {
  const store = getStore()
  const post = (store.posts || []).find(p => p.id === postId)
  const newSub = {
    id: Date.now(),
    author_name: store.user?.username || 'lavender_nation',
    time_ago: '剛剛',
    ...payload
  }
  if (post) {
    if (!post.submissions) post.submissions = []
    post.submissions.push(newSub)
    post.result_count = (post.result_count || 0) + 1
    saveStore(store)
  }
  return newSub
}

// ---------------------------------------------------------
// My Tasks (Private, Public, Completed)
// ---------------------------------------------------------
export async function fetchMyTasks() {
  const store = getStore()
  return store.myTasks || { private: [], public: [], completed: [] }
}

export async function createMyTask(payload) {
  const store = getStore()
  const newTask = {
    id: Date.now(),
    ...payload,
    is_completed: false
  }
  if (!store.myTasks) store.myTasks = { private: [], public: [], completed: [] }
  if (payload.task_type === 'public') {
    store.myTasks.public.unshift(newTask)
  } else {
    store.myTasks.private.unshift(newTask)
  }
  saveStore(store)
  return newTask
}

export async function completeMyTask(id) {
  const store = getStore()
  if (!store.myTasks) return { success: true }

  let found = null
  let pIdx = store.myTasks.private.findIndex(t => t.id === id)
  if (pIdx !== -1) {
    found = store.myTasks.private.splice(pIdx, 1)[0]
  } else {
    let pubIdx = store.myTasks.public.findIndex(t => t.id === id)
    if (pubIdx !== -1) {
      found = store.myTasks.public.splice(pubIdx, 1)[0]
    }
  }

  if (found) {
    found.is_completed = true
    if (!store.myTasks.completed) store.myTasks.completed = []
    store.myTasks.completed.unshift(found)
    store.profile.stats.outcomes++
    saveStore(store)
  }
  return { success: true }
}

export async function updateTaskProgress(id, note) {
  const store = getStore()
  const task = (store.myTasks?.private || []).find(t => t.id === id)
  if (task) {
    task.progress_note = note
    saveStore(store)
  }
  return { success: true }
}

// ---------------------------------------------------------
// Inspirations
// ---------------------------------------------------------
export async function fetchInspirations(search = '') {
  const store = getStore()
  let list = store.inspirations || []
  if (search) {
    const q = search.toLowerCase()
    list = list.filter(i => (i.title || '').toLowerCase().includes(q))
  }
  return list
}

export async function createInspiration(payload) {
  const store = getStore()
  const newItem = {
    id: Date.now(),
    date_group: '2026/06/24',
    is_pinned: false,
    ...payload
  }
  if (!store.inspirations) store.inspirations = []
  store.inspirations.unshift(newItem)
  store.profile.stats.inspirations++
  saveStore(store)
  return newItem
}

export async function deleteInspiration(id) {
  const store = getStore()
  store.inspirations = (store.inspirations || []).filter(i => i.id !== id)
  saveStore(store)
}

// ---------------------------------------------------------
// Profile
// ---------------------------------------------------------
export async function fetchProfile() {
  const store = getStore()
  return store.profile
}

export async function updateProfile(data) {
  const store = getStore()
  if (data.bio !== undefined) store.profile.user.bio = data.bio
  if (data.weekly_info !== undefined) store.profile.user.weekly_info = data.weekly_info
  if (store.user) {
    if (data.bio !== undefined) store.user.bio = data.bio
    if (data.weekly_info !== undefined) store.user.weekly_info = data.weekly_info
  }
  saveStore(store)
  return store.profile
}
