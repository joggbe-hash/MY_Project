<template>
  <div class="settings-view">
    <!-- Top Header Bar -->
    <header class="settings-header-bar">
      <div class="header-left">
        <button class="back-link-btn" @click="$emit('back')">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <polyline points="15 18 9 12 15 6"/>
          </svg>
          <span>返回個人檔案</span>
        </button>
        <h1 class="settings-page-title">編輯帳號設定</h1>
      </div>

      <div class="header-right">
        <button class="btn-save-primary" :disabled="isSaving" @click="handleSave">
          <svg v-if="!isSaving" width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2">
            <polyline points="20 6 9 17 4 12"/>
          </svg>
          <span>{{ isSaving ? '儲存中...' : '儲存變更' }}</span>
        </button>
      </div>
    </header>

    <!-- Success Toast Alert -->
    <transition name="fade">
      <div v-if="showSuccessToast" class="success-toast">
        <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="#2E7D32" stroke-width="2.5">
          <polyline points="20 6 9 17 4 12"/>
        </svg>
        <span>個人帳號設定已成功更新！</span>
      </div>
    </transition>

    <div class="settings-content-wrapper">
      <!-- Section 1: 大頭貼與個人基本檔案 -->
      <section class="settings-card">
        <div class="section-header">
          <div class="section-icon">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
              <circle cx="12" cy="7" r="4"/>
            </svg>
          </div>
          <div>
            <h2 class="section-title">個人檔案與頭像</h2>
            <p class="section-desc">更新您的公開名稱、個人簡介與頭貼外觀</p>
          </div>
        </div>

        <!-- Avatar Customizer Row -->
        <div class="avatar-customizer-box">
          <div 
            class="avatar-preview-circle" 
            :style="{ backgroundColor: form.avatar ? 'transparent' : form.avatar_color }"
          >
            <img v-if="form.avatar" :src="form.avatar" alt="Avatar Preview" class="avatar-preview-img" />
            <span v-else class="avatar-preview-char">
              {{ (form.nickname || form.username || 'U').charAt(0).toUpperCase() }}
            </span>
          </div>

          <div class="avatar-controls-col">
            <div class="avatar-title-row">
              <span class="control-label-bold">個人大頭貼</span>
              <span class="control-hint">支援 JPG、PNG 圖片或選擇風格代表色</span>
            </div>

            <div class="avatar-actions-row">
              <label class="btn-upload-avatar">
                <input type="file" accept="image/*" class="file-hidden-input" @change="handleFileUpload" />
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"/>
                  <circle cx="12" cy="13" r="4"/>
                </svg>
                <span>上傳照片</span>
              </label>

              <button v-if="form.avatar" class="btn-remove-avatar" @click="removeCustomAvatar">
                移除照片
              </button>
            </div>

            <!-- Color Palette Presets (When no photo or for initials) -->
            <div class="avatar-colors-row">
              <span class="palette-label">底色風格：</span>
              <div class="color-palette-options">
                <button
                  v-for="color in presetColors"
                  :key="color.hex"
                  type="button"
                  class="color-dot"
                  :class="{ active: form.avatar_color === color.hex }"
                  :style="{ backgroundColor: color.hex }"
                  :title="color.name"
                  @click="form.avatar_color = color.hex"
                ></button>
              </div>
            </div>
          </div>
        </div>

        <div class="form-grid">
          <!-- Username (Handle) -->
          <div class="form-group">
            <label class="form-label">
              <span>帳號名稱 (ID)</span>
              <span class="form-required">*</span>
            </label>
            <div class="input-with-prefix">
              <span class="input-prefix">@</span>
              <input 
                v-model="form.username" 
                type="text" 
                placeholder="例如 mine_account" 
                class="form-input" 
                maxlength="30"
              />
            </div>
            <span class="form-field-hint">社群中的唯一識別碼，其他人可透過搜尋找到您</span>
          </div>

          <!-- Display Nickname -->
          <div class="form-group">
            <label class="form-label">
              <span>顯示暱稱</span>
            </label>
            <input 
              v-model="form.nickname" 
              type="text" 
              placeholder="例如 我的帳號" 
              class="form-input" 
              maxlength="30"
            />
            <span class="form-field-hint">顯示在個人卡片與動態牆上的稱呼</span>
          </div>

          <!-- Bio / Motto -->
          <div class="form-group full-width">
            <div class="label-with-counter">
              <label class="form-label">個人標語 / 座右銘</label>
              <span class="char-count">{{ form.bio?.length || 0 }} / 120</span>
            </div>
            <textarea 
              v-model="form.bio" 
              rows="3" 
              placeholder="分享一句座右銘或您對記錄生活的想法..." 
              class="form-textarea"
              maxlength="120"
            ></textarea>
          </div>

          <!-- Weekly Status / Goal Info -->
          <div class="form-group full-width">
            <label class="form-label">本週目標與進度摘要</label>
            <input 
              v-model="form.weekly_info" 
              type="text" 
              placeholder="例如 本週完成 5 個任務，收藏 8 則靈感" 
              class="form-input" 
            />
            <span class="form-field-hint">會顯示在您的個人檔案主頁，督促自己達成目標</span>
          </div>
        </div>
      </section>

      <!-- Section 2: 帳號資訊與安全性 -->
      <section class="settings-card">
        <div class="section-header">
          <div class="section-icon">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <rect x="3" y="11" width="18" height="11" rx="2" ry="2"/>
              <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
            </svg>
          </div>
          <div>
            <h2 class="section-title">帳號與安全性</h2>
            <p class="section-desc">管理您的聯絡信箱與登入密碼保護</p>
          </div>
        </div>

        <div class="form-grid">
          <!-- Email Address -->
          <div class="form-group full-width">
            <label class="form-label">電子信箱</label>
            <div class="input-with-badge">
              <input 
                v-model="form.email" 
                type="email" 
                placeholder="user@example.com" 
                class="form-input" 
              />
              <span class="verified-badge">
                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
                  <polyline points="20 6 9 17 4 12"/>
                </svg>
                已認證
              </span>
            </div>
          </div>

          <!-- Password Change Accordion / Inputs -->
          <div class="form-group">
            <label class="form-label">目前密碼</label>
            <input 
              v-model="passwordForm.current" 
              :type="showPassword ? 'text' : 'password'" 
              placeholder="••••••••" 
              class="form-input" 
            />
          </div>

          <div class="form-group">
            <label class="form-label">新密碼</label>
            <input 
              v-model="passwordForm.next" 
              :type="showPassword ? 'text' : 'password'" 
              placeholder="至少 6 位字元" 
              class="form-input" 
            />
          </div>

          <div class="form-group full-width password-toggle-row">
            <label class="checkbox-label">
              <input v-model="showPassword" type="checkbox" />
              <span>顯示密碼文字</span>
            </label>
          </div>
        </div>
      </section>

      <!-- Section 3: 隱私與通知偏好 -->
      <section class="settings-card">
        <div class="section-header">
          <div class="section-icon">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
            </svg>
          </div>
          <div>
            <h2 class="section-title">隱私權限與偏好</h2>
            <p class="section-desc">自訂帳號公開狀態與社群推播提醒</p>
          </div>
        </div>

        <div class="toggle-list">
          <!-- Private Account Toggle -->
          <div class="toggle-item">
            <div class="toggle-text-col">
              <span class="toggle-title">設為不公開帳號（私人）</span>
              <span class="toggle-desc">開啟後，僅有您核准的追蹤者能看見您的任務打卡成果與靈感清單</span>
            </div>
            <label class="switch-control">
              <input v-model="form.is_private" type="checkbox" />
              <span class="slider"></span>
            </label>
          </div>

          <!-- Show Albums Toggle -->
          <div class="toggle-item">
            <div class="toggle-text-col">
              <span class="toggle-title">公開成果專輯</span>
              <span class="toggle-desc">允許其他使用者在您的個人檔案中瀏覽公開的成果相簿</span>
            </div>
            <label class="switch-control">
              <input v-model="form.show_albums" type="checkbox" />
              <span class="slider"></span>
            </label>
          </div>

          <!-- Daily Task Reminder -->
          <div class="toggle-item">
            <div class="toggle-text-col">
              <span class="toggle-title">每日任務進度推播</span>
              <span class="toggle-desc">於每日晚間 20:00 提醒您未完成的每日任務打卡</span>
            </div>
            <label class="switch-control">
              <input v-model="form.notify_tasks" type="checkbox" />
              <span class="slider"></span>
            </label>
          </div>

          <!-- Interaction Notifications -->
          <div class="toggle-item">
            <div class="toggle-text-col">
              <span class="toggle-title">貼文與靈感互動通知</span>
              <span class="toggle-desc">當其他使用者收藏您的靈感、或加入您發起的任務時通知您</span>
            </div>
            <label class="switch-control">
              <input v-model="form.notify_likes" type="checkbox" />
              <span class="slider"></span>
            </label>
          </div>
        </div>
      </section>

      <!-- Section 4: 帳號管理與危險區域 -->
      <section class="settings-card danger-zone-card">
        <div class="section-header">
          <div class="section-icon danger-icon">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="#D32F2F" stroke-width="2">
              <circle cx="12" cy="12" r="10"/>
              <line x1="12" y1="8" x2="12" y2="12"/>
              <line x1="12" y1="16" x2="12.01" y2="16"/>
            </svg>
          </div>
          <div>
            <h2 class="section-title text-danger">危險區域</h2>
            <p class="section-desc">管理登出與帳號註銷作業</p>
          </div>
        </div>

        <div class="danger-actions-list">
          <div class="danger-action-row">
            <div>
              <div class="danger-title">登出此帳號</div>
              <div class="danger-desc">登出後需重新輸入帳號密碼方可再次登入</div>
            </div>
            <button class="btn-danger-outline" @click="$emit('logout')">
              登出帳號
            </button>
          </div>

          <div class="danger-action-row">
            <div>
              <div class="danger-title">暫時停用帳號</div>
              <div class="danger-desc">隱藏個人檔案與所有貼文，隨時重新登入即可復原</div>
            </div>
            <button class="btn-danger-subtle" @click="handleDeactivate">
              停用帳號
            </button>
          </div>
        </div>
      </section>

      <!-- Bottom Sticky Action Buttons -->
      <div class="bottom-action-bar">
        <button class="btn-cancel" @click="$emit('back')">
          取消
        </button>
        <button class="btn-save-primary large-save-btn" :disabled="isSaving" @click="handleSave">
          <svg v-if="!isSaving" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2">
            <polyline points="20 6 9 17 4 12"/>
          </svg>
          <span>{{ isSaving ? '正在儲存中...' : '確認儲存所有設定' }}</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'

const props = defineProps({
  currentUser: {
    type: Object,
    default: () => ({})
  },
  profile: {
    type: Object,
    default: () => ({})
  }
})

const emit = defineEmits(['back', 'save', 'logout'])

const isSaving = ref(false)
const showSuccessToast = ref(false)
const showPassword = ref(false)

const presetColors = [
  { name: '經典黑曜', hex: '#141417' },
  { name: '天際寶藍', hex: '#2563EB' },
  { name: '薰衣草紫', hex: '#7C3AED' },
  { name: '活力陽光', hex: '#EA580C' },
  { name: '森林翡翠', hex: '#059669' },
  { name: '玫瑰乾燥', hex: '#E11D48' }
]

const form = reactive({
  username: props.currentUser?.username || props.profile?.user?.username || 'mine_account',
  nickname: props.currentUser?.nickname || props.profile?.user?.nickname || '我的帳號',
  bio: props.currentUser?.bio || props.profile?.user?.bio || '把靈感變成行動，把行動留下成果',
  weekly_info: props.currentUser?.weekly_info || props.profile?.user?.weekly_info || '本週完成 5 個任務，收藏 8 則靈感',
  avatar: props.currentUser?.avatar || props.profile?.user?.avatar || '',
  avatar_color: props.currentUser?.avatar_color || props.profile?.user?.avatar_color || '#141417',
  email: props.currentUser?.email || props.profile?.user?.email || 'user@example.com',
  is_private: props.currentUser?.is_private || false,
  show_albums: props.currentUser?.show_albums ?? true,
  notify_tasks: props.currentUser?.notify_tasks ?? true,
  notify_likes: props.currentUser?.notify_likes ?? true
})

const passwordForm = reactive({
  current: '',
  next: ''
})

watch(() => props.currentUser, (newVal) => {
  if (newVal) {
    if (newVal.username) form.username = newVal.username
    if (newVal.nickname) form.nickname = newVal.nickname
    if (newVal.bio) form.bio = newVal.bio
    if (newVal.weekly_info) form.weekly_info = newVal.weekly_info
    if (newVal.avatar !== undefined) form.avatar = newVal.avatar
    if (newVal.avatar_color) form.avatar_color = newVal.avatar_color
    if (newVal.email) form.email = newVal.email
    if (newVal.is_private !== undefined) form.is_private = newVal.is_private
  }
}, { deep: true })

function handleFileUpload(e) {
  const file = e.target.files?.[0]
  if (!file) return

  if (file.size > 2 * 1024 * 1024) {
    alert('圖片大小建議在 2MB 以內')
    return
  }

  const reader = new FileReader()
  reader.onload = (event) => {
    form.avatar = event.target.result
  }
  reader.readAsDataURL(file)
}

function removeCustomAvatar() {
  form.avatar = ''
}

function handleDeactivate() {
  const confirmed = confirm('確定要暫時停用帳號嗎？停用後其他使用者將暫時無法搜尋到您。')
  if (confirmed) {
    alert('已為您暫時隱藏個人資料，重新登入即可隨時復原。')
    emit('logout')
  }
}

async function handleSave() {
  if (!form.username.trim()) {
    alert('請輸入帳號名稱')
    return
  }

  isSaving.value = true
  try {
    const payload = {
      username: form.username.trim().replace(/^@/, ''),
      nickname: form.nickname.trim(),
      bio: form.bio.trim(),
      weekly_info: form.weekly_info.trim(),
      avatar: form.avatar,
      avatar_color: form.avatar_color,
      email: form.email.trim(),
      is_private: form.is_private,
      show_albums: form.show_albums,
      notify_tasks: form.notify_tasks,
      notify_likes: form.notify_likes
    }

    emit('save', payload)

    showSuccessToast.value = true
    setTimeout(() => {
      showSuccessToast.value = false
      emit('back')
    }, 900)
  } finally {
    isSaving.value = false
  }
}
</script>

<style scoped>
.settings-view {
  width: 100%;
  max-width: 680px;
  margin: 0 auto;
  padding: 0 16px 80px;
  box-sizing: border-box;
  font-family: 'Noto Sans TC', sans-serif;
  color: #141417;
}

/* Header Bar */
.settings-header-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 18px 0 16px;
  border-bottom: 1px solid #ECECE8;
  margin-bottom: 20px;
  position: sticky;
  top: 0;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(8px);
  z-index: 20;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.back-link-btn {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  padding: 6px 12px;
  border-radius: 20px;
  border: 1px solid #DBDBD9;
  background: #FFFFFF;
  font-size: 13px;
  font-weight: 500;
  color: #141417;
  cursor: pointer;
  transition: all 0.15s ease;
}

.back-link-btn:hover {
  background: #F4F4F1;
  border-color: #BFBFBD;
}

.settings-page-title {
  margin: 0;
  font-size: 20px;
  font-weight: 700;
  color: #141417;
}

.btn-save-primary {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  background: #141417;
  color: #FFFFFF;
  border: none;
  border-radius: 20px;
  padding: 7px 16px;
  font-size: 13px;
  font-weight: 600;
  font-family: inherit;
  cursor: pointer;
  transition: all 0.15s ease;
}

.btn-save-primary:hover:not(:disabled) {
  background: #2E2E33;
  transform: translateY(-1px);
}

.btn-save-primary:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

/* Toast */
.success-toast {
  display: flex;
  align-items: center;
  gap: 8px;
  background: #E8F5E9;
  border: 1px solid #C8E6C9;
  color: #2E7D32;
  padding: 10px 16px;
  border-radius: 10px;
  font-size: 13px;
  font-weight: 600;
  margin-bottom: 16px;
  box-shadow: 0 4px 12px rgba(46, 125, 50, 0.12);
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

/* Content Container */
.settings-content-wrapper {
  display: flex;
  flex-direction: column;
  gap: 18px;
}

/* Settings Cards */
.settings-card {
  background: #FFFFFF;
  border: 1px solid #DBDBD9;
  border-radius: 12px;
  padding: 20px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.02);
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.section-header {
  display: flex;
  align-items: center;
  gap: 12px;
  border-bottom: 1px solid #F0F0ED;
  padding-bottom: 14px;
}

.section-icon {
  width: 36px;
  height: 36px;
  border-radius: 10px;
  background: #F4F4F1;
  color: #141417;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.section-title {
  margin: 0;
  font-size: 16px;
  font-weight: 700;
  color: #141417;
}

.section-desc {
  margin: 2px 0 0;
  font-size: 12px;
  color: #75757A;
}

/* Avatar Customizer */
.avatar-customizer-box {
  display: flex;
  align-items: center;
  gap: 18px;
  background: #FAFAFA;
  border: 1px solid #EBEBE8;
  border-radius: 10px;
  padding: 14px;
}

.avatar-preview-circle {
  width: 72px;
  height: 72px;
  border-radius: 50%;
  border: 2px solid #FFFFFF;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  display: flex;
  align-items: center;
  justify-content: center;
  color: #FFFFFF;
  font-size: 26px;
  font-weight: 700;
  overflow: hidden;
  flex-shrink: 0;
}

.avatar-preview-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar-controls-col {
  display: flex;
  flex-direction: column;
  gap: 8px;
  flex: 1;
}

.avatar-title-row {
  display: flex;
  flex-direction: column;
}

.control-label-bold {
  font-size: 13px;
  font-weight: 700;
  color: #141417;
}

.control-hint {
  font-size: 11px;
  color: #75757A;
}

.avatar-actions-row {
  display: flex;
  align-items: center;
  gap: 8px;
}

.btn-upload-avatar {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 5px 12px;
  border-radius: 16px;
  background: #FFFFFF;
  border: 1px solid #DBDBD9;
  font-size: 12px;
  font-weight: 500;
  color: #141417;
  cursor: pointer;
  transition: all 0.15s ease;
}

.btn-upload-avatar:hover {
  background: #F0F0ED;
}

.file-hidden-input {
  display: none;
}

.btn-remove-avatar {
  background: none;
  border: none;
  font-size: 12px;
  color: #D32F2F;
  cursor: pointer;
  padding: 4px 8px;
}

.btn-remove-avatar:hover {
  text-decoration: underline;
}

.avatar-colors-row {
  display: flex;
  align-items: center;
  gap: 8px;
}

.palette-label {
  font-size: 11px;
  color: #75757A;
}

.color-palette-options {
  display: flex;
  gap: 6px;
}

.color-dot {
  width: 20px;
  height: 20px;
  border-radius: 50%;
  border: 2px solid transparent;
  cursor: pointer;
  padding: 0;
  transition: transform 0.15s ease, border-color 0.15s ease;
}

.color-dot:hover {
  transform: scale(1.15);
}

.color-dot.active {
  border-color: #141417;
  transform: scale(1.15);
}

/* Forms */
.form-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 14px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.form-group.full-width {
  grid-column: 1 / -1;
}

.form-label {
  font-size: 12px;
  font-weight: 600;
  color: #2D2D30;
  display: flex;
  align-items: center;
  gap: 3px;
}

.form-required {
  color: #E53935;
}

.form-field-hint {
  font-size: 11px;
  color: #8C8C91;
}

.label-with-counter {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.char-count {
  font-size: 11px;
  color: #8C8C91;
}

.form-input {
  height: 38px;
  padding: 0 12px;
  border-radius: 8px;
  border: 1px solid #DBDBD9;
  background: #FFFFFF;
  font-size: 13px;
  font-family: inherit;
  color: #141417;
  outline: none;
  transition: border-color 0.15s ease, box-shadow 0.15s ease;
}

.form-input:focus {
  border-color: #141417;
  box-shadow: 0 0 0 2px rgba(20, 20, 23, 0.08);
}

.input-with-prefix {
  display: flex;
  align-items: center;
  border: 1px solid #DBDBD9;
  border-radius: 8px;
  background: #FFFFFF;
  overflow: hidden;
  transition: border-color 0.15s ease;
}

.input-with-prefix:focus-within {
  border-color: #141417;
  box-shadow: 0 0 0 2px rgba(20, 20, 23, 0.08);
}

.input-prefix {
  padding: 0 10px;
  font-size: 13px;
  color: #75757A;
  background: #F7F7F5;
  border-right: 1px solid #DBDBD9;
  height: 38px;
  display: flex;
  align-items: center;
}

.input-with-prefix .form-input {
  border: none;
  flex: 1;
}

.input-with-prefix .form-input:focus {
  box-shadow: none;
}

.input-with-badge {
  position: relative;
  display: flex;
  align-items: center;
}

.input-with-badge .form-input {
  width: 100%;
  padding-right: 76px;
}

.verified-badge {
  position: absolute;
  right: 10px;
  display: inline-flex;
  align-items: center;
  gap: 3px;
  background: #E8F5E9;
  color: #2E7D32;
  font-size: 11px;
  font-weight: 600;
  padding: 2px 6px;
  border-radius: 6px;
}

.form-textarea {
  padding: 10px 12px;
  border-radius: 8px;
  border: 1px solid #DBDBD9;
  background: #FFFFFF;
  font-size: 13px;
  font-family: inherit;
  color: #141417;
  resize: vertical;
  outline: none;
  transition: border-color 0.15s ease;
}

.form-textarea:focus {
  border-color: #141417;
  box-shadow: 0 0 0 2px rgba(20, 20, 23, 0.08);
}

.password-toggle-row {
  margin-top: -6px;
}

.checkbox-label {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: #75757A;
  cursor: pointer;
}

/* Toggles List */
.toggle-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.toggle-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding-bottom: 12px;
  border-bottom: 1px solid #F5F5F2;
}

.toggle-item:last-child {
  border-bottom: none;
  padding-bottom: 0;
}

.toggle-text-col {
  display: flex;
  flex-direction: column;
}

.toggle-title {
  font-size: 13px;
  font-weight: 600;
  color: #141417;
}

.toggle-desc {
  font-size: 11px;
  color: #75757A;
  margin-top: 2px;
}

/* Switch */
.switch-control {
  position: relative;
  display: inline-block;
  width: 44px;
  height: 24px;
  flex-shrink: 0;
}

.switch-control input {
  opacity: 0;
  width: 0;
  height: 0;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #D6D6D4;
  transition: 0.2s;
  border-radius: 24px;
}

.slider:before {
  position: absolute;
  content: "";
  height: 18px;
  width: 18px;
  left: 3px;
  bottom: 3px;
  background-color: white;
  transition: 0.2s;
  border-radius: 50%;
  box-shadow: 0 1px 3px rgba(0,0,0,0.2);
}

input:checked + .slider {
  background-color: #141417;
}

input:checked + .slider:before {
  transform: translateX(20px);
}

/* Danger Zone */
.danger-zone-card {
  border-color: #FFCDD2;
  background: #FFFBFB;
}

.danger-icon {
  background: #FFEBEE;
}

.text-danger {
  color: #D32F2F;
}

.danger-actions-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.danger-action-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding: 10px 12px;
  background: #FFFFFF;
  border: 1px solid #FFEBEE;
  border-radius: 8px;
}

.danger-title {
  font-size: 13px;
  font-weight: 600;
  color: #141417;
}

.danger-desc {
  font-size: 11px;
  color: #75757A;
}

.btn-danger-outline {
  padding: 6px 14px;
  border-radius: 16px;
  border: 1px solid #D32F2F;
  background: #FFFFFF;
  color: #D32F2F;
  font-size: 12px;
  font-weight: 600;
  cursor: pointer;
  white-space: nowrap;
  transition: all 0.15s ease;
}

.btn-danger-outline:hover {
  background: #D32F2F;
  color: #FFFFFF;
}

.btn-danger-subtle {
  padding: 6px 14px;
  border-radius: 16px;
  border: 1px solid #E0E0E0;
  background: #FFFFFF;
  color: #75757A;
  font-size: 12px;
  cursor: pointer;
  white-space: nowrap;
}

.btn-danger-subtle:hover {
  background: #F5F5F5;
  color: #141417;
}

/* Bottom Action Bar */
.bottom-action-bar {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 12px;
  padding-top: 10px;
}

.btn-cancel {
  padding: 8px 18px;
  border-radius: 20px;
  border: 1px solid #DBDBD9;
  background: #FFFFFF;
  font-size: 13px;
  font-weight: 500;
  color: #141417;
  cursor: pointer;
}

.btn-cancel:hover {
  background: #F4F4F1;
}

.large-save-btn {
  padding: 8px 24px;
  font-size: 14px;
}

/* RWD Mobile styles */
@media (max-width: 640px) {
  .form-grid {
    grid-template-columns: 1fr;
  }

  .avatar-customizer-box {
    flex-direction: column;
    text-align: center;
    align-items: center;
  }

  .avatar-actions-row, .avatar-colors-row {
    justify-content: center;
  }

  .danger-action-row {
    flex-direction: column;
    align-items: flex-start;
    gap: 10px;
  }

  .danger-action-row button {
    align-self: flex-end;
  }
}
</style>
