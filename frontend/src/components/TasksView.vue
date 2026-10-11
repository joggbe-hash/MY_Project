<template>
  <div class="tasks-view">
    <!-- View Header (Figma Screen 05: 我的任務) -->
    <div class="view-header">
      <div class="view-header-title">我的任務</div>
    </div>

    <!-- Section 1: 今日私人任務 -->
    <div class="task-category-card">
      <div class="section-title-row">
        <h3 class="section-title">今日私人任務</h3>
        <button class="btn-add-mini" @click="showAddPrivate = !showAddPrivate">
          {{ showAddPrivate ? '取消' : '＋ 新增' }}
        </button>
      </div>

      <!-- Inline Add Form -->
      <div v-if="showAddPrivate" class="inline-add-row">
        <input 
          v-model="newPrivateTitle" 
          placeholder="輸入私人任務名稱..." 
          class="inline-input"
          @keyup.enter="submitPrivateTask"
        />
        <button class="btn-dark-inline" @click="submitPrivateTask">確認</button>
      </div>

      <!-- Private Tasks List -->
      <div class="task-items-container">
        <div 
          v-for="task in privateTasks" 
          :key="task.id" 
          class="task-item-box"
        >
          <div class="task-info-col">
            <div class="task-title-text">{{ task.title }}</div>
            <div class="task-sub-text">{{ task.description || '私人任務，會保留紀錄' }}</div>
          </div>

          <div class="task-btn-row">
            <button 
              class="btn-execute-pill" 
              @click="$emit('execute-task', task)"
            >
              執行任務
            </button>
          </div>
        </div>

        <div v-if="privateTasks.length === 0" class="empty-task-hint">
          目前沒有私人任務，點擊上方「＋ 新增」即可建立！
        </div>
      </div>
    </div>

    <!-- Section 2: 公共任務 -->
    <div class="task-category-card">
      <div class="section-title-row">
        <h3 class="section-title">公共任務</h3>
      </div>

      <!-- Public Tasks List -->
      <div class="task-items-container">
        <div 
          v-for="task in publicTasks" 
          :key="task.id" 
          class="task-item-box"
        >
          <div class="task-info-col">
            <div class="task-title-text">{{ task.title }}</div>
            <div class="task-sub-text">{{ task.description || '來自任務貼文，完成後需分享成果' }}</div>
          </div>

          <div class="task-btn-row">
            <button 
              class="btn-execute-pill" 
              @click="$emit('execute-task', task)"
            >
              執行任務
            </button>
          </div>
        </div>

        <div v-if="publicTasks.length === 0" class="empty-task-hint">
          尚未加入任何公共任務，前往貼文區點擊「我也要」即可加入！
        </div>
      </div>
    </div>

    <!-- Section 3: 完成紀錄 -->
    <div class="task-category-card">
      <div class="completed-section-header" @click="isCompletedExpanded = !isCompletedExpanded">
        <h3 class="section-title">完成紀錄</h3>
        <span class="count-tag">{{ completedTasks.length }} 則紀錄 {{ isCompletedExpanded ? '▲' : '▼' }}</span>
      </div>

      <div v-if="isCompletedExpanded && completedTasks.length > 0" class="completed-list">
        <div 
          v-for="task in completedTasks" 
          :key="task.id" 
          class="completed-item-row"
        >
          <div class="completed-info">
            <span class="completed-check-icon">✓</span>
            <span class="completed-title">{{ task.title }}</span>
          </div>
          <span class="completed-status-tag">已完成</span>
        </div>
      </div>

      <div v-else-if="isCompletedExpanded && completedTasks.length === 0" class="empty-task-hint">
        尚無完成紀錄，完成任務後成果將自動收錄於此！
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
  privateTasks: {
    type: Array,
    default: () => []
  },
  publicTasks: {
    type: Array,
    default: () => []
  },
  completedTasks: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['execute-task', 'create-private-task', 'complete-task'])

const showAddPrivate = ref(false)
const newPrivateTitle = ref('')
const isCompletedExpanded = ref(true)

function submitPrivateTask() {
  if (!newPrivateTitle.value.trim()) return
  emit('create-private-task', newPrivateTitle.value.trim())
  newPrivateTitle.value = ''
  showAddPrivate.value = false
}
</script>

<style scoped>
.tasks-view {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.view-header {
  background: #FFFFFF;
  padding: 12px 18px;
  border-radius: 8px;
  border: 1px solid #DBDBD9;
}

.view-header-title {
  color: #141417;
  font-size: 24px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  line-height: 1.2;
}

.task-category-card {
  width: 100%;
  padding: 16px;
  background: #FFFFFF;
  border-radius: 8px;
  outline: 1px solid #DBDBD9;
  outline-offset: -1px;
  display: flex;
  flex-direction: column;
  gap: 14px;
  box-sizing: border-box;
}

.section-title-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.section-title {
  color: #141417;
  font-size: 16px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  margin: 0;
}

.btn-add-mini {
  padding: 4px 10px;
  border-radius: 12px;
  border: 1px solid #DBDBD9;
  background: transparent;
  color: #141417;
  font-size: 12px;
  cursor: pointer;
  transition: all 0.15s ease;
}

.btn-add-mini:hover {
  background: #F0F0ED;
}

.inline-add-row {
  display: flex;
  gap: 8px;
}

.inline-input {
  flex: 1;
  padding: 8px 12px;
  border: 1px solid #DBDBD9;
  border-radius: 6px;
  font-size: 13px;
  outline: none;
}

.btn-dark-inline {
  padding: 6px 14px;
  background: #141417;
  color: #FFFFFF;
  border: none;
  border-radius: 6px;
  font-size: 12px;
  cursor: pointer;
}

.task-items-container {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.task-item-box {
  padding: 12px 14px;
  background: #FFFFFF;
  border-radius: 8px;
  outline: 1px solid #DBDBD9;
  outline-offset: -1px;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 10px;
  box-sizing: border-box;
  transition: all 0.15s ease;
}

.task-item-box:hover {
  outline-color: #B8B8B8;
}

.task-info-col {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 3px;
}

.task-title-text {
  color: #141417;
  font-size: 14px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  line-height: 1.3;
}

.task-sub-text {
  color: #75757A;
  font-size: 12px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 400;
  line-height: 1.3;
}

.task-btn-row {
  display: flex;
  justify-content: flex-end;
}

.btn-execute-pill {
  width: 88px;
  padding: 7px 12px;
  background: #FFFFFF;
  border-radius: 18px;
  outline: 1px solid #E5E5E5;
  outline-offset: -1px;
  border: none;
  color: #141417;
  font-size: 13px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.15s ease;
  display: flex;
  justify-content: center;
  align-items: center;
}

.btn-execute-pill:hover {
  background: #141417;
  color: #FFFFFF;
  outline-color: #141417;
}

.completed-section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  cursor: pointer;
  user-select: none;
}

.count-tag {
  color: #75757A;
  font-size: 12px;
}

.completed-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
  padding-top: 4px;
}

.completed-item-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 12px;
  background: #F8F8F6;
  border-radius: 6px;
}

.completed-info {
  display: flex;
  align-items: center;
  gap: 8px;
}

.completed-check-icon {
  color: #141417;
  font-weight: 700;
}

.completed-title {
  color: #141417;
  font-size: 13px;
}

.completed-status-tag {
  color: #75757A;
  font-size: 11px;
}

.empty-task-hint {
  padding: 16px 0;
  text-align: center;
  color: #75757A;
  font-size: 12px;
}
</style>
