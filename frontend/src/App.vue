<template>
  <div class="app-wrapper">
    <div class="desktop-app-container">
      <!-- 1. Left Desktop Sidebar (電腦版左側導航列) -->
      <Sidebar 
        :current-tab="currentTab"
        :current-user="currentUser"
        @change-tab="setTab"
        @open-create-post="setTab('create')"
        @open-daily-task="setTab('onboarding')"
        @logout="handleLogout"
      />

      <!-- 2. Center Main Flow (電腦版中間主內容流) -->
      <main class="main-feed-flow">
        <!-- 01. 首次詢問 (今天是否有任務？) -->
        <DailyTaskView 
          v-if="currentTab === 'onboarding'"
          :tasks="dailyTasks"
          @add-task="handleAddDailyTask"
          @toggle-task="handleToggleDailyTask"
          @delete-task="handleDeleteDailyTask"
          @finish-setup="finishDailyTaskSetup"
        />

        <!-- 02. 貼文區 (貼文串 & 展開打卡成果) -->
        <FeedView 
          v-else-if="currentTab === 'feed'"
          :posts="posts"
          @join-post="handleJoinPost"
          @submit-post-result="handleSubmitResult"
          @change-tab="setTab"
        />

        <!-- 04. 我也要發文 (發文頁面) -->
        <CreatePostView 
          v-else-if="currentTab === 'create'"
          @publish="handlePublishPost"
        />

        <!-- 05. 我的任務 (私人任務 + 公共任務 + 完成紀錄) -->
        <TasksView 
          v-else-if="currentTab === 'tasks'"
          :private-tasks="myTasks.private"
          :public-tasks="myTasks.public"
          :completed-tasks="myTasks.completed"
          @create-private-task="handleCreatePrivateTask"
          @complete-task="handleCompleteTask"
          @update-progress="handleUpdateTaskProgress"
        />

        <!-- 06. 靈感清單 (蒐集靈感、速記、日期分組) -->
        <InspirationsView 
          v-else-if="currentTab === 'inspirations'"
          :inspirations="filteredInspirations"
          @add-inspiration="handleAddInspiration"
          @delete-inspiration="handleDeleteInspiration"
          @search="handleSearchInspirations"
        />

        <!-- 07. 個人 (九宮格作品、個人資訊與統計) -->
        <ProfileView 
          v-else-if="currentTab === 'profile'"
          :profile="profileData"
          @update-profile="handleUpdateProfile"
          @logout="handleLogout"
        />
      </main>

      <!-- 3. Right Desktop Panel (電腦版右側工具面板) -->
      <RightPanel 
        :daily-tasks="dailyTasks"
        :posts="posts"
        @add-daily-task="handleAddDailyTask"
        @toggle-daily-task="handleToggleDailyTask"
        @open-daily-task="setTab('onboarding')"
        @change-tab="setTab"
        @add-inspiration="handleAddInspirationText"
      />

      <!-- 4. Mobile Bottom Navigation (手機尺寸時自動浮現) -->
      <BottomNav 
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
import InspirationsView from './components/InspirationsView.vue'
import ProfileView from './components/ProfileView.vue'
import * as api from './api/client.js'

// Active Tab ('feed' | 'onboarding' | 'create' | 'tasks' | 'inspirations' | 'profile')
const currentTab = ref('feed')

// Current User State (default lavender_nation from Figma)
const currentUser = ref({
  id: 1,
  username: 'lavender_nation',
  nickname: '薰衣草國度',
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
])

const myTasks = ref({
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
})

const inspirations = ref([
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
])

const inspirationSearch = ref('')
const filteredInspirations = computed(() => {
  if (!inspirationSearch.value) return inspirations.value
  const q = inspirationSearch.value.toLowerCase()
  return inspirations.value.filter(i => i.title.toLowerCase().includes(q))
})

const profileData = ref({
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

  const prof = await api.fetchProfile()
  if (prof) {
    profileData.value = prof
  }
}

function setTab(tab) {
  currentTab.value = tab
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

// 01 Onboarding Handler: 點擊「完成，進入主頁」或「跳過，進入主頁」
function finishDailyTaskSetup() {
  setTab('feed')
}

// Handler actions
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

// 02 Feed Action: 點擊「我也要」
async function handleJoinPost(post) {
  post.is_joined = true
  await api.joinPost(post.id)

  // 加入我的任務 -> 公共任務
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

// 02 Feed Action: 提交成果打卡
async function handleSubmitResult({ postId, content }) {
  const post = posts.value.find(p => p.id === postId)
  const authorName = currentUser.value?.username || profileData.value.user?.username || 'lavender_nation'
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

  // 同步完成公共任務
  const taskIndex = myTasks.value.public.findIndex(t => t.post_id === postId)
  if (taskIndex !== -1) {
    const task = myTasks.value.public.splice(taskIndex, 1)[0]
    task.is_completed = true
    myTasks.value.completed.unshift(task)
    await api.completeMyTask(task.id)
  }
}

// 04 發文 Action: 發布貼文
async function handlePublishPost(payload) {
  const authorName = currentUser.value?.username || profileData.value.user?.username || 'lavender_nation'
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
    submissions: []
  }

  if (isTask) {
    createdPost.is_joined = true
    // 發起人自動加入任務列表
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
  profileData.value.stats.launches++
}

// 05 任務 Action: 新增私人任務
async function handleCreatePrivateTask(title) {
  const res = await api.createMyTask({
    title,
    description: '私人任務，永久留在自己的清單',
    task_type: 'private'
  })
  const newTask = res || {
    id: Date.now(),
    title,
    description: '私人任務，永久留在自己的清單',
    is_completed: false
  }
  myTasks.value.private.unshift(newTask)
}

// 05 任務 Action: 完成任務
async function handleCompleteTask(id) {
  // 檢查私人任務
  let taskIndex = myTasks.value.private.findIndex(t => t.id === id)
  if (taskIndex !== -1) {
    const task = myTasks.value.private.splice(taskIndex, 1)[0]
    task.is_completed = true
    myTasks.value.completed.unshift(task)
    profileData.value.stats.outcomes++
    await api.completeMyTask(id)
    return
  }

  // 檢查公共任務
  taskIndex = myTasks.value.public.findIndex(t => t.id === id)
  if (taskIndex !== -1) {
    const task = myTasks.value.public.splice(taskIndex, 1)[0]
    task.is_completed = true
    myTasks.value.completed.unshift(task)
    profileData.value.stats.outcomes++
    await api.completeMyTask(id)

    // 若對應貼文存在，同步留下完成打卡記錄
    if (task.post_id) {
      const post = posts.value.find(p => p.id === task.post_id)
      if (post) {
        if (!post.submissions) post.submissions = []
        const authorName = currentUser.value?.username || profileData.value.user?.username || 'lavender_nation'
        post.submissions.push({
          id: Date.now(),
          author_name: authorName,
          time_ago: '剛剛',
          content: '完成打卡！'
        })
        post.result_count = (post.result_count || 0) + 1
        await api.addSubmission(task.post_id, { content: '完成打卡！' })
      }
    }
  }
}

function handleUpdateTaskProgress({ id, note }) {
  const task = myTasks.value.private.find(t => t.id === id)
  if (task) task.progress_note = note
  api.updateTaskProgress(id, note)
}

// 06 靈感 Action
async function handleAddInspiration(payload) {
  const todayStr = '2026/06/24'
  const item = await api.createInspiration(payload)
  inspirations.value.unshift(item || { id: Date.now(), date_group: todayStr, ...payload })
  profileData.value.stats.inspirations++
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

// 07 個人 Action
async function handleUpdateProfile(data) {
  if (profileData.value.user) {
    profileData.value.user.bio = data.bio
    profileData.value.user.weekly_info = data.weekly_info
  }
  await api.updateProfile(data)
}

function handleLogout() {
  currentTab.value = 'feed'
}
</script>
