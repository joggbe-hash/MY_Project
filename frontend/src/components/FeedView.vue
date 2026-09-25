<template>
  <div class="feed-view">
    <!-- View Header (Figma Screen 02) -->
    <div class="feed-header">
      <h1 class="view-title">貼文區</h1>
      
      <!-- Category Tabs (所有貼文, 任務貼文, 靈感貼文) -->
      <div class="category-tabs-row">
        <button 
          class="cat-tab-pill" 
          :class="{ active: currentFilter === 'all' }"
          @click="currentFilter = 'all'"
        >
          所有貼文
        </button>
        <button 
          class="cat-tab-pill" 
          :class="{ active: currentFilter === 'task' }"
          @click="currentFilter = 'task'"
        >
          任務貼文
        </button>
        <button 
          class="cat-tab-pill" 
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
        <!-- Top Joined / Completed Badge (Figma Screen 02_展開 / Screen 03) -->
        <div v-if="post.is_completed" class="joined-badge completed-badge">
          任務：已完成
        </div>
        <div v-else-if="post.is_joined" class="joined-badge">
          任務：已加入
        </div>

        <!-- Author Row -->
        <div class="card-author-row">
          <div class="avatar">
            <span class="avatar-letter">{{ post.author_name ? post.author_name.charAt(0).toUpperCase() : 'U' }}</span>
          </div>
          <div class="author-meta">
            <div class="author-name">{{ post.author_name }}</div>
            <div class="author-subtitle">{{ post.expires_info }}</div>
          </div>
        </div>

        <!-- Post Content -->
        <h3 class="post-title">{{ post.title }}</h3>
        <p class="post-desc">{{ post.content }}</p>

        <!-- Preview / Image if applicable -->
        <div v-if="post.image_url && post.image_url !== 'uploaded-photo-preview'" class="preview-box">
          <span class="preview-placeholder-text">照片預覽 / 色票紀錄</span>
        </div>

        <!-- Action Row -->
        <div class="card-bottom-row">
          <div class="action-btn-col">
            <button 
              v-if="!post.is_joined && !post.is_completed" 
              class="btn-dark join-btn"
              @click="handleJoin(post)"
            >
              我也要
            </button>
            <button 
              v-else-if="post.is_joined && !post.is_completed" 
              class="btn-outline finish-btn" 
              @click="openSubmitResult(post)"
            >
              執行任務
            </button>
            <button 
              v-else-if="post.is_completed" 
              class="btn-outline completed-btn" 
              @click="openSubmitResult(post)"
            >
              ✓任務完成
            </button>
          </div>

          <!-- Stats Row (Figma: 成果 3, 720) -->
          <div class="card-stats-col">
            <div class="stat-bubble" v-if="post.post_type === 'task'" @click="toggleThread(post.id)">
              <!-- Chat icon from Figma Screen 02 -->
              <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/>
              </svg>
              <span>成果 {{ (post.submissions?.length ?? post.result_count) || 0 }}</span>
            </div>

            <div class="stat-bubble">
              <!-- Upload/Share icon from Figma Screen 02 -->
              <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"/>
                <polyline points="16 6 12 2 8 6"/>
                <line x1="12" y1="2" x2="12" y2="15"/>
              </svg>
              <span>{{ post.likes_count ?? 720 }}</span>
            </div>
          </div>
        </div>

        <!-- Timeline Thread (Figma Screen 02_展開) -->
        <div v-if="post.is_joined || expandedThreadId === post.id" class="thread-timeline">
          <!-- Dynamic submissions list -->
          <div 
            v-for="(sub, sIdx) in (post.submissions || [])" 
            :key="sub.id || sIdx" 
            class="timeline-item"
          >
            <div class="timeline-line-col">
              <div class="avatar avatar-sm">
                <span>{{ (sub.author_name || 'U').charAt(0).toUpperCase() }}</span>
              </div>
              <div v-if="sIdx < (post.submissions.length - 1)" class="timeline-line"></div>
            </div>
            <div class="timeline-content">
              <div class="timeline-author-header">
                <span class="name">{{ sub.author_name }}</span>
                <span class="time">{{ sub.time_ago || '剛剛' }}</span>
              </div>
              <div class="timeline-body">{{ sub.content }}</div>
              <div v-if="sub.image_url" class="submission-photo-box">
                <span>成果照片</span>
              </div>
            </div>
          </div>

          <!-- Quick reply row when joined -->
          <div v-if="post.is_joined" class="reply-input-row">
            <input 
              v-model="replyText[post.id]" 
              placeholder="分享你的成果心得..." 
              class="reply-input"
              @keyup.enter="submitReply(post.id)"
            />
            <button class="btn-dark btn-xs" @click="submitReply(post.id)">送出</button>
          </div>
        </div>
      </article>

      <div v-if="filteredPosts.length === 0" class="empty-feed">
        目前沒有符合分類的貼文唷！
      </div>
    </div>

    <!-- Result Submission Modal -->
    <div v-if="showResultModal" class="modal-backdrop" @click.self="showResultModal = false">
      <div class="modal-card">
        <h3 class="card-title">分享「{{ activePostToSubmit?.title }}」打卡成果</h3>
        <p class="card-desc">分享一張照片與一句話，記錄你今天的完成成果。</p>
        <textarea 
          v-model="resultNote" 
          class="textarea-control" 
          placeholder="今天的心得、感想或紀錄..."
        ></textarea>
        <div class="preview-box mock-upload-trigger" @click="mockAttachPhoto">
          <span>{{ hasAttachedPhoto ? '✓ 已附加成果照片' : '＋ 點擊加入成果照片' }}</span>
        </div>
        <div class="modal-footer">
          <button class="btn-outline" @click="showResultModal = false">取消</button>
          <button class="btn-dark" @click="confirmSubmitResult">發布成果</button>
        </div>
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

const emit = defineEmits(['join-post', 'submit-post-result', 'change-tab'])

const currentFilter = ref('all')
const expandedThreadId = ref(null)
const replyText = ref({})
const showResultModal = ref(false)
const activePostToSubmit = ref(null)
const resultNote = ref('')
const hasAttachedPhoto = ref(false)

const filteredPosts = computed(() => {
  if (currentFilter.value === 'task') {
    return props.posts.filter(p => p.post_type === 'task')
  }
  if (currentFilter.value === 'inspiration') {
    return props.posts.filter(p => p.post_type === 'inspiration')
  }
  return props.posts
})

function handleJoin(post) {
  emit('join-post', post)
  post.is_joined = true
  expandedThreadId.value = post.id
}

function toggleThread(postId) {
  expandedThreadId.value = expandedThreadId.value === postId ? null : postId
}

function openSubmitResult(post) {
  activePostToSubmit.value = post
  resultNote.value = ''
  hasAttachedPhoto.value = false
  showResultModal.value = true
}

function mockAttachPhoto() {
  hasAttachedPhoto.value = !hasAttachedPhoto.value
}

function confirmSubmitResult() {
  if (!resultNote.value.trim() && !activePostToSubmit.value) return
  emit('submit-post-result', {
    postId: activePostToSubmit.value.id,
    content: resultNote.value.trim() || '完成今日打卡！',
    imageUrl: hasAttachedPhoto.value ? 'placeholder-user-sky' : ''
  })
  showResultModal.value = false
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
  gap: 16px;
}

.feed-header {
  display: flex;
  flex-direction: column;
  gap: 12px;
  padding: 4px 2px;
}

.view-title {
  font-size: 24px;
  font-weight: 700;
  color: var(--text-main);
}

.category-tabs-row {
  display: flex;
  gap: 8px;
}

.cat-tab-pill {
  padding: 7px 18px;
  border-radius: var(--radius-pill);
  font-size: 13px;
  font-weight: 500;
  border: 1px solid var(--border-light);
  background: var(--bg-card);
  color: var(--text-main);
  cursor: pointer;
  transition: var(--transition);
}

.cat-tab-pill.active {
  background: var(--bg-dark);
  color: var(--text-white);
  border-color: var(--bg-dark);
  font-weight: 700;
}

.posts-list {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.post-card {
  padding: 18px 20px;
  gap: 12px;
}

.joined-badge {
  display: inline-flex;
  align-items: center;
  padding: 4px 12px;
  background: var(--bg-dark);
  color: #fff;
  font-size: 11px;
  font-weight: 700;
  border-radius: var(--radius-pill);
  align-self: flex-start;
}

.card-author-row {
  display: flex;
  align-items: center;
  gap: 10px;
}

.avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background-color: var(--avatar-default);
  display: flex;
  align-items: center;
  justify-content: center;
}

.avatar-letter {
  font-size: 14px;
  font-weight: 700;
  color: #555;
}

.author-meta {
  display: flex;
  flex-direction: column;
}

.author-name {
  font-size: 14px;
  font-weight: 700;
  color: var(--text-main);
}

.author-subtitle {
  font-size: 12px;
  color: var(--text-secondary);
}

.post-title {
  font-size: 16px;
  font-weight: 700;
  color: var(--text-main);
  line-height: 1.4;
}

.post-desc {
  font-size: 13px;
  color: var(--text-main);
  line-height: 1.5;
}

.preview-box {
  width: 100%;
  height: 120px;
  background: var(--bg-subtle);
  border-radius: var(--radius-md);
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text-secondary);
  font-size: 13px;
  border: 1px solid var(--border-card);
}

.card-bottom-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 4px;
}

.join-btn {
  padding: 6px 18px;
  font-size: 13px;
}

.finish-btn {
  padding: 6px 16px;
  font-size: 12px;
}

.card-stats-col {
  display: flex;
  align-items: center;
  gap: 14px;
  color: var(--text-secondary);
  font-size: 12px;
}

.stat-bubble {
  display: flex;
  align-items: center;
  gap: 5px;
  cursor: pointer;
}

.thread-timeline {
  display: flex;
  flex-direction: column;
  margin-top: 6px;
  border-top: 1px solid var(--border-light);
  padding-top: 14px;
}

.timeline-item {
  display: flex;
  gap: 12px;
  padding-bottom: 14px;
}

.timeline-line-col {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 28px;
  flex-shrink: 0;
}

.avatar-sm {
  width: 26px;
  height: 26px;
  border-radius: 50%;
  background: var(--avatar-default);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 11px;
}

.timeline-line {
  width: 1.5px;
  flex: 1;
  min-height: 20px;
  background: var(--border-main);
  margin-top: 4px;
}

.timeline-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.timeline-author-header {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
}

.timeline-author-header .name {
  font-weight: 700;
  color: var(--text-main);
}

.timeline-author-header .time {
  color: var(--text-secondary);
  font-size: 11px;
}

.timeline-body {
  font-size: 13px;
  color: var(--text-main);
  line-height: 1.4;
}

.submission-photo-box {
  width: 100%;
  height: 90px;
  background: #EAEAEA;
  border-radius: var(--radius-md);
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text-secondary);
  font-size: 12px;
  margin-top: 4px;
}

.reply-input-row {
  display: flex;
  gap: 8px;
  margin-top: 8px;
  padding-top: 10px;
  border-top: 1px dashed var(--border-light);
}

.reply-input {
  flex: 1;
  padding: 6px 12px;
  border-radius: var(--radius-pill);
  border: 1px solid var(--border-main);
  background: var(--bg-card);
  font-size: 12px;
  outline: none;
}

.reply-input:focus {
  border-color: var(--bg-dark);
}

.empty-feed {
  padding: 40px;
  text-align: center;
  color: var(--text-secondary);
  font-size: 13px;
  background: var(--bg-card);
  border-radius: var(--radius-md);
  border: 1px dashed var(--border-main);
}

.btn-xs {
  padding: 4px 12px;
  font-size: 11px;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  margin-top: 8px;
}

.mock-upload-trigger {
  cursor: pointer;
}
</style>
