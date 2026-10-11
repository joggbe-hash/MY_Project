<template>
  <aside class="desktop-right-panel">
    <!-- Today's Tasks Widget -->
    <div class="widget-card">
      <div class="widget-header">
        <div class="widget-title">今日任務</div>
        <button class="widget-action-link" @click="$emit('open-daily-task')">管理</button>
      </div>
      <p class="widget-subtitle">設定好今天想做的小事，都完成了再打卡進入主頁。</p>

      <!-- Quick Add Input -->
      <div class="quick-add-form" @submit.prevent="submitTask">
        <input 
          v-model="newTaskTitle" 
          type="text" 
          placeholder="輸入想完成的小事..." 
          class="quick-input"
          @keyup.enter="submitTask"
        />
        <button class="quick-add-btn" :disabled="!newTaskTitle.trim()" @click="submitTask">
          ＋
        </button>
      </div>

      <!-- Task Items List -->
      <div class="task-mini-list">
        <div 
          v-for="(task, idx) in dailyTasks" 
          :key="task.id" 
          class="task-mini-item"
          :class="{ done: task.is_completed }"
          @click="$emit('toggle-daily-task', task.id)"
        >
          <div class="task-checkbox" :class="{ checked: task.is_completed }">
            <svg v-if="task.is_completed" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3">
              <polyline points="20 6 9 17 4 12"/>
            </svg>
          </div>
          <span class="task-text">{{ idx + 1 }}. {{ cleanTitle(task.title) }}</span>
        </div>

        <div v-if="dailyTasks.length === 0" class="task-empty-hint">
          今天還沒有任何任務唷~
        </div>
      </div>
    </div>

    <!-- Quick Inspiration Note Widget -->
    <div class="widget-card">
      <div class="widget-header">
        <div class="widget-title">靈感即時速記</div>
        <button class="widget-action-link" @click="$emit('change-tab', 'inspirations')">全部</button>
      </div>
      <p class="widget-subtitle">今天有甚麼想法嗎？馬上記錄下來吧！</p>
      
      <div class="quick-note-box">
        <textarea 
          v-model="quickInspiration" 
          placeholder="捕捉生活中的微小靈感..."
          class="quick-textarea"
          rows="2"
        ></textarea>
        <div class="quick-note-footer">
          <span class="note-tip">隨手記錄保留 30 天</span>
          <button 
            class="btn-dark btn-sm" 
            :disabled="!quickInspiration.trim()"
            @click="submitInspiration"
          >
            記下
          </button>
        </div>
      </div>
    </div>

    <!-- Trending / Community Tasks -->
    <div class="widget-card trending-card">
      <div class="widget-header">
        <div class="widget-title">社群熱門挑戰</div>
      </div>
      <div class="trending-list">
        <div 
          v-for="post in trendingPosts" 
          :key="post.id" 
          class="trending-item" 
          @click="$emit('change-tab', 'feed')"
        >
          <div class="trending-tag" :class="{ 'tag-inspiration': post.post_type === 'inspiration' }">
            {{ post.post_type === 'task' ? '任務' : '靈感' }}
          </div>
          <div class="trending-info">
            <div class="trending-title">{{ post.title }}</div>
            <div class="trending-stats">
              <span v-if="post.post_type === 'task'">
                {{ post.likes_count ?? 0 }} 人也在做 · 成果 {{ (post.submissions?.length ?? post.result_count) || 0 }}
              </span>
              <span v-else>
                {{ post.likes_count ?? 0 }} 人喜歡 · {{ post.saves_count ?? 0 }} 人收藏
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </aside>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  dailyTasks: {
    type: Array,
    default: () => []
  },
  posts: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['add-daily-task', 'toggle-daily-task', 'open-daily-task', 'change-tab', 'add-inspiration'])

const newTaskTitle = ref('')
const quickInspiration = ref('')

const trendingPosts = computed(() => {
  if (props.posts && props.posts.length > 0) {
    return [...props.posts]
      .sort((a, b) => (b.likes_count || 0) - (a.likes_count || 0))
      .slice(0, 2)
  }
  return [
    { id: 1, post_type: 'task', title: '拍一張今天的天空', likes_count: 720, result_count: 3 },
    { id: 2, post_type: 'inspiration', title: '把通勤路上的顏色做成色票', likes_count: 86, saves_count: 12 }
  ]
})

function cleanTitle(title) {
  return (title || '').replace(/^\d+\.\s*/, '')
}

function submitTask() {
  if (!newTaskTitle.value.trim()) return
  emit('add-daily-task', newTaskTitle.value.trim())
  newTaskTitle.value = ''
}

function submitInspiration() {
  if (!quickInspiration.value.trim()) return
  emit('add-inspiration', quickInspiration.value.trim())
  quickInspiration.value = ''
}
</script>

<style scoped>
.widget-card {
  background: var(--bg-card);
  border-radius: 16px;
  border: 1px solid var(--border-main);
  padding: 18px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  box-shadow: var(--shadow-sm);
}

.widget-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.widget-title {
  font-size: 15px;
  font-weight: 700;
  color: var(--text-main);
}

.widget-action-link {
  font-size: 12px;
  color: var(--text-secondary);
  background: none;
  border: none;
  cursor: pointer;
  padding: 2px 6px;
  border-radius: 4px;
}

.widget-action-link:hover {
  color: var(--text-main);
  background: var(--bg-secondary);
}

.widget-subtitle {
  font-size: 12px;
  color: var(--text-secondary);
  line-height: 1.4;
  margin-top: -6px;
}

.quick-add-form {
  display: flex;
  gap: 6px;
}

.quick-input {
  flex: 1;
  padding: 8px 12px;
  border-radius: 8px;
  border: 1px solid var(--border-main);
  background: var(--bg-app);
  font-size: 13px;
  font-family: inherit;
  outline: none;
}

.quick-input:focus {
  border-color: var(--bg-dark);
}

.quick-add-btn {
  width: 34px;
  height: 34px;
  border-radius: 8px;
  background: var(--bg-dark);
  color: #fff;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  font-weight: 700;
  flex-shrink: 0;
  transition: var(--transition);
}

.quick-add-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.task-mini-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
  max-height: 160px;
  overflow-y: auto;
}

.task-mini-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 8px 10px;
  border-radius: 8px;
  border: 1px solid var(--border-card);
  background: var(--bg-app);
  cursor: pointer;
  font-size: 13px;
  transition: var(--transition);
}

.task-mini-item:hover {
  border-color: var(--border-main);
}

.task-mini-item.done {
  opacity: 0.6;
}

.task-mini-item.done .task-text {
  text-decoration: line-through;
}

.task-checkbox {
  width: 18px;
  height: 18px;
  border-radius: 4px;
  border: 1.5px solid var(--border-strong);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.task-checkbox.checked {
  background: var(--bg-dark);
  border-color: var(--bg-dark);
  color: #fff;
}

.task-text {
  flex: 1;
  color: var(--text-main);
  font-size: 13px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.task-empty-hint {
  text-align: center;
  color: var(--text-secondary);
  font-size: 12px;
  padding: 12px 0;
}

/* Quick Note */
.quick-note-box {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.quick-textarea {
  width: 100%;
  padding: 10px;
  border-radius: 8px;
  border: 1px solid var(--border-main);
  background: var(--bg-app);
  font-size: 13px;
  font-family: inherit;
  resize: none;
  outline: none;
}

.quick-textarea:focus {
  border-color: var(--bg-dark);
}

.quick-note-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.note-tip {
  font-size: 11px;
  color: var(--text-muted);
}

.btn-sm {
  padding: 5px 12px;
  font-size: 12px;
}

/* Trending */
.trending-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.trending-item {
  display: flex;
  align-items: flex-start;
  gap: 10px;
  padding: 10px;
  border-radius: 8px;
  border: 1px solid var(--border-light);
  cursor: pointer;
  transition: var(--transition);
}

.trending-item:hover {
  background: var(--bg-subtle);
  border-color: var(--border-main);
}

.trending-tag {
  font-size: 10px;
  font-weight: 700;
  padding: 2px 6px;
  border-radius: 4px;
  background: var(--bg-dark);
  color: #fff;
  flex-shrink: 0;
  margin-top: 2px;
}

.trending-tag.tag-inspiration {
  background: var(--text-secondary);
}

.trending-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.trending-title {
  font-size: 13px;
  font-weight: 700;
  color: var(--text-main);
  line-height: 1.3;
}

.trending-stats {
  font-size: 11px;
  color: var(--text-secondary);
}
</style>
