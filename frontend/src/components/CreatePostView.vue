<template>
  <div class="create-post-view">
    <!-- Header (Figma Screen 04) -->
    <div class="view-header">
      <div class="view-header-title">我也要發文</div>
      <div class="view-header-subtitle">發布任務或靈感，讓別人也能回應</div>
    </div>

    <!-- Main Card Form -->
    <div class="create-form-card">
      <div class="type-selection-title">選擇貼文類型</div>

      <!-- Type Pills Row -->
      <div class="type-buttons-row">
        <button 
          type="button"
          class="type-btn"
          :class="{ active: postType === 'task' }"
          @click="postType = 'task'"
        >
          任務貼文
        </button>
        <button 
          type="button"
          class="type-btn"
          :class="{ active: postType === 'inspiration' }"
          @click="postType = 'inspiration'"
        >
          靈感貼文
        </button>
      </div>

      <!-- Title Input Box -->
      <div class="title-input-box">
        <input 
          v-model="postTitle" 
          placeholder="貼文標題:" 
          class="title-input"
        />
      </div>

      <!-- Content Textarea Box -->
      <div class="content-textarea-box">
        <textarea 
          v-model="postContent" 
          placeholder="寫下你想邀請別人一起做的事，或記錄一則靈感。"
          class="content-textarea"
          rows="6"
        ></textarea>
      </div>

      <!-- Photo Upload Box (Screen 04: 可加入照片) -->
      <label class="photo-upload-container">
        <input type="file" accept="image/*" class="hidden-file-input" @change="handleImagePick" />
        <div v-if="uploadedPreview" class="photo-has-preview">
          <img :src="uploadedPreview" class="preview-img" />
          <span class="preview-remove-hint">點擊更換照片</span>
        </div>
        <div v-else class="upload-placeholder-content">
          <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="#4B4B50" stroke-width="2">
            <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/>
            <circle cx="8.5" cy="8.5" r="1.5"/>
            <polyline points="21 15 16 10 5 21"/>
          </svg>
          <span class="upload-text">可加入照片</span>
        </div>
      </label>

      <!-- Publish Button -->
      <button 
        class="publish-submit-btn" 
        :disabled="!postTitle.trim()"
        @click="handlePublish"
      >
        發布貼文
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const emit = defineEmits(['publish'])

const postType = ref('task')
const postTitle = ref('')
const postContent = ref('')
const uploadedPreview = ref('')

function handleImagePick(e) {
  const file = e.target.files?.[0]
  if (!file) return
  const reader = new FileReader()
  reader.onload = (event) => {
    uploadedPreview.value = event.target.result
  }
  reader.readAsDataURL(file)
}

function handlePublish() {
  if (!postTitle.value.trim()) return
  emit('publish', {
    post_type: postType.value,
    title: postTitle.value.trim(),
    content: postContent.value.trim() || (postType.value === 'task' ? '一起來挑戰看看！' : '看到的好點子，分享給大家。'),
    image_url: uploadedPreview.value || ''
  })
  postTitle.value = ''
  postContent.value = ''
  uploadedPreview.value = ''
}
</script>

<style scoped>
.create-post-view {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.view-header {
  background: #FFFFFF;
  padding: 12px 18px;
  border-radius: 8px;
  border: 1px solid #DBDBD9;
  display: flex;
  flex-direction: column;
  gap: 3px;
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

.create-form-card {
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

.type-selection-title {
  color: #141417;
  font-size: 16px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
}

.type-buttons-row {
  display: flex;
  gap: 8px;
}

.type-btn {
  width: 130px;
  padding: 7px 12px;
  background: #FFFFFF;
  border-radius: 18px;
  outline: 1px solid #DBDBD9;
  outline-offset: -1px;
  border: none;
  color: #141417;
  font-size: 13px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 400;
  cursor: pointer;
  transition: all 0.15s ease;
  display: flex;
  justify-content: center;
  align-items: center;
}

.type-btn.active {
  background: #141417;
  outline-color: #141417;
  color: #FFFFFF;
  font-weight: 700;
}

.title-input-box {
  width: 100%;
  padding: 12px;
  background: #FFFFFF;
  border-radius: 8px;
  outline: 1px solid #DBDBD9;
  outline-offset: -1px;
  box-sizing: border-box;
}

.title-input {
  width: 100%;
  border: none;
  outline: none;
  font-size: 14px;
  font-family: inherit;
  color: #141417;
}

.title-input::placeholder {
  color: #75757A;
}

.content-textarea-box {
  width: 100%;
  height: 170px;
  padding: 12px;
  background: #FFFFFF;
  border-radius: 8px;
  outline: 1px solid #DBDBD9;
  outline-offset: -1px;
  box-sizing: border-box;
}

.content-textarea {
  width: 100%;
  height: 100%;
  border: none;
  outline: none;
  font-size: 14px;
  font-family: inherit;
  color: #141417;
  resize: none;
}

.content-textarea::placeholder {
  color: #75757A;
  line-height: 1.5;
}

.photo-upload-container {
  width: 100%;
  height: 110px;
  background: #F0F0ED;
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: background 0.15s ease;
}

.photo-upload-container:hover {
  background: #EAEAE6;
}

.hidden-file-input {
  display: none;
}

.upload-placeholder-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
}

.upload-text {
  color: #75757A;
  font-size: 13px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 400;
}

.photo-has-preview {
  width: 100%;
  height: 100%;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.preview-remove-hint {
  position: absolute;
  bottom: 6px;
  padding: 2px 8px;
  background: rgba(0, 0, 0, 0.6);
  color: white;
  border-radius: 10px;
  font-size: 11px;
}

.publish-submit-btn {
  width: 100%;
  padding: 8px 12px;
  background: #141417;
  border-radius: 18px;
  outline: 1px solid #141417;
  outline-offset: -1px;
  border: none;
  color: #FFFFFF;
  font-size: 13px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.15s ease;
  display: flex;
  justify-content: center;
  align-items: center;
}

.publish-submit-btn:hover {
  background: #25252b;
}

.publish-submit-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
</style>
