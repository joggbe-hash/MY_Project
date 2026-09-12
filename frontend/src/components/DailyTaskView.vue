<template>
  <div class="daily-task-view">
    <!-- View Header (Figma Screen 01) -->
    <div class="view-header">
      <div class="view-header-title">我也要</div>
      <div class="view-header-subtitle">先把今天想做的事列出來，完成後進入貼文區</div>
    </div>

    <!-- Main Card -->
    <div class="figma-card daily-task-card">
      <h2 class="card-heading">今天是否有任務？</h2>
      <p class="card-subheading">可以連續加入多個今日任務，都設定好了再按完成。</p>

      <!-- Input Row -->
      <div class="task-input-wrapper">
        <input 
          v-model="inputTask" 
          type="text" 
          placeholder="輸入一件今天想完成的小事" 
          class="input-control"
          @keyup.enter="handleAddTask"
        />
      </div>

      <!-- Add Task Button -->
      <div>
        <button 
          class="btn-dark add-task-btn" 
          :disabled="!inputTask.trim()"
          @click="handleAddTask"
        >
          加入任務
        </button>
      </div>

      <!-- Tasks Section Header -->
      <div class="tasks-section-header">已加入今日任務</div>

      <!-- Task Items or Empty State -->
      <div class="daily-tasks-list">
        <div 
          v-for="(task, idx) in tasks" 
          :key="task.id" 
          class="daily-task-row"
        >
          <div class="daily-task-title">
            {{ idx + 1 }}. {{ cleanTitle(task.title) }}
          </div>
          <button class="task-delete-icon" title="刪除任務" @click="$emit('delete-task', task.id)">
            <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polyline points="3 6 5 6 21 6"/>
              <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
            </svg>
          </button>
        </div>

        <!-- Figma Screen 01 Empty State (無任務) -->
        <div v-if="tasks.length === 0" class="empty-tasks-box">
          <p class="empty-tasks-text">今天還沒有任何任務唷~</p>
        </div>
      </div>

      <!-- Bottom Completion / Skip Button -->
      <div class="task-card-footer">
        <button 
          v-if="tasks.length > 0" 
          class="btn-dark finish-all-btn" 
          @click="$emit('finish-setup')"
        >
          完成，進入主頁
        </button>
        <button 
          v-else 
          class="btn-outline skip-btn" 
          @click="$emit('finish-setup')"
        >
          跳過，進入主頁
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
  tasks: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['add-task', 'delete-task', 'finish-setup'])

const inputTask = ref('')

function handleAddTask() {
  if (!inputTask.value.trim()) return
  emit('add-task', inputTask.value.trim())
  inputTask.value = ''
}

function cleanTitle(title) {
  return (title || '').replace(/^\d+\.\s*/, '')
}
</script>

<style scoped>
.daily-task-view {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.daily-task-card {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.card-heading {
  font-size: 18px;
  font-weight: 700;
  color: var(--text-main);
}

.card-subheading {
  font-size: 13px;
  color: var(--text-secondary);
  line-height: 1.4;
  margin-top: -6px;
}

.task-input-wrapper {
  width: 100%;
}

.add-task-btn {
  padding: 8px 18px;
  font-size: 13px;
}

.tasks-section-header {
  font-size: 15px;
  font-weight: 700;
  color: var(--text-main);
  padding-top: 4px;
}

.daily-tasks-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.daily-task-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 14px;
  background: var(--bg-card);
  border: 1px solid var(--border-main);
  border-radius: var(--radius-md);
}

.daily-task-title {
  font-size: 14px;
  color: var(--text-main);
  line-height: 1.3;
}

.task-delete-icon {
  background: none;
  border: none;
  color: var(--icon-color);
  cursor: pointer;
  padding: 4px;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.task-delete-icon:hover {
  color: #FA5252;
  background: var(--bg-subtle);
}

.empty-tasks-box {
  padding: 30px 16px;
  text-align: center;
}

.empty-tasks-text {
  font-size: 13px;
  color: var(--text-secondary);
}

.task-card-footer {
  display: flex;
  justify-content: flex-end;
  padding-top: 12px;
}

.finish-all-btn {
  width: 100%;
  padding: 12px;
  font-size: 15px;
  font-weight: 700;
}

.skip-btn {
  padding: 8px 24px;
  border-radius: var(--radius-pill);
  font-size: 13px;
}
</style>
