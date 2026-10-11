<template>
  <div class="web-app-root">
    <!-- CASE A: Dedicated Auth Page (登入 / 建立帳號 / 忘記密碼 - 獨立全頁面，無側邊欄干擾，桌面與手機自適應) -->
    <div v-if="currentTab === 'auth'" class="auth-page-wrapper">
      <div class="auth-card-shell">
        <LoginView @login-success="handleLoginSuccess" />
      </div>
    </div>

    <!-- CASE B: Main Application with Full Desktop & Mobile RWD -->
    <div v-else class="app-layout-container">
      <!-- 1. 電腦版左側導航列（手機與窄螢幕 < 768px 自動隱藏） -->
      <Sidebar 
        :current-tab="currentTab"
        :current-user="currentUser"
        @change-tab="setTab"
        @open-create-post="setTab('create')"
        @open-daily-task="setTab('onboarding')"
        @logout="handleLogout"
      />

      <!-- 2. 中間主內容區（電腦版居中舒適閱讀寬度 640px，手機版 100% 滿版 RWD） -->
      <main class="main-flow-content">
        <!-- 01 & 08 & 10. 今日任務引導 -->
        <DailyTaskView 
          v-if="currentTab === 'onboarding'"
          :tasks="dailyTasks"
          @add-task="handleAddDailyTask"
          @delete-task="handleDeleteDailyTask"
          @toggle-task="handleToggleDailyTask"
          @finish-setup="finishDailyTaskSetup"
        />

        <!-- 02 & 03 & 16. 貼文區 -->
        <FeedView 
          v-else-if="currentTab === 'feed'"
          :posts="posts"
          @join-post="handleJoinPost"
          @submit-post-result="handleSubmitResult"
          @execute-task="startExecuteTaskFromPost"
          @view-user-profile="openUserProfile"
        />

        <!-- 04. 我也要發文 -->
        <CreatePostView 
          v-else-if="currentTab === 'create'"
          @publish="handlePublishPost"
        />

        <!-- 05. 我的任務 -->
        <TasksView 
          v-else-if="currentTab === 'tasks'"
          :private-tasks="myTasks.private"
          :public-tasks="myTasks.public"
          :completed-tasks="myTasks.completed"
          @execute-task="handleStartExecuteTask"
          @create-private-task="handleCreatePrivateTask"
          @complete-task="handleCompleteTask"
        />

        <!-- 09. 任務執行 / 過程記錄 -->
        <TaskExecuteView 
          v-else-if="currentTab === 'task-execute'"
          :task="activeExecutingTask"
          @back="handleBackFromExecute"
          @save-progress="handleSaveTaskProgress"
          @complete-task="handleFinishTaskExecution"
        />

        <!-- 06. 靈感清單 -->
        <InspirationsView 
          v-else-if="currentTab === 'inspirations'"
          :inspirations="filteredInspirations"
          @add-inspiration="handleAddInspiration"
          @delete-inspiration="handleDeleteInspiration"
          @search="handleSearchInspirations"
        />

        <!-- 07 & 11. 個人主頁 -->
        <ProfileView 
          v-else-if="currentTab === 'profile'"
          :profile="profileData"
          :initial-is-other-user="viewingOtherUser"
          @update-profile="handleUpdateProfile"
          @navigate-auth="handleLogout"
          @navigate-onboarding="setTab('onboarding')"
        />
      </main>

      <!-- 3. 電腦版右側面板（大螢幕顯示，小於 1120px 與手機自動隱藏） -->
      <RightPanel 
        :daily-tasks="dailyTasks"
        :posts="posts"
        @add-daily-task="handleAddDailyTask"
        @toggle-daily-task="handleToggleDailyTask"
        @open-daily-task="setTab('onboarding')"
        @change-tab="setTab"
        @add-inspiration="handleAddInspirationText"
      />

      <!-- 4. 手機版底部導航列（小於 768px 時固定浮現，電腦版自動隱藏） -->
      <BottomNav 
        v-if="currentTab !== 'onboarding' && currentTab !== 'task-execute'"
        :current-tab="currentTab"
        @change-tab="setTab"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import Sidebar from './components/Sidebar.vue'
import RightPanel from './components/RightPanel.vue'
import BottomNav from './components/BottomNav.vue'
import DailyTaskView from './components/DailyTaskView.vue'
import FeedView from './components/FeedView.vue'
import CreatePostView from './components/CreatePostView.vue'
import TasksView from './components/TasksView.vue'
import TaskExecuteView from './components/TaskExecuteView.vue'
import InspirationsView from './components/InspirationsView.vue'
import ProfileView from './components/ProfileView.vue'
import LoginView from './components/LoginView.vue'
import * as api from './api/client.js'

// Active Tab ('feed' | 'onboarding' | 'create' | 'tasks' | 'task-execute' | 'inspirations' | 'profile' | 'auth')
const currentTab = ref('feed')
const isLoggedIn = ref(true)
const viewingOtherUser = ref(false)

// Active task being executed
const activeExecutingTask = ref({
  id: 201,
  title: '拍一張今天的天空',
  description: '用照片或文字逐步記錄執行過程，完成後發布成果。'
})

// Current User State
const currentUser = ref({
  id: 1,
  username: 'mine_account',
  nickname: '我的帳號',
  bio: '把靈感變成行動，把行動留下成果',
  weekly_info: '本週完成 5 個任務，收藏 8 則靈感'
})

// Data states
const dailyTasks = ref([
  { id: 1, title: '完成專題使用流程圖', is_completed: false },
  { id: 2, title: '拍一張今天的天空', is_completed: false },
  { id: 3, title: '整理書桌 10 分鐘', is_completed: false }
])

const posts = ref([
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
    is_completed: false,
    submissions: [
      { id: 1, author_name: 'kb_q_p', time_ago: '4小時', content: '應該長這樣', image_url: 'sky-1' },
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
    comments_count: 4,
    is_joined: false,
    is_completed: false,
    submissions: []
  },
  {
    id: 3,
    author_name: 'lavender_nation',
    post_type: 'task',
    title: '拍一張今天的天空',
    content: '用一張照片記錄今天的狀態，完成後分享一句話。',
    expires_info: '任務貼文 · 剩 18 小時',
    result_count: 3,
    likes_count: 720,
    is_joined: true,
    is_completed: false,
    submissions: []
  },
  {
    id: 4,
    author_name: 'aaa',
    post_type: 'task',
    title: '整理書桌10分鐘',
    content: '去整理啦',
    expires_info: '任務貼文 · 剩 18 小時',
    result_count: 3,
    likes_count: 720,
    is_joined: false,
    is_completed: false,
    submissions: []
  }
])

const myTasks = ref({
  private: [
    {
      id: 101,
      title: '完成企劃草稿',
      description: '私人任務，會保留紀錄',
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
      post_id: 4,
      title: '公共任務：整理書桌 10 分鐘',
      description: '來自任務貼文，完成後需分享成果',
      is_completed: false
    }
  ],
  completed: []
})

const inspirations = ref([
  {
    id: 1,
    title: '去租書店把柯南漫畫裡的壞人都圈出來。',
    date_group: '2026/06/24'
  },
  {
    id: 2,
    title: '蛇的排遺出口只有一個，尿液有可能結晶堵塞。',
    preview_content: '貼文連結 / 照片預覽',
    date_group: '2026/06/24'
  },
  {
    id: 3,
    title: '舉辦社區吃麻糬大賽解決人口老化問題。',
    date_group: '2026/06/24'
  }
])

const inspirationSearch = ref('')
const filteredInspirations = computed(() => {
  if (!inspirationSearch.value) return inspirations.value
  const q = inspirationSearch.value.toLowerCase()
  return inspirations.value.filter(i => i.title.toLowerCase().includes(q))
})

const profileData = ref({
  user: {
    username: 'mine_account',
    bio: '把靈感變成行動，把行動留下成果',
    weekly_info: '本週完成 5 個任務，收藏 8 則靈感'
  },
  stats: {
    outcomes: 12,
    followers: 28,
    albums: 4
  }
})

// Lifecycle
onMounted(async () => {
  loadDataFromBackend()
})

async function loadDataFromBackend() {
  const dt = await api.fetchDailyTasks()
  if (dt && dt.length) dailyTasks.value = dt

  const p = await api.fetchPosts()
  if (p && p.length) posts.value = p

  const mt = await api.fetchMyTasks()
  if (mt) myTasks.value = mt

  const ins = await api.fetchInspirations()
  if (ins && ins.length) inspirations.value = ins
}

function setTab(tab) {
  if (!isLoggedIn.value && tab !== 'auth') {
    currentTab.value = 'auth'
    return
  }
  currentTab.value = tab
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

function handleLogout() {
  isLoggedIn.value = false
  currentTab.value = 'auth'
}

function finishDailyTaskSetup() {
  setTab('feed')
}

// Daily Task actions
async function handleAddDailyTask(title) {
  const task = await api.createDailyTask(title)
  dailyTasks.value.push(task || { id: Date.now(), title, is_completed: false })
}

async function handleDeleteDailyTask(id) {
  dailyTasks.value = dailyTasks.value.filter(t => t.id !== id)
  await api.deleteDailyTask(id)
}

async function handleToggleDailyTask(id) {
  const task = dailyTasks.value.find(t => t.id === id)
  if (task) {
    task.is_completed = !task.is_completed
    await api.toggleDailyTask(id)
  }
}

// Feed Actions
async function handleJoinPost(post) {
  post.is_joined = true
  await api.joinPost(post.id)

  const exists = myTasks.value.public.some(t => t.post_id === post.id)
  if (!exists) {
    myTasks.value.public.unshift({
      id: Date.now(),
      post_id: post.id,
      title: '公共任務：' + post.title,
      description: '來自任務貼文，完成後需分享成果',
      is_completed: false
    })
  }
}

function startExecuteTaskFromPost(post) {
  activeExecutingTask.value = {
    id: post.id,
    title: post.title,
    description: post.content || '用照片或文字逐步記錄執行過程，完成後發布成果。'
  }
  setTab('task-execute')
}

function openUserProfile(username) {
  if (username === 'mine_account' || username === currentUser.value.username) {
    viewingOtherUser.value = false
  } else {
    viewingOtherUser.value = true
  }
  setTab('profile')
}

async function handleSubmitResult({ postId, content }) {
  const post = posts.value.find(p => p.id === postId)
  const authorName = currentUser.value?.username || 'mine_account'
  if (post) {
    if (!post.submissions) post.submissions = []
    post.submissions.push({
      id: Date.now(),
      author_name: authorName,
      time_ago: '剛剛',
      content
    })
    post.result_count = (post.result_count || 0) + 1
  }

  await api.addSubmission(postId, { content })
  profileData.value.stats.outcomes++
}

// Create Post Action
async function handlePublishPost(payload) {
  const authorName = currentUser.value?.username || 'mine_account'
  const isTask = payload.post_type === 'task'
  const newPost = await api.createPost(payload)
  const createdPost = newPost || {
    id: Date.now(),
    ...payload,
    author_name: authorName,
    expires_info: isTask ? '任務貼文 · 剩 24 小時' : '靈感貼文 · 收藏保留 30 天',
    likes_count: 0,
    result_count: 0,
    is_joined: isTask,
    is_completed: false,
    submissions: []
  }

  if (isTask) {
    createdPost.is_joined = true
    const exists = myTasks.value.public.some(t => t.post_id === createdPost.id)
    if (!exists) {
      myTasks.value.public.unshift({
        id: Date.now(),
        post_id: createdPost.id,
        title: '公共任務：' + createdPost.title,
        description: '來自任務貼文，完成後需分享成果',
        is_completed: false
      })
    }
  }

  posts.value.unshift(createdPost)
  setTab('feed')
}

// Tasks Actions
function handleStartExecuteTask(task) {
  activeExecutingTask.value = task
  setTab('task-execute')
}

async function handleCreatePrivateTask(title) {
  const res = await api.createMyTask({
    title,
    description: '私人任務，會保留紀錄',
    task_type: 'private'
  })
  const newTask = res || {
    id: Date.now(),
    title,
    description: '私人任務，會保留紀錄',
    is_completed: false
  }
  myTasks.value.private.unshift(newTask)
}

async function handleCompleteTask(id) {
  let taskIndex = myTasks.value.private.findIndex(t => t.id === id)
  if (taskIndex !== -1) {
    const task = myTasks.value.private.splice(taskIndex, 1)[0]
    task.is_completed = true
    myTasks.value.completed.unshift(task)
    profileData.value.stats.outcomes++
    await api.completeMyTask(id)
    return
  }

  taskIndex = myTasks.value.public.findIndex(t => t.id === id)
  if (taskIndex !== -1) {
    const task = myTasks.value.public.splice(taskIndex, 1)[0]
    task.is_completed = true
    myTasks.value.completed.unshift(task)
    profileData.value.stats.outcomes++
    await api.completeMyTask(id)

    if (task.post_id) {
      const post = posts.value.find(p => p.id === task.post_id)
      if (post) {
        post.is_completed = true
      }
    }
  }
}

// Task Execute actions
function handleBackFromExecute() {
  setTab('tasks')
}

function handleSaveTaskProgress(payload) {
  // Save progress notes or steps
}

function handleFinishTaskExecution({ taskId, taskTitle, steps }) {
  // Mark task completed
  handleCompleteTask(taskId)

  // Find corresponding post and mark as completed
  const post = posts.value.find(p => p.id === taskId || p.title === taskTitle || p.title === activeExecutingTask.value?.title)
  if (post) {
    post.is_completed = true
    const authorName = currentUser.value?.username || 'mine_account'
    const lastStep = steps && steps.length > 0 ? steps[steps.length - 1] : null
    post.submissions.unshift({
      id: Date.now(),
      author_name: authorName,
      time_ago: '剛剛',
      content: lastStep?.text || '完成打卡！',
      image_url: lastStep?.imageUrl || 'uploaded-photo-preview'
    })
  }

  // Navigate to feed with completed state
  setTab('feed')
}

// Inspiration Actions
async function handleAddInspiration(payload) {
  const todayStr = '2026/06/24'
  const item = await api.createInspiration(payload)
  inspirations.value.unshift(item || { id: Date.now(), date_group: todayStr, ...payload })
  profileData.value.stats.outcomes++
}

function handleAddInspirationText(text) {
  handleAddInspiration({ title: text })
}

async function handleDeleteInspiration(id) {
  inspirations.value = inspirations.value.filter(i => i.id !== id)
  await api.deleteInspiration(id)
}

function handleSearchInspirations(q) {
  inspirationSearch.value = q
}

// Profile Actions
async function handleUpdateProfile(data) {
  if (profileData.value.user) {
    profileData.value.user.bio = data.bio
    profileData.value.user.weekly_info = data.weekly_info
  }
  await api.updateProfile(data)
}

function handleLoginSuccess(user) {
  isLoggedIn.value = true
  currentUser.value.username = user.username
  if (profileData.value.user) {
    profileData.value.user.username = user.username
  }
  currentTab.value = 'feed'
}
</script>

<style scoped>
.app-root-wrapper {
  width: 100%;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}
</style>
