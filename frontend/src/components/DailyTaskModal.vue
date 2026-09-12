<template>
  <div class="modal-backdrop" @click.self="$emit('close')">
    <div class="modal-card daily-task-modal">
      <!-- Modal Header -->
      <div class="modal-top">
        <h2 class="modal-title">今天是否有任務？</h2>
        <button class="close-btn" @click="$emit('close')">✕</button>
      </div>
      <p class="modal-subtitle">可以連續加入多個今日任務，都設定好了再按完成。</p>

      <!-- Input Row -->
      <div class="input-row">
        <input 
          v-model="inputTask" 
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

      <!-- Tasks Section -->
      <div class="tasks-list-header">已加入今日任務</div>

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
          <button class="task-delete-icon" @click="$emit('delete-task', task.id)">
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polyline points="3 6 5 6 21 6"/>
              <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
            </svg>
          </button>
        </div>

        <!-- Figma Screen 8 Empty State -->
        <div v-if="tasks.length === 0" class="empty-tasks-box">
          <p class="empty-tasks-text">今天還沒有任何任務唷~</p>
        </div>
      </div>

      <!-- Bottom Completion / Skip Button -->
      <div class="modal-bottom-action">
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

const emit = defineEmits(['close', 'add-task', 'delete-task', 'finish-setup'])

const inputTask = ref('')

function handleAddTask() {
  if (!inputTask.value.trim()) return
  emit('add-task', inputTask.value.trim())
  inputTask.value = ''
}

function cleanTitle(title) {
  // strip existing leading number prefix if user inputs plain text
  return title.replace(/^\d+\.\s*/, '')
}
</script>

<style scoped>
.daily-task-modal {
  max-width: 420px;
  gap: 14px;
}

.modal-top {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-title {
  font-size: 18px;
  font-weight: 700;
  color: var(--text-main);
}

.modal-subtitle {
  font-size: 13px;
  color: var(--text-secondary);
  line-height: 1.4;
  margin-top: -6px;
}

.close-btn {
  background: none;
  border: none;
  font-size: 18px;
  color: var(--text-secondary);
  cursor: pointer;
}

.add-task-btn {
  padding: 8px 16px;
  font-size: 13px;
}

.tasks-list-header {
  font-size: 15px;
  font-weight: 700;
  color: var(--text-main);
  padding-top: 4px;
}

.daily-tasks-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
  max-height: 220px;
  overflow-y: auto;
}

.daily-task-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 14px;
  background: var(--bg-card);
  border: 1px solid var(--border-main);
  border-radius: var(--radius-md);
}

.daily-task-title {
  font-size: 13px;
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
  padding: 24px;
  text-align: center;
}

.empty-tasks-text {
  font-size: 13px;
  color: var(--text-secondary);
}

.modal-bottom-action {
  display: flex;
  justify-content: center;
  padding-top: 10px;
}

.finish-all-btn {
  width: 100%;
  padding: 10px;
}

.skip-btn {
  padding: 8px 24px;
  border-radius: var(--radius-pill);
}
</style>
