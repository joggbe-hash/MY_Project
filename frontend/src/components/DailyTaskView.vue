<template>
  <div class="daily-task-view">
    <!-- View Header (Figma Screen 01 & 08 & 10) -->
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
          placeholder="輸入一件今天想完成的事情" 
          class="input-control"
          @keyup.enter="handleAddTask"
        />
      </div>

      <!-- Add Task Button -->
      <div class="add-btn-wrapper">
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
          :class="{ 'is-completed': task.is_completed }"
        >
          <div 
            class="daily-task-title"
            :class="{ 'title-done': task.is_completed }"
            @click="$emit('toggle-task', task.id)"
          >
            {{ idx + 1 }}. {{ cleanTitle(task.title) }}
          </div>

          <div class="task-action-icons">
            <!-- Toggle complete check icon -->
            <button 
              class="task-icon-btn check-icon-btn" 
              :title="task.is_completed ? '標記為未完成' : '標記為已完成'" 
              @click="$emit('toggle-task', task.id)"
            >
              <svg v-if="task.is_completed" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#141417" stroke-width="2.5">
                <polyline points="20 6 9 17 4 12"></polyline>
              </svg>
              <svg v-else width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#75757A" stroke-width="2">
                <circle cx="12" cy="12" r="9"></circle>
              </svg>
            </button>

            <!-- Delete icon -->
            <button 
              class="task-icon-btn delete-icon-btn" 
              title="刪除任務" 
              @click="$emit('delete-task', task.id)"
            >
              <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="#75757A" stroke-width="2">
                <polyline points="3 6 5 6 21 6"/>
                <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
              </svg>
            </button>
          </div>
        </div>

        <!-- Figma Screen 08 Empty State (無任務) -->
        <div v-if="tasks.length === 0" class="empty-tasks-box">
          <p class="empty-tasks-text">今天還沒有任何任務唷~</p>
        </div>
      </div>

      <!-- Bottom Completion / Skip Button -->
      <div class="task-card-footer" :class="{ 'align-right': tasks.length === 0 }">
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

const emit = defineEmits(['add-task', 'delete-task', 'toggle-task', 'finish-setup'])

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
  gap: 12px;
}

.view-header {
  background: #FFFFFF;
  border-radius: 8px;
  border: 1px solid #DBDBD9;
  padding: 12px 18px;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.view-header-title {
  color: #141417;
  font-size: 24px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  line-height: 1.2;
}

.view-header-subtitle {
  color: #75757A;
  font-size: 12px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 400;
}

.daily-task-card {
  background: #FFFFFF;
  border-radius: 8px;
  outline: 1px solid #DBDBD9;
  outline-offset: -1px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.card-heading {
  font-size: 18px;
  font-weight: 700;
  color: #141417;
}

.card-subheading {
  font-size: 13px;
  color: #75757A;
  margin-top: -4px;
  line-height: 1.4;
}

.task-input-wrapper {
  width: 100%;
}

.input-control {
  width: 100%;
  height: 48px;
  padding: 10px 12px;
  background: #FFFFFF;
  border: 1px solid #DBDBD9;
  border-radius: 8px;
  font-size: 14px;
  font-family: inherit;
  color: #141417;
  outline: none;
  box-sizing: border-box;
}

.input-control:focus {
  border-color: #141417;
}

.input-control::placeholder {
  color: #75757A;
}

.add-btn-wrapper {
  display: flex;
}

.add-task-btn {
  width: 124px;
  padding: 7px 12px;
  background: #141417;
  border-radius: 18px;
  border: 1px solid #141417;
  color: #FFFFFF;
  font-size: 13px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.15s ease;
}

.add-task-btn:hover {
  background: #25252b;
}

.add-task-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.tasks-section-header {
  font-size: 15px;
  font-weight: 700;
  color: #141417;
  padding-top: 4px;
}

.daily-tasks-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.daily-task-row {
  height: 46px;
  padding: 12px 14px;
  background: #FFFFFF;
  border-radius: 8px;
  outline: 1px solid #DBDBD9;
  outline-offset: -1px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 10px;
  box-sizing: border-box;
  transition: all 0.15s ease;
}

.daily-task-row:hover {
  outline-color: #B8B8B8;
}

.daily-task-title {
  flex: 1;
  color: #141417;
  font-size: 13px;
  font-weight: 400;
  cursor: pointer;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.title-done {
  color: #B8B8B8 !important;
  text-decoration: line-through !important;
}

.task-action-icons {
  display: flex;
  align-items: center;
  gap: 6px;
}

.task-icon-btn {
  background: none;
  border: none;
  padding: 4px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 4px;
  transition: all 0.15s ease;
}

.task-icon-btn:hover {
  background: #F0F0ED;
}

.empty-tasks-box {
  width: 100%;
  padding: 24px 0;
  text-align: center;
}

.empty-tasks-text {
  color: #75757A;
  font-size: 13px;
}

.task-card-footer {
  display: flex;
  margin-top: 8px;
}

.task-card-footer.align-right {
  justify-content: flex-end;
}

.finish-all-btn {
  width: 100%;
  padding: 7px 12px;
  background: #141417;
  border-radius: 18px;
  border: 1px solid #141417;
  color: #FFFFFF;
  font-size: 13px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.15s ease;
}

.finish-all-btn:hover {
  background: #25252b;
}

.skip-btn {
  width: 124px;
  padding: 7px 12px;
  background: #FFFFFF;
  border-radius: 18px;
  outline: 1px solid #DBDBD9;
  outline-offset: -1px;
  border: none;
  color: #141417;
  font-size: 13px;
  font-weight: 400;
  cursor: pointer;
  transition: all 0.15s ease;
}

.skip-btn:hover {
  background: #F5F5F7;
  outline-color: #B8B8B8;
}
</style>
