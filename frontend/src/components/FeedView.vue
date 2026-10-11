<template>
  <div class="feed-view">
    <!-- View Header (Figma Screen 02 & 03) -->
    <div class="feed-header">
      <div class="feed-header-top">
        <h1 class="view-title">貼文區</h1>
        <div class="header-icon-box">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#4B4B50" stroke-width="2">
            <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"></path>
            <path d="M13.73 21a2 2 0 0 1-3.46 0"></path>
          </svg>
        </div>
      </div>

      <!-- 搜尋列：搜尋使用者帳號、貼文 -->
      <div class="feed-search-bar">
        <div class="search-input-inner">
          <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="#75757A" stroke-width="2">
            <circle cx="11" cy="11" r="8"></circle>
            <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
          </svg>
          <input 
            v-model="searchQuery"
            type="text" 
            placeholder="搜尋使用者帳號、貼文..." 
            class="feed-search-input"
          />
          <button v-if="searchQuery" class="clear-search-btn" @click="searchQuery = ''">
            ✕
          </button>
        </div>

        <!-- 搜尋使用者下拉選單 -->
        <div v-if="searchQuery.trim()" class="user-search-results">
          <div class="user-search-title">使用者帳號 (點擊查看他人檔案)</div>
          <div v-if="matchingUsers.length === 0" class="user-search-empty">
            查無相符使用者，下方為貼文過濾結果
          </div>
          <div 
            v-for="user in matchingUsers" 
            :key="user.username" 
            class="user-search-item"
            @click="handleSelectUser(user.username)"
          >
            <div class="user-search-avatar">
              {{ user.username.charAt(0).toUpperCase() }}
            </div>
            <div class="user-search-meta">
              <div class="user-search-name">@{{ user.username }}</div>
              <div class="user-search-bio">{{ user.bio }}</div>
            </div>
            <div class="user-search-action">
              <span>查看檔案</span>
              <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <polyline points="9 18 15 12 9 6"/>
              </svg>
            </div>
          </div>
        </div>
      </div>
      
      <!-- Filters Row -->
      <div class="feed-filters-bar">
        <!-- 來源下拉選單 -->
        <div class="source-dropdown-wrap" @click="isSourceDropdownOpen = !isSourceDropdownOpen">
          <span class="source-text">{{ currentSource }}</span>
          <svg width="12" height="6" viewBox="0 0 12 6" fill="none" stroke="#141417" stroke-width="1.8">
            <polyline points="1 1 6 5 11 1"></polyline>
          </svg>

          <!-- Dropdown menu -->
          <div v-if="isSourceDropdownOpen" class="source-menu">
            <div 
              v-for="source in ['所有來源', '追蹤中', '熱門推薦']" 
              :key="source"
              class="source-menu-item"
              :class="{ active: currentSource === source }"
              @click.stop="selectSource(source)"
            >
              {{ source }}
            </div>
          </div>
        </div>

        <!-- 貼文類別 Pills -->
        <button 
          class="filter-pill" 
          :class="{ active: currentFilter === 'all' }"
          @click="currentFilter = 'all'"
        >
          所有貼文
        </button>
        <button 
          class="filter-pill" 
          :class="{ active: currentFilter === 'task' }"
          @click="currentFilter = 'task'"
        >
          任務貼文
        </button>
        <button 
          class="filter-pill" 
          :class="{ active: currentFilter === 'inspiration' }"
          @click="currentFilter = 'inspiration'"
        >
          靈感貼文
        </button>
      </div>
    </div>

    <!-- Posts Feed List -->
    <div class="posts-list">
      <article 
        v-for="post in filteredPosts" 
        :key="post.id" 
        class="figma-card post-card"
      >
        <!-- Top Tag Badge (Figma: 任務：已加入 / 任務：已完成) -->
        <div v-if="post.is_completed" class="status-badge badge-completed">
          任務：已完成
        </div>
        <div v-else-if="post.is_joined" class="status-badge badge-joined">
          任務：已加入
        </div>

        <!-- Author Row -->
        <div class="card-author-row" @click="$emit('view-user-profile', post.author_name)">
          <div class="author-avatar">
            <span class="avatar-letter">{{ (post.author_name || 'U').charAt(0).toUpperCase() }}</span>
          </div>
          <div class="author-meta">
            <div class="author-name">{{ post.author_name }}</div>
            <div class="author-subtitle">{{ post.expires_info }}</div>
          </div>
        </div>

        <!-- Post Content -->
        <h3 class="post-title">{{ post.title }}</h3>
        <p class="post-desc">{{ post.content }}</p>

        <!-- Preview / Image if applicable (Figma Screen 02 & 03: 照片預覽) -->
        <div v-if="post.post_type === 'inspiration' || post.image_url" class="photo-preview-container">
          <span class="preview-text">照片預覽</span>
        </div>

        <!-- Action & Stats Row -->
        <div class="card-bottom-row">
          <!-- Button States -->
          <div class="action-btn-col">
            <!-- 1. Not joined: 「我也要」 -->
            <button 
              v-if="!post.is_joined && !post.is_completed" 
              class="btn-dark-pill"
              @click="handleJoin(post)"
            >
              我也要
            </button>

            <!-- 2. Joined and not completed: 「執行任務」 (Figma Screen 16) -->
            <button 
              v-else-if="post.is_joined && !post.is_completed" 
              class="btn-outline-pill execute-btn" 
              @click="$emit('execute-task', post)"
            >
              執行任務
            </button>

            <!-- 3. Completed: 「✓任務完成」 (Figma Screen 03) -->
            <button 
              v-else-if="post.is_completed" 
              class="btn-outline-pill completed-btn" 
              @click="toggleThread(post.id)"
            >
              ✓任務完成
            </button>
          </div>

          <!-- Stats Counters (成果 3, 720 / 12, 4, 86) -->
          <div class="card-stats-col">
            <!-- If inspiration: 12 (bookmark), 4 (chat), 86 (share) -->
            <template v-if="post.post_type === 'inspiration'">
              <div class="stat-pill">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="#4B4B50" stroke-width="2">
                  <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
                </svg>
                <span>{{ post.saves_count ?? 12 }}</span>
              </div>
              <div class="stat-pill" @click="toggleThread(post.id)">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="#4B4B50" stroke-width="2">
                  <path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/>
                </svg>
                <span>{{ post.comments_count ?? 4 }}</span>
              </div>
              <div class="stat-pill">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="#4B4B50" stroke-width="2">
                  <path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"/>
                  <polyline points="16 6 12 2 8 6"/>
                  <line x1="12" y1="2" x2="12" y2="15"/>
                </svg>
                <span>{{ post.likes_count ?? 86 }}</span>
              </div>
            </template>

            <!-- If task: 成果 3 (chat), 720 (share) -->
            <template v-else>
              <div class="stat-pill" @click="toggleThread(post.id)">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="#4B4B50" stroke-width="2">
                  <path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/>
                </svg>
                <span>成果 {{ (post.submissions?.length ?? post.result_count) || 3 }}</span>
              </div>
              <div class="stat-pill">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="#4B4B50" stroke-width="2">
                  <path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"/>
                  <polyline points="16 6 12 2 8 6"/>
                  <line x1="12" y1="2" x2="12" y2="15"/>
                </svg>
                <span>{{ post.likes_count ?? 720 }}</span>
              </div>
            </template>
          </div>
        </div>

        <!-- Thread Timeline / Submissions (Figma Screen 03: 成果展示串接) -->
        <div v-if="post.is_completed || expandedThreadId === post.id" class="thread-timeline-section">
          <div class="timeline-divider"></div>

          <div 
            v-for="(sub, sIdx) in (post.submissions || defaultSubmissions)" 
            :key="sub.id || sIdx" 
            class="timeline-row-item"
          >
            <!-- Avatar and connected vertical line -->
            <div class="timeline-avatar-col">
              <div class="timeline-avatar">
                <span>{{ (sub.author_name || 'U').charAt(0).toUpperCase() }}</span>
              </div>
              <div v-if="sIdx < ((post.submissions || defaultSubmissions).length - 1)" class="timeline-vertical-line"></div>
            </div>

            <!-- Content -->
            <div class="timeline-content-col">
              <div class="timeline-header-info">
                <span class="timeline-user-name">{{ sub.author_name }}</span>
                <span class="timeline-time">{{ sub.time_ago || '2小時' }}</span>
              </div>
              <div class="timeline-text">{{ sub.content }}</div>

              <!-- Photo preview box in submission (e.g. kb_q_p 成果照片) -->
              <div v-if="sub.image_url" class="submission-photo-card">
                <span>成果照片</span>
              </div>
            </div>
          </div>

          <!-- Quick comment/reply form -->
          <div class="timeline-reply-row">
            <input 
              v-model="replyText[post.id]" 
              placeholder="分享一句話或打卡成果..." 
              class="reply-input"
              @keyup.enter="submitReply(post.id)"
            />
            <button class="reply-submit-btn" @click="submitReply(post.id)">
              送出
            </button>
          </div>
        </div>
      </article>

      <div v-if="filteredPosts.length === 0" class="empty-feed">
        目前沒有符合分類的貼文唷！
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  posts: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['join-post', 'submit-post-result', 'execute-task', 'view-user-profile'])

const currentFilter = ref('all')
const currentSource = ref('所有來源')
const isSourceDropdownOpen = ref(false)
const expandedThreadId = ref(null)
const replyText = ref({})
const searchQuery = ref('')

const defaultSubmissions = [
  { id: 1, author_name: 'kb_q_p', time_ago: '4小時', content: '應該長這樣', image_url: 'sky-1' },
  { id: 2, author_name: '同學A', time_ago: '2小時', content: '今天的天空是藍灰色，很適合當背景。' },
  { id: 3, author_name: '同學B', time_ago: '2小時', content: '我看倒像綠豆糕。' }
]

const knownUsers = computed(() => {
  const map = new Map()
  map.set('others_account', { username: 'others_account', bio: '把靈感變成行動，把行動留下成果' })
  map.set('bbb', { username: 'bbb', bio: '生活狀態打卡者 · 今日天空任務發起人' })
  map.set('dafuu0000', { username: 'dafuu0000', bio: '配色與靈感創作者 · 色票愛好者' })
  map.set('lavender_nation', { username: 'lavender_nation', bio: '天空觀察愛好者' })
  map.set('aaa', { username: 'aaa', bio: '生活整理打卡' })
  map.set('kb_q_p', { username: 'kb_q_p', bio: '成果紀錄達人' })

  if (props.posts) {
    props.posts.forEach(p => {
      if (p.author_name && !map.has(p.author_name)) {
        map.set(p.author_name, { username: p.author_name, bio: '社群創作者' })
      }
    })
  }
  return Array.from(map.values())
})

const matchingUsers = computed(() => {
  if (!searchQuery.value.trim()) return []
  const q = searchQuery.value.trim().toLowerCase()
  return knownUsers.value.filter(u => u.username.toLowerCase().includes(q))
})

function handleSelectUser(username) {
  emit('view-user-profile', username)
  searchQuery.value = ''
}

const filteredPosts = computed(() => {
  let list = props.posts
  if (currentFilter.value === 'task') {
    list = list.filter(p => p.post_type === 'task')
  } else if (currentFilter.value === 'inspiration') {
    list = list.filter(p => p.post_type === 'inspiration')
  }
  if (searchQuery.value.trim()) {
    const q = searchQuery.value.trim().toLowerCase()
    list = list.filter(p => 
      (p.title && p.title.toLowerCase().includes(q)) || 
      (p.content && p.content.toLowerCase().includes(q)) ||
      (p.author_name && p.author_name.toLowerCase().includes(q))
    )
  }
  return list
})

function selectSource(source) {
  currentSource.value = source
  isSourceDropdownOpen.value = false
}

function handleJoin(post) {
  emit('join-post', post)
  post.is_joined = true
}

function toggleThread(postId) {
  expandedThreadId.value = expandedThreadId.value === postId ? null : postId
}

function submitReply(postId) {
  const text = replyText.value[postId]
  if (!text || !text.trim()) return
  emit('submit-post-result', {
    postId,
    content: text.trim(),
    imageUrl: ''
  })
  replyText.value[postId] = ''
}
</script>

<style scoped>
.feed-view {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.feed-header {
  background: #FFFFFF;
  padding: 10px 16px 12px;
  border-radius: 8px;
  border: 1px solid #DBDBD9;
  display: flex;
  flex-direction: column;
  gap: 7px;
}

.feed-header-top {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.feed-search-bar {
  position: relative;
  width: 100%;
  margin: 4px 0 2px;
}

.search-input-inner {
  display: flex;
  align-items: center;
  gap: 8px;
  background: #F7F7F5;
  border: 1px solid #DBDBD9;
  border-radius: 20px;
  padding: 6px 14px;
  transition: all 0.15s ease;
}

.search-input-inner:focus-within {
  background: #FFFFFF;
  border-color: #141417;
  box-shadow: 0 0 0 2px rgba(20, 20, 23, 0.06);
}

.feed-search-input {
  flex: 1;
  border: none;
  background: transparent;
  outline: none;
  font-size: 13px;
  font-family: inherit;
  color: #141417;
}

.feed-search-input::placeholder {
  color: #9E9EA4;
}

.clear-search-btn {
  background: none;
  border: none;
  color: #75757A;
  font-size: 12px;
  cursor: pointer;
  padding: 0 4px;
}

.user-search-results {
  position: absolute;
  top: calc(100% + 6px);
  left: 0;
  right: 0;
  background: #FFFFFF;
  border: 1px solid #DBDBD9;
  border-radius: 12px;
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.08);
  z-index: 50;
  padding: 8px;
  display: flex;
  flex-direction: column;
  gap: 4px;
  max-height: 260px;
  overflow-y: auto;
}

.user-search-title {
  font-size: 11px;
  font-weight: 700;
  color: #75757A;
  padding: 4px 8px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.user-search-empty {
  font-size: 12px;
  color: #9E9EA4;
  padding: 8px;
  text-align: center;
}

.user-search-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 8px 10px;
  border-radius: 8px;
  cursor: pointer;
  transition: background 0.15s ease;
}

.user-search-item:hover {
  background: #F4F4F1;
}

.user-search-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: #D1D6DB;
  color: #141417;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 13px;
  font-weight: 700;
  flex-shrink: 0;
}

.user-search-meta {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.user-search-name {
  font-size: 13px;
  font-weight: 700;
  color: #141417;
}

.user-search-bio {
  font-size: 11px;
  color: #75757A;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.user-search-action {
  display: flex;
  align-items: center;
  gap: 3px;
  font-size: 12px;
  color: #141417;
  font-weight: 500;
}

.view-title {
  color: #141417;
  font-size: 24px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  margin: 0;
}

.header-icon-box {
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.feed-filters-bar {
  display: flex;
  align-items: center;
  gap: 8px;
  overflow-x: auto;
  scrollbar-width: none;
}
.feed-filters-bar::-webkit-scrollbar {
  display: none;
}

.source-dropdown-wrap {
  width: 84px;
  padding: 7px 8px;
  background: #FFFFFF;
  border-radius: 6px;
  outline: 1px solid #DBDBD9;
  outline-offset: -1px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 4px;
  cursor: pointer;
  position: relative;
  user-select: none;
}

.source-text {
  color: #141417;
  font-size: 13px;
  font-weight: 500;
  white-space: nowrap;
}

.source-menu {
  position: absolute;
  top: 100%;
  left: 0;
  margin-top: 4px;
  width: 100px;
  background: #FFFFFF;
  border-radius: 6px;
  border: 1px solid #DBDBD9;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  z-index: 100;
  overflow: hidden;
}

.source-menu-item {
  padding: 8px 10px;
  font-size: 12px;
  color: #141417;
  cursor: pointer;
  transition: background 0.15s ease;
}

.source-menu-item:hover, .source-menu-item.active {
  background: #F0F0ED;
  font-weight: 600;
}

.filter-pill {
  flex: 1;
  padding: 7px 12px;
  border-radius: 18px;
  outline: 1px solid #E5E5E5;
  outline-offset: -1px;
  border: none;
  background: transparent;
  color: #141417;
  font-size: 13px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 400;
  cursor: pointer;
  white-space: nowrap;
  transition: all 0.15s ease;
  display: flex;
  justify-content: center;
  align-items: center;
}

.filter-pill.active {
  background: #141417;
  color: #FFFFFF;
  outline-color: #141417;
  font-weight: 700;
}

.posts-list {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.post-card {
  background: #FFFFFF;
  border-radius: 8px;
  outline: 1px solid #DBDBD9;
  outline-offset: -1px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  box-sizing: border-box;
}

.status-badge {
  align-self: flex-start;
  padding: 7px 12px;
  border-radius: 18px;
  font-size: 13px;
  font-weight: 700;
  color: #FFFFFF;
  background: #141417;
  outline: 1px solid #141417;
  outline-offset: -1px;
}

.card-author-row {
  display: flex;
  align-items: center;
  gap: 9px;
  cursor: pointer;
}

.author-avatar {
  width: 34px;
  height: 34px;
  background: #D1D6DB;
  border-radius: 9999px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 14px;
  color: #4B4B50;
  flex-shrink: 0;
}

.author-meta {
  display: flex;
  flex-direction: column;
  gap: 1px;
}

.author-name {
  color: #141417;
  font-size: 14px;
  font-weight: 700;
  line-height: 1.2;
}

.author-subtitle {
  color: #75757A;
  font-size: 11px;
  font-weight: 400;
}

.post-title {
  color: #141417;
  font-size: 16px;
  font-weight: 700;
  margin: 0;
  line-height: 1.3;
}

.post-desc {
  color: #141417;
  font-size: 13px;
  font-weight: 400;
  line-height: 1.4;
  margin: 0;
}

.photo-preview-container {
  width: 100%;
  height: 110px;
  background: #F0F0ED;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.preview-text {
  color: #75757A;
  font-size: 13px;
  font-weight: 400;
}

.card-bottom-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 4px;
}

.action-btn-col {
  display: flex;
  align-items: center;
}

.btn-dark-pill {
  width: 88px;
  padding: 7px 12px;
  background: #141417;
  border-radius: 18px;
  border: 1px solid #141417;
  color: #FFFFFF;
  font-size: 13px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.15s ease;
  display: flex;
  justify-content: center;
  align-items: center;
}

.btn-dark-pill:hover {
  background: #25252b;
}

.btn-outline-pill {
  width: 88px;
  padding: 7px 12px;
  border-radius: 18px;
  outline: 1px solid #E5E5E5;
  outline-offset: -1px;
  background: #FFFFFF;
  border: none;
  color: #141417;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.15s ease;
  display: flex;
  justify-content: center;
  align-items: center;
}

.btn-outline-pill:hover {
  outline-color: #B8B8B8;
  background: #FBFBF9;
}

.card-stats-col {
  display: flex;
  align-items: center;
  gap: 16px;
}

.stat-pill {
  display: flex;
  align-items: center;
  gap: 5px;
  color: #75757A;
  font-size: 13px;
  cursor: pointer;
  transition: color 0.15s ease;
}

.stat-pill:hover {
  color: #141417;
}

.thread-timeline-section {
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-top: 8px;
}

.timeline-divider {
  width: 100%;
  height: 1px;
  background: #DBDBD9;
}

.timeline-row-item {
  display: flex;
  gap: 12px;
  padding: 4px 0;
}

.timeline-avatar-col {
  width: 34px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
}

.timeline-avatar {
  width: 34px;
  height: 34px;
  background: #D1D6DB;
  border-radius: 9999px;
  border: 1px solid #DBDBD9;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 13px;
  font-weight: 700;
  color: #4B4B50;
  flex-shrink: 0;
}

.timeline-vertical-line {
  width: 2px;
  flex: 1;
  min-height: 40px;
  background: #DBDBD9;
}

.timeline-content-col {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.timeline-header-info {
  display: flex;
  align-items: center;
  gap: 6px;
}

.timeline-user-name {
  color: #141417;
  font-size: 14px;
  font-weight: 700;
  line-height: 20px;
}

.timeline-time {
  color: #7A7A80;
  font-size: 13px;
  line-height: 18px;
}

.timeline-text {
  color: #141417;
  font-size: 14px;
  line-height: 20px;
}

.submission-photo-card {
  width: 100%;
  height: 92px;
  background: #F2F2F0;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #7A7A80;
  font-size: 13px;
  margin-top: 4px;
}

.timeline-reply-row {
  display: flex;
  gap: 8px;
  margin-top: 6px;
}

.reply-input {
  flex: 1;
  padding: 8px 12px;
  background: #FFFFFF;
  border: 1px solid #DBDBD9;
  border-radius: 18px;
  font-size: 13px;
  outline: none;
}

.reply-submit-btn {
  padding: 6px 14px;
  background: #141417;
  color: #FFFFFF;
  border: none;
  border-radius: 18px;
  font-size: 12px;
  font-weight: 600;
  cursor: pointer;
}

.empty-feed {
  padding: 30px;
  text-align: center;
  color: #75757A;
  background: #FFFFFF;
  border-radius: 8px;
  border: 1px solid #DBDBD9;
}
</style>
