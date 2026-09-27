<template>
  <div class="task-execution-view">
    <!-- View Header (Figma Screen 10: 執行記錄) -->
    <div class="view-header execution-header-bar">
      <button class="back-btn" @click="$emit('back')" title="返回我的任務">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
          <polyline points="15 18 9 12 15 6"/>
        </svg>
      </button>
      <h1 class="view-header-title">執行記錄</h1>
    </div>

    <!-- Active Task Card -->
    <div class="figma-card task-hero-card">
      <div class="status-pill">進行中</div>
      <h2 class="execution-task-title">{{ currentTask.title }}</h2>
      <p class="execution-task-desc">
        {{ currentTask.description || '用照片或文字逐步記錄執行過程，完成後發布成果。' }}
      </p>
    </div>

    <!-- Process Section Header -->
    <div class="execution-section-header">
      <h3 class="section-heading">過程記錄</h3>
      <p class="section-subtext">依序加入每個步驟，可使用文字或照片</p>
    </div>

    <!-- Execution Steps List -->
    <div class="execution-steps-container">
      <!-- Dynamic Steps Loop -->
      <div 
        v-for="(step, index) in steps" 
        :key="step.id || index" 
        class="step-card figma-card"
      >
        <div class="step-badge">{{ index + 1 }}</div>

        <!-- Step content: photo or text -->
        <div v-if="step.type === 'photo'" class="step-media-box">
          <img v-if="step.image_url" :src="step.image_url" class="step-preview-img" alt="步驟相片" />
          <div v-else class="step-media-placeholder">
            <span>{{ step.content || '一張相片' }}</span>
          </div>
        </div>

        <div v-else-if="step.type === 'text'" class="step-text-box">
          <p class="step-text-content">{{ step.content }}</p>
        </div>

        <div class="step-footer-row">
          <span class="step-time">{{ step.time_text || '記錄時間 今天 14:20' }}</span>
          <button 
            v-if="steps.length > 1" 
            class="step-delete-btn" 
            @click="removeStep(index)" 
            title="刪除此步驟"
          >
            ✕
          </button>
        </div>
      </div>

      <!-- Current Adding Step Card (Matching Step 2 in Figma Mockup) -->
      <div class="step-card figma-card active-add-step-card">
        <div class="step-badge">{{ steps.length + 1 }}</div>
        
        <div class="step-actions-row">
          <!-- Add Text Button Box -->
          <button class="step-add-box" @click="openAddTextModal" type="button">
            <span class="plus-icon">＋</span>
            <span class="add-box-label">新增文字</span>
          </button>

          <!-- Add Photo Button Box -->
          <button class="step-add-box" @click="triggerPhotoPicker" type="button">
            <span class="plus-icon">＋</span>
            <span class="add-box-label">拍照或從相簿選擇</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Hidden native file input for photo selection -->
    <input 
      ref="fileInputRef" 
      type="file" 
      accept="image/*" 
      style="display: none" 
      @change="handlePhotoUpload"
    />

    <!-- Bottom Floating Mascot & Action Controls -->
    <div class="execution-bottom-container">
      <!-- Floating Dragon / Mascot Bubble -->
      <div class="mascot-wrapper">
        <div 
          class="floating-mascot-bubble" 
          @click="toggleMascotTip" 
          title="靈感小夥伴"
        >
          <!-- Mascot Illustration SVG (Cute Golden Companion) -->
          <svg class="mascot-svg" width="36" height="36" viewBox="0 0 100 100" fill="none">
            <!-- Glow & Body -->
            <circle cx="50" cy="50" r="46" fill="#FFF9E6" />
            <!-- Little dragon ears/horns -->
            <path d="M30 28C26 18 36 14 40 22" stroke="#F59E0B" stroke-width="4" stroke-linecap="round"/>
            <path d="M70 28C74 18 64 14 60 22" stroke="#F59E0B" stroke-width="4" stroke-linecap="round"/>
            <!-- Body & Face -->
            <ellipse cx="50" cy="54" rx="26" ry="24" fill="#FBBF24" />
            <!-- Cheeks -->
            <circle cx="36" cy="56" r="4" fill="#F87171" opacity="0.8"/>
            <circle cx="64" cy="56" r="4" fill="#F87171" opacity="0.8"/>
            <!-- Eyes -->
            <circle cx="42" cy="50" r="3.5" fill="#1F2937" />
            <circle cx="58" cy="50" r="3.5" fill="#1F2937" />
            <circle cx="43" cy="48.5" r="1.2" fill="#FFFFFF" />
            <circle cx="59" cy="48.5" r="1.2" fill="#FFFFFF" />
            <!-- Smile -->
            <path d="M47 57Q50 61 53 57" stroke="#1F2937" stroke-width="2.5" stroke-linecap="round" fill="none"/>
            <!-- Little golden paws -->
            <ellipse cx="37" cy="72" rx="6" ry="4" fill="#F59E0B"/>
            <ellipse cx="63" cy="72" rx="6" ry="4" fill="#F59E0B"/>
            <!-- Sparkle -->
            <path d="M78 26L80 32L86 34L80 36L78 42L76 36L70 34L76 32Z" fill="#F59E0B"/>
          </svg>
        </div>

        <!-- Speech bubble tooltip from mascot -->
        <transition name="fade">
          <div v-if="showMascotTip" class="mascot-speech-bubble" @click="showMascotTip = false">
            <span>{{ mascotQuote }}</span>
          </div>
        </transition>
      </div>

      <!-- Action Buttons Row (儲存目前進度 / 完成任務) -->
      <div class="execution-actions-row">
        <button class="btn-outline execution-btn save-btn" @click="handleSaveProgress">
          儲存目前進度
        </button>
        <button class="btn-dark execution-btn finish-btn" @click="handleFinishTask">
          完成任務
        </button>
      </div>
    </div>

    <!-- Inline Modal for Adding Text Step -->
    <div v-if="showTextModal" class="modal-backdrop" @click.self="showTextModal = false">
      <div class="modal-card add-text-modal">
        <h3 class="modal-title">新增文字步驟</h3>
        <p class="modal-subtitle">寫下執行這個步驟的想法或進度描述</p>
        <textarea 
          v-model="newStepText" 
          placeholder="例如：第一步先在頂樓找到最佳角度觀察今天天空的雲層流動..." 
          class="step-textarea"
          rows="4"
          autofocus
        ></textarea>
        <div class="modal-actions-row">
          <button class="btn-outline" @click="showTextModal = false">取消</button>
          <button class="btn-dark" :disabled="!newStepText.trim()" @click="confirmAddTextStep">
            確認加入
          </button>
        </div>
      </div>
    </div>

    <!-- RUICHIEN 設計註釋: 完成任務「然後要去哪裡？」 -->
    <div v-if="showCompletionModal" class="modal-backdrop" @click.self="showCompletionModal = false">
      <div class="modal-card destination-modal">
        <div class="destination-icon">🎉</div>
        <h3 class="destination-title">恭喜完成任務！</h3>
        <p class="destination-subtitle">
          「{{ currentTask.title }}」執行紀錄已成功保存為成果。接下來你想去哪裡？
        </p>

        <div class="destination-options-list">
          <button class="destination-opt-card" @click="handleNavigateTo('profile')">
            <div class="opt-icon">🏆</div>
            <div class="opt-text">
              <div class="opt-title">前往個人檔案</div>
              <div class="opt-desc">在成果專區查看剛發布的成就記錄</div>
            </div>
            <div class="opt-arrow">›</div>
          </button>

          <button class="destination-opt-card" @click="handleNavigateTo('feed')">
            <div class="opt-icon">💬</div>
            <div class="opt-text">
              <div class="opt-title">前往動態貼文串</div>
              <div class="opt-desc">查看大家的回應與其他夥伴的打卡成果</div>
            </div>
            <div class="opt-arrow">›</div>
          </button>

          <button class="destination-opt-card" @click="handleNavigateTo('tasks')">
            <div class="opt-icon">📋</div>
            <div class="opt-text">
              <div class="opt-title">回到我的任務</div>
              <div class="opt-desc">繼續挑選或安排今日的下一個目標</div>
            </div>
            <div class="opt-arrow">›</div>
          </button>
        </div>
      </div>
    </div>

    <!-- RUICHIEN 設計註釋: 儲存進度「然後要去哪裡？」 -->
    <div v-if="showSaveModal" class="modal-backdrop" @click.self="showSaveModal = false">
      <div class="modal-card destination-modal">
        <div class="destination-icon">💾</div>
        <h3 class="destination-title">進度已儲存！</h3>
        <p class="destination-subtitle">
          目前步驟已完整記錄。接下來你想做甚麼？
        </p>

        <div class="destination-options-list">
          <button class="destination-opt-card" @click="showSaveModal = false">
            <div class="opt-icon">✏️</div>
            <div class="opt-text">
              <div class="opt-title">留在本頁繼續記錄</div>
              <div class="opt-desc">繼續新增文字或拍照上傳步驟</div>
            </div>
            <div class="opt-arrow">›</div>
          </button>

          <button class="destination-opt-card" @click="handleNavigateTo('tasks')">
            <div class="opt-icon">📋</div>
            <div class="opt-text">
              <div class="opt-title">返回我的任務清單</div>
              <div class="opt-desc">隨時可以再次點擊「執行任務」接續記錄</div>
            </div>
            <div class="opt-arrow">›</div>
          </button>
        </div>
      </div>
    </div>

    <!-- Feedback Toast -->
    <transition name="toast-fade">
      <div v-if="toastMessage" class="feedback-toast">
        {{ toastMessage }}
      </div>
    </transition>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  task: {
    type: Object,
    default: () => ({
      id: 1,
      title: '拍一張今天的天空',
      description: '用照片或文字逐步記錄執行過程，完成後發布成果。'
    })
  }
})

const emit = defineEmits(['back', 'navigate', 'save-progress', 'finish-task'])

const currentTask = computed(() => {
  return props.task || {
    id: 1,
    title: '拍一張今天的天空',
    description: '用照片或文字逐步記錄執行過程，完成後發布成果。'
  }
})

// Initial steps array matching the Figma Screen 10 mockup:
// Step 1: "一張相片", "記錄時間 今天 14:20"
const steps = ref([
  {
    id: 1,
    type: 'photo',
    content: '一張相片',
    image_url: null,
    time_text: '記錄時間 今天 14:20'
  }
])

const fileInputRef = ref(null)
const showTextModal = ref(false)
const showCompletionModal = ref(false)
const showSaveModal = ref(false)
const newStepText = ref('')
const toastMessage = ref('')
const showMascotTip = ref(false)

const mascotQuotes = [
  '加油！一步一腳印，隨手記錄下成果吧～',
  '天空每天都不同，就像每天的你一樣特別！',
  '已經開始執行了，真的很棒！',
  '每一步都是成長的足跡 ✨'
]
const mascotQuote = ref(mascotQuotes[0])

function toggleMascotTip() {
  const randomIdx = Math.floor(Math.random() * mascotQuotes.length)
  mascotQuote.value = mascotQuotes[randomIdx]
  showMascotTip.value = !showMascotTip.value
}

function triggerPhotoPicker() {
  if (fileInputRef.value) {
    fileInputRef.value.click()
  }
}

function handlePhotoUpload(event) {
  const file = event.target.files?.[0]
  if (!file) return

  const reader = new FileReader()
  reader.onload = (e) => {
    const now = new Date()
    const hours = String(now.getHours()).padStart(2, '0')
    const minutes = String(now.getMinutes()).padStart(2, '0')
    const timeText = `記錄時間 今天 ${hours}:${minutes}`

    steps.value.push({
      id: Date.now(),
      type: 'photo',
      content: file.name,
      image_url: e.target.result,
      time_text: timeText
    })

    showToast('已新增相片記錄！')
  }
  reader.readAsDataURL(file)
  event.target.value = ''
}

function openAddTextModal() {
  newStepText.value = ''
  showTextModal.value = true
}

function confirmAddTextStep() {
  if (!newStepText.value.trim()) return

  const now = new Date()
  const hours = String(now.getHours()).padStart(2, '0')
  const minutes = String(now.getMinutes()).padStart(2, '0')
  const timeText = `記錄時間 今天 ${hours}:${minutes}`

  steps.value.push({
    id: Date.now(),
    type: 'text',
    content: newStepText.value.trim(),
    time_text: timeText
  })

  showTextModal.value = false
  newStepText.value = ''
  showToast('已新增文字記錄！')
}

function removeStep(index) {
  steps.value.splice(index, 1)
  showToast('已移除步驟')
}

function handleSaveProgress() {
  emit('save-progress', {
    taskId: currentTask.value.id,
    steps: steps.value,
    note: steps.value.map((s, i) => `${i + 1}. ${s.content}`).join('; ')
  })
  showToast('已儲存目前進度！')
  showSaveModal.value = true
}

function handleFinishTask() {
  emit('finish-task', {
    taskId: currentTask.value.id,
    taskTitle: currentTask.value.title,
    steps: steps.value,
    submissionContent: steps.value.map(s => s.content).join('，')
  })
  showCompletionModal.value = true
}

// RUICHIEN 設計註釋: 完成任務「然後要去哪裡？」
function handleNavigateTo(tab) {
  showCompletionModal.value = false
  showSaveModal.value = false
  emit('navigate', tab)
}

function showToast(msg) {
  toastMessage.value = msg
  setTimeout(() => {
    toastMessage.value = ''
  }, 2200)
}
</script>

<style scoped>
.task-execution-view {
  display: flex;
  flex-direction: column;
  gap: 16px;
  position: relative;
  padding-bottom: 24px;
}

/* Header */
.execution-header-bar {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 18px;
}

.back-btn {
  background: none;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text-main);
  padding: 4px;
  border-radius: 6px;
  transition: var(--transition);
}

.back-btn:hover {
  background: var(--bg-subtle);
}

/* Task Hero Card */
.task-hero-card {
  padding: 14px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 7px;
  background: #FFFFFF;
  border: 1px solid #D6D9DB;
  border-radius: 6px;
}

.status-pill {
  display: inline-block;
  padding: 5px 10px;
  background: #1F1F21;
  color: #FFFFFF;
  font-size: 11px;
  font-weight: 500;
  line-height: 16px;
  border-radius: 14px;
}

.execution-task-title {
  font-size: 17px;
  font-weight: 500;
  line-height: 25px;
  color: #1A1A1C;
  margin-top: 2px;
}

.execution-task-desc {
  font-size: 12px;
  font-weight: 400;
  line-height: 17px;
  color: #616369;
}

/* Process Section */
.execution-section-header {
  display: flex;
  flex-direction: column;
  gap: 1px;
  padding: 0 2px;
}

.section-heading {
  font-size: 16px;
  font-weight: 500;
  line-height: 23px;
  color: #1A1A1C;
}

.section-subtext {
  font-size: 11px;
  font-weight: 400;
  line-height: 16px;
  color: #616369;
}

/* Steps List */
.execution-steps-container {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.step-card {
  padding: 14px;
  background: #FFFFFF;
  border: 1px solid #D6D9DB;
  border-radius: 6px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  position: relative;
}

.step-badge {
  width: 26px;
  height: 26px;
  background: #1F1F21;
  color: #FFFFFF;
  border-radius: 13px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 500;
  line-height: 17px;
}

/* Step media / preview */
.step-media-box {
  height: 112px;
  background: #F5F5F7;
  border-radius: 5px;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.step-media-placeholder {
  color: #616369;
  font-size: 13px;
}

.step-preview-img {
  width: 100%;
  max-height: 240px;
  object-fit: cover;
  border-radius: 8px;
}

.step-text-box {
  min-height: 60px;
  padding: 12px 14px;
  background: #F4F4F6;
  border-radius: 8px;
  display: flex;
  align-items: center;
}

.step-text-content {
  font-size: 13px;
  color: var(--text-main);
  line-height: 1.5;
  white-space: pre-wrap;
}

.step-footer-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 11px;
  color: var(--text-secondary);
}

.step-delete-btn {
  background: none;
  border: none;
  color: #999;
  cursor: pointer;
  font-size: 12px;
  padding: 2px 6px;
  border-radius: 4px;
}

.step-delete-btn:hover {
  color: #D32F2F;
  background: #FFEBEE;
}

/* Action step card (Step 2 in screenshot) */
.active-add-step-card {
  background: var(--bg-card);
}

.step-actions-row {
  display: flex;
  gap: 10px;
}

.step-add-box {
  flex: 1;
  height: 112px;
  background: #F5F5F7;
  border: none;
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 5px;
  color: #616369;
  cursor: pointer;
  transition: var(--transition);
  outline: none;
}

.step-add-box:hover {
  background: #ECECF0;
}

.plus-icon {
  font-size: 24px;
  line-height: 35px;
  color: #616369;
}

.add-box-label {
  font-size: 12px;
  line-height: 17px;
  color: #616369;
}

/* Bottom Controls & Floating Mascot */
.execution-bottom-container {
  display: flex;
  flex-direction: column;
  gap: 12px;
  position: relative;
  margin-top: 4px;
}

/* Mascot */
.mascot-wrapper {
  position: absolute;
  right: 12px;
  top: -64px;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  z-index: 20;
}

.floating-mascot-bubble {
  width: 52px;
  height: 52px;
  border-radius: 50%;
  background: #FFFFFF;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.14);
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: transform 0.2s cubic-bezier(0.34, 1.56, 0.64, 1);
  border: 1px solid rgba(0, 0, 0, 0.05);
}

.floating-mascot-bubble:hover {
  transform: scale(1.08) translateY(-2px);
}

.mascot-speech-bubble {
  position: absolute;
  bottom: 60px;
  right: 0;
  background: #141417;
  color: #FFFFFF;
  padding: 8px 12px;
  border-radius: 10px;
  font-size: 12px;
  white-space: nowrap;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  cursor: pointer;
}

.mascot-speech-bubble::after {
  content: '';
  position: absolute;
  bottom: -6px;
  right: 20px;
  width: 0;
  height: 0;
  border-left: 6px solid transparent;
  border-right: 6px solid transparent;
  border-top: 6px solid #141417;
}

/* Execution Action Buttons */
.execution-actions-row {
  display: flex;
  gap: 10px;
}

.execution-btn {
  flex: 1;
  height: 46px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 13px;
  font-weight: 500;
  line-height: 19px;
  border-radius: 6px;
  cursor: pointer;
  transition: var(--transition);
}

.save-btn {
  background: #FFFFFF;
  border: 1px solid #1A1A1C;
  color: #1A1A1C;
}

.save-btn:hover {
  background: #F4F4F6;
}

.finish-btn {
  background: #1F1F21;
  border: none;
  color: #FFFFFF;
}

.finish-btn:hover {
  background: #2A2A30;
}

/* Add Text Modal */
.add-text-modal {
  max-width: 400px;
  width: 90%;
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.step-textarea {
  width: 100%;
  border: 1px solid var(--border-main);
  border-radius: 8px;
  padding: 10px 12px;
  font-family: inherit;
  font-size: 13px;
  resize: vertical;
  outline: none;
}

.step-textarea:focus {
  border-color: #141417;
}

.modal-actions-row {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

/* Feedback Toast */
.feedback-toast {
  position: fixed;
  bottom: 80px;
  left: 50%;
  transform: translateX(-50%);
  background: rgba(20, 20, 23, 0.92);
  color: #FFFFFF;
  padding: 10px 18px;
  border-radius: 20px;
  font-size: 13px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
  z-index: 1000;
  pointer-events: none;
}

.toast-fade-enter-active,
.toast-fade-leave-active,
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s, transform 0.25s;
}

.toast-fade-enter-from,
.toast-fade-leave-to {
  opacity: 0;
  transform: translate(-50%, 10px);
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(6px);
}

/* RUICHIEN 設計註釋: 完成任務/儲存進度去向彈窗樣式 */
.destination-modal {
  max-width: 360px;
  width: 90%;
  text-align: center;
  padding: 24px 20px;
  border-radius: 20px;
}

.destination-icon {
  font-size: 40px;
  margin-bottom: 8px;
}

.destination-title {
  font-size: 18px;
  font-weight: 700;
  color: var(--text-main);
  margin-bottom: 6px;
}

.destination-subtitle {
  font-size: 13px;
  color: var(--text-secondary);
  line-height: 1.5;
  margin-bottom: 20px;
}

.destination-options-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.destination-opt-card {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 14px;
  border-radius: 12px;
  border: 1px solid var(--border-main);
  background: var(--bg-card);
  cursor: pointer;
  text-align: left;
  transition: var(--transition);
}

.destination-opt-card:hover {
  background: var(--bg-subtle);
  border-color: var(--text-main);
  transform: translateY(-1px);
}

.opt-icon {
  font-size: 24px;
  flex-shrink: 0;
}

.opt-text {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.opt-title {
  font-size: 14px;
  font-weight: 700;
  color: var(--text-main);
}

.opt-desc {
  font-size: 11px;
  color: var(--text-secondary);
}

.opt-arrow {
  font-size: 18px;
  color: var(--text-secondary);
  font-weight: 600;
}
</style>
