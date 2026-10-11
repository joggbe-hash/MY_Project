<template>
  <div class="task-execute-view">
    <!-- Top Header -->
    <div class="execute-header">
      <div class="header-back-row">
        <button class="back-btn" @click="$emit('back')">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <polyline points="15 18 9 12 15 6"></polyline>
          </svg>
        </button>
        <h1 class="execute-title">{{ currentTask?.title || '拍一張今天的天空' }}</h1>
      </div>
    </div>

    <!-- Main Content Container -->
    <div class="execute-content-body">
      <!-- Task Description Card -->
      <div class="task-summary-card">
        <div class="summary-label">(任務內文)</div>
        <p class="summary-desc">
          {{ currentTask?.description || '用照片或文字逐步記錄執行過程，完成後發布成果。' }}
        </p>
      </div>

      <!-- Section Title: 過程記錄 -->
      <div class="steps-heading-row">
        <div class="steps-heading-title">過程記錄</div>
        <div class="steps-heading-sub">依序加入每個步驟，可使用文字或照片</div>
      </div>

      <!-- Dynamic Step List -->
      <div class="steps-list">
        <!-- Render existing recorded steps -->
        <div 
          v-for="(step, index) in recordedSteps" 
          :key="step.id || index" 
          class="step-card"
        >
          <div class="step-num-badge">
            {{ index + 1 }}
          </div>

          <!-- Step with image -->
          <div v-if="step.type === 'photo' || step.imageUrl" class="step-photo-preview">
            <img v-if="step.imageUrl && step.imageUrl.startsWith('data:')" :src="step.imageUrl" class="uploaded-img" />
            <div v-else class="step-photo-placeholder">
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#75757A" stroke-width="1.8">
                <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/>
                <circle cx="8.5" cy="8.5" r="1.5"/>
                <polyline points="21 15 16 10 5 21"/>
              </svg>
              <span>{{ step.text || '一張相片' }}</span>
            </div>
          </div>

          <!-- Step with text only -->
          <div v-else class="step-text-content">
            <p>{{ step.text }}</p>
          </div>

          <div class="step-time-label">
            記錄時間 {{ step.time || '今天 14:20' }}
          </div>
        </div>

        <!-- Next Step Creator Card (matching Figma Screen 09 Step 2) -->
        <div class="step-card next-step-builder">
          <div class="step-num-badge">
            {{ recordedSteps.length + 1 }}
          </div>

          <!-- If currently typing text for new step -->
          <div v-if="isAddingText" class="inline-text-form">
            <textarea 
              v-model="newStepText" 
              placeholder="輸入步驟過程心得..." 
              class="step-textarea"
              rows="3"
            ></textarea>
            <div class="form-btn-row">
              <button class="btn-cancel" @click="isAddingText = false">取消</button>
              <button class="btn-confirm" @click="confirmAddTextStep">確認儲存步驟</button>
            </div>
          </div>

          <!-- Default: two action boxes side by side -->
          <div v-else class="step-actions-split">
            <!-- Box 1: 新增文字 -->
            <button class="action-box" @click="isAddingText = true">
              <span class="action-plus">＋</span>
              <span class="action-text">新增文字</span>
            </button>

            <!-- Box 2: 拍照或從相簿選擇 -->
            <label class="action-box upload-label-btn">
              <input type="file" accept="image/*" class="hidden-file-input" @change="handleFileUpload" />
              <span class="action-plus">＋</span>
              <span class="action-text">拍照或從相簿選擇</span>
            </label>
          </div>
        </div>
      </div>
    </div>

    <!-- Bottom Action Footer -->
    <div class="execute-footer-bar">
      <button class="btn-save-progress" @click="handleSaveProgress">
        儲存目前進度
      </button>
      <button class="btn-complete-task" @click="handleCompleteTask">
        完成任務
      </button>
    </div>

    <!-- Success Toast Alert -->
    <div v-if="toastMessage" class="toast-popup">
      {{ toastMessage }}
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
  task: {
    type: Object,
    default: () => ({
      id: 201,
      title: '拍一張今天的天空',
      description: '用照片或文字逐步記錄執行過程，完成後發布成果。'
    })
  }
})

const emit = defineEmits(['back', 'save-progress', 'complete-task'])

const currentTask = ref(props.task)
const isAddingText = ref(false)
const newStepText = ref('')
const toastMessage = ref('')

const recordedSteps = ref([
  {
    id: 1,
    type: 'photo',
    text: '一張相片',
    time: '今天 14:20',
    imageUrl: ''
  }
])

function showToast(msg) {
  toastMessage.value = msg
  setTimeout(() => {
    toastMessage.value = ''
  }, 2200)
}

function confirmAddTextStep() {
  if (!newStepText.value.trim()) return
  const now = new Date()
  const timeStr = `今天 ${String(now.getHours()).padStart(2, '0')}:${String(now.getMinutes()).padStart(2, '0')}`
  recordedSteps.value.push({
    id: Date.now(),
    type: 'text',
    text: newStepText.value.trim(),
    time: timeStr
  })
  newStepText.value = ''
  isAddingText.value = false
  showToast('步驟已加入！')
}

function handleFileUpload(e) {
  const file = e.target.files?.[0]
  if (!file) return
  const reader = new FileReader()
  reader.onload = (event) => {
    const now = new Date()
    const timeStr = `今天 ${String(now.getHours()).padStart(2, '0')}:${String(now.getMinutes()).padStart(2, '0')}`
    recordedSteps.value.push({
      id: Date.now(),
      type: 'photo',
      text: file.name || '成果相片',
      time: timeStr,
      imageUrl: event.target.result
    })
    showToast('照片已上傳加入步驟！')
  }
  reader.readAsDataURL(file)
}

function handleSaveProgress() {
  emit('save-progress', {
    taskId: currentTask.value?.id,
    steps: recordedSteps.value
  })
  showToast('進度已成功儲存！')
}

function handleCompleteTask() {
  emit('complete-task', {
    taskId: currentTask.value?.id,
    taskTitle: currentTask.value?.title,
    steps: recordedSteps.value
  })
}
</script>

<style scoped>
.task-execute-view {
  width: 100%;
  background: #FBFBF9;
  display: flex;
  flex-direction: column;
  position: relative;
  box-sizing: border-box;
}

.execute-header {
  width: 100%;
  height: 60px;
  padding: 0 16px;
  background: #FFFFFF;
  border-bottom: 1px solid #DBDBD9;
  display: flex;
  align-items: center;
  box-sizing: border-box;
}

.header-back-row {
  display: flex;
  align-items: center;
  gap: 10px;
}

.back-btn {
  background: none;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #141417;
  padding: 4px;
  border-radius: 6px;
}

.back-btn:hover {
  background: #F0F0ED;
}

.execute-title {
  color: #141417;
  font-size: 20px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  line-height: 28px;
  margin: 0;
}

.execute-content-body {
  padding: 14px 16px;
  display: flex;
  flex-direction: column;
  gap: 14px;
  box-sizing: border-box;
}

.task-summary-card {
  padding: 14px;
  background: #FFFFFF;
  border-radius: 8px;
  outline: 1px solid #D6D9DB;
  outline-offset: -1px;
  display: flex;
  flex-direction: column;
  gap: 7px;
}

.summary-label {
  color: #1A1A1C;
  font-size: 16px;
  font-weight: 700;
  line-height: 24px;
}

.summary-desc {
  color: #616369;
  font-size: 13px;
  line-height: 19px;
  margin: 0;
}

.steps-heading-row {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.steps-heading-title {
  color: #1A1A1C;
  font-size: 16px;
  font-weight: 700;
}

.steps-heading-sub {
  color: #616369;
  font-size: 12px;
}

.steps-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.step-card {
  padding: 14px;
  background: #FFFFFF;
  border-radius: 8px;
  outline: 1px solid #D6D9DB;
  outline-offset: -1px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.step-num-badge {
  width: 26px;
  height: 26px;
  background: #1F1F21;
  border-radius: 13px;
  color: #FFFFFF;
  font-size: 12px;
  font-weight: 700;
  display: flex;
  align-items: center;
  justify-content: center;
}

.step-photo-preview {
  width: 100%;
  height: 120px;
  background: #F5F5F7;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.uploaded-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.step-photo-placeholder {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  color: #616369;
  font-size: 12px;
}

.step-text-content {
  padding: 10px 12px;
  background: #F5F5F7;
  border-radius: 6px;
  font-size: 13px;
  color: #141417;
  line-height: 1.5;
}

.step-time-label {
  color: #616369;
  font-size: 11px;
}

.step-actions-split {
  display: flex;
  gap: 10px;
  width: 100%;
}

.action-box {
  flex: 1;
  height: 112px;
  background: #F5F5F7;
  border-radius: 6px;
  border: 1px dashed #D6D9DB;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 6px;
  cursor: pointer;
  transition: all 0.15s ease;
  user-select: none;
}

.action-box:hover {
  background: #EAEAEF;
  border-color: #B8B8B8;
}

.action-plus {
  font-size: 24px;
  color: #616369;
  line-height: 1;
}

.action-text {
  font-size: 12px;
  color: #616369;
  text-align: center;
}

.upload-label-btn {
  cursor: pointer;
  position: relative;
}

.hidden-file-input {
  display: none;
}

.inline-text-form {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.step-textarea {
  width: 100%;
  padding: 10px;
  border: 1px solid #DBDBD9;
  border-radius: 6px;
  font-family: inherit;
  font-size: 13px;
  resize: vertical;
  box-sizing: border-box;
}

.form-btn-row {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
}

.btn-cancel {
  padding: 6px 14px;
  border: 1px solid #DBDBD9;
  background: white;
  border-radius: 14px;
  font-size: 12px;
  cursor: pointer;
}

.btn-confirm {
  padding: 6px 14px;
  background: #141417;
  color: white;
  border: none;
  border-radius: 14px;
  font-size: 12px;
  cursor: pointer;
}

.execute-footer-bar {
  width: 100%;
  height: 78px;
  padding: 16px 20px;
  background: #FFFFFF;
  border-top: 1px solid #DBDBD9;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
  box-sizing: border-box;
}

.btn-save-progress {
  flex: 1;
  height: 46px;
  background: #FFFFFF;
  border-radius: 6px;
  border: 1px solid #1A1A1C;
  color: #1A1A1C;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.15s ease;
}

.btn-save-progress:hover {
  background: #F5F5F7;
}

.btn-complete-task {
  flex: 1;
  height: 46px;
  background: #1F1F21;
  border-radius: 6px;
  border: none;
  color: #FFFFFF;
  font-size: 13px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.15s ease;
}

.btn-complete-task:hover {
  background: #333338;
}

.toast-popup {
  position: fixed;
  top: 70px;
  left: 50%;
  transform: translateX(-50%);
  background: rgba(20, 20, 23, 0.92);
  color: #FFFFFF;
  padding: 9px 18px;
  border-radius: 20px;
  font-size: 13px;
  z-index: 9999;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  animation: fadeInDown 0.2s ease;
}

@keyframes fadeInDown {
  from {
    opacity: 0;
    transform: translate(-50%, -10px);
  }
  to {
    opacity: 1;
    transform: translate(-50%, 0);
  }
}
</style>
